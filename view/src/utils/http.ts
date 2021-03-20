import axios from 'axios';
import { Message } from 'element-ui';
import router from '../router/index';
import { cookies } from '@/utils/model';
import Session from '@/utils/session';

const http = axios.create({ timeout: 1000 * 12 });
const session = new Session();

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
        let statusCode = err.response;

        switch (statusCode) {
            case 401:
                Message.error('User session expired, please retry login');

                console.log('Remove Session at http.ts');
                session.remove(cookies.navpage);
                session.remove(cookies.sessionId);

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
