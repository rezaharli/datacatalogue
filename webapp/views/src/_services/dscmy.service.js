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
            res.DataFlat = _.cloneDeep(res.Data);

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
            res.Data = _.map(res.Data, function(v){
                var keys = Object.keys(v);

                keys.forEach(key => {
                    v[key] = v[key].toString().trim() ? v[key] : "-";
                })
                
                return v;
            });

            res.DataFlat = _.cloneDeep(res.Data);
            
            var tmp = _.groupBy(res.Data, "CDE")
            res.Data = _.map(Object.keys(tmp), function(v1, i1){
                var tmpTmp = _.cloneDeep(tmp[v1]);

                var ret         = tmpTmp[0];
                ret.ID          = i1;
                ret.CDEsVal = tmpTmp;
                ret.CDE     = v1;

                var tmp2 = _.groupBy(tmp[v1], "UPSTREAM_SYSTEM");  
                ret.UPSTREAM_SYSTEMs = _.map(Object.keys(tmp2), function(v2, i2){
                    var tmpTmp2 = _.cloneDeep(tmp2[v2]);

                    var ret             = tmpTmp2[0];
                    ret.UPSTREAM_SYSTEMID    = i2;
                    ret.UPSTREAM_SYSTEMsVal  = tmpTmp2;
                    ret.UPSTREAM_SYSTEM      = v2;

                    var tmp3 = _.groupBy(tmp2[v2], "TABLE_NAME");
                    ret.TABLE_NAMEs = _.map(Object.keys(tmp3), function(v3, i3){
                        var tmpTmp3 = _.cloneDeep(tmp3[v3]);

                        var ret                 = tmpTmp3[0];
                        ret.TABLE_NAMEID   = i3;
                        ret.TABLE_NAMEsVal = tmpTmp3;
                        ret.TABLE_NAME     = v3;

                        var tmp4 = _.groupBy(tmp3[v3], "COLUMN_NAME");
                        ret.COLUMN_NAMEs = _.map(Object.keys(tmp4), function(v4, i4){
                            var tmpTmp4 = _.cloneDeep(tmp4[v4]);

                            var ret             = tmpTmp4[0];
                            ret.COLUMN_NAMEID    = i4;
                            ret.COLUMN_NAMEsVal  = tmpTmp4;
                            ret.COLUMN_NAME      = v4;

                            var tmp5 = _.groupBy(tmp4[v4], "DSP_NAME");
                            ret.DSP_NAMEs = _.map(Object.keys(tmp5), function(v5, i5){
                                var tmpTmp5 = _.cloneDeep(tmp5[v5]);
    
                                var ret                 = tmpTmp5[0];
                                ret.DSP_NAMEID       = i5;
                                ret.DSP_NAMEsVal     = tmpTmp5;
                                ret.DSP_NAME         = v5;
    
                                return ret;
                            });

                            return ret;
                        });

                        return ret;
                    });

                    return ret;
                });

                return ret;
            });

            res.Data.forEach(v => {
                v.UPSTREAM_SYSTEMs.forEach(w => {
                    w.TABLE_NAMEs.forEach(x => {
                        x.COLUMN_NAMEs.forEach(y => {
                            y.DSP_NAMEs.shift();
                        });
    
                        if(x.COLUMN_NAMEs[0].DSP_NAMEs.length == 0){
                            x.COLUMN_NAMEs.shift();
                        }
                    });

                    if(w.TABLE_NAMEs[0].COLUMN_NAMEs.length == 0){
                        w.TABLE_NAMEs.shift();
                    }
                });

                if(v.UPSTREAM_SYSTEMs[0].TABLE_NAMEs.length == 0){
                    v.UPSTREAM_SYSTEMs.shift();
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
                    ret.TMTID       = tmpTmp2[0].TMTID;
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
    return fetchWHeader(`/dsc/getinterfacestable`, param).then(
        res => {
            res.DataFlat = _.cloneDeep(res.Data);
            
            var tmp = _.groupBy(res.Data, "IMM_INTERFACE")
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmpTmp = _.cloneDeep(tmp[v]);

                var ret         = tmpTmp[0];
                ret.ID          = i;
                ret.IMM_INTERFACEsVal     = tmpTmp;
                ret.IMM_INTERFACE         = v;

                var tmp2 = _.groupBy(tmp[v], "PROCESS_OWNER");  
                ret.Owners = _.map(Object.keys(tmp2), function(w, i){
                    var tmpTmp2 = _.cloneDeep(tmp2[w]);

                    var ret         = tmpTmp2[0];
                    ret.TMTID       = tmpTmp2[0].TMTID;
                    ret.OwnersVal   = tmpTmp2;
                    ret.PROCESS_OWNER  = w;

                    return ret;
                });

                return ret;
            });

            res.Data.forEach(v => {
                v.Owners.shift();
            });
            
            return res;
        }
    );
}

function getInterfacesCdeTable(param) {
    return fetchWHeader(`/dsc/getinterfacescdetable`, param).then(
        res => {
            res.Data = res.Data.concat(resData)

            res.DataFlat = _.cloneDeep(res.Data);
            
            var tmp = _.groupBy(res.Data, "CDE")
            res.Data = _.map(Object.keys(tmp), function(v1, i1){
                var tmpTmp = _.cloneDeep(tmp[v1]);

                var ret         = tmpTmp[0];
                ret.ID          = i1;
                ret.CDEsVal = tmpTmp;
                ret.CDE     = v1;

                var tmp2 = _.groupBy(tmp[v1], "TABLE_NAME");  
                ret.TABLE_NAMEs = _.map(Object.keys(tmp2), function(v2, i2){
                    var tmpTmp2 = _.cloneDeep(tmp2[v2]);

                    var ret             = tmpTmp2[0];
                    ret.TABLE_NAMEID    = i2;
                    ret.TABLE_NAMEsVal  = tmpTmp2;
                    ret.TABLE_NAME      = v2;

                    var tmp3 = _.groupBy(tmp2[v2], "COLUMN_NAME");
                    ret.COLUMN_NAMEs = _.map(Object.keys(tmp3), function(v3, i3){
                        var tmpTmp3 = _.cloneDeep(tmp3[v3]);

                        var ret                 = tmpTmp3[0];
                        ret.COLUMN_NAMEID   = i3;
                        ret.COLUMN_NAMEsVal = tmpTmp3;
                        ret.COLUMN_NAME     = v3;

                        var tmp4 = _.groupBy(tmp3[v3], "PRECEDING_SYSTEM");
                        ret.PRECEDING_SYSTEMs = _.map(Object.keys(tmp4), function(v4, i4){
                            var tmpTmp4 = _.cloneDeep(tmp4[v4]);

                            var ret             = tmpTmp4[0];
                            ret.PRECEDING_SYSTEMID    = i4;
                            ret.PRECEDING_SYSTEMsVal  = tmpTmp4;
                            ret.PRECEDING_SYSTEM      = v4;

                            var tmp5 = _.groupBy(tmp4[v4], "SUCCEDING_SYSTEM");
                            ret.SUCCEDING_SYSTEMs = _.map(Object.keys(tmp5), function(v5, i5){
                                var tmpTmp5 = _.cloneDeep(tmp5[v5]);
    
                                var ret                 = tmpTmp5[0];
                                ret.SUCCEDING_SYSTEMID       = i5;
                                ret.SUCCEDING_SYSTEMsVal     = tmpTmp5;
                                ret.SUCCEDING_SYSTEM         = v5;
    
                                return ret;
                            });

                            return ret;
                        });

                        return ret;
                    });

                    return ret;
                });

                return ret;
            });

            res.Data.forEach(v => {
                v.TABLE_NAMEs.forEach(w => {
                    w.COLUMN_NAMEs.forEach(x => {
                        x.PRECEDING_SYSTEMs.forEach(y => {
                            y.SUCCEDING_SYSTEMs.shift();
                        });
    
                        if(x.PRECEDING_SYSTEMs[0].SUCCEDING_SYSTEMs.length == 0){
                            x.PRECEDING_SYSTEMs.shift();
                        }
                    });

                    if(w.COLUMN_NAMEs[0].PRECEDING_SYSTEMs.length == 0){
                        w.COLUMN_NAMEs.shift();
                    }
                });

                if(v.TABLE_NAMEs[0].COLUMN_NAMEs.length == 0){
                    v.TABLE_NAMEs.shift();
                }
            });
            
            return res;
        }
    );
}

function getDdTable(param) {
    return fetchWHeader(`/dsc/getddtable`, param);
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
    return fetchWHeader(`/dsc/getdetails`, param);
}