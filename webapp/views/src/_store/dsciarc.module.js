import { dscMyService } from '../_services/dscmy.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        tabName: '',
        filters: {
            left: {},
            right: {}
        },
        system: '',
        left: newTableObject(),
        exportDatas: [],
        leftHeaders: [
            // { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Sytem Name', value: 'SYSTEM_NAME' },
            // { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'ITAM ID', value: 'ITAM_ID' },
            // { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Table Name', value: 'TABLE_NAME' },
            // { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Column Name', value: 'COLUMN_NAME' },
            // { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Alias Name', value: 'BUSINESS_ALIAS_NAME' },
            // { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Alias Description', value: 'BUSINESS_ALIAS_DESCRIPTION' },
            // { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'CDE (Yes/No)', value: 'CDE_YES_NO' },
            
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Information Asset Names', value: 'INFORMATION_ASSET_NAMES' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Information Asset Description', value: 'INFORMATION_ASSET_DESCRIPTION' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'C - Confidentiality', value: 'CONFIDENTIALITY' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'I - Integrity', value: 'INTEGRITY' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'A - Availability', value: 'AVAILABILITY' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Overall CIA Rating', value: 'OVERALL_CIA_RATING' },
            // { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Record Categories', value: 'RECORD_CATEGORIES' },
            // { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'PII Flag', value: 'PII_FLAG' },
        ],
        isRightTable: false,
        DDSource: [],
        detailsLoading: true,
        detailsSource: [],
        error: null
    }
};

const actions = {
    exportData({ commit }) {
        commit('getExportDataRequest');

        Object.keys(state.all.filters.left).map(function(key, index) {
            state.all.filters.left[key] = (typeof(state.all.filters.left[key]) == "object") ? state.all.filters.left[key] : (state.all.filters.left[key] ? state.all.filters.left[key].toString() : "");
        });

        var param = {
            System: state.all.system,
            Filters: state.all.filters.left,
            Pagination: _.cloneDeep(state.all.left.pagination)
        }

        param.Pagination.rowsPerPage = -1;

        return dscMyService.getIarcTable(param)
            .then(
                res => commit('getExportDataSuccess', res),
                error => commit('getExportDataFailure', error)
            );
    },
    getLeftTable({ commit }, system) {
        commit('getLeftTableRequest');

        Object.keys(state.all.filters.left).map(function(key, index) {
            state.all.filters.left[key] = (typeof(state.all.filters.left[key]) == "object") ? state.all.filters.left[key] : (state.all.filters.left[key] ? state.all.filters.left[key].toString() : "");
        });

        var param = {
            System: state.all.system,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return dscMyService.getIarcTable(param)
            .then(
                res => commit('getLeftTableSuccess', res),
                error => commit('getLeftTableFailure', error)
            );
    },
};

const mutations = {
    getExportDataRequest(state) {
        state.all.left.isLoading = true;
    },
    getExportDataSuccess(state, res) {
        state.all.exportDatas = res.Data.Flat;

        state.all.left.isLoading = false;
    },
    getExportDataFailure(state, error) {
        state.all.left.isLoading = false;
        state.all.error = error;
    },
    getLeftTableRequest(state) {
        state.all.left.isLoading = true;
        state.all.left.source = [];
        state.all.left.display = [];
        state.all.left.totalItems = 0;
    },
    getLeftTableSuccess(state, res) {
        state.all.left.source = res.Data.Flat;
        state.all.left.display = res.Data.Grouped;
        state.all.left.totalItems = res.Data.Flat[0] ? res.Data.Flat[0].RESULT_COUNT : 0;

        state.all.left.isLoading = false;
    },
    getLeftTableFailure(state, error) {
        state.all.left.isLoading = false;
        state.all.error = error;
    },
};

export const dsciarc = {
    namespaced: true,
    state,
    actions,
    mutations
};