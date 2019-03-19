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
    return fetchWHeader(`/ddo/getrighttable`, param)
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