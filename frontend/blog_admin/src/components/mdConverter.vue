<template>
    <div>
      <el-row class="title">
        <el-col class="writer">
            {{title}}
        </el-col>
        <el-col class="mder" v-html="mdTitle">
        </el-col>
      </el-row>
      <el-row class="content">
        <el-col class="writer">
            {{content}}
        </el-col>
        <el-col class="mder" v-html="mdContent">
        </el-col>
      </el-row>
    </div>    
</template>
<script>
export default {
    props:['title', 'content'],
    data(){
        return {
            mdTitle:'',
            mdContent:''
        }
    },
    watch:{
        title:function(val){
            console.log(val)
            let converter = new this.$showdown.Converter()
            this.mdTitle = converter.makeHtml(val)
            console.log(this.mdTitle)
        }
    },
    mounted(){
        console.log(this.title)
        let converter = new this.$showdown.Converter()
        this.mdTitle = converter.makeHtml(this.title)
        console.log(this.mdTitle)
    }
}
</script>
<style lang="stylus" scoped>

</style>