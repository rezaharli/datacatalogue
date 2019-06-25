import { fetchWHeader } from '../_helpers/auth-header';

export const ddoMyService = {
    getLeftTable,
    getHomepageCounts,
    getBusinesstermTable,
    getSystemsTable,
    getSystemsBusinesstermTable,
    getDownstreamTable,
    getDownstreamBusinesstermTable,
    getRightTable,
    getDetails
};

function getLeftTable(param) {
    return fetchWHeader(`/ddo/getlefttable`, param).then(
        res => {
            var tmp = _.groupBy(res.Data, "SUB_DOMAINS")
            
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var ret = tmp[v][0];
                ret.SUB_DOMAINS = v;
                ret.Values = tmp[v];

                return ret;
            });

            return res;
        }
    )
}

function getHomepageCounts(param) {
    return fetchWHeader(`/ddo/gethomepagecounts`, param).then(res => {
        res.Data = res.Data[0] ? res.Data[0] : null;
        return res;
    });
}

function getBusinesstermTable(param) {
    return fetchWHeader(`/ddo/getbusinesstermtable`, param);
}

function getSystemsTable(param) {
    return fetchWHeader(`/ddo/getsystemstable`, param);
}

function getSystemsBusinesstermTable(param) {
    return fetchWHeader(`/ddo/getsystemsbusinesstermtable`, param).then(
        res => {
            res.DataFlat = _.cloneDeep(res.Data);
            
            var tmp = _.groupBy(res.Data, "BT_NAME")
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmpTmp = _.cloneDeep(tmp[v]);

                var ret         = tmpTmp[0];
                ret.ID          = i;
                ret.BT_NAMEsVal = tmpTmp;
                ret.BT_NAME     = v;

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

function getDownstreamTable(param) {
    return fetchWHeader(`/ddo/getdownstreamtable`, param);
}

function getDownstreamBusinesstermTable(param) {
    return fetchWHeader(`/ddo/getdownstreambusinesstermtable`, param).then(
        res => {
            res.Data = _.orderBy(res.Data, ["GOLDEN_SOURCE"], ["desc"])

            res.Data = _.map(res.Data, function(v){
                v.GOLDEN_SOURCE = v.GOLDEN_SOURCE.toLowerCase();
                v.GOLDEN_SOURCE = v.GOLDEN_SOURCE.charAt(0).toUpperCase() + v.GOLDEN_SOURCE.slice(1);
                var keys = Object.keys(v);

                keys.forEach(key => {
                    v[key] = v[key] ? v[key] : "NA";
                })
                
                return v;
            });

            res.DataFlat = _.cloneDeep(res.Data);
            
            var tmp = _.groupBy(res.Data, "BT_NAME")
            res.Data = _.map(Object.keys(tmp), function(v1, i1){
                var tmpTmp = _.cloneDeep(tmp[v1]);

                var ret         = tmpTmp[0];
                ret.ID          = i1;
                ret.BT_NAMEsVal = tmpTmp;
                ret.BT_NAME     = v1;

                var tmp2 = _.groupBy(tmp[v1], "ALIAS_NAME");  
                ret.ALIAS_NAMEs = _.map(Object.keys(tmp2), function(v2, i2){
                    var tmpTmp2 = _.cloneDeep(tmp2[v2]);

                    var ret             = tmpTmp2[0];
                    ret.ALIAS_NAMEID    = i2;
                    ret.ALIAS_NAMEsVal  = tmpTmp2;
                    ret.ALIAS_NAME      = v2;

                    var tmp3 = _.groupBy(tmp2[v2], "SYSTEM_CONSUMED");
                    ret.SYSTEM_CONSUMEDs = _.map(Object.keys(tmp3), function(v3, i3){
                        var tmpTmp3 = _.cloneDeep(tmp3[v3]);

                        var ret                 = tmpTmp3[0];
                        ret.SYSTEM_CONSUMEDID   = i3;
                        ret.SYSTEM_CONSUMEDsVal = tmpTmp3;
                        ret.SYSTEM_CONSUMED     = v3;

                        var tmp4 = _.groupBy(tmp3[v3], "TABLE_NAME");
                        ret.TABLE_NAMEs = _.map(Object.keys(tmp4), function(v4, i4){
                            var tmpTmp4 = _.cloneDeep(tmp4[v4]);

                            var ret             = tmpTmp4[0];
                            ret.TABLE_NAMEID    = i4;
                            ret.TABLE_NAMEsVal  = tmpTmp4;
                            ret.TABLE_NAME      = v4;

                            var tmp5 = _.groupBy(tmp4[v4], "COLUMN_NAME");
                            ret.COLUMN_NAMEs = _.map(Object.keys(tmp5), function(v5, i5){
                                var tmpTmp5 = _.cloneDeep(tmp5[v5]);
    
                                var ret                 = tmpTmp5[0];
                                ret.COLUMN_NAMEID       = i5;
                                ret.COLUMN_NAMEsVal     = tmpTmp5;
                                ret.COLUMN_NAME         = v5;

                                var tmp6 = _.groupBy(tmp5[v5], "GOLDEN_SOURCE");
                                ret.GOLDEN_SOURCEs = _.map(Object.keys(tmp6), function(v6, i6){
                                    var tmpTmp6 = _.cloneDeep(tmp6[v6]);
        
                                    var ret                 = tmpTmp6[0];
                                    ret.GOLDEN_SOURCEID     = i6;
                                    ret.GOLDEN_SOURCEsVal   = tmpTmp6;
                                    ret.GOLDEN_SOURCE       = v6;

                                    var tmp7 = _.groupBy(tmp6[v6], "GS_SYSTEM_NAME");
                                    ret.GS_SYSTEM_NAMEs = _.map(Object.keys(tmp7), function(v7, i7){
                                        var tmpTmp7 = _.cloneDeep(tmp7[v7]);
            
                                        var ret                 = tmpTmp7[0];
                                        ret.GS_SYSTEM_NAMEID     = i7;
                                        ret.GS_SYSTEM_NAMEsVal   = tmpTmp7;
                                        ret.GS_SYSTEM_NAME       = v7;

                                        var tmp8 = _.groupBy(tmp7[v7], "GS_TABLE_NAME");
                                        ret.GS_TABLE_NAMEs = _.map(Object.keys(tmp8), function(v8, i8){
                                            var tmpTmp8 = _.cloneDeep(tmp8[v8]);
                
                                            var ret                 = tmpTmp8[0];
                                            ret.GS_TABLE_NAMEID     = i8;
                                            ret.GS_TABLE_NAMEsVal   = tmpTmp8;
                                            ret.GS_TABLE_NAME       = v8;

                                            var tmp9 = _.groupBy(tmp8[v8], "GS_COLUMN_NAME");
                                            ret.GS_COLUMN_NAMEs = _.map(Object.keys(tmp9), function(v9, i9){
                                                var tmpTmp9 = _.cloneDeep(tmp9[v9]);
                    
                                                var ret                 = tmpTmp9[0];
                                                ret.GS_COLUMN_NAMEID     = i9;
                                                ret.GS_COLUMN_NAMEsVal   = tmpTmp9;
                                                ret.GS_COLUMN_NAME       = v9;
                    
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

                            return ret;
                        });

                        return ret;
                    });

                    return ret;
                });

                return ret;
            });

            res.Data.forEach(v => {
                v.ALIAS_NAMEs.forEach(w => {
                    w.SYSTEM_CONSUMEDs.forEach(x => {
                        x.TABLE_NAMEs.forEach(y => {
                            y.COLUMN_NAMEs.forEach(z => {
                                z.GOLDEN_SOURCEs.forEach(ww => {
                                    ww.GS_SYSTEM_NAMEs.forEach(wx => {
                                        wx.GS_TABLE_NAMEs.forEach(wy => {
                                            wy.GS_COLUMN_NAMEs.shift();
                                        });
                    
                                        if(wx.GS_TABLE_NAMEs[0].GS_COLUMN_NAMEs.length == 0){
                                            wx.GS_TABLE_NAMEs.shift();
                                        }
                                    });
                
                                    if(ww.GS_SYSTEM_NAMEs[0].GS_TABLE_NAMEs.length == 0){
                                        ww.GS_SYSTEM_NAMEs.shift();
                                    }
                                });
            
                                if(z.GOLDEN_SOURCEs[0].GS_SYSTEM_NAMEs.length == 0){
                                    z.GOLDEN_SOURCEs.shift();
                                }
                            });
        
                            if(y.COLUMN_NAMEs[0].GOLDEN_SOURCEs.length == 0){
                                y.COLUMN_NAMEs.shift();
                            }
                        });
    
                        if(x.TABLE_NAMEs[0].COLUMN_NAMEs.length == 0){
                            x.TABLE_NAMEs.shift();
                        }
                    });

                    if(w.SYSTEM_CONSUMEDs[0].TABLE_NAMEs.length == 0){
                        w.SYSTEM_CONSUMEDs.shift();
                    }
                });

                if(v.ALIAS_NAMEs[0].SYSTEM_CONSUMEDs.length == 0){
                    v.ALIAS_NAMEs.shift();
                }
            });
            
            return res;
        }
    );
}


function getRightTable(param) {
    return fetchWHeader(`/ddo/getrighttable`, param).then(
        res => {
            res.Data = _.map(res.Data, function(v){
                v.CDE_YES_NO = v.CDE_YES_NO == 0 ? "No" : "Yes";

                v.LEFTID = v.TSCID;
                return v;
            });

            return res;
        }
    )
}

function getDetails(param) {
    return fetchWHeader(`/ddo/getdetails`, param);
}