<template>
    <div :ref="containerRef"></div>
</template>

<script>
import lottie from 'lottie-web';

export default {
    props: {
        animationData: {
            type: Object, // Adjust the type based on the actual structure of your animation data
            required: true
        }
    },
    data() {
        return {
            containerRef: 'lottieContainer',
            animation: null
        };
    },
    mounted() {
        console.log(this.animationData);
        this.renderLottieAnimation();
    },
    methods: {
        renderLottieAnimation() {
            var animationData = this.animationData;
            this.animation = lottie.loadAnimation({
                container: this.$refs.lottieContainer,
                renderer: 'svg',
                loop: true,
                autoplay: true,
                animationData
            });
        }
    },
    beforeDestroy() {
        if (this.animation) {
            this.animation.destroy();
        }
    }
};
</script>

<style scoped>
/* Add any styles you need for the container here */
</style>
