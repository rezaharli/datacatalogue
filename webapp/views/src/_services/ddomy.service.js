import { fetchWHeader } from '../_helpers/auth-header';

export const ddoMyService = {
    getLeftTable,
    getHomepageCounts,
    getBusinesstermTable,
    getSystemsTable,
    getSystemsBusinesstermTable,
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
    return fetchWHeader(`/ddo/getsystemsbusinesstermtable`, param);
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
    // var urls = [
    //     { name: 'DetailsBusinessMetadata', url: `/ddo/getdetailsbusinessmetadatafromdomain` },
    //     { name: 'DDSourceBusinessMetadata', url: `/ddo/getddsourcebusinessmetadatafromdomain` },
    //     { name: 'DetailsDownstreamUsage', url: `/ddo/getdetailsdownstreamusageofbusinessterm` }, 
    //     { name: 'DDSourceDownstreamUsage', url: `/ddo/getddsourcedownstreamusageofbusinessterm` },
    //     { name: 'DetailsBTResiding', url: `/ddo/getdetailsbtresiding` },
    //     { name: 'DDSourceBTResiding', url: `/ddo/getddsourcebtresiding` },
    // ]

    // return Promise.all(
    //     urls.map(v => fetchWHeader(v.url, param).then(resp => resp.Data))
    // ).then(results => {
    //     var res = {}
    //     res.Data = {}

    //     urls.forEach((v, i) => { res.Data[v.name] = results[i] });

    //     var tmp = _.groupBy(res.Data.DetailsBusinessMetadata, "ID")
    //     res.Data.DetailsBusinessMetadata = _.map(Object.keys(tmp), (v, i) => {
    //         var ret = tmp[v][0];
    //         ret.ID = v;
    //         ret.Values = tmp[v];

    //         return ret;
    //     });

    //     var tmp = _.groupBy(res.Data.DetailsDownstreamUsage, "ID")
    //     res.Data.DetailsDownstreamUsage = _.map(Object.keys(tmp), (v, i) => {
    //         var ret = tmp[v][0];
    //         ret.ID = v;
    //         ret.Values = tmp[v];

    //         return ret;
    //     });

    //     var tmp = _.groupBy(res.Data.DetailsBTResiding, "ID")
    //     res.Data.DetailsBTResiding = _.map(Object.keys(tmp), (v, i) => {
    //         var ret = tmp[v][0];
    //         ret.ID = v;
    //         ret.Values = tmp[v];

    //         return ret;
    //     });

    //     return res;
    // });

    return fetchWHeader(`/ddo/getdetails`, param).then(
        res => {
            var tmp = _.groupBy(res.Data.DetailsBusinessMetadata, "ID")
            res.Data.DetailsBusinessMetadata = _.map(Object.keys(tmp), function(v, i){
                var ret = tmp[v][0];
                ret.ID = v;
                ret.Values = tmp[v];

                return ret;
            });

            var tmp = _.groupBy(res.Data.DetailsDownstreamUsage, "ID")
            res.Data.DetailsDownstreamUsage = _.map(Object.keys(tmp), function(v, i){
                var ret = tmp[v][0];
                ret.ID = v;
                ret.Values = tmp[v];

                return ret;
            });

            var tmp = _.groupBy(res.Data.DetailsBTResiding, "ID")
            res.Data.DetailsBTResiding = _.map(Object.keys(tmp), function(v, i){
                var ret = tmp[v][0];
                ret.ID = v;
                ret.Values = tmp[v];

                return ret;
            });

            return res;
        }
    )

}