import VueRouter from 'vue-router'
import Login from '../router/login.vue'
import Main from '../router/main.vue'

export default new VueRouter({
    routes:[
        {
            path: '/', name: 'login', component: Login,
        },
        {
            path: '/main', name: 'main', component: Main
        }
    ]
})