import { fetchWHeader } from '../_helpers/auth-header';

export const configService = {
    getConfig
};

function getConfig(param) {
    return fetchWHeader(`/dashboard/getconfig`, param);
}