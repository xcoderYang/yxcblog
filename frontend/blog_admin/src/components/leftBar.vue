<template>
    <div class="leftBarWrap">
        <el-row class="leftBar">
            <el-col>
                <el-menu
                default-active="bloglist"
                class="el-menu-vertical-demo"
                @select="itemSelect"
                >
                <el-menu-item index="1">
                    <i class="el-icon-s-data"></i>
                    <span slot="title">总览</span>
                </el-menu-item>
                <el-menu-item index="bloglist">
                    <i class="el-icon-menu"></i>
                    <span slot="title">博文列表</span>
                </el-menu-item>
                <el-menu-item index="picture">
                    <i class="el-icon-picture"></i>
                    <span slot="title">图片墙</span>
                </el-menu-item>
                <el-menu-item index="userManage">
                    <i class="el-icon-user-solid"></i>
                    <span slot="title">用户管理</span>
                </el-menu-item>
                </el-menu>
            </el-col>
        </el-row>
    </div>  
</template>
<script>
export default {
    data(){
        return {

        }
    },
    methods:{
        itemSelect(to, path){
            this.$emit("toggle", to, path)
            // console.log(arguments)
        }
    }
}
</script>
<style lang="stylus" scoped>
.leftBarWrap
    width 100%
    height 100%
    .leftBar
        width 100%
        height 100%
</style>