import { fetchWHeader } from '../_helpers/auth-header';

export const ddoMyService = {
    getLeftTable,
    getRightTable,
    getDetails
};

function getLeftTable(param) {
    return fetchWHeader(`/ddo/getlefttable`, param)
}

function getRightTable(param) {
    return fetchWHeader(`/ddo/getrighttable`, param).then(
        res => {
            res.Data = _.map(res.Data, function(v){
                v.CDE_YES_NO = v.CDE_YES_NO == 0 ? "No" : "Yes";
                return v;
            });

            return res;
        }
    )
}

function getDetails(param) {
    return fetchWHeader(`/ddo/getdetails`, param).then(
        res => {
            var tmp = _.groupBy(res.Data.Detail, "ID")
            
            res.Data.Detail = _.map(Object.keys(tmp), function(v, i){
                var ret = tmp[v][0];
                ret.ID = v;
                ret.Values = tmp[v];

                return ret;
            });

            return res;
        }
    )
}