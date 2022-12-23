import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
const path = require('path')
export default defineConfig({
  plugins: [vue()],

  // [vuestic-ui] Add alias for ~normalize.css.
  resolve: {
    alias: [
        { find: /^~(.*)$/, replacement: '$1' },
        {find: '@', replacement: path.resolve(__dirname,'/src')}
    ],
  },
})
