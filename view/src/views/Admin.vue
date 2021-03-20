<template>
    <el-container>
        <el-header>Header</el-header>
        <el-container>
            <el-aside width="200px">Aside</el-aside>
            <el-main>Main</el-main>
        </el-container>
    </el-container>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
    name: 'Admin',
    data() {
        return {};
    },
    mounted() {
        let _this = this;
    },
    methods: {}
});
</script>

<style lang="scss">
</style>
