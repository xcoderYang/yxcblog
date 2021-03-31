<template>
    <div class="bgWrap">
        <Login class="login" @login="login" @registry="registry" :onlyLogin="false"></Login>
    </div>
</template>
<script>
import Login from "../components/login.vue";
export default {
    name: 'test',
    components:{
        Login
    },
    data(){
        return {
            
        }
    },
    methods:{
        login(data){
            this.$axios.get(this.serverUrl+"/api/user/login", {
                params:data
            })
            .then((res)=>{
                if(res.data.success){
                    console.log(res)
                    let cookie = document.cookie
                    let sessionId = cookie.slice(cookie.indexOf("sessionId")).split("; ")[0].split("=")[1]
                    let userInfo = JSON.parse(res.data.data)
                    userInfo.sessionId = sessionId
                    this.$store.commit("USER_LOGIN", userInfo)
                    this.$router.push("/main")
                }else{
                    alert(res.data.msg)
                }
            })
            .catch(()=>{
                console.log("No")
            })
        },
        registry(data){
            this.$axios.get(this.serverUrl+"/api/user/registry", {
                params:data
            })
            .then(()=>{
                console.log("Yes")
            })
            .catch(()=>{
                console.log("No")
            })
        }
    }
}
</script>
<style lang="stylus" scoped>
.bgWrap
    height 100%
    width 100%
    .login
        position relative
        top 45%
        left 60%
        transform translateY(-50%)
.test
    width 100%
    height 100%
    background-color black
</style>