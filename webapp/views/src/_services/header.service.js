import { fetchWHeader } from '../_helpers/auth-header';

export const headerService = {
    getOpts
};

function getOpts(param) {
    return fetchWHeader(`/base/getheaderopts`, param);
}