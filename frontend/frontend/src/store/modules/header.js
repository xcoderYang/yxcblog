const state=()=>({
    headerVision:true,
})

const getters={}

const mutations = {
    HEADER_HIDE(state){
        state.headerVision = false
    },
    HEADER_SHOW(state){
        state.headerVision = true
    }
    
}

const actions = {
    HEADER_SHOW({commit}){
        commit("HEADER_SHOW")
    },
    HEADER_HIDE({commit}){
        commit("HEADER_HIDE")
    }
}

export default{
    state,
    getters,
    actions,
    mutations
}