<template>
    <div class="loginWrap" :class="status+'Size'">
        <div class="loginBox">
            <section class="title">{{status == 'login'?'登录':'注册'}}</section>
            <section v-if="status==='login'" class="loginForm subForm" :class="statusType">
                <el-form ref="loginForm" :model="loginForm" :rules="rules" label-width="80px" key="login">
                    <el-form-item label="用户名" prop="username" class="username">
                        <el-input v-model="loginForm.username" class="input"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="password" class="password">
                        <el-input show-password v-model="loginForm.password" class="input"></el-input>
                    </el-form-item>
                    <!-- <el-form-item label="验证码" prop="verCode" v-loading="loginForm.loginVerCodeLoading">
                        <el-input v-model="loginForm.verCode"></el-input>
                    </el-form-item> -->
                    <!-- <el-form-item>
                        <el-button type="primary" @click="loginSubmit" class="loginBtn">登录</el-button>
                        <el-button type="primary" @click="status = 'registry'" class="cancel">注册</el-button>
                    </el-form-item> -->
                </el-form>
                <section class="btnGroup" style="margin-top:80px;">
                    <el-button type="primary" @click="submit('login', loginForm)" class="commit">登录</el-button>
                    <el-button v-if="!onlyLogin" type="primary" @click="toggleStatus('registry')" class="cancel">注册</el-button>
                </section>
            </section>
            <section v-else-if="!onlyLogin" class="registryForm subForm" :class="statusType">
                <el-form ref="registryForm" :model="registryForm" :rules="rules" label-width="150px" key="registry">
                    <el-form-item label="用户名" prop="username" class="username">
                        <el-input v-model="registryForm.username" class="input"></el-input>
                    </el-form-item>
                    <el-form-item label="电子邮箱" prop="email" class="email">
                        <el-input v-model="registryForm.email" class="input"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="password" class="password">
                        <el-input show-password v-model="registryForm.password" class="input"></el-input>
                    </el-form-item>
                    <el-form-item label="请再次输入密码" prop="password" class="repassword">
                        <el-input show-password v-model="registryForm.repassword" class="input"></el-input>
                    </el-form-item>
                    <!-- <el-form-item label="验证码" prop="verCode" v-loading="loginForm.registryVerCodeLoading">
                        <el-input v-model="registryForm.verCode"></el-input>
                    </el-form-item> -->
                    <!-- <el-form-item>
                        <el-button type="primary" @click="regSubmit" class="registryBtn">注册</el-button>
                        <el-button type="primary" @click="status = 'login'" class="cancel">取消</el-button>
                    </el-form-item> -->
                </el-form>
                <section class="btnGroup">
                    <el-button type="primary" @click="submit('registry', registryForm)" class="commit">注册</el-button>
                    <el-button type="primary" @click="toggleStatus('login')" class="cancel">取消</el-button>
                </section>
            </section>
        </div>
    </div>
</template>
<script>
export default {
    name: 'Login',
    props:['onlyLogin'],
    data(){
        // 验证码校验
        // let verCodeCheck = (rule, value, callback)=>{
        //     this.getVerCode().
        //      then(()=>{
        //          callback()
        //      },()=>{
        //          callback(new Error("服务器错误"))
        //      })
        // }
        // let serverCheckCode = (rule,value,callback)=>{
        //     this.serverCheckVerCode().
        //          then(()=>{

        //          },()=>{
        //              callback(new Error(""))
        //          })
        // }
        return {
            status: 'login',
            statusType: 'login',
            loginForm:{
                username: '',
                password: '',
                verCode:  '',
                loginVerCodeLoading: false,
            },
            registryForm:{
                username: '',
                email: '',
                password: '',
                repassword: '',
                verCode: '',
                registryVerCodeLoading: false,
            },
            rules:{
                username:[
                    {required: true, message:'请输入用户名',trigger: 'blur'},
                    {min:6, max: 18, message:'用户名长度在6~18个字符',trigger: 'blur'},
                ],
                password:[
                    {required: true, message:'请输入密码',trigger: 'blur'},
                    {min:6, max: 18, message:'密码长度在6~18个字符',trigger: 'blur'},
                ],
                email:[
                    {required: true, message:'请输入注册邮箱',trigger: 'blur'},
                    {type: 'email', message: '邮箱格式不正确', trigger: 'blur'},
                ],
                // verCode:[
                //     {required: true, message:'请输入验证码',trigger: 'blur'},
                //     {validator: verCodeCheck, trigger: 'blur'}
                // ]
            }
        }
    },
    mounted(){
        console.log(this.serverUrl)
    },
    methods:{
        submit(formName, form){
            this.$refs[formName+'Form'].validate((valid)=>{
                if (valid){
                    this.$emit(formName, form)
                }
            })
        },
        toggleStatus(status){
            this.statusType = 'noshow'
            this.status = status
            setTimeout(()=>{this.statusType = 'show'}, 500)
        }
        // getVerCode(){
        //     return new Promise((resolve, reject)=>{
        //         this.$axios.get(this.serverUrl+"/api/getVerCode")
        //         .then((res)=>{
        //             resolve(res)
        //         })
        //         .catch(()=>{
        //             reject("服务器错误")
        //         })
        //     })
        // },
        // serverCheckVerCode(){
        //     return new Promise((resolve, reject)=>{

        //     })
        // }
    }
}
</script>
<style lang="stylus" scoped>
.loginSize
    width 350px
    height 450px
.registrySize
    width 450px
    height 520px
.loginWrap
    transition all .5s
    padding 10px
    border 1px solid gray
    border-radius 10px
    box-sizing border-box
    .noshow
        display none
    .title
        padding 10px
        padding-bottom 15px
        padding-left 0
        margin-bottom 40px
        font-size 20px
        text-align center
        box-sizing border-box
        border-bottom 1px solid gray
    .subForm
        margin-top 20px
        .username, .password, .email, .repassword
            margin-bottom 40px
        .input
            width 220px
    .btnGroup
        text-align center
        .cancel
            margin-left 30px

.test
    width 100%
    height 100%
    background-color black
</style>