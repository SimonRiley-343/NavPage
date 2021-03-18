<template>
    <el-container>
        <el-main class="login-main">
            <el-row>
                <el-col :span="8" :offset="8" class="login-col-title">NavPage Web Admin</el-col>
                <el-col :span="8" :offset="8" class="login-col-passwd">
                    <el-input placeholder="Password" v-model="password" show-password></el-input>
                </el-col>
                <el-col class="login-col-btn">
                    <el-button type="primary" round>Login</el-button>
                </el-col>
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
        return {
            password: ''
        };
    },
    mounted() {
        if (localStorage.getItem('sessionId')) {
            this.login();
        }
    },
    methods: {
        login() {
            const reqData = {
                passwd: this.password
            };
            login(reqData)
                .then((res) => {
                    if (res.status) {
                        this.$router.push({
                            name: 'Admin'
                        });
                    }
                })
                .catch((e) => {
                    console.log(e);
                });
        }
    }
});
</script>

<style lang="scss">
.login-main {
    margin-top: 100px;
}

.login-col-title {
    text-align: center;
    font-size: 34px;
    font-weight: 600;
}

.login-col-passwd {
    margin-top: 80px;
}

.login-col-btn {
    margin-top: 40px;
    text-align: center;
}
</style>
