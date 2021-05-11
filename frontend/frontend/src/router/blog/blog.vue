<template>
    <div class="blogWrap">
        <div class="container">
            <section class="article">
                <LEFTRECOMMEND class="leftRecommend"></LEFTRECOMMEND>
                <RIGHTRECOMMEND class="rightRecommend"></RIGHTRECOMMEND>
                <section class="title" @click="test" v-html="article.title">
                </section>
                <section class="label">
                    <span v-for="(label,i) in article.labels" :key="i" class="labelItem">
                        {{label}}
                    </span>
                </section>
                <section class="content" v-html="article.content">
                </section>
                <section class="comment">
                </section>
                <el-backtop target=".article" :bottom="300" :right="200"></el-backtop>
            </section>
        </div>
    </div>
</template>
<script>
import LEFTRECOMMEND from "../../components/left_recommend"
import RIGHTRECOMMEND from "../../components/right_recommend"
export default {
    components:{
        LEFTRECOMMEND,
        RIGHTRECOMMEND,
    },
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
        console.log(this.$store)
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
        },
        goBack(){
            
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
    .article
        width 100%
        height auto
        overflow auto
        padding-bottom 50px
        .title
            width 800px
            text-align center
            margin auto
            box-sizing border-box
            font-size 30px
        .label
            text-align center
            padding-bottom 20px
            .labelItem
                font-size 12px
                margin 0 5px
                padding 0 5px
                background-color #FFDEAD
                border-radius 5px
        .content
            width 800px
            margin auto
        .leftRecommend
            display none
            position fixed
            top 30%
            left 200px
            border 1px solid black
        .rightRecommend
            display none
            position fixed
            top 30%
            right 200px
            border 1px solid black
        .scrollTop
            position fixed
            bottom 20%
            right 10%

</style>