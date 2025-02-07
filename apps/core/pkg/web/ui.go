package web

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"rogchap.com/v8go"
)

//go:embed static/*
var static embed.FS
var serverEntryContent string
var indexHTMLContent string
var isolatePool isoPool
var jsCache = make(map[string]string)

func init() {
	serverEntry, err := static.ReadFile("static/website/server/entry-server.js")
	if err != nil {
		log.Panicln("entry-server.js does not exist", err)
	}
	serverEntryContent = string(serverEntry)

	indexHTML, err := static.ReadFile("static/website/client/index.html")
	if err != nil {
		log.Panicln("index.html does not exist", err)
	}
	indexHTMLContent = string(indexHTML)

	isolatePool = isoPool{
		pool: make(chan isoCtx, 10),
	}
}

func ServeUi(mux *http.ServeMux) {
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		filepath := fmt.Sprintf("static/website/client%s", r.RequestURI)

		if r.RequestURI != "/" && fileExists(filepath) {
			if strings.HasPrefix(r.RequestURI, "/assets/index-") && strings.HasSuffix(r.RequestURI, ".js") {
				if _, ok := jsCache[r.RequestURI]; !ok {
					content, _ := static.ReadFile(filepath)
					jsCache[r.RequestURI] = replaceViteEnv(string(content))
				}

				w.Header().Set("Content-Type", "application/javascript")
				w.WriteHeader(200)
				w.Write([]byte(jsCache[r.RequestURI]))
				return
			}

			r.URL.Path = filepath

			http.FileServer(http.FS(static)).ServeHTTP(w, r)
			return
		}

		ic := isolatePool.Get()

		renderCmd := fmt.Sprintf(`JSON.stringify(render("%s"))`, r.URL.Path)
		val, err := ic.ctx.RunScript(renderCmd, "entry-server.js")
		if err != nil {
			if jsErr, ok := err.(*v8go.JSError); ok {
				err = fmt.Errorf("JavaScript Error: %v", jsErr.StackTrace)
			}
			log.Panic(err)
		}

		var result map[string]string
		if err := json.Unmarshal([]byte(val.String()), &result); err != nil {
			log.Panicln("Can not parse ssr result", err)
		}

		finalHTML := strings.Replace(indexHTMLContent, "<!--app-head-->", result["head"], 1)
		finalHTML = strings.Replace(finalHTML, "<!--app-html-->", result["html"], 1)

		if u := GetSessionUser(r); u != nil {
			finalHTML = strings.Replace(finalHTML, "<!--app-js-->", fmt.Sprintf(`<script>
				globalThis.user = {
					email: "%s"
				}
			</script>`, u.Email), 1)
		}

		isolatePool.Put(ic)

		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(finalHTML))
	})
}

func fileExists(filename string) bool {
	_, err := static.Open(filename)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return false
		}
	}
	return true
}

func replaceViteEnv(val string) string {
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 && strings.HasPrefix(parts[0], "VITE_") {
			key := parts[0]
			value := parts[1]
			val = strings.ReplaceAll(val, key, value)
		}
	}

	return val
}

type isoCtx struct {
	iso *v8go.Isolate
	ctx *v8go.Context
}

type isoPool struct {
	pool chan isoCtx
}

func (i *isoPool) Get() isoCtx {
	select {
	case ic := <-i.pool:
		return ic
	default:
		iso := v8go.NewIsolate()
		ctx := v8go.NewContext(iso)

		script, err := iso.CompileUnboundScript(serverEntryContent, "entry-server.js", v8go.CompileOptions{})
		if err != nil {
			log.Panicln("Can not compile entry-server.js", err)
		}

		_, err = script.Run(ctx)
		if err != nil {
			if jsErr, ok := err.(*v8go.JSError); ok {
				err = fmt.Errorf("JavaScript Error: %v", jsErr.StackTrace)
			}
			log.Panic(err)
		}

		return isoCtx{
			iso: iso,
			ctx: ctx,
		}
	}
}

func (i *isoPool) Put(ic isoCtx) {
	select {
	case i.pool <- ic:
		break
	default:
		ic.ctx.Close()
		ic.iso.Dispose()
	}
}
