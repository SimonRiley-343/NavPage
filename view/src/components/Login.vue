<template>
    <el-container>
        <el-main class="login-main">
            <el-row>
                <el-col :span="8" :offset="8" class="login-col-title">NavPage Web Admin</el-col>
                <el-col :span="8" :offset="8" class="login-col-passwd">
                    <el-input placeholder="Password" v-model="password" show-password></el-input>
                </el-col>
                <el-col class="login-col-btn">
                    <el-button type="primary" round @click="login" :disabled="btnStatus.disabled"
                        :loading="btnStatus.loading">Login</el-button>
                </el-col>
            </el-row>
        </el-main>
    </el-container>
</template>

<script lang="ts">
import Vue from 'vue';
import { Message } from 'element-ui';
import { login } from '@/utils/api';
import { cookies, reqCode } from '@/utils/model';
import { removeSession, isSessionExist } from '@/utils/session';

export default Vue.extend({
    name: 'Login',
    data() {
        return {
            password: '',
            btnStatus: {
                disabled: true,
                loading: false
            }
        };
    },
    mounted() {
        let _this = this;
        if (isSessionExist(_this, cookies.navpage) && isSessionExist(_this, cookies.sessionId)) {
            _this.login();
        }
    },
    methods: {
        login() {
            let _this = this;
            _this.btnStatus.loading = true;
            login({
                passwd: _this.password
            })
                .then((res) => {
                    let resData = res.data as loginResObj;

                    switch (resData.code) {
                        case reqCode.success:
                            if (resData.login) {
                                _this.$router.push({
                                    name: 'Admin'
                                });
                                return;
                            }
                            break;
                        case reqCode.paramErr:
                            Message.error('Login - Please input admin password');
                            break;
                        case reqCode.loginFailed:
                            Message.error('Login - Login failed, please check password');
                            break;
                        default:
                            Message.error('Login - Unknown error, please contact site admin');
                            break;
                    }

                    removeSession(_this, cookies.navpage);
                    removeSession(_this, cookies.sessionId);
                })
                .catch((e) => {
                    console.log(e);
                })
                .finally(() => {
                    _this.btnStatus.loading = false;
                });
        }
    },
    watch: {
        password(value) {
            let _this = this;
            if (value.length === 0) {
                _this.btnStatus.disabled = true;
            } else {
                _this.btnStatus.disabled = false;
            }
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
