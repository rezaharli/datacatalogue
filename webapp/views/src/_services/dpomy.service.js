import { fetchWHeader } from '../_helpers/auth-header';

export const dpoMyService = {
    getAllSystem,
    getTableName,
    getDetails
};

function getAllSystem() {
    return fetchWHeader(`/dpo/getlefttable`, {})
}

function getTableName(processID) {
    return fetchWHeader(`/dpo/getrighttable`, { ProcessID: parseInt(processID) })
}

function getDetails(param) {
    return fetchWHeader(`/dpo/getdetails`, { LeftParam: parseInt(param.left), RightParam: parseInt(param.right) })
}