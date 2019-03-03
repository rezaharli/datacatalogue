import { fetchWHeader } from '../_helpers/auth-header';

export const dpoMyService = {
    getLeftTable,
    getRightTable,
    getDetails
};

function getLeftTable(param) {
    return fetchWHeader(`/dpo/getlefttable`, param)
}

function getRightTable(param) {
    return fetchWHeader(`/dpo/getrighttable`, param)
}

function getDetails(param) {
    return fetchWHeader(`/dpo/getdetails`, { LeftParam: parseInt(param.left), RightParam: parseInt(param.right) })
}