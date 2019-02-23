import { fetchWHeader } from '../_helpers/auth-header';

export const dpoMyService = {
    getAllSystem,
    getTableName,
};

function getAllSystem() {
    return fetchWHeader(`/dpo/getlefttable`, {})
}

function getTableName(processID) {
    return fetchWHeader(`/dpo/getrighttable`, { ProcessID: parseInt(processID) })
}