<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>

export default {
  name: 'App',
  components: {
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
</style>