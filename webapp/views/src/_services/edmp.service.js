import { fetchWHeader } from '../_helpers/auth-header';

export const edmpService = {
    getHomepageCounts,
    getTechnicalTable,
    getBusinessTable,
    getConsumptionTable,
    getDropdownOpts,
    getIarcPersonal
};

function getHomepageCounts(param) {
    return fetchWHeader(`/dsc/gethomepagecounts`, param).then(res => {
        res.Data = res.Data[0] ? res.Data[0] : null;
        return res;
    });
}

function getTechnicalTable(param) {
    return fetchWHeader(`/dsc/getedmpddtechnicaltable`, param);
}

function getBusinessTable(param) {
    return fetchWHeader(`/dsc/getedmpddbusinesstable`, param);
}

function getConsumptionTable(param) {
    return fetchWHeader(`/dsc/getedmpddconsumptiontable`, param);
}

function getDropdownOpts(param) {
    return fetchWHeader(`/dsc/getedmpdddropdowns`, param);
}

function getIarcPersonal(param) {
    return fetchWHeader(`/dsc/getedmpiarcpersonaltable`, param);
}
