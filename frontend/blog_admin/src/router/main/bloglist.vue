<template>
    <div class="">
        <el-table 
        :data="tableData" 
        style="width: 90%;margin-left:5%;margin-top:80px;overflow:auto;"
        height="600px"
        stripe
        border
        :header-cell-style="headStyle"
        :cell-style="rowStyle"
        >
            <el-table-column
                prop="id"
                label="序号"
                width="180">
            </el-table-column>
            <el-table-column
                prop="title"
                label="标题"
                width="180">
            </el-table-column>
            <el-table-column
                prop="content"
                label="内容">
            </el-table-column>
            <el-table-column
                prop="type"
                label="类型">
            </el-table-column>
            <el-table-column
                prop="label"
                label="标签">
            </el-table-column>
            <el-table-column
                prop="visitedN"
                label="访问量">
            </el-table-column>
            <el-table-column
                prop="commentN"
                label="评论数">
            </el-table-column>
            <el-table-column
                label="操作">
                <template slot-scope="scope">
                    <el-button
                        @click.native.prevent="detail(scope.$index)"
                        type="text"
                        size="small"
                    >
                    查看详情
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-dialog 
        title="博文详情"
        :visible.sync="showDetail"
        >
        <div>
            <el-form ref="blogForm" :model="blogForm" label-width="100px" v-loading="updateLoading">
                <el-form-item label="标题">
                    <el-input v-model="blogForm.title"></el-input>
                </el-form-item>
                <el-form-item label="内容">
                    <el-input type="textarea" :row="10" v-model="blogForm.content"></el-input>
                </el-form-item>
                <el-form-item label="类型">
                    <el-checkbox-group v-model="blogForm.type" @change="selectType">
                        <el-checkbox v-for="(type, i) in allType" :label="type" :key="i"></el-checkbox>
                    </el-checkbox-group>
                </el-form-item>
                <el-form-item label="标签">
                    <el-checkbox-group v-model="blogForm.label" @change="selectLabel">
                        <el-checkbox v-for="(label, i) in allLabel" :label="label" :key="i"></el-checkbox>
                    </el-checkbox-group>
                </el-form-item>
                <el-form-item label="访问量">
                    <span>{{blogForm.visitedN}}</span>
                </el-form-item>
                <el-form-item label="评论数">
                    <span>{{blogForm.commentN}}</span>
                </el-form-item>
            </el-form>
            <section style="text-align:center">
                <el-button type="primary" @click="blogUpdate">
                    确定修改
                </el-button>
            </section>
        </div>
        </el-dialog>
    </div>
</template>
<script>
export default {
    data(){
        return {
            tableData:[],
            showDetail: false,
            detailInfo: {},
            blogForm:{
                blogId: '',
                title: 'title',
                content: 'content',
                visitedN: 100,
                commentN: 1000,

                type: [],
                label: [],
            },
            allType: ['a','b','c','d','e','f','g'],
            allLabel: ['A', 'B', 'C', 'D', 'E', 'F'],
            updateLoading: false
        }
    },
    methods:{
        rowStyle(){
            return "text-align:center"
        },
        headStyle(){
            return "text-align:center"
        },
        detail(index){
            Object.assign(this.blogForm, this.tableData[index])
            this.showDetail = true
        },
        typeChange(){
            console.log(arguments)
        },
        blogUpdate(){
            this.updateLoading = true
            let blogForm = this.blogForm
            blogForm.type = blogForm.type.join('|')
            blogForm.label = blogForm.label.join('|')
            console.log(blogForm)
            this.$axios.post("/api/blog/updateBlogById", {
                blogForm: this.blogForm
            })
            .then(()=>{
                setTimeout(()=>{
                    this.$router.go(0)
                }, 500)
            })
            .catch((err)=>{
                setTimeout(()=>{
                    this.updateLoading = false
                    alert(err)
                }, 500)
            })
        },
        selectType(){
            console.log(arguments)
        },
        selectLabel(){
            console.log(arguments)
        },
        dataInit(){
            this.$axios.get("/api/blog/list", {
                params:{
                    pageNum:1
                }
            })
            .then((data)=>{
                if(data.data.success){
                    let blogData = data.data.data
                    let defaultData = []
                    blogData.forEach((blog,i)=>{
                        defaultData.push({
                            blogId: blog.blogId,
                            id: i+1,
                            title: blog.title,
                            content: blog.content,
                            type: blog.type.split('|')[0] == ""?[]:blog.type.split('|'),
                            label: blog.label.split('|')[0] == ""?[]:blog.label.split('|'),
                            visitedN: blog.visitedN,
                            commentN: blog.commentN
                        })
                    })
                    this.tableData = defaultData
                }
            })
            .catch((err)=>{
                console.log(err)
            })
        }
    },
    mounted(){
        this.dataInit()
    }
}
</script>
<style lang="stylus" scoped>
</style>
