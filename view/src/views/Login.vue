<template>
    <el-container>
        <el-header class="login-header" height="140px">
            <h1 class="login-title">NavPage Web Admin</h1>
            <h3 class="login-subtitle">Document Everything</h3>
        </el-header>
        <el-main class="login-main">
            <el-row>
                <el-col :span="16" :offset="4">
                    <el-input placeholder="Password" v-model="password" show-password></el-input>
                </el-col>
                <el-col :span="6" :offset="9" class="login-col-btn">
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
import Api from '@/utils/api';
import { cookies, reqCode } from '@/utils/model';
import Session from '@/utils/session';

const session = new Session();
const api = new Api();

export default Vue.extend({
    name: 'Login',
    data() {
        return {
            password: '',
            btnStatus: {
                disabled: true,
                loading: false
            },
            jumpName: 'Admin'
        };
    },
    mounted() {
        let _this = this;
        if (session.exist(cookies.navpage) && session.exist(cookies.sessionId)) {
            _this.login();
        }
        if (this.$route.query.redirect) {
            _this.jumpName = this.$route.query.redirect;
        }
    },
    methods: {
        login() {
            let _this = this;
            _this.btnStatus.loading = true;
            api.login({
                passwd: _this.password
            })
                .then((res) => {
                    let resData = res.data as loginResObj;

                    switch (resData.code) {
                        case reqCode.success:
                            if (resData.login) {
                                _this.$router.push({
                                    name: _this.jumpName
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

                    session.remove(cookies.navpage);
                    session.remove(cookies.sessionId);
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
    $min-width: 400px;
    margin: 10vh calc((100% - #{$min-width}) / 2);
    padding: 60px 10px !important;
    border-radius: 10px;
    background-color: rgba($color: #fefefe, $alpha: 1);
    box-shadow: 0 2px 12px 0 rgba($color: #000000, $alpha: 0.2);
}

.login-header {
    text-align: center;
    margin-top: 6vh;
}

.login-title {
    font-size: 2.5em;
    font-weight: 600;
    margin-block-start: 1em;
    margin-block-end: 0.3em;
}

.login-subtitle {
    font-size: 1.5em;
    font-weight: 300;
    margin-block-start: 0;
    margin-block-end: 0;
}

.login-col-btn {
    margin-top: 40px;
    text-align: center;
}
</style>
