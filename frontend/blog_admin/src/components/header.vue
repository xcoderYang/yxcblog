<template>
    <div class="headerWrap">
        <el-row class="header">
            <el-col class="logo" :span="6">asdf</el-col>
            <el-col class="gap" :span="10">asdf</el-col>
            <el-col class="setting" :span="8">
                <el-col :span="14">dd</el-col>
                <el-col :span="5">
                    <span>username</span>
                </el-col>
                <el-col :span="5">
                    <span>退出登录</span>
                </el-col>
            </el-col>
        </el-row>
    </div>
</template>
<script>
export default {
    name: "header-top",
    data(){
        return {

        }
    },
    methods:{
        goback(){

        }
    }
}
</script>
<style lang="stylus" scoped>
.headerWrap
    height 60px
    width 100%
    display flex
    align-items center
    .header
        width 100%
</style>