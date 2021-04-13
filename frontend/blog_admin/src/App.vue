<template>
  <div id="app">
    <router-view></router-view>
    <el-dialog 
    :visible.sync="showLoginBox"
    center
    width="400px"
    :close-on-click-modal=false
    :close-on-press-escape=false
    :show-close=false
    >
        <LOGIN :onlyLogin="true" @login="dialogLogin"></LOGIN>
    </el-dialog>
  </div>
</template>

<script>

import LOGIN from "./components/login"

export default {
  name: 'App',
  components: {
    LOGIN
  },
  methods:{
    dialogLogin(data){
      this.$axios.get(this.serverUrl+"/api/user/login", {
          params:data
      })
      .then((res)=>{
          if(res.data.success){
              let cookie = document.cookie
              let sessionId = cookie.slice(cookie.indexOf("sessionId")).split("; ")[0].split("=")[1]
              let userInfo = JSON.parse(res.data.data)
              userInfo.sessionId = sessionId
              this.$store.commit("USER_LOGIN", userInfo)
              this.$router.go(0)
              // this.showLoginBox = false
              // if(this.blockedNextUrl){
              //   this.$router.push(this.blockedNextUrl)
              //   this.blockedNextUrl = ""
              // }
          }else{
              alert(res.data.msg)
          }
      })
      .catch((err)=>{
          alert(err.response.data.msg)
      })
    }
  },
  data(){
    return {
      showLoginBox: false,
      blockedNextUrl: "",
    }
  },
  // 这里要用 created是因为，父组件的 mounted在子组件的 mounted之后，如果在其后添加拦截，这里的拦截不会生效
  // 而父组件的 created在子组件的 created,mounted之前，所以在这里添加拦截
  created(){
    let that = this
    window.that = this
    this.$store.commit("USER_INFO_SYNC")
    this.$axios.interceptors.response.use((response)=>{
      return response;
    }, (error)=>{
      if(error.response.status === 401 && that.$route.path!="/"){
        that.showLoginBox = true
      }
      return Promise.reject(error);
    });
    this.$router.beforeEach((to, from ,next)=>{
      let isLogin = that.$store.state.user.userLogin
      if(!isLogin){
        // 从login页面跳往其他页面
        if(from.path == "/" && to.path != "/"){
          next("/")
          return
        }
        // 其他页面之间互相跳转
        else if(to.path != "/" && from.path!="/"){
          // 用户登录后可继续跳转
          that.blockedNextUrl = to.path
          that.showLoginBox = true
          return
        }
      }
      next()
    })
  },
  mounted(){
  }
}
</script>

<style lang="stylus">
html, body
  width 100%
  height 100%
  min-width 1000px
  padding 0
  margin 0
  overflow hidden
i
  font-style normal
#wrap
  width 100%
  height 100%
#app
  width 100%
  height 100%
  display flex
  align-items center
  justify-content center
</style>