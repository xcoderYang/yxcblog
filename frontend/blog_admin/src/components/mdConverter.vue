<template>
    <div>
      <el-row class="title">
        <el-col :span="10" :offset="1" class="writer">
            <el-input type="textarea" v-model="title" :rows="20"></el-input>
        </el-col>
        <el-col :span="10" :offset="2" class="mder" v-html="mdTitle">
        </el-col>
      </el-row>
      <el-row class="content">
        <el-col :span="12" class="writer">
            {{content}}
        </el-col>
        <el-col :span="12" class="mder" v-html="mdContent">
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