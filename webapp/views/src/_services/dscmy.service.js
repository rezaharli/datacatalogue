import { fetchWHeader } from '../_helpers/auth-header';

export const dscMyService = {
    getLeftTable,
    getRightTable,
    getInterfacesRightTable,
    getDetails
};

function getLeftTable(param) {
    return fetchWHeader(`/dsc/getallsystems`, param)
}

function getRightTable(param) {
    return fetchWHeader(`/dsc/gettablename`, param).then(
        res => {
            var tmp = _.groupBy(res.Data, "TABLE_NAME")
            
            res.Data = _.map(Object.keys(tmp), function(v, i){
                return {
                    ID: tmp[v][0].ID,
                    TSID: tmp[v][0].TSID,
                    TABLE_NAME: v,
                    Columns: tmp[v],
                    RESULT_COUNT: Object.keys(tmp).length
                }
            });

            return res;
        }
    )
}

function getInterfacesRightTable(param) {
    return fetchWHeader(`/dsc/getinterfacesrighttable`, param)
}

function getDetails(param) {
    return fetchWHeader(`/dsc/getdetails`, param)
}