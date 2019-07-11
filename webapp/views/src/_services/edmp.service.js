import { fetchWHeader } from '../_helpers/auth-header';

export const edmpService = {
    getHomepageCounts
};

function getHomepageCounts(param) {
    return fetchWHeader(`/dsc/gethomepagecounts`, param).then(res => {
        res.Data = res.Data[0] ? res.Data[0] : null;
        return res;
    });
}