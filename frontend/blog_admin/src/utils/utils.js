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
// https://github.com/Robbendebiene/Sliding-Scroll/blob/master/sliding-scroll.js 大佬写的滑动逻辑
function scrollToY(ele, y, duration = 1000){
    // cancel if already on top

    if (ele.scrollTop === y) return;

    const cosParameter = (ele.scrollTop-y) / 2;
    let scrollCount = 0, oldTimestamp = null;

    function step (newTimestamp) {
        if (oldTimestamp !== null) {
            // if duration is 0 scrollCount will be Infinity
            scrollCount += Math.PI * (newTimestamp - oldTimestamp) / duration;
            if (scrollCount >= Math.PI) return ele.scrollTop = y;
            ele.scrollTop = cosParameter + y + cosParameter * Math.cos(scrollCount);
        }
        oldTimestamp = newTimestamp;
        window.requestAnimationFrame(step);
    }
    window.requestAnimationFrame(step);
}

export default{
    debounce,
    scrollToY
}