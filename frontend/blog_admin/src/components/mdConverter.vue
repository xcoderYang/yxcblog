<template>
    <div class="converter">
      <el-row class="title">
        <el-col :span="10" :offset="1" class="writer">
            <el-input placeholder="请输入博文标题" v-model="curTitle"></el-input>
        </el-col>
        <el-col :span="10" :offset="2" class="mder" v-html="mdTitle">
        </el-col>
      </el-row>
      <el-row style="margin-top:30px;" class="content">
        <el-col :span="10" :offset="1" class="writer">
            <textarea placeholder="请输入博文内容" v-model="curContent" rows="40" class="blogInput" @scroll="blogScroll">
            </textarea>
            <!-- <el-input placeholder="请输入博文内容" type="textarea" v-model="curContent" :rows="30" resize="none" @scroll.native="test"></el-input> -->
        </el-col>
        <el-col :span="10" :offset="2" class="mder blogPreview" v-html="mdContent">
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
            debounceCache: "",
            scrollPercent: ""
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
        },
        scrollPercent: function(val){
            let ele = document.getElementsByClassName("blogPreview")[0]
            this.$utils.scrollToY(ele, ele.scrollHeight*val, 500)
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
        blogScroll(args){
            if (this.debounceCache!==""){
                this.debounceCache(args)
                return
            }
            this.debounceCache = this.$utils.debounce(function(args){
                this.scrollPercent = args.target.scrollTop/args.target.scrollHeight
            }, 150, this)
            this.debounceCache(args)
        },
    }
}
</script>
<style lang="stylus" scoped>
.converter
    .blogInput
        resize none
        width 100%
        padding 20px
        font-size 18px
        color #333
        height 650px
        box-sizing border-box
        line-height 30px
    .blogPreview
        height 650px
        overflow auto
        font-size 18px
        padding 20px
        border 1px solid black
</style>