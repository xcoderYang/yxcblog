<template>
    <div class="bgWrap">
        <Login class="login" @login="login" @registry="registry"></Login>
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
            this.$router.push("/main")
            this.$axios.get(this.serverUrl+"/api/user/login", {
                params:data
            })
            .then((ans)=>{
                console.log(ans)
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
        left 70%
        transform translateY(-50%)
.test
    width 100%
    height 100%
    background-color black
</style>