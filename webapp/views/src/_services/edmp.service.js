import { fetchWHeader } from '../_helpers/auth-header';

export const edmpService = {
    getHomepageCounts,
    getTechnicalTable,
    getBusinessTable,
    getConsumptionTable,
    getDdDropdownOpts,
    getIarcDropdownOpts,
    getIarcPersonalTable
};

function getHomepageCounts(param) {
    return fetchWHeader(`/dsc/gethomepagecounts`, param).then(res => {
        res.Data = res.Data[0] ? res.Data[0] : null;
        return res;
    });
}

function getTechnicalTable(param) {
    return fetchWHeader(`/dsc/getedmpddtechnicaltable`, param).then(
        res => {
            res.DataFlat = _.cloneDeep(res.Data);

            var tmp = _.groupBy(res.Data, "TABLE_NAME")
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmpTmp = _.cloneDeep(tmp[v]);

                var ret         = tmpTmp[0];
                ret.ID          = i;
                ret.TABLE_NAMEsVal     = tmpTmp;
                ret.TABLE_NAME         = v;

                var tmp2 = _.groupBy(tmp[v], "COLUMN_NAME");  
                ret.Tables = _.map(Object.keys(tmp2), function(w, i){
                    var tmpTmp2 = _.cloneDeep(tmp2[w]);

                    var ret         = tmpTmp2[0];
                    ret.TMTID       = tmpTmp2[0].TMTID;
                    ret.TablesVal   = tmpTmp2;
                    ret.COLUMN_NAME  = w;

                    return ret;
                });

                return ret;
            });

            res.Data.forEach(v => {
                v.Tables.shift();
            });
            
            return res;
        }
    );
}

function getBusinessTable(param) {
    return fetchWHeader(`/dsc/getedmpddbusinesstable`, param).then(
        res => {
            res.DataFlat = _.cloneDeep(res.Data);

            var tmp = _.groupBy(res.Data, "TABLE_NAME")
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmpTmp = _.cloneDeep(tmp[v]);

                var ret         = tmpTmp[0];
                ret.ID          = i;
                ret.TABLE_NAMEsVal     = tmpTmp;
                ret.TABLE_NAME         = v;

                var tmp2 = _.groupBy(tmp[v], "COLUMN_NAME");  
                ret.Tables = _.map(Object.keys(tmp2), function(w, i){
                    var tmpTmp2 = _.cloneDeep(tmp2[w]);

                    var ret         = tmpTmp2[0];
                    ret.TMTID       = tmpTmp2[0].TMTID;
                    ret.TablesVal   = tmpTmp2;
                    ret.COLUMN_NAME  = w;

                    return ret;
                });

                return ret;
            });

            res.Data.forEach(v => {
                v.Tables.shift();
            });
            
            return res;
        }
    );
}

function getConsumptionTable(param) {
    return fetchWHeader(`/dsc/getedmpddconsumptiontable`, param).then(
        res => {
            res.DataFlat = _.cloneDeep(res.Data);

            var tmp = _.groupBy(res.Data, "TABLE_NAME")
            res.Data = _.map(Object.keys(tmp), function(v, i){
                var tmpTmp = _.cloneDeep(tmp[v]);

                var ret         = tmpTmp[0];
                ret.ID          = i;
                ret.TABLE_NAMEsVal     = tmpTmp;
                ret.TABLE_NAME         = v;

                var tmp2 = _.groupBy(tmp[v], "COLUMN_NAME");  
                ret.Tables = _.map(Object.keys(tmp2), function(w, i){
                    var tmpTmp2 = _.cloneDeep(tmp2[w]);

                    var ret         = tmpTmp2[0];
                    ret.TMTID       = tmpTmp2[0].TMTID;
                    ret.TablesVal   = tmpTmp2;
                    ret.COLUMN_NAME  = w;

                    return ret;
                });

                return ret;
            });

            res.Data.forEach(v => {
                v.Tables.shift();
            });
            
            return res;
        }
    );
}

function getDdDropdownOpts(param) {
    return fetchWHeader(`/dsc/getedmpdddropdowns`, param);
}

function getIarcDropdownOpts(param) {
    return fetchWHeader(`/dsc/getedmpiarcdropdowns`, param);
}

function getIarcPersonalTable(param) {
    return fetchWHeader(`/dsc/getedmpiarcpersonaltable`, param).then(
        res => {
            res.DataFlat = _.cloneDeep(res.Data);
            return res;
        }
    );
}
