import { fetchWHeader } from '../_helpers/auth-header';

export const ddoMyService = {
    getAllSystem,
    getTableName
};

function getAllSystem() {
    return fetchWHeader(`/rfo/getlefttable`, {})
}

function getTableName(processID) {
    return fetchWHeader(`/rfo/getrighttable`, { ProcessID: parseInt(processID) })
}