import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import Vuex from 'vuex'
import RouterCfg from './config/router.js'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import Showdown from 'showdown'
import Axios from "axios"
import Moment from "moment"

console.log(Showdown)
Vue.config.productionTip = false
Vue.use(VueRouter)
Vue.use(ElementUI)
Vue.use(Showdown)
Vue.use(Vuex)

Vue.prototype.serverUrl = "http://localhost:8082"
Vue.prototype.$showdown = Showdown
Vue.prototype.$axios = Axios
Vue.prototype.$moment = Moment

new Vue({
  render: h => h(App),
  router: RouterCfg,
  store: Vuex,  
}).$mount('#app')
