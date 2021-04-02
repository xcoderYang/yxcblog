<template>
    <div class="mainWrap">
        <div class="headerPanel">
            <HEADERTOP></HEADERTOP>
        </div>
        <el-row>
            <el-col class="leftBarPanel" :span="3">
                <LEFTBAR @toggle="sideToggle"></LEFTBAR>
            </el-col>
            <el-col :span="21">
                <router-view></router-view>
            </el-col>
        </el-row>
    </div>
</template>
<script>
import HEADERTOP from "../components/header"
import LEFTBAR from "../components/leftBar"
export default {
    name: 'mainPage',
    components:{
        HEADERTOP,
        LEFTBAR
    },
    data(){
        return{

        }
    },
    mounted(){
        
    },
    methods:{
        logout(){
            this.$store.commit("USER_LOGOUT")
            this.$router.push("/")
        },
        sideToggle(to){
            if (this.$route.fullPath === "/main/"+to){
                return
            }else{
                this.$router.push("/main/"+to)
            }
        }
    }
}
</script>
<style lang="stylus" scoped>
.mainWrap
    width 100%
    height 100%
</style>