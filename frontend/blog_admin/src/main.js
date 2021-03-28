// import Vue from 'vue'
// import App from './App.vue'

// Vue.config.productionTip = false

// new Vue({
//   render: h => h(App),
// }).$mount('#app')


import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import Vuex from 'vuex'
import RouterCfg from './config/router.js'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import Showdown from 'showdown'
import Axios from "axios"

Vue.config.productionTip = false
Vue.use(VueRouter)
Vue.use(ElementUI)
Vue.use(Showdown)
Vue.use(Vuex)

Vue.prototype.$showdown = Showdown
Vue.prototype.serverUrl = "http://localhost:8080"
Vue.prototype.$axios = Axios

Vue.prototype.$utils = {
    async handlePromise(promise){
        promise.then(data=>[data, null]).catch(err=>[null, err])
    }
}

new Vue({
  render: h => h(App),
  router: RouterCfg,
  store: Vuex,  
}).$mount('#app')
