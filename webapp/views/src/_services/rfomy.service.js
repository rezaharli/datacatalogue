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

                var tmp2 = _.groupBy(tmp[v], "RISK_SUB_TYPE");  
                ret.RISK_SUB_TYPEs = _.map(Object.keys(tmp2), function(w, j){
                    var tmpTmp2 = _.cloneDeep(tmp2[w]);

                    var ret = tmpTmp2[0];
                    ret.ID = j;
                    ret.RISK_SUB_TYPEsVal = tmpTmp2;
                    ret.RISK_SUB_TYPE = w;

                    return ret;
                });

                return ret;
            });

            res.Data.forEach(v => {
                v.RISK_SUB_TYPEs.shift();
            });
            
            return res;
        },
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
                    var tmpTmp2 = (tmp2[w]);

                    var ret = tmpTmp2[0];
                    ret.ID = j;
                    ret.PRIORITY_REPORT_RATIONALEsVal = tmpTmp2;
                    ret.PRIORITY_REPORT_RATIONALE = w;

                    return ret;
                });

                var tmp3 = _.groupBy(tmp[v], "CRM_NAME");  
                ret.CRM_NAMEs = _.map(Object.keys(tmp3), function(w, j){
                    var tmpTmp3 = (tmp3[w]);

                    var ret = tmpTmp3[0];
                    ret.ID = j;
                    ret.CRM_NAMEsVal = tmpTmp3;
                    ret.CRM_NAME = w;

                    return ret;
                });

                var tmp4 = _.groupBy(tmp[v], "CRM_RATIONALE");  
                ret.CRM_RATIONALEs = _.map(Object.keys(tmp4), function(w, j){
                    var tmpTmp4 = (tmp4[w]);

                    var ret = tmpTmp4[0];
                    ret.ID = j;
                    ret.CRM_RATIONALEsVal = tmpTmp4;
                    ret.CRM_RATIONALE = w;

                    return ret;
                });

                var tmp5 = _.groupBy(tmp[v], "CDE_NAME");  
                ret.CDE_NAMEs = _.map(Object.keys(tmp5), function(w, j){
                    var tmpTmp5 = (tmp5[w]);

                    var ret = tmpTmp5[0];
                    ret.ID = j;
                    ret.CDE_NAMEsVal = tmpTmp5;
                    ret.CDE_NAME = w;

                    return ret;
                });

                return ret;
            });

            res.Data.forEach(v => {
                v.PRIORITY_REPORT_RATIONALEsVal = _.cloneDeep(v.PRIORITY_REPORT_RATIONALEs);
                v.CRM_NAMEsVal = _.cloneDeep(v.CRM_NAMEs);
                v.CRM_RATIONALEsVal = _.cloneDeep(v.CRM_RATIONALEs);
                v.CDE_NAMEsVal = _.cloneDeep(v.CDE_NAMEs);
                
                v.PRIORITY_REPORT_RATIONALEs.shift();
                v.CRM_NAMEs.shift();
                v.CRM_RATIONALEs.shift();
                v.CDE_NAMEs.shift();

                var lengths = [v.PRIORITY_REPORT_RATIONALEs.length + 1, v.CRM_NAMEs.length + 1, v.CRM_RATIONALEs.length + 1, v.CDE_NAMEs.length + 1];
                v.rowspanAcuan = lcm(lengths);

                v.theMostLength = _.max([v.PRIORITY_REPORT_RATIONALEs, v.CRM_NAMEs, v.CRM_RATIONALEs, v.CDE_NAMEs], v => v.length);

                v.childrenRow = [];
                v.theMostLength.forEach((w, i) => { // 012
                    var PRIORITY_REPORT_RATIONALEsSkip = (v.rowspanAcuan / (v.PRIORITY_REPORT_RATIONALEs.length + 1)) - 1;
                    var CRM_NAMEsSkip = (v.rowspanAcuan / (v.CRM_NAMEs.length + 1)) - 1;
                    var CRM_RATIONALEsSkip = (v.rowspanAcuan / (v.CRM_RATIONALEs.length + 1)) - 1;
                    var CDE_NAMEsSkip = (v.rowspanAcuan / (v.CDE_NAMEs.length + 1)) - 1;

                    v.childrenRow.push({
                        PRIORITY_REPORT_RATIONALE: i >= PRIORITY_REPORT_RATIONALEsSkip ? (v.PRIORITY_REPORT_RATIONALEs[i - (PRIORITY_REPORT_RATIONALEsSkip)] ? v.PRIORITY_REPORT_RATIONALEs[i - (PRIORITY_REPORT_RATIONALEsSkip)].PRIORITY_REPORT_RATIONALE : null) : null,
                        CRM_NAME: i >= CRM_NAMEsSkip ? (v.CRM_NAMEs[i - (CRM_NAMEsSkip)] ? v.CRM_NAMEs[i - (CRM_NAMEsSkip)].CRM_NAME : null) : null,
                        CRM_RATIONALE: i >= CRM_RATIONALEsSkip ? (v.CRM_RATIONALEs[i - (CRM_RATIONALEsSkip)] ? v.CRM_RATIONALEs[i - (CRM_RATIONALEsSkip)].CRM_RATIONALE : null) : null,
                        CDE_NAME: i >= CDE_NAMEsSkip ? (v.CDE_NAMEs[i - (CDE_NAMEsSkip)] ? v.CDE_NAMEs[i - (CDE_NAMEsSkip)].CDE_NAME : null) : null,
                        CDE_RATIONALE: i >= CDE_NAMEsSkip ? (v.CDE_NAMEs[i - (CDE_NAMEsSkip)] ? v.CDE_NAMEs[i - (CDE_NAMEsSkip)].CDE_RATIONALE : null) : null,
                    })
                });
                
                v.expanded = false;
            });
            
            return res;
        },
    );
}