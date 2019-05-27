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
                v.DATA_SLA = v.DATA_SLA == 0 ? "No" : "Yes";
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

function getDatalineageTable(param) {
    return fetchWHeader(`/dpo/getdatalineagetable`, param);
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
    return fetchWHeader(`/dpo/getdetails`, param).then(
        res => {
            var tmp = _.groupBy(res.Data.DetailsImmediatePrecedingSystem, "ID")
            res.Data.DetailsImmediatePrecedingSystem = _.map(Object.keys(tmp), function(v, i){
                var ret = tmp[v][0];
                ret.ID = v;
                ret.Values = tmp[v];

                return ret;
            });

            var tmp = _.groupBy(res.Data.DetailsUltimateSourceSystem, "ID")
            res.Data.DetailsUltimateSourceSystem = _.map(Object.keys(tmp), function(v, i){
                var ret = tmp[v][0];
                ret.ID = v;
                ret.Values = tmp[v];

                return ret;
            });

            var tmp = _.groupBy(res.Data.DetailsDomainView, "ID")
            res.Data.DetailsDomainView = _.map(Object.keys(tmp), function(v, i){
                var ret = tmp[v][0];
                ret.ID = v;
                ret.Values = tmp[v];

                return ret;
            });

            var tmp = _.groupBy(res.Data.DetailsDataStandards, "ID")
            res.Data.DetailsDataStandards = _.map(Object.keys(tmp), function(v, i){
                var ret = tmp[v][0];
                ret.ID = v;
                ret.Values = tmp[v];

                return ret;
            });

            return res;
        }
    )

    return fetchWHeader(`/dpo/getdetails`, param).then(
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