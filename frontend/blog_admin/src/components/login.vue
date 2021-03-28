<template>
    <div class="loginWrap">
        <div class="loginBox">
            <section class="title">title</section>
            <section v-if="status==='login'" class="loginForm">
                <el-form ref="loginForm" :model="loginForm" :rules="rules" label-width="80px" key="login">
                    <el-form-item label="用户名" prop="username">
                        <el-input v-model="loginForm.username"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="password">
                        <el-input v-model="loginForm.password"></el-input>
                    </el-form-item>
                    <el-form-item label="验证码" prop="verCode" v-loading="loginForm.loginVerCodeLoading">
                        <el-input v-model="loginForm.verCode"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="loginSubmit">登录</el-button>
                        <el-button type="primary" @click="status = 'registry'">注册</el-button>
                    </el-form-item>
                </el-form>
            </section>
            <section v-else class="registryForm">
                <el-form ref="registryForm" :model="registryForm" :rules="rules" label-width="80px" key="registry">
                    <el-form-item label="用户名" prop="username">
                        <el-input v-model="registryForm.username"></el-input>
                    </el-form-item>
                    <el-form-item label="电子邮箱" prop="email">
                        <el-input v-model="registryForm.email"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="password">
                        <el-input v-model="registryForm.password"></el-input>
                    </el-form-item>
                    <el-form-item label="验证码" prop="verCode" v-loading="loginForm.registryVerCodeLoading">
                        <el-input v-model="registryForm.verCode"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="regSubmit">注册</el-button>
                        <el-button type="primary" @click="status = 'login'">取消</el-button>
                    </el-form-item>
                </el-form>
            </section>
            <section class="btnGroup">btn</section>
        </div>
    </div>
</template>
<script>
export default {
    name: 'Login',
    data(){
        //  let getVerCode = (rule, value, callback)=>{
        //      this.getVerCode().
        //      then(()=>{
        //          callback()
        //      },()=>{
        //          callback(new Error("服务器错误"))
        //      })
        // }
        let verCodeCheck = (rule, value, callback)=>{
            callback()
        }
        return {
            status: 'login',
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
                verCode:[
                    {required: true, message:'请输入验证码',trigger: 'blur'},
                    {validator: verCodeCheck, trigger: 'blur'}
                ]
            }
        }
    },
    mounted(){
        console.log(this.serverUrl)
    },
    methods:{
        loginSubmit(){
            this.$axios.get(this.serverUrl+"/api/login", {
                params:this.loginForm
            })
            .then(()=>{
                console.log("Yes")
            })
            .catch(()=>{
                console.log("No")
            })
            console.log(this.$axios)
        },
        regSubmit(){
            
        },
        getVerCode(){
            return new Promise((resolve, reject)=>{
                this.$axios.get(this.serverUrl+"/api/getVerCode")
                .then((res)=>{
                    resolve(res)
                })
                .catch(()=>{
                    reject("服务器错误")
                })
            })
        },
        // serverCheckVerCode(){
        //     return new Promise((resolve, reject)=>{

        //     })
        // }
    }
}
</script>
<style lang="stylus" scoped>
.test
    width 100%
    height 100%
    background-color black
</style>