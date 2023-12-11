import React from 'react'
import ReactDOM from 'react-dom/client'


import App from './App.tsx'
import "normalize.css/normalize.css"
import '@/assets/styles/global.css'
import "@/assets/styles/tsx-module.scss"

(async () => {

  // 初始化系统基础配置信息（确保所有模块的基础数据加载后，再创建ui）

  // 初始化store

  // 初始化路由
  



  // 初始化ui
  ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
      <App />
    </React.StrictMode>,
  )
})()
