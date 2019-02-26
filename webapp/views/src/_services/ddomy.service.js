import { fetchWHeader } from '../_helpers/auth-header';

export const ddoMyService = {
    getAllSystem,
    getTableName,
};

function getAllSystem() {
    return fetchWHeader(`/ddo/getlefttable`, {})
}

function getTableName(processID) {
    return fetchWHeader(`/ddo/getrighttable`, { ProcessID: parseInt(processID) })
}