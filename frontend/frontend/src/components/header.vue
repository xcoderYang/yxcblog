<template>
<div class="header">
    <el-row style="width:100%">
        <el-col :span="11" :offset="1" class="leftbar">
            <i class="icon">这是 logo</i>
        </el-col>
        <el-col :span="12" class="rightbar">
            <el-row>
                <el-col :offset="4" :span="6" class="blog">
                <el-link :underline="false" class="link">blog</el-link>
                </el-col>
                <el-col :span="6" class="images">
                <el-link :underline="false" class="link">images</el-link>
                </el-col>
                <el-col :span="6" class="diary">
                <el-link :underline="false" class="link">diary</el-link>
                </el-col>
            </el-row>
        </el-col>
    </el-row>
</div>    
</template>

<script>
export default {
    
}
</script>

<style lang="stylus" scoped>
.header
    display flex
    height 40px
    align-items center
    .rightbar
        .link
            font-size 18px
</style>