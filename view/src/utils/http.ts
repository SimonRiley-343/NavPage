import axios from 'axios';
import { Message } from 'element-ui';
import router from '../router/index';
import { cookies } from '@/utils/model';
import { removeSession } from '@/utils/session';

let http = axios.create({ timeout: 1000 * 12 });

http.interceptors.request.use(
    (config) => {
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
        let _this = this;
        let statusCode = err.response;

        switch (statusCode) {
            case 401:
                Message.error('User session expired, please retry login');

                console.log('Remove Session at http.ts');
                removeSession(_this, cookies.navpage);
                removeSession(_this, cookies.sessionId);

                router.push({
                    name: 'Login'
                });
                break;

            default:
                Message.error('Unknown error');
                break;
        }

        return Promise.reject(err);
    }
);

export default http;
