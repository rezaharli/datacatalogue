import { fetchFile } from '../_helpers/auth-header';

export const exportService = {
    doExport
};

function doExport(param) {
    return fetchFile(`/base/exporttocsv`, param);
}