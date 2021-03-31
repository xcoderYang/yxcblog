

const state=()=>({
    userInfo:{},
    userLogin:false
})

const getters={}

const mutations = {
    USER_LOGIN(state, user){
        sessionStorage.setItem("userInfo", JSON.stringify(user))
        sessionStorage.setItem("userLogin", "true")
        state.userInfo = user
        state.userLogin = true
    },
    USER_LOGOUT(state){
        state.userInfo = {}
        state.userLogin = false
        sessionStorage.setItem("userInfo", "")
        sessionStorage.setItem("userLogin", "false")
    },
    USER_INFO_SYNC(state){
        state.userInfo = sessionStorage.getItem("userInfo")
        state.userLogin = JSON.parse(sessionStorage.getItem("userLogin"))
    }
}

const actions = {
    USER_LOGIN({commit}, user){
        commit("USER_LOGIN", user)
    },
    USER_LOGOUT({commit}){
        commit("USER_LOGOUT")
    },
    USER_INFO_SYNC({commit}){
        commit("USER_INFO_SYNC")
    }
}

export default{
    state,
    getters,
    actions,
    mutations
}