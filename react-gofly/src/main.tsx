import React from 'react'
import ReactDOM from 'react-dom/client'


import App from './App.tsx'

// global styles
import "normalize.css/normalize.css"
import '@/assets/styles/global.css'
import "@/assets/styles/tsx-module.scss"


 // 初始化ui
 ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
