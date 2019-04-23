import { fetchWHeader } from '../_helpers/auth-header';

export const dscMyService = {
    getLeftTable,
    getHomepageCounts,
    getCdeTable,
    getCdpTable,
    getCdpCdeTable,
    getInterfacesTable,
    getInterfacesCdeTable,
    getRightTable,
    getInterfacesRightTable,
    getDdTable,
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

function getHomepageCounts(param) {
    return fetchWHeader(`/dsc/gethomepagecounts`, param).then(res => {
        res.Data = res.Data[0] ? res.Data[0] : null;
        return res;
    });
}

function getCdeTable(param) {
    return fetchWHeader(`/dsc/getcdetable`, param).then(
        res => {
            res.DataFlat = _.cloneDeep(res.Data);
            
            var tmp = _.groupBy(res.Data, "CDE")
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmpTmp = _.cloneDeep(tmp[v]);

                var ret         = tmpTmp[0];
                ret.ID          = i;
                ret.CDEsVal     = tmpTmp;
                ret.CDE         = v;

                var tmp2 = _.groupBy(tmp[v], "TABLE_NAME");  
                ret.Tables = _.map(Object.keys(tmp2), function(w, i){
                    var tmpTmp2 = _.cloneDeep(tmp2[w]);

                    var ret         = tmpTmp2[0];
                    ret.TMTID       = tmpTmp2[0].ID;
                    ret.TablesVal   = tmpTmp2;
                    ret.TABLE_NAME  = w;

                    var tmp3 = _.groupBy(tmp2[w], "COLUMN_NAME");
                    ret.Columns = _.map(Object.keys(tmp3), function(x, i){
                        var tmpTmp3 = _.cloneDeep(tmp3[x]);

                        var ret         = tmpTmp3[0];
                        ret.COLID       = tmpTmp3[0].COLID;
                        ret.ColumnsVal  = tmpTmp3;
                        ret.COLUMN_NAME = x;

                        var tmp4 = _.groupBy(tmp3[x], "DSP_NAME");  
                        ret.Dsps = _.map(Object.keys(tmp4), function(x, i){
                            var tmpTmp4 = _.cloneDeep(tmp4[x]);

                            var ret         = tmpTmp4[0];
                            ret.DSPID       = i;
                            ret.DspsVal     = tmpTmp4;
                            ret.DSP_NAME    = x;

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
                        x.Dsps.shift();
                    });

                    if(w.Columns[0].Dsps.length == 0){
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

function getCdpTable(param) {
    return fetchWHeader(`/dsc/getcdptable`, param);
}

function getCdpCdeTable(param) {
    return fetchWHeader(`/dsc/getcdpcdetable`, param).then(
        res => {
            res.DataFlat = _.cloneDeep(res.Data);
            
            var tmp = _.groupBy(res.Data, "CDE")
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmpTmp = _.cloneDeep(tmp[v]);

                var ret         = tmpTmp[0];
                ret.ID          = i;
                ret.CDEsVal     = tmpTmp;
                ret.CDE         = v;

                var tmp2 = _.groupBy(tmp[v], "TABLE_NAME");  
                ret.Tables = _.map(Object.keys(tmp2), function(w, i){
                    var tmpTmp2 = _.cloneDeep(tmp2[w]);

                    var ret         = tmpTmp2[0];
                    ret.TMTID       = tmpTmp2[0].ID;
                    ret.TablesVal   = tmpTmp2;
                    ret.TABLE_NAME  = w;

                    var tmp3 = _.groupBy(tmp2[w], "COLUMN_NAME");
                    ret.Columns = _.map(Object.keys(tmp3), function(x, i){
                        var tmpTmp3 = _.cloneDeep(tmp3[x]);

                        var ret         = tmpTmp3[0];
                        ret.COLID       = tmpTmp3[0].COLID;
                        ret.ColumnsVal  = tmpTmp3;
                        ret.COLUMN_NAME = x;

                        return ret;
                    });

                    return ret;
                });

                return ret;
            });

            res.Data.forEach(v => {
                v.Tables.forEach(w => {
                    w.Columns.shift();
                });

                if(v.Tables[0].Columns.length == 0){
                    v.Tables.shift();
                }
            });
            
            return res;
        }
    );
}

function getInterfacesTable(param) {
    return fetchWHeader(`/dsc/getinterfacestable`, param);
}

function getInterfacesCdeTable(param) {
    return fetchWHeader(`/dsc/getinterfacescdetable`, param).then(
        res => {
            res.DataFlat = _.cloneDeep(res.Data);
            
            var tmp = _.groupBy(res.Data, "CDE")
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmpTmp = _.cloneDeep(tmp[v]);

                var ret         = tmpTmp[0];
                ret.ID          = i;
                ret.CDEsVal     = tmpTmp;
                ret.CDE         = v;

                var tmp2 = _.groupBy(tmp[v], "TABLE_NAME");  
                ret.Tables = _.map(Object.keys(tmp2), function(w, i){
                    var tmpTmp2 = _.cloneDeep(tmp2[w]);

                    var ret         = tmpTmp2[0];
                    ret.TMTID       = tmpTmp2[0].ID;
                    ret.TablesVal   = tmpTmp2;
                    ret.TABLE_NAME  = w;

                    var tmp3 = _.groupBy(tmp2[w], "COLUMN_NAME");
                    ret.Columns = _.map(Object.keys(tmp3), function(x, i){
                        var tmpTmp3 = _.cloneDeep(tmp3[x]);

                        var ret         = tmpTmp3[0];
                        ret.COLID       = tmpTmp3[0].COLID;
                        ret.ColumnsVal  = tmpTmp3;
                        ret.COLUMN_NAME = x;

                        return ret;
                    });

                    return ret;
                });

                return ret;
            });

            res.Data.forEach(v => {
                v.Tables.forEach(w => {
                    w.Columns.shift();
                });

                if(v.Tables[0].Columns.length == 0){
                    v.Tables.shift();
                }
            });
            
            return res;
        }
    );
}

function getDdTable(param) {
    return fetchWHeader(`/dsc/getddtable`, param).then(
        res => {
            res.DataFlat = _.cloneDeep(res.Data);

            res.Data = _.map(res.Data, function(v){
                v.CDE_YES_NO = v.CDE_YES_NO == 0 ? "No" : "Yes";
                return v;
            });

            return res;
        }
    );
}

function getRightTable(param) {
    return fetchWHeader(`/dsc/gettablename`, param).then(
        res => {
            res.DataFlat = _.cloneDeep(res.Data);

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
            res.DataFlat = _.cloneDeep(res.Data);

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