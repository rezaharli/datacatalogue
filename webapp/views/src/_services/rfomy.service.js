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
            res.Data = [
                { PRIORITY_REPORT: "FM Exposure", PRIORITY_REPORT_RATIONALE: "ZXCV", CRM_NAME: "PFE", CRM_RATIONALE: "GGKJH", CDE_NAME: "Legal Name", CDE_RATIONALE: "Rationale 1" },
                { PRIORITY_REPORT: "FM Exposure", PRIORITY_REPORT_RATIONALE: "ZXCV", CRM_NAME: "CAT1", CRM_RATIONALE: "GGKJH", CDE_NAME: "Incorporation Date", CDE_RATIONALE: "Rationale 2" },
                { PRIORITY_REPORT: "FM Exposure", PRIORITY_REPORT_RATIONALE: "ZXCV", CRM_NAME: "EEPE", CRM_RATIONALE: "GGKJH", CDE_NAME: "Legal Name", CDE_RATIONALE: "Rationale 1" },
                // { PRIORITY_REPORT: "FIN Rep", PRIORITY_REPORT_RATIONALE: "MM abcde", CRM_NAME: "xx", CRM_RATIONALE: "GGKJH", CDE_NAME: "Legal Name", CDE_RATIONALE: "Rationale 1" },
                // { PRIORITY_REPORT: "FIN Rep", PRIORITY_REPORT_RATIONALE: "fghijk lmnop", CRM_NAME: "yy", CRM_RATIONALE: "GGKJH", CDE_NAME: "", CDE_RATIONALE: "" },
            ];

            res.DataFlat = _.cloneDeep(res.Data);
            
            res.Data = _.groupBy(res.Data, "PRIORITY_REPORT");
            
            return res;
        },
        err => {
            var res = {};
            res.Data = [
                { PRIORITY_REPORT: "FM Exposure", PRIORITY_REPORT_RATIONALE: "ZXCV", CRM_NAME: "PFE", CRM_RATIONALE: "GGKK LIIHHH", CDE_NAME: "Legal Name", CDE_RATIONALE: "Rationale 1" },
                { PRIORITY_REPORT: "FM Exposure", PRIORITY_REPORT_RATIONALE: "ZXCV", CRM_NAME: "CAT1", CRM_RATIONALE: "GGKK LIIHHH", CDE_NAME: "Incorporation Date", CDE_RATIONALE: "Rationale 2" },
                { PRIORITY_REPORT: "FM Exposure", PRIORITY_REPORT_RATIONALE: "ZXCV", CRM_NAME: "EEPE", CRM_RATIONALE: "GGKK LIIHHH", CDE_NAME: "Legal Name", CDE_RATIONALE: "Rationale 1" },
                // { PRIORITY_REPORT: "FIN Rep", PRIORITY_REPORT_RATIONALE: "MM abcde", CRM_NAME: "xx", CRM_RATIONALE: "GGG", CDE_NAME: "", CDE_RATIONALE: "" },
                // { PRIORITY_REPORT: "FIN Rep", PRIORITY_REPORT_RATIONALE: "fghijk lmnop", CRM_NAME: "yy", CRM_RATIONALE: "GGG", CDE_NAME: "", CDE_RATIONALE: "" },
            ];

            res.DataFlat = _.cloneDeep(res.Data);

            var tmp = _.groupBy(res.Data, "PRIORITY_REPORT")
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmpTmp = _.cloneDeep(tmp[v]);

                var ret = tmpTmp[0];
                ret.ID = i;
                ret.PRIORITY_REPORTsVal = tmpTmp;
                ret.PRIORITY_REPORT = v;

                var tmp2 = _.groupBy(tmp[v], "PRIORITY_REPORT_RATIONALE");  
                ret.PRIORITY_REPORT_RATIONALEs = _.map(Object.keys(tmp2), function(w, j){
                    var tmpTmp2 = _.cloneDeep(tmp2[w]);

                    ret.PRIORITY_REPORT_RATIONALEsVal = tmpTmp2;
                    ret.PRIORITY_REPORT_RATIONALE = w;

                    return ret;
                });

                var tmp3 = _.groupBy(tmp[v], "CRM_NAME");  
                ret.CRM_NAMEs = _.map(Object.keys(tmp3), function(w, j){
                    var tmpTmp3 = _.cloneDeep(tmp3[w]);

                    ret.CRM_NAMEsVal = tmpTmp3;
                    ret.CRM_NAME = w;

                    return ret;
                });

                var tmp4 = _.groupBy(tmp[v], "CRM_RATIONALE");  
                ret.CRM_RATIONALEs = _.map(Object.keys(tmp4), function(w, j){
                    var tmpTmp4 = _.cloneDeep(tmp4[w]);

                    ret.CRM_RATIONALEsVal = tmpTmp4;
                    ret.CRM_RATIONALE = w;

                    return ret;
                });

                return ret;
            });

            res.Data.forEach(v => {
                v.PRIORITY_REPORT_RATIONALEs.shift();
                v.CRM_NAMEs.shift();
                v.CRM_RATIONALEs.shift();
            });
            
            console.log(res.Data);
            return res;
        }
    );
}