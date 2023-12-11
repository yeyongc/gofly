import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import {resolve} from "path"
import { atRule } from 'postcss'
import postCssPxToRem from "postcss-pxtorem"
import autoprefixer from "autoprefixer"

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve:{
    alias: {
      "@": resolve(__dirname, "src")
    },
    extensions: [".js", ".ts", ".tsx", ".jsx", ".json", ".mjs"]
  },
  css: {
    preprocessorOptions:{
      scss: {
        additionalData: `@use "@/assets/styles/global-scss-var.scss" as *;`,
      }
    },
    postcss:{
      plugins:[
        autoprefixer({
          overrideBrowserslist:[
            "Android 4.1",
            "iOS 7.1",
            "Chrome > 31",
            "ie >= 8",
            "ff > 31",
            "> 1%"
          ],
          grid: true
        }),
        {
          postcssPlugin: "internal:charset-removal",
          AtRule:{
            charset: (atRule) => {
              if (atRule.name === 'charset'){
                atRule.remove()
              }
            }
          }
        },
        postCssPxToRem({
          rootValue: 100,
          propList: ['*'],
          selectorBlackList: ["norem"],
          exclude: /node_modules/i
        })
      ]
    }   
  },
  server:{
    port: 80
  }
})
