import { fetchWHeader } from '../_helpers/auth-header';

import moment from 'moment'

export const cryptoService = {
    encrypt, decrypt
};

function encrypt(param) {
    return fetchWHeader(`/dashboard/encrypt`, param);
}

function decrypt(param) {
    return fetchWHeader(`/dashboard/decrypt`, param);
}