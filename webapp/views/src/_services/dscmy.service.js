import { fetchWHeader } from '../_helpers/auth-header';

export const dscMyService = {
    getLeftTable,
    getRightTable,
    getInterfacesRightTable,
    getDetails
};

function getLeftTable(param) {
    return fetchWHeader(`/dsc/getallsystems`, param).then(
        res => {
            var tmp = _.groupBy(res.Data, "SYSTEM_NAME")
            
            res.Data = _.map(Object.keys(tmp), function(v, i){
                return {
                    ID: tmp[v][0].ID,
                    SYSTEM_NAME: v,
                    Custodians: tmp[v],
                    RESULT_COUNT: Object.keys(tmp).length
                }
            });

            return res;
        }
    )
}

function getRightTable(param) {
    return fetchWHeader(`/dsc/gettablename`, param).then(
        res => {
            res.Data = _.map(res.Data, function(v){
                v.CDE = v.CDE == 0 ? "No" : "Yes";
                return v;
            });

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
    return fetchWHeader(`/dsc/getinterfacesrighttable`, param).then(
        res => {
            res.Data = _.map(res.Data, function(v){
                v.IMM_PREC_SYSTEM_SLA = v.IMM_PREC_SYSTEM_SLA == 0 ? "No" : "Yes";
                v.IMM_PREC_SYSTEM_OLA = v.IMM_PREC_SYSTEM_OLA == 0 ? "No" : "Yes";
                v.IMM_SUCC_SYSTEM_SLA = v.IMM_SUCC_SYSTEM_SLA == 0 ? "No" : "Yes";
                v.IMM_SUCC_SYSTEM_OLA = v.IMM_SUCC_SYSTEM_OLA == 0 ? "No" : "Yes";
                return v;
            });

            return res;
        }
    )
}

function getDetails(param) {
    return fetchWHeader(`/dsc/getdetails`, param)
}