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
            res.Data = _.map(res.Data, function(v){
                var keys = Object.keys(v);

                keys.forEach(key => {
                    v[key] = v[key] ? v[key] : "NA";
                })
                
                return v;
            });

            res.DataFlat = _.cloneDeep(res.Data);
            
            var tmp = _.groupBy(res.Data, "BT_NAME")
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmpTmp = _.cloneDeep(tmp[v]);

                var ret         = tmpTmp[0];
                ret.ID          = i;
                ret.BT_NAMEsVal = tmpTmp;
                ret.BT_NAME     = v;

                var tmp2 = _.groupBy(tmp[v], "SYSTEM_CONSUMED");  
                ret.Systems = _.map(Object.keys(tmp2), function(w, j){
                    var tmpTmp2 = _.cloneDeep(tmp2[w]);

                    var ret         = tmpTmp2[0];
                    ret.SYSID       = j;
                    ret.SystemsVal   = tmpTmp2;
                    ret.SYSTEM_CONSUMED  = w;

                    var tmp3 = _.groupBy(tmp2[w], "TABLE_NAME");
                    ret.Tables = _.map(Object.keys(tmp3), function(x, k){
                        var tmpTmp3 = _.cloneDeep(tmp3[x]);

                        var ret         = tmpTmp3[0];
                        ret.TMTID       = k;
                        ret.TablesVal  = tmpTmp3;
                        ret.TABLE_NAME = x;

                        var tmp4 = _.groupBy(tmp3[x], "COLUMN_NAME");
                        ret.Columns = _.map(Object.keys(tmp4), function(y, l){
                            var tmpTmp4 = _.cloneDeep(tmp4[y]);

                            var ret         = tmpTmp4[0];
                            ret.COLID       = l;
                            ret.ColumnsVal  = tmpTmp4;
                            ret.COLUMN_NAME = y;

                            return ret;
                        });

                        return ret;
                    });

                    return ret;
                });

                var tmp4 = _.groupBy(tmp[v], "GS_SYSTEM_NAME")
                ret.GSSystems = _.map(Object.keys(tmp4), function(w, j){
                    var tmpTmp4 = _.cloneDeep(tmp4[w]);

                    var ret             = tmpTmp4[0];
                    ret.ID              = j;
                    ret.GSSystemsVal    = tmpTmp4;
                    ret.GS_SYSTEM_NAME  = w;

                    var tmp5 = _.groupBy(tmp4[w], "GS_TABLE_NAME");  
                    ret.GSTables = _.map(Object.keys(tmp5), function(x, k){
                        var tmpTmp5 = _.cloneDeep(tmp5[x]);

                        var ret             = tmpTmp5[0];
                        ret.TMTID           = k;
                        ret.GSTablesVal     = tmpTmp5;
                        ret.GS_TABLE_NAME   = x;

                        var tmp6 = _.groupBy(tmp5[x], "GS_COLUMN_NAME");
                        ret.GSColumns = _.map(Object.keys(tmp6), function(y, l){
                            var tmpTmp6 = _.cloneDeep(tmp6[y]);

                            var ret             = tmpTmp6[0];
                            ret.COLID           = l;
                            ret.GSColumnsVal    = tmpTmp6;
                            ret.GS_COLUMN_NAME  = y;

                            return ret;
                        });

                        return ret;
                    });

                    return ret;
                });


                return ret;
            });

            res.Data.forEach(v => {
                v.Systems.forEach(w => {
                    w.Tables.forEach(x => {
                        x.Columns.shift();
                    });

                    if(w.Tables[0].Columns.length == 0){
                        w.Tables.shift();
                    }
                });

                if(v.Systems[0].Tables.length == 0){
                    v.Systems.shift();
                }

                v.GSSystems.forEach(w => {
                    w.GSTables.forEach(x => {
                        x.GSColumns.shift();
                    });

                    if(w.GSTables[0].GSColumns.length == 0){
                        w.GSTables.shift();
                    }
                });
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
    return fetchWHeader(`/ddo/getdetails`, param).then(
        res => {
            var tmp = _.groupBy(res.Data.Detail, "ID")
            
            res.Data.Detail = _.map(Object.keys(tmp), function(v, i){
                var ret = tmp[v][0];
                ret.ID = v;
                ret.Values = tmp[v];

                return ret;
            });

            return res;
        }
    )
}