<template>
    <div class="wrap">
        <div class="triangle" @click="scrollTop"></div>
        <!-- <el-backtop target=".triangle"></el-backtop> -->
    </div>
</template>
<script>
export default {
    name: "scrollTop",
    data(){
        return {

        }
    },
    methods:{
        scrollTop(){
            this.$emit("scrollTop")
        }
    }
}
</script>
<style lang="stylus" scoped>
.wrap
    width 50px
    height 50px
    border-radius 50%
    box-shadow 0 0 5px gray
    cursor pointer
    background white
    .triangle
        margin auto
        margin-top 17px
        width 0
        height 0
        border-left 10px solid transparent
        border-right 10px solid transparent
        border-bottom 16px solid gray
</style>