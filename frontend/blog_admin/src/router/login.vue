<template>
    <div class="loginWrap">
        <Login></Login>
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
    }
}
</script>
<style lang="stylus" scoped>
.test
    width 100%
    height 100%
    background-color black
</style>