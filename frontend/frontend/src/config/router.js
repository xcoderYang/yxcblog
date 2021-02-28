import VueRouter from 'vue-router'
import Index from '../router/index.vue'
import Test from '../router/test.vue'
import Error404 from '../router/error404.vue'
import Diary from '../router/diary.vue'
import Images from '../router/images.vue'
import Blog from '../router/blog.vue'

export default new VueRouter({
    routes:[
        {
            path: '/', name: 'index', component: Index,
        },
        {
            path: '/blog', name: 'blog', component: Blog,
        },
        {
            path: '/images', name: 'images', component: Images,
        },
        {
            path: '/diary', name: 'diary', component: Diary,
        },
        {
            path: '*', name: 'error404', component: Error404,
        },
        {
            path: 'test', name: 'test', component: Test,
        }
    ]
})