<template>
    <div id="app">
        <div class="bg"></div>
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

<style>
body {
    margin: 0;
}

#app {
    font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
}

.bg {
    background-image: url("assets/bg.png");
    background-position: 50% 50%;
    background-size: cover;
    background-repeat: no-repeat;
    position: absolute;
    top: -200px;
    bottom: -200px;
    left: 0;
    right: 0;
    z-index: -1;
    min-height: 1200px;
}

#navpage {
    max-width: 1200px;
    margin: 0 auto;
}
</style>
