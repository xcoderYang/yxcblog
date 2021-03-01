<template>
    <div class="blogWrap">
        <section class="title" @click="test">
            {{article.title}}
        </section>
        <section class="content">
            {{article.content}}
        </section>    
    </div>
</template>
<script>
export default {
    data(){
        return {
            article:{
                "createTime": "now",
                "title": "112233",
                "content": "oneonetwotwothreethree",
                "tag": "number",
                "blogId": "123",
                "lastUpdate": "last",
            },
        }
    },
    methods: {
        test: function(){
            console.log(this.$route.params)
            console.log(this.$showdown)
        }
    },
}
</script>
<style lang="stylus" scoped>
.blogWrap
    width 100%
    display flex
    flex-direction column
    align-items center
    background-color #F5F7FA
    .title
        width 1000px
        margin-top 70px
        text-align center
        box-sizing border-box
        font-size 30px
    .content
        width 1000px
        

</style>