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
                        @click.native.prevent="showDetailDialog(scope.$index)"
                        type="text"
                        size="small"
                    >
                    查看详情
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
        <div style="text-align:center;margin-top:50px;">
            <el-button @click="showBlogAddDialog">
                新增博文
            </el-button>
        </div>
        <!-- 博文编辑框 -->
        <el-dialog
        title="博文详情"
        :visible.sync="showDetail"
        :close-on-click-modal=false
        >
            <div>
                <el-form ref="blogForm" :rules="blogRules" :model="blogForm" label-width="100px" v-loading="updateLoading">
                    <el-form-item label="标题" prop="title">
                        <el-input v-model="blogForm.title"></el-input>
                    </el-form-item>
                    <el-form-item label="内容" prop="content">
                        <el-input type="textarea" :row="10" v-model="blogForm.content"></el-input>
                    </el-form-item>
                    <el-form-item label="类型">
                        <el-checkbox-group v-model="blogForm.type">
                            <el-checkbox v-for="(type, i) in allType" :label="type" :key="i"></el-checkbox>
                        </el-checkbox-group>
                    </el-form-item>
                    <el-form-item label="标签">
                        <el-checkbox-group v-model="blogForm.label">
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
                    <el-button style="margin-right:50px;" type="primary" @click="blogUpdate">
                        确定修改
                    </el-button>
                    <el-button type="danger" @click="blogDelete">
                        删除博文
                    </el-button>
                </section>
            </div>
        </el-dialog>
        <!-- 新增博文框 -->
        <el-dialog 
        title="新增博文"
        :visible.sync="showBlogAdd"
        :close-on-click-modal=false
        >
            <div>
                <el-form ref="blogAddForm" :rules="blogRules" :model="blogAddForm" label-width="100px" v-loading="addLoading">
                    <el-form-item label="标题" prop="title">
                        <el-input v-model="blogAddForm.title"></el-input>
                    </el-form-item>
                    <el-form-item label="内容" prop="content">
                        <el-input type="textarea" :row="10" v-model="blogAddForm.content"></el-input>
                    </el-form-item>
                    <el-form-item label="类型">
                        <el-checkbox-group v-model="blogAddForm.type">
                            <el-checkbox v-for="(type, i) in allType" :label="type" :key="i"></el-checkbox>
                        </el-checkbox-group>
                    </el-form-item>
                    <el-form-item label="标签">
                        <el-checkbox-group v-model="blogAddForm.label">
                            <el-checkbox v-for="(label, i) in allLabel" :label="label" :key="i"></el-checkbox>
                        </el-checkbox-group>
                    </el-form-item>
                    <el-form-item label="访问量">
                        <span>{{blogAddForm.visitedN}}</span>
                    </el-form-item>
                    <el-form-item label="评论数">
                        <span>{{blogAddForm.commentN}}</span>
                    </el-form-item>
                </el-form>
                <section style="text-align:center">
                    <el-button type="primary" @click="blogAdd">
                        新增博文
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
            showBlogAdd: false,
            blogForm:{
                blogId: '',
                title: 'title',
                content: 'content',
                visitedN: 100,
                commentN: 1000,

                type: [],
                label: [],
            },
            blogAddForm:{
                blogId: '',
                title: '',
                content: '',
                visitedN: 0,
                commentN: 0,
                type: [],
                label: [],
            },
            allType: ['a','b','c','d','e','f','g'],
            allLabel: ['A', 'B', 'C', 'D', 'E', 'F'],
            updateLoading: false,
            addLoading: false,
            blogRules:{
                title:[
                    {required:true,message:'请输入博文标题', trigger: 'blur'},
                    {min:3,max:100,message:'标题长度请控制在3-100之间', trigger:'blur'}
                ],
                content:[
                    {required:true,message:'请输入博文内容',trigger:'blur'}
                ]
            }
        }
    },
    methods:{
        rowStyle(){
            return "text-align:center"
        },
        headStyle(){
            return "text-align:center"
        },
        showDetailDialog(index){
            Object.assign(this.blogForm, this.tableData[index])
            this.showDetail = true
        },
        showBlogAddDialog(){
            this.showBlogAdd = true
        },
        typeChange(){
            console.log(arguments)
        },
        blogUpdate(){
            this.$refs['blogForm'].validate((valid)=>{
                if(valid){
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
                }
            })
        },
        blogAdd(){
            this.$refs['blogAddForm'].validate((valid)=>{
                if(valid){
                    this.addLoading = true
                    let blogAddForm = this.blogAddForm
                    blogAddForm.type = blogAddForm.type.join('|')
                    blogAddForm.label = blogAddForm.label.join('|')
                    console.log(blogAddForm)
                    this.$axios.post("/api/blog/createBlog", {
                        blogAddForm: this.blogAddForm
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
                }else{
                    return false
                }
            })
        },
        blogDelete(){
            this.$axios.post("/api/blog/deleteBlogById", {
                blogId: this.blogForm.blogId
            })
            .then(()=>{
                console.log("yes")
            })
            .catch(()=>{
                console.log("no")
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
