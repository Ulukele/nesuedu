import { createApp } from 'vue'
import App from './modules/index/App.vue'
import { createVuestic } from 'vuestic-ui'
import router from './router/router'
import 'vuestic-ui/css'
import 'material-design-icons-iconfont/dist/material-design-icons.min.css'

import '../style.css'

createApp(App).use(createVuestic({
    config: {
      colors: {
        variables: {
            mainBlack: "#434343",
            customBlue: "#9ED5F0",
            customGreen: "#87C644",
            customRed: "#FF6A6A",
            customGrey: "#989898",
            backgroundColor: "#D9D9D9",
        }
      },
    },
  })).use(router).mount('#app')
