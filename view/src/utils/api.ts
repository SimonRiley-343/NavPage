import http from '@/utils/http';
import qs from 'qs';
import base from '@/utils/base';

export function login(data = {}) {
    return http({
        url: `${base.url}/login`,
        method: 'post',
        data: qs.stringify(data)
    });
}

export function pages(data = {}) {
    return http({
        url: `${base.url}/pages`,
        method: 'post',
        data: qs.stringify(data)
    });
}

export function addPage(data = {}) {
    return http({
        url: `${base.url}/addPage`,
        method: 'post',
        data: qs.stringify(data)
    });
}
