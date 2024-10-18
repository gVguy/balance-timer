import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import autoprefixer from 'autoprefixer'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  css: {
    postcss: {
      plugins: [autoprefixer()]
    }
  },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, "./src"),
      '@w': path.resolve(__dirname, "./wailsjs"),
    },
  }
})
