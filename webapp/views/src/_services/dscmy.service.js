import { fetchWHeader } from '../_helpers/auth-header';

export const dscMyService = {
    getAllSystem,
    getTableName,
};

function getAllSystem() {
    return fetchWHeader(`/dsc/getallsystems`, {})
}

function getTableName(system) {
    return fetchWHeader(`/dsc/gettablename`, { SystemID: parseInt(system) })
}