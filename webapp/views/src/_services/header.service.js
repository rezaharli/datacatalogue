import { fetchWHeader } from '../_helpers/auth-header';

export const headerService = {
    getOpts,
    getRowCount
};

function getOpts(param) {
    return fetchWHeader(`/base/getheaderopts`, param);
}

function getRowCount(param) {
    return fetchWHeader(`/base/getrowcount`, param);
}