import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  css: {
    preprocessorOptions: {
      scss: {
        api: "modern",
      },
    },
  }, server: {
    hmr: {
      host: 'localhost',
      protocol: 'ws',
    },
  },
  resolve: {
    alias: [
      { find: "@wails", replacement: path.resolve(__dirname, "./wailsjs/go/") },
      { find: "@views", replacement: path.resolve(__dirname, "./src/views/") },
      { find: "@router", replacement: path.resolve(__dirname, "./src/router/") },
      { find: "@comp", replacement: path.resolve(__dirname, "./src/components/") },
    ]
  },
  define:{
    '__VUE_PROD_HYDRATION_MISMATCH_DETAILS__': false,
  }
})
