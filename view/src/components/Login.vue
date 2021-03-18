<template>
    <el-container>
        <el-main>
            <el-row>
                <el-col :span="8" :offset="8"></el-col>
            </el-row>
        </el-main>
    </el-container>
</template>

<script lang="ts">
import Vue from 'vue';
import { login } from '@/utils/api';

export default Vue.extend({
    name: 'Login',
    data() {
        return {};
    },
    mounted() {
        // this.login();
    },
    methods: {
        login() {
            login()
                .then((res) => {
                    // TODO: Check login status
                })
                .catch((e) => {
                    console.log(e);
                });
        }
    }
});
</script>

<style lang="scss">
</style>
