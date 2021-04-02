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
                    <el-checkbox-group v-model="selectedType">
                        <el-checkbox v-for="(type, i) in typeList" :label="type" :key="i"></el-checkbox>
                    </el-checkbox-group>
                </el-form-item>
                <el-form-item label="标签">
                    <el-checkbox-group v-model="selectedLabel">
                        <el-checkbox v-for="(label, i) in labelList" :label="label" :key="i"></el-checkbox>
                    </el-checkbox-group>
                </el-form-item>
                <el-form-item label="访问量">
                    <span>{{blogForm.visitedN}}</span>
                </el-form-item>
                <el-form-item label="评论数">
                    <span>{{blogForm.commentN}}</span>
                </el-form-item>
            </el-form>
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
                type: "abc",
                label: "cba",
                visitedN: 100,
                commentN: 1000,
            },{
                id:0,
                title: "test",
                content: "test",
                type: "abc",
                label: "cba",
                visitedN: 100,
                commentN: 1000,
            }],
            showDetail: false,
            detailInfo: {}
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
            this.detailInfo = this.tableData[index]
            this.showDetail = true
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