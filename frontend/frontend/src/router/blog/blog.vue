<template>
    <div class="blogWrap">
        <section class="title">
            abcc
        </section>
        <section class="content">
        saksjdlksjadkl
        </section>    
    </div>
</template>
<script>
export default {
    data(){
        return {
            lists: [{
                "createTime": "2020/11/11",
                "title": "this is title this is title this is title this is title this is title this is title",
                "tags": ["golang", "mysql", "分布式"],
                "blogId": "",
            },{
                "createTime": "2020/11/11",
                "title": "this is title this is title this is title this is title this is title this is title",
                "tags": ["golang", "mysql", "分布式"],
                "blogId": "",
            },{
                "createTime": "2020/11/11",
                "title": "this is title this is title this is title this is title this is title this is title",
                "tags": ["golang", "mysql", "分布式"],
                "blogId": "",
            },{
                "createTime": "2020/11/11",
                "title": "this is title this is title this is title this is title this is title this is title",
                "tags": ["golang", "mysql", "分布式"],
                "blogId": "",
            }]
        }
    },
    methods: {
        test: function(){
            console.log(this.$route.params)
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

</style>