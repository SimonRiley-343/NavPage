import axios from 'axios';
import { Message } from 'element-ui';
import router from '../router/index';

const http = axios.create({ timeout: 1000 * 12 });

http.interceptors.request.use(
    (config) => {
        if (localStorage.getItem('sessionId')) {
            config.headers.Authorization = localStorage.getItem('sessionId');
        }

        return config;
    },
    (err) => {
        return Promise.reject(err);
    }
);

http.interceptors.response.use(
    (res) => {
        return res;
    },
    (err) => {
        const statusCode = err.response;

        switch (statusCode) {
            case 401:
                Message.error('用户 Session 过期，请重新登录');
                break;

            default:
                Message.error('未知错误');
                break;
        }

        localStorage.removeItem('sessionId');
        router.push('/');

        return Promise.reject(err);
    }
);

export default http;
