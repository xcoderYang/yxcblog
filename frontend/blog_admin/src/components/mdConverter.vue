<template>
    <div>
      <el-row class="title">
        <el-col :span="10" :offset="1" class="writer">
            <el-input placeholder="请输入博文标题" v-model="curTitle"></el-input>
        </el-col>
        <el-col :span="10" :offset="2" class="mder" v-html="mdTitle">
        </el-col>
      </el-row>
      <el-row style="margin-top:30px;" class="content">
        <el-col :span="10" :offset="1" class="writer">
            <textarea placeholder="请输入博文内容" v-model="curContent" rows="40" style="resize:none;width:100%;padding:20px;font-size:18px;color:#333;height:650px;box-sizing:border-box" @scroll="test">
            </textarea>
            <!-- <el-input placeholder="请输入博文内容" type="textarea" v-model="curContent" :rows="30" resize="none" @scroll.native="test"></el-input> -->
        </el-col>
        <el-col :span="10" :offset="2" class="mder" v-html="mdContent" style="height:650px;overflow:auto;font-size:18px;padding:20px;border:1px solid black;">
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
        },
        test(){
            console.log(arguments)
        }
    }
}
</script>
<style lang="stylus" scoped>

</style>