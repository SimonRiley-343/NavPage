<template>
    <div id="app">
        <NavPage></NavPage>
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
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
