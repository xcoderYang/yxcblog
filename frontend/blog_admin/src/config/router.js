import VueRouter from 'vue-router'
import Login from '../router/login.vue'
import Main from '../router/main.vue'
import BlogList from "../router/main/bloglist"

export default new VueRouter({
    routes:[
        {
            path: '/', name: 'login', component: Login,
        },
        {
            path: '/main', name: 'main', component: Main,redirect:'/main/bloglist',
            children:[{
                path: 'bloglist',
                component: BlogList
            },{
                path: 'picture',
                component: ()=>import("../router/main/picture")
            },{
                path: 'userManage',
                component: ()=>import("../router/main/userManage")
            }]
        }
    ]
})