import Vue from 'vue';
import http from '@/utils/http';
import qs from 'qs';
import base from '@/utils/base';

export default class Api extends Vue {
    public login(data = {}) {
        return http({
            url: `${base.url}/login`,
            method: 'post',
            data: qs.stringify(data)
        });
    }

    public pages(data = {}) {
        return http({
            url: `${base.url}/pages`,
            method: 'post',
            data: qs.stringify(data)
        });
    }

    public addPage(data = {}) {
        return http({
            url: `${base.url}/addPage`,
            method: 'post',
            data: qs.stringify(data)
        });
    }
}
