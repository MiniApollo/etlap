import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

const hostURL = "http://localhost:8080"

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    // Redirect API request to web server
    proxy: {
      '/food': {
        target: hostURL,
        changeOrigin: true,
      },
      '/order': {
        target: hostURL,
        changeOrigin: true,
      },
      '/auth': {
        target: hostURL,
        changeOrigin: true,
      },
      '/customer': {
        target: hostURL,
        changeOrigin: true,
      },
    },
  },
})
