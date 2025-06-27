package command

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
)

type CommandMode string

const (
	CommandSync  CommandMode = "sync"
	CommandAsync CommandMode = "async"
)

type CommandHandler interface {
	Handle(context.Context, Command) (any, error)
}

type Command interface {
	GetID() string
	GetCreatedAt() time.Time
	GetMetadata() Metadata
	GetMode() CommandMode
	Marshal() ([]byte, error)
}

type Metadata map[string]string

type BaseCommand struct {
	ID        string      `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	Type      string      `json:"type"`
	Meta      Metadata    `json:"meta,omitempty"`
	Mode      CommandMode `json:"mode,omitempty"`
}

func NewBaseCommand(cmd Command, mode *CommandMode, meta *Metadata) BaseCommand {
	if mode == nil {
		defaultMode := CommandSync
		mode = &defaultMode
	}

	if meta == nil {
		meta = &Metadata{}
	}

	return BaseCommand{
		ID:        uuid.NewString(),
		CreatedAt: time.Now().UTC(),
		Meta:      *meta,
		Mode:      *mode,
		Type:      QualifiedType(cmd),
	}
}

func (c BaseCommand) GetID() string {
	return c.ID
}

func (c BaseCommand) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c BaseCommand) GetMetadata() Metadata {
	return c.Meta
}

func (c BaseCommand) GetMode() CommandMode {
	return c.Mode
}

func (c *BaseCommand) SetMode(mode CommandMode) {
	c.Mode = mode
}

func (c *BaseCommand) AddMetadata(key, val string) {
	c.Meta[key] = val
}

func (c BaseCommand) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

func QualifiedType(cmd any) string {
	t := reflect.TypeOf(cmd)
	if t == nil {
		return "unknown.Unknown"
	}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	pkgPath := t.PkgPath()
	pkgParts := strings.Split(pkgPath, "/")

	typeName := t.Name()
	prefix := "unknown"

	if len(pkgParts) >= 2 && pkgParts[len(pkgParts)-1] == "" {
		prefix = pkgParts[len(pkgParts)-2]
		return prefix + "." + typeName
	}

	return t.PkgPath() + "." + typeName
}
