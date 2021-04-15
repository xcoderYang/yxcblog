<template>
    <div class="blogWrap">
        <section class="blogTitle">
            所有博文
        </section>
        <section class="bloglist">
            <ul class="blogul">
                <li v-for="(blog,index) in lists" :key="index">
                    <span class="createTime">{{blog.createTime}}</span>
                    <span class="title" @click="toBlog(blog.blogId)">{{blog.title}}</span>
                    <span class="label">
                        <span v-for="(label,i) in blog.labels" :key="i" class="labelItem">
                            {{label}}
                        </span>
                    </span>
                </li>
            </ul>
        </section>
</div>
    
</template>
<script>
export default {
    data(){
        return {
            lists: [{}]
        }
    },
    mounted(){
        this.$axios.get(this.serverUrl + "/server/blog/list",{
            params:{
                pageNum:1
            }
        })
        .then((res)=>{
            this.data2Page(res.data.data)
        })
        .catch((err)=>{
            console.log(err)
        })
    },
    methods:{
        toBlog: function(blogId){
            this.$router.push({name:"blog", params:{blogId}})
        },
        data2Page(data){
            let moment = this.$moment
            this.lists = data.map((e)=>{
                return {
                    createTime: moment(e.createAt).format("YYYY/MM/DD"),
                    title: e.title,
                    labels: e.label.split('|'),
                    blogId: e.blogId
                }
            })
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
    .blogTitle
        width 1000px
        margin-top 70px
        padding-left 40px
        text-align left
        box-sizing border-box
        font-size 24px
        color #8968CD
    .bloglist
        width 1000px
        .blogul
            font-size 16px
            line-height 1.8
            letter-spacing 10px
            li
                height 50px
                display flex
                align-items center
                list-style none
                span
                    display inline-block
                    letter-spacing 2px
                .createTime
                    width 110px
                .title
                    width 570px
                    overflow hidden
                    white-space nowrap
                    text-overflow ellipsis
                    margin-left 20px
                    cursor pointer
                    &:hover
                        color #CD8500
                .label
                    padding 0 20px
                    .labelItem
                        font-size 12px
                        margin 0 5px
                        padding 0 5px
                        background-color #FFDEAD
                        border-radius 5px

</style>