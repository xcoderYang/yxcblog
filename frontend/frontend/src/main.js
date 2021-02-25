import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import RouterCfg from './config/router.js'

Vue.config.productionTip = false
Vue.use(VueRouter)
new Vue({
  render: h => h(App),
  router: RouterCfg
}).$mount('#app')
