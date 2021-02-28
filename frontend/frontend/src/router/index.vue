<template>
    <div class="outerWrap">
        <div class="wrap">
            <div class="innerWrap">
                <figure class="title">
                    欢迎来到我的小屋
                </figure>
            </div>
        </div>
        <div>
            testestestest
        </div>
    </div>
</template>

<script>
export default{
    name: 'index',
    data(){
        return {
            
        }
    }
}
</script>
<style scoped lang="stylus">
.outerWrap
    width 100%
    height 100%
    .wrap
        height 100%
        width 100%
        background url("../assets/index.jpg")
        background-size cover
        .innerWrap
            display flex
            justify-content center
            align-items center
            width 100%
            height 100%
            background-color rgba(0, 0, 0, 0.5)
            .title
                font-size 48px
                color white
</style>