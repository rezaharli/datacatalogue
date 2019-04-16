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
            var tmp = _.groupBy(res.Data, "TABLE_NAME")
            
            res.Data = _.map(Object.keys(tmp), function(v){
                var tmp2 = _.groupBy(tmp[v], "COLUMN_NAME");  

                var columns = _.map(Object.keys(tmp2), function(w, i){
                    var tmp3 = _.groupBy(tmp2[w], "DSP_NAME");  

                    var dsps = _.map(Object.keys(tmp3), function(x, i){
                        var ret = tmp3[x][0];
                        ret.DSPID = i;
                        ret.DSP_NAME = x;
                        ret.DspsVal = tmp3[x];

                        return ret;
                    });

                    var ret = tmp2[w][0];
                    ret.COLID = tmp2[w][0].COLID;
                    ret.COLUMN_NAME = w;
                    ret.Dsps = dsps;
                    ret.ColumnsVal = tmp2[w];

                    ret.Dsps.shift();

                    return ret;
                });

                var ret = tmp[v][0];
                ret.TABLE_NAME = v;
                ret.TMTID = tmp[v][0].TMTID;
                ret.Columns = columns;
                ret.TablesVal = tmp[v];

                ret.Columns.shift();

                return ret;
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
            var tmp = _.groupBy(res.Data, "TABLE_NAME")
            
            res.Data = _.map(Object.keys(tmp), function(v){
                var tmp2 = _.groupBy(tmp[v], "COLUMN_NAME");  

                var columns = _.map(Object.keys(tmp2), function(w, i){
                    var ret = tmp2[w][0];
                    ret.COLID = tmp2[w][0].COLID;
                    ret.COLUMN_NAME = w;
                    ret.ColumnsVal = tmp2[w];

                    return ret;
                });

                var ret = tmp[v][0];
                ret.TABLE_NAME = v;
                ret.TMTID = tmp[v][0].TMTID;
                ret.Columns = columns;
                ret.TablesVal = tmp[v];

                ret.Columns.shift();

                return ret;
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
            var tmp = _.groupBy(res.Data, "TABLE_NAME")
            
            res.Data = _.map(Object.keys(tmp), function(v){
                var tmp2 = _.groupBy(tmp[v], "COLUMN_NAME");  

                var columns = _.map(Object.keys(tmp2), function(w, i){
                    var ret = tmp2[w][0];
                    ret.COLID = tmp2[w][0].COLID;
                    ret.COLUMN_NAME = w;
                    ret.ColumnsVal = tmp2[w];

                    return ret;
                });

                var ret = tmp[v][0];
                ret.TABLE_NAME = v;
                ret.TMTID = tmp[v][0].TMTID;
                ret.Columns = columns;
                ret.TablesVal = tmp[v];

                ret.Columns.shift();

                return ret;
            });
            
            return res;
        }
    );
}

function getRightTable(param) {
    return fetchWHeader(`/dsc/gettablename`, param).then(
        res => {
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