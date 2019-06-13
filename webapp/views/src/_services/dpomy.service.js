import { fetchWHeader } from '../_helpers/auth-header';

export const dpoMyService = {
    getLeftTable,
    getHomepageCounts,
    getDataelementsTable,
    getDatalineageTable,
    getRightTable,
    getDetails
};

function getLeftTable(param) {
    return fetchWHeader(`/dpo/getlefttable`, param)
}

function getHomepageCounts(param) {
    return fetchWHeader(`/dpo/gethomepagecounts`, param).then(res => {
        res.Data = res.Data[0] ? res.Data[0] : null;
        return res;
    });
}

function getDataelementsTable(param) {
    return fetchWHeader(`/dpo/getdataelementstable`, param).then(
        res => {
            res.DataFlat = _.cloneDeep(res.Data);

            res.Data = _.map(res.Data, function(v){
                v.CDE = v.CDE == 0 ? "No" : "Yes";
                
                if(v.DATA_SLA){
                    v.DATA_SLA = v.DATA_SLA == 0 ? "No" : "Yes";
                }else{
                    v.DATA_SLA = "NA";
                }

                return v;
            });
            
            var tmp = _.groupBy(res.Data, "ALIAS_NAME")
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmpTmp = _.cloneDeep(tmp[v]);

                var ret         = tmpTmp[0];
                ret.ID          = i;
                ret.BT_NAMEsVal = tmpTmp;
                ret.BT_NAME     = v;

                var tmp2 = _.groupBy(tmp[v], "TABLE_NAME");  
                ret.Tables = _.map(Object.keys(tmp2), function(w, j){
                    var tmpTmp2 = _.cloneDeep(tmp2[w]);

                    var ret         = tmpTmp2[0];
                    ret.TMTID       = j;
                    ret.TablesVal   = tmpTmp2;
                    ret.TABLE_NAME  = w;

                    var tmp3 = _.groupBy(tmp2[w], "COLUMN_NAME");
                    ret.Columns = _.map(Object.keys(tmp3), function(x, k){
                        var tmpTmp3 = _.cloneDeep(tmp3[x]);

                        var ret         = tmpTmp3[0];
                        ret.COLID       = k;
                        ret.ColumnsVal  = tmpTmp3;
                        ret.COLUMN_NAME = x;

                        var tmp4 = _.groupBy(tmp3[x], "ULT_SYSTEM_NAME");
                        ret.ULT_SYSTEM_NAMEs = _.map(Object.keys(tmp4), function(y, l){
                            var tmpTmp4 = _.cloneDeep(tmp4[y]);

                            var ret         = tmpTmp4[0];
                            ret.ULTSYSID       = l;
                            ret.ULT_SYSTEM_NAMEsVal  = tmpTmp4;
                            ret.ULT_SYSTEM_NAME = y;

                            return ret;
                        });

                        return ret;
                    });

                    return ret;
                });

                return ret;
            });

            res.Data.forEach(v => {
                v.Tables.forEach(w => {
                    w.Columns.forEach(x => {
                        x.ULT_SYSTEM_NAMEs.shift();
                    });

                    if(w.Columns[0].ULT_SYSTEM_NAMEs.length == 0){
                        w.Columns.shift();
                    }
                });

                if(v.Tables[0].Columns.length == 0){
                    v.Tables.shift();
                }
            });
            
            return res;
        }
    );
}

function getDatalineageTable(param) {
    return fetchWHeader(`/dpo/getdatalineagetable`, param).then(
        res => {
            res.Data = res.Data.filter(v => v.PR_NAME != "");
            res.DataFlat = _.cloneDeep(res.Data);

            var tmp = _.groupBy(res.Data, "SYSTEM_NAME")
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmpTmp = _.cloneDeep(tmp[v]);

                var ret = tmpTmp[0];
                ret.ID = i;
                ret.SYSTEM_NAMEsVal = tmpTmp;
                ret.SYSTEM_NAME = v;

                var tmp2 = _.groupBy(tmp[v], "PR_NAME");  
                ret.PR_NAMEs = _.map(Object.keys(tmp2), function(w, j){
                    var tmpTmp2 = _.cloneDeep(tmp2[w]);

                    var ret = tmpTmp2[0];
                    ret.ID = j;
                    ret.PR_NAMEsVal = tmpTmp2;
                    ret.PR_NAME = w;

                    return ret;
                });

                return ret;
            });

            res.Data.forEach(v => {
                v.SYSTEM_NAMEsVal.shift();
            });
            
            return res;
        },
    );
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
    return fetchWHeader(`/dpo/getdetails`, param);
}