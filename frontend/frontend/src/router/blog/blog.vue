<template>
    <div class="blogWrap">
        <div class="container">
            <section class="title" @click="test" v-html="article.title">
            </section>
            <section class="label">
                <span v-for="(label,i) in article.labels" :key="i" class="labelItem">
                    {{label}}
                </span>
            </section>
            <section class="content" v-html="article.content">
            </section>    
        </div>
    </div>
</template>
<script>
export default {
    // props:['blogId'],
    data(){
        return {
            article:{
            },
            blogId: ''
        }
    },
    mounted(){
        this.blogId = this.$route.params.blogId
        this.$axios.get(this.serverUrl + "/server/blog/blogItem", {
            params:{
                blogId: this.blogId
            }
        })
        .then((res)=>{
            console.log(res)
            this.data2Page(res.data.data)
        })
        .catch((err)=>{
            console.log("no")
            console.log(err)
        })
    },
    methods: {
        test(){
            console.log("test")
        },
        data2Page(data){
            let moment = this.$moment
            let showdown = this.$showdown
            let converter = new showdown.Converter()

            this.article = {
                createTime: moment(data.createAt).format("YYYY/MM/DD"),
                title: converter.makeHtml(data.title),
                content: converter.makeHtml(data.content),
                labels: data.label.split("|"),
                blogId: data.blogId,
                lastUpdate: moment(data.updateAt).format("YYYY/MM/DD")
            }
        }
    },
}
</script>
<style lang="stylus" scoped>
.blogWrap
    .container
        width 100%
        display flex
        flex-direction column
        align-items center
        background-color #F5F7FA
    .title
        width 1000px
        margin-top 20px
        text-align center
        box-sizing border-box
        font-size 30px
    .content
        width 1000px
        overflow auto
    .label
        padding 0 20px
        margin 10px 0
        .labelItem
            font-size 12px
            margin 0 5px
            padding 0 5px
            background-color #FFDEAD
            border-radius 5px
        

</style>