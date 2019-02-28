import { fetchWHeader } from '../_helpers/auth-header';

export const dscMyService = {
    getAllSystem,
    getTableName,
    getInterfacesRightTable,
    getDetails
};

function getAllSystem() {
    return fetchWHeader(`/dsc/getallsystems`, {})
}

function getTableName(system) {
    return fetchWHeader(`/dsc/gettablename`, { SystemID: parseInt(system) }).then(
        res => {
            var tmp = _.groupBy(res.Data, "TABLE_NAME")
            
            res.Data = _.map(Object.keys(tmp), function(v, i){
                return {
                    ID: tmp[v][0].ID,
                    TSID: tmp[v][0].TSID,
                    TABLE_NAME: v,
                    Columns: tmp[v]
                }
            });

            return res;
        }
    )
}

function getInterfacesRightTable(system) {
    return fetchWHeader(`/dsc/getinterfacesrighttable`, { SystemID: parseInt(system) })
}

function getDetails(param) {
    return fetchWHeader(`/dsc/getdetails`, { LeftParam: parseInt(param.left), RightParam: parseInt(param.right) })
}