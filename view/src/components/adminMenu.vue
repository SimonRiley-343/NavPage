<template>
    <el-menu class="el-menu-vertical-demo" :collapse="isCollapse" @select="handleSelect">
        <el-menu-item index="overview">
            <i class="el-icon-setting"></i>
            <span slot="title">Test</span>
        </el-menu-item>
    </el-menu>
</template>

<script>
import Vue from 'vue';

export default Vue.extend({
    name: 'adminMenu',
    props: {
        isCollapse: {
            type: Boolean,
            default: false
        },
        selected: {
            type: String,
            default: ''
        }
    },
    data() {
        return {
            menuSelect: this.selected
        };
    },
    methods: {
        handleSelect(index) {
            this.menuSelect = index;
            this.$emit('update:selected', this.menuSelect);
        }
    },
    watch: {
        selected(val) {
            this.menuSelect = val;
        }
    }
});
</script>

<style lang="scss">
.el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    min-height: 400px;
}
</style>
