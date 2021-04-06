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
            <el-form ref="blogForm" :model="blogForm" label-width="100px">
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
            tableData:[{
                id:0,
                title: "test",
                content: "test",
                type: ['a','b','c'],
                label: ['C','B','A'],
                visitedN: 100,
                commentN: 1000,
            },{
                id:1,
                title: "test",
                content: "test",
                type: ['a','b','c'],
                label: ['C','B','A'],
                visitedN: 100,
                commentN: 1000,
            }],
            showDetail: false,
            detailInfo: {},
            blogForm:{
                title: 'title',
                content: 'content',
                visitedN: 100,
                commentN: 1000,

                type: [],
                label: [],
            },
            allType: ['a','b','c','d','e','f','g'],
            allLabel: ['A', 'B', 'C', 'D', 'E', 'F']
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
            let detailInfo = this.tableData[index]
            this.detailInfo = detailInfo
            this.showDetail = true
            Object.assign(this.blogForm, detailInfo)
        },
        typeChange(){
            console.log(arguments)
        },
        blogUpdate(){
            this.$axios.post("/api/blog/updateBlogById", {
                blogForm: this.blogForm
            })
            .then((res)=>{
                console.log(res)
            })
            .catch((err)=>{
                console.log(err)
            })
        },
        selectType(){
            console.log(arguments)
        },
        selectLabel(){
            console.log(arguments)
        }
    },
    mounted(){
        for(let i=0; i<3; i++){
            this.tableData = this.tableData.concat(this.tableData)
        }
    }
}
</script>
<style lang="stylus" scoped>
</style>