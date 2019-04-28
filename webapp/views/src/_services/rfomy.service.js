import { fetchWHeader } from '../_helpers/auth-header';
import { lcm } from '../_helpers/helper';

export const rfoMyService = {
    getLeftTable,
    getRightTable,
    getAllRisk,
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

function getAllRisk(param) {
    return fetchWHeader(`/rfo/getlefttable`, param).then(
        res => {
            res.DataFlat = _.cloneDeep(res.Data);

            var tmp = _.groupBy(res.Data, "PRINCIPAL_RISK_TYPES")
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmpTmp = _.cloneDeep(tmp[v]);

                var ret = tmpTmp[0];
                ret.ID = i;
                ret.PRINCIPAL_RISK_TYPESsVal = tmpTmp;
                ret.PRINCIPAL_RISK_TYPES = v;

                return ret;
            });

            // res.Data.forEach(v => {
            //     v.RISK_SUB_TYPEs.shift();
            //     v.RISK_FRAMEWORK_OWNERs.shift();
            //     v.RISK_REPORTING_LEADs.shift();

            //     v.theMostLength = _.max([v.RISK_SUB_TYPEs, v.RISK_FRAMEWORK_OWNERs, v.RISK_REPORTING_LEADs], v => v.length);
            // });
            
            return res;
        },
        err => {
            var res = {};
            res.Data = [
                { PRINCIPAL_RISK_TYPES: "Traded", RISK_SUB_TYPE: "Credit", RISK_FRAMEWORK_OWNER: "John Doe", RISK_REPORTING_LEAD: "Mark R.", PRIORITY_REPORTS_UNIQUE_COUNT: "5", CRITICAL_RISK_MEASURES_UNIQUE_COUNT: "24", CRITICAL_DATA_ELEMENT_UNIQUE_COUNT: "40" },
                { PRINCIPAL_RISK_TYPES: "Traded", RISK_SUB_TYPE: "Market", RISK_FRAMEWORK_OWNER: "John Doe", RISK_REPORTING_LEAD: "Chris E.", PRIORITY_REPORTS_UNIQUE_COUNT: "7 Date", CRITICAL_RISK_MEASURES_UNIQUE_COUNT: "30", CRITICAL_DATA_ELEMENT_UNIQUE_COUNT: "35" },
                { PRINCIPAL_RISK_TYPES: "Capital & Liquidity", RISK_SUB_TYPE: "Capital", RISK_FRAMEWORK_OWNER: "Jess How", RISK_REPORTING_LEAD: "Robert D.J.", PRIORITY_REPORTS_UNIQUE_COUNT: "9", CRITICAL_RISK_MEASURES_UNIQUE_COUNT: "14", CRITICAL_DATA_ELEMENT_UNIQUE_COUNT: "50" },
                { PRINCIPAL_RISK_TYPES: "Capital & Liquidity", RISK_SUB_TYPE: "Liquidity", RISK_FRAMEWORK_OWNER: "Jess How", RISK_REPORTING_LEAD: "Scarlett J.", PRIORITY_REPORTS_UNIQUE_COUNT: "4", CRITICAL_RISK_MEASURES_UNIQUE_COUNT: "23", CRITICAL_DATA_ELEMENT_UNIQUE_COUNT: "43" },
            ];

            res.DataFlat = _.cloneDeep(res.Data);

            var tmp = _.groupBy(res.Data, "PRINCIPAL_RISK_TYPES")
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmpTmp = _.cloneDeep(tmp[v]);

                var ret = tmpTmp[0];
                ret.ID = i;
                ret.PRINCIPAL_RISK_TYPESsVal = tmpTmp;
                ret.PRINCIPAL_RISK_TYPES = v;

                return ret;
            });

            res.Data.forEach(v => {
                v.RISK_SUB_TYPEs.shift();
                v.RISK_FRAMEWORK_OWNERs.shift();
                v.RISK_REPORTING_LEADs.shift();

                v.theMostLength = _.max([v.RISK_SUB_TYPEs, v.RISK_FRAMEWORK_OWNERs, v.RISK_REPORTING_LEADs], v => v.length);
            });
            
            return res;
        }
    );
}

function getPriorityTable(param) {
    return fetchWHeader(`/rfo/getprioritytable`, param).then(
        res => {
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

                    var ret = tmpTmp2[0];
                    ret.ID = j;
                    ret.PRIORITY_REPORT_RATIONALEsVal = tmpTmp2;
                    ret.PRIORITY_REPORT_RATIONALE = w;

                    return ret;
                });

                var tmp3 = _.groupBy(tmp[v], "CRM_NAME");  
                ret.CRM_NAMEs = _.map(Object.keys(tmp3), function(w, j){
                    var tmpTmp3 = _.cloneDeep(tmp3[w]);

                    var ret = tmpTmp3[0];
                    ret.ID = j;
                    ret.CRM_NAMEsVal = tmpTmp3;
                    ret.CRM_NAME = w;

                    return ret;
                });

                var tmp4 = _.groupBy(tmp[v], "CRM_RATIONALE");  
                ret.CRM_RATIONALEs = _.map(Object.keys(tmp4), function(w, j){
                    var tmpTmp4 = _.cloneDeep(tmp4[w]);

                    var ret = tmpTmp4[0];
                    ret.ID = j;
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

                v.theMostLength = _.max([v.PRIORITY_REPORT_RATIONALEs, v.CRM_NAMEs, v.CRM_RATIONALEs], v => v.length);
                var lengths = [v.PRIORITY_REPORT_RATIONALEs.length + 1, v.CRM_NAMEs.length + 1, v.CRM_RATIONALEs.length + 1];
                
                v.rowspanAcuan = lcm(lengths);
                v.expanded = false;
            });
            
            console.log(res.Data);
            return res;
        },
    );
}