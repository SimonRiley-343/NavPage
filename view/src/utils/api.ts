import http from '@/utils/http';
import qs from 'qs';
import base from '@/utils/base';

export function pages(data = '') {
    return http({
        url: `${base.url}/pages`,
        method: 'post',
        data: qs.stringify(data)
    });
}
