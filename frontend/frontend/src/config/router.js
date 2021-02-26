import Index from '../router/index.vue'
import Test from '../router/test.vue'
import Error404 from '../router/error404.vue'
import VueRouter from 'vue-router'

export default new VueRouter({
    routes:[
        {
            path: '/', name: 'index', component: Index,
        },
        {
            path: '/blog', name: 'test', component: Test,
        },
        {
            path: '/images', name: 'images', component: Test,
        },
        {
            path: '/diary', name: 'diary', component: Test,
        },
        {
            path: '*', name: 'error404', component: Error404,
        }
    ]
})