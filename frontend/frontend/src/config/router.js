import VueRouter from 'vue-router'
import Index from '../router/index.vue'
import Test from '../router/test.vue'
import Error404 from '../router/error404.vue'
import Diary from '../router/diary/diary.vue'
import Images from '../router/images/images.vue'
import Bloglist from '../router/blog/bloglist.vue'
import Blog from '../router/blog/blog.vue'

export default new VueRouter({
    routes:[
        {
            path: '/', name: 'index', component: Index,
        },
        {
            path: '/bloglist',
            name: 'bloglist',
            component: Bloglist,
        },
        {
            path: '/blog/:blogId',
            name: 'blog',
            component: Blog,
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
            path: '/test', name: 'test', component: Test,
        }
    ]
})