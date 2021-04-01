<template>
    <div class="headerWrap">
        <el-page-header @back="goback" content="abc页面"></el-page-header>
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
</style>