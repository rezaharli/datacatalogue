import { fetchWHeader } from '../_helpers/auth-header';

export const dpoMyService = {
    getLeftTable,
    getRightTable,
    getDetails
};

function getLeftTable(param) {
    return fetchWHeader(`/dpo/getlefttable`, param)
}

function getRightTable(param) {
    return fetchWHeader(`/dpo/getrighttable`, param).then(
        res => {
            res.Data = _.map(res.Data, function(v){
                v.LEFTID = v.TDPID;
                return v;
            });

            return res;
        }
    )
}

function getDetails(param) {
    return fetchWHeader(`/dpo/getdetails`, { LeftParam: parseInt(param.left), RightParam: parseInt(param.right) })
}