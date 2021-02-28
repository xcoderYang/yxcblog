<template>
<div class="headerWrap">
    <div class="header">
        <el-col :span="11" :offset="1" class="leftbar">
            <i class="icon mainLogo" @click="toHome"></i>
        </el-col>
        <el-col :span="12" class="rightbar">
            <el-row style="height:100%">
                <el-col :offset="4" :span="6" class="blog linkbtn">
                    <div class="link" @click="toBlog">
                        blog
                    </div>
                </el-col>
                <el-col :span="6" class="images linkbtn">
                    <div class="link" @click="toImages">
                        images
                    </div>
                </el-col>
                <el-col :span="6" class="diary linkbtn">
                    <div class="link" @click="toDiary">
                        diary
                    </div>
                </el-col>
            </el-row>
        </el-col>
    </div>    
</div>
</template>

<script>
export default {
    data(){
        return {

        }
    },
    methods:{
        toHome: function(){
            if (this.$router.currentRoute.fullPath=="/"){
                return
            }
            this.$router.push("/home")
        },
        toBlog: function(){
            this.$router.push("/blog")
        },
        toImages: function(){
            this.$router.push("/images")
        },
        toDiary: function(){
            this.$router.push("/diary")
        }
    }
}
</script>

<style lang="stylus" scoped>
.mainLogo
    width 264px
    height 40px
    cursor pointer
    background url("../assets/main.jpg")
    background-size 100% 100%
.headerWrap
    height 50px
.header
    width 100%
    background-color white
    position fixed
    display flex
    height 50px
    align-items center
    .leftbar
        display flex
        justify-content left
        align-items center
    .rightbar
        .linkbtn
            .link
                display flex
                justify-content center
                align-items center
                font-size 18px
                cursor pointer
                &:hover
                    color white
                    background-color #409EFF
</style>