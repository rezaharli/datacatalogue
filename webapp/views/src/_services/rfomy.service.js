import { fetchWHeader } from '../_helpers/auth-header';

export const rfoMyService = {
    getLeftTable,
    getRightTable,
    getPriorityTable,
};

function getLeftTable(param) {
    return fetchWHeader(`/rfo/getlefttable`, param)
}

function getRightTable(param) {
    return fetchWHeader(`/rfo/getrighttable`, param).then(
        res => {
            res.Data = _.map(res.Data, function(v){
                v.LEFTID = v.TPRID;
                return v;
            });
            var tmp = _.groupBy(res.Data, "PRINCIPAL_RISK");

            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmp2 = _.groupBy(tmp[v], "RISK_SUB");

                var RiskSubTypes = _.map(Object.keys(tmp2), function(v, i){
                    var tmp3 = _.groupBy(tmp2[v], "PR_NAME");

                    var PriorityReports = _.map(Object.keys(tmp3), function(v, i){
                        var tmp4 = _.groupBy(tmp3[v], "CRM_NAME");

                        var CRMNames = _.map(Object.keys(tmp4), function(v, i){
                            var tmp5 = _.groupBy(tmp4[v], "ASSOC_CDES");

                            var CDEs = _.map(Object.keys(tmp5), function(v, i){
                                var ret = tmp5[v][0];
                                ret.ID = i
                                ret.ASSOC_CDES = v;
                                ret.Rationales = tmp5[v];

                                return ret;
                            });

                            var ret = tmp4[v][0];
                            ret.ID = i
                            ret.CRM_NAME = v;
                            ret.CDEs = CDEs;
                            ret.CDEsVal = tmp4[v];

                            return ret;
                        });

                        var ret = tmp3[v][0];
                        ret.ID = i
                        ret.PR_NAME = v;
                        ret.CRMNames = CRMNames;
                        ret.CRMNamesVal = tmp3[v];

                        return ret;
                    });

                    var ret = tmp2[v][0];
                    ret.ID = i
                    ret.RISK_SUB = v;
                    ret.PriorityReports = PriorityReports;
                    ret.PriorityReportsVal = tmp2[v];

                    return ret;
                });

                var ret = tmp[v][0];
                ret.LeftID = tmp[v][0].TPRID;
                ret.PRINCIPAL_RISK = v;
                ret.RiskSubTypes = RiskSubTypes;
                ret.RiskSubTypesVal = tmp[v];

                return ret;
            });

            return res;
        }
    )
}

function getPriorityTable(param) {
    return fetchWHeader(`/dsc/getcdetable`, param).then(
        res => {
            res.DataFlat = _.cloneDeep(res.Data);
            
            var tmp = _.groupBy(res.Data, "CDE")
            
            return res;
        }
    );
}