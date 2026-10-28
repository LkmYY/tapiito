import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 5173,
    proxy: {
      // 本地联调：Go 后端分成公开(:8080)与管理(:8081)双端口，
      // devgateway(:8090) 按路径+方法分流，等价于生产的 nginx 配置。
      '/api': {
        target: 'http://localhost:8090',
        changeOrigin: true
      },
      '/uploads': {
        target: 'http://localhost:8090',
        changeOrigin: true
      }
    }
  }
})