import { fetchWHeader } from '../_helpers/auth-header';

export const ddoMyService = {
    getLeftTable,
    getRightTable,
};

function getLeftTable(param) {
    return fetchWHeader(`/rfo/getlefttable`, param)
}

function getRightTable(param) {
    return fetchWHeader(`/rfo/getrighttable`, param)
}