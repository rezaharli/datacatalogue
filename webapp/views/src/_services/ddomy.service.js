import { fetchWHeader } from '../_helpers/auth-header';

export const ddoMyService = {
    getLeftTable,
    getRightTable,
    getDetails
};

function getLeftTable(param) {
    return fetchWHeader(`/ddo/getlefttable`, param)
}

function getRightTable(param) {
    return fetchWHeader(`/ddo/getrighttable`, param)
}

function getDetails(param) {
    return fetchWHeader(`/ddo/getdetails`, { LeftParam: parseInt(param.left), RightParam: parseInt(param.right) })
}