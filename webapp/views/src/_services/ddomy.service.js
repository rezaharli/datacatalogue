import { fetchWHeader } from '../_helpers/auth-header';

export const ddoMyService = {
    getAllSystem,
    getTableName,
    getDetails
};

function getAllSystem() {
    return fetchWHeader(`/ddo/getlefttable`, {})
}

function getTableName(processID) {
    return fetchWHeader(`/ddo/getrighttable`, { ProcessID: parseInt(processID) })
}

function getDetails(param) {
    return fetchWHeader(`/ddo/getdetails`, { LeftParam: parseInt(param.left), RightParam: parseInt(param.right) })
}