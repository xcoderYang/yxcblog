import Index from '../router/index.vue'
import Test from '../router/test.vue'
import VueRouter from 'vue-router'

export default new VueRouter({
    routes:[
        {
            path: '/', name: 'index', component: Index,
        },
        {
            path: '/test', name: 'test', component: Test,
        }
    ]
})