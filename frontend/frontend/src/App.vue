<template>
  <div id="wrap">
    <Header :class="{hide:!headerShow}"></Header>
    <div id="app">
    <router-view></router-view>
  </div>
  </div>
</template>

<script>
import Header from './components/header.vue'

export default {
  name: 'App',
  components:{
    Header
  },
  data(){
    return {
      headerShow: true,
      hideRouter:["blog"]
    }
  },
  mounted(){
    this.headerShow = this.$store.state.header.headerVision
    console.log(this.headerShow)
    console.log(this.$route.fullPath)
    if(this.hideRouter.includes(this.$route.name)){
      this.headerShow = false
    }else{
      this.headerShow = true
    }
  },
  watch:{
    "$route"(to){
      console.log(to)
      if(this.hideRouter.includes(to.name)){
        this.headerShow = false
      }else{
        this.headerShow = true
      }
    }
  }
}
</script>

<style lang="stylus">
html, body
  width 100%
  height 100%
  padding 0
  margin 0
  overflow hidden
i
  font-style normal
#wrap
  width 100%
  height 100%
#app
  width 100%
  height 100%
  display flex
  align-items center
  justify-content center
div
  width 100%
  height 100%
.hide
  display none
</style>
