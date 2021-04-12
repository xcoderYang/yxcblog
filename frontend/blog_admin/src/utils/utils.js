function debounce(fn, delay, ctx){
    let movement = null
    return function(){
        let args = arguments
        clearTimeout(movement)
        movement = setTimeout(function(){
            fn.apply(ctx, args)
        }, delay)
    }
}

export default{
    debounce
}