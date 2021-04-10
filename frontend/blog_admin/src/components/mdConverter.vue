<template>
    <div>
      <el-row class="title">
        <el-col :span="10" :offset="1" class="writer">
            <el-input placeholder="请输入博文标题" v-model="curTitle"></el-input>
        </el-col>
        <el-col :span="10" :offset="2" class="mder" v-html="mdTitle">
        </el-col>
      </el-row>
      <el-row style="margin-top:40px;" class="content">
        <el-col :span="10" :offset="1" class="writer">
            <el-input placeholder="请输入博文内容" type="textarea" v-model="curContent" :rows="20"></el-input>
        </el-col>
        <el-col :span="10" :offset="2" class="mder" v-html="mdContent">
        </el-col>
      </el-row>
      <el-row style="text-align:center;margin-top:50px;" class="btnGroup">
          <el-button @click="confirm" style="margin-right:100px;">
              确认修改
          </el-button>
          <el-button @click="cancel">
              取消修改
          </el-button>
      </el-row>
    </div>   
</template>
<script>
export default {
    props:['title', 'content'],
    data(){
        return {
            curTitle: '',
            curContent: '',
            mdTitle:'',
            mdContent:'',
        }
    },
    watch:{
        title: function(){
            this.curTitle = this.title
        },
        content: function(){
            this.curContent = this.content
        },
        curTitle:function(val){
            let converter = new this.$showdown.Converter()
            this.mdTitle = converter.makeHtml(val)
        },
        curContent:function(val){
            let converter = new this.$showdown.Converter()
            this.mdContent = converter.makeHtml(val)
        }
    },
    mounted(){
        this.curTitle = this.title
        this.curContent = this.content
    },
    methods:{
        confirm(){
            this.$emit("confirm", {
                title: this.curTitle,
                content: this.curContent
            })
        },
        cancel(){
            this.$emit("cancel")
        }
    }
}
</script>
<style lang="stylus" scoped>

</style>