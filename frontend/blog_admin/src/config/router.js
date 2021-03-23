import VueRouter from 'vue-router'
import Login from '../router/login.vue'

export default new VueRouter({
    routes:[
        {
            path: '/', name: 'login', component: Login,
        },
    ]
})