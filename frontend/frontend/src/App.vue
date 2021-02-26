<template>
  <div id="wrap">
    <Header></Header>
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
</style>
