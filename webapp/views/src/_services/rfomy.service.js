import { fetchWHeader } from '../_helpers/auth-header';

export const ddoMyService = {
    getLeftTable,
    getRightTable,
};

function getLeftTable(param) {
    return fetchWHeader(`/rfo/getlefttable`, param)
}

function getRightTable(param) {
    return fetchWHeader(`/rfo/getrighttable`, param).then(
        res => {
            var tmp = _.groupBy(res.Data, "PRINCIPAL_RISK_TYPE");

            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmp2 = _.groupBy(tmp[v], "RISK_SUB_TYPE");


                var RiskSubTypes = _.map(Object.keys(tmp2), function(v, i){
                    var tmp3 = _.groupBy(tmp2[v], "PRIORITY_REPORT_RATIONALE");

                    var PriorityReports = _.map(Object.keys(tmp3), function(v, i){
                        var tmp4 = _.groupBy(tmp3[v], "CRM_NAME");

                        var CRMNames = _.map(Object.keys(tmp4), function(v, i){
                            var tmp5 = _.groupBy(tmp4[v], "ASSOCIATED_CDES");

                            var CDEs = _.map(Object.keys(tmp5), function(v, i){
                                return {
                                    ID: i,
                                    LeftID: tmp5[v][0].TPRID,
                                    ASSOCIATED_CDES: v,
                                    Rationales: tmp5[v],
                                    RESULT_COUNT: Object.keys(tmp5).length
                                }
                            });

                            return {
                                ID: i,
                                LeftID: tmp4[v][0].TPRID,
                                CRM_NAME: v,
                                CDEs: CDEs,
                                CDEsVal: tmp4[v],
                                RESULT_COUNT: Object.keys(tmp4).length
                            }
                        });

                        return {
                            ID: i,
                            LeftID: tmp3[v][0].TPRID,
                            PRIORITY_REPORT_RATIONALE: v,
                            CRMNames: CRMNames,
                            CRMNamesVal: tmp3[v],
                            RESULT_COUNT: Object.keys(tmp3).length
                        }
                    });

                    return {
                        ID: i,
                        LeftID: tmp2[v][0].TPRID,
                        RISK_SUB_TYPE: v,
                        PriorityReports: PriorityReports,
                        PriorityReportsVal: tmp2[v],
                        RESULT_COUNT: Object.keys(tmp2).length
                    }
                });

                return {
                    ID: tmp[v][0].ID,
                    LeftID: tmp[v][0].TPRID,
                    PRINCIPAL_RISK_TYPE: v,
                    RiskSubTypes: RiskSubTypes,
                    RiskSubTypesVal: tmp[v],
                    RESULT_COUNT: Object.keys(tmp).length
                }
            });

            return res;
        }
    )
}