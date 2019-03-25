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
                var ret = tmp[v][0];
                ret.SYSTEM_NAME = v;
                ret.Custodians = tmp[v];

                return ret;
            });

            return res;
        }
    )
}

function getRightTable(param) {
    return fetchWHeader(`/dsc/gettablename`, param).then(
        res => {
            res.Data = _.map(res.Data, function(v){
                v.CDE_YES_NO = v.CDE_YES_NO == 0 ? "No" : "Yes";
                
                v.LEFTID = v.TSID;
                return v;
            });

            var tmp = _.groupBy(res.Data, "TABLE_NAME")
            
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var ret = tmp[v][0];
                ret.TABLE_NAME = v;
                ret.Columns = tmp[v];

                return ret;
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
                
                v.LEFTID = v.TSID;
                return v;
            });

            var tmp = _.groupBy(res.Data, "LIST_OF_CDE")
            
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var ret = tmp[v][0];
                ret.LIST_OF_CDE = v;
                ret.Values = tmp[v];

                return ret;
            });

            return res;
        }
    )
}

function getDetails(param) {
    return fetchWHeader(`/dsc/getdetails`, param).then(
        res => {
            var tmp = _.groupBy(res.Data.Detail, "ID")
            
            res.Data.Detail = _.map(Object.keys(tmp), function(v, i){
                var ret = tmp[v][0];
                ret.ID = v;
                ret.Values = tmp[v];

                return ret
            });

            return res;
        }
    )
}