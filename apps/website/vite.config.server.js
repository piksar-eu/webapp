import { defineConfig } from 'vite'
import config from './vite.config'

export default defineConfig({
    ...config,
    build: {
        copyPublicDir: false
    },
    ssr: {
        noExternal: true,
    },
})