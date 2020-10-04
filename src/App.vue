<template>
    <div id="app">
        <NavPage id="navpage"></NavPage>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import NavPage from '@/components/NavPage.vue'

export default Vue.extend({
    components: {
        NavPage
    }
})
</script>

<style lang="scss">
html, body, #app {
    height: 100%;
}

body {
    margin: 0;
}

#app {
    font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    background-image: url("assets/bg.png");
    background-position: 50% 50%;
    background-size: cover;
    background-repeat: no-repeat;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    z-index: -1;
    min-height: 900px;
}

#navpage {
    max-width: 1200px;
    margin: 0 auto;
}
</style>
