<template>
    <el-container>
        <el-header class="admin-header" height="3em">
            <h1 class="admin-title">NavPage Web Admin</h1>
        </el-header>
        <el-container>
            <el-aside width="240px">
                <admin-menu :selected.sync="menuSelected"></admin-menu>
            </el-aside>
            <el-main>Main</el-main>
        </el-container>
    </el-container>
</template>

<script lang="ts">
import Vue from 'vue';
import adminMenu from '@/components/adminMenu.vue';

export default Vue.extend({
    name: 'Admin',
    components: {
        adminMenu
    },
    data() {
        return {
            menuSelected: ''
        };
    },
    mounted() {
        let _this = this;
    },
    methods: {},
    watch: {
        menuSelected(val) {
            console.log(val);
        }
    }
});
</script>

<style lang="scss">
.admin-header {
    padding: 0.75em 20px !important;
    background-color: #409eff;
    color: #f2f6fc;
}

.admin-title {
    margin-block-start: 0;
    margin-block-end: 0;
}
</style>
