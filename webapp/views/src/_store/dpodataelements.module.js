import { dpoMyService } from '../_services/dpomy.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        filename: "dpo.sql", 
        queryname: "dpo-dataelements",
        param: {},
        tabName: '',
        filters: {
            left: {},
            right: {}
        },
        dspname: '',
        left: newTableObject(),
        exportDatas: [],
        leftHeaders: [
            { align: 'left', display: true, filterable: false, exportable: false, displayCount: false, sortable: false, text: 'Details', value: 'Details' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'System / EUC used by Process', value: 'SYSTEM_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'ITAM ID', value: 'ITAM_ID' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'CDE NAME', value: 'ALIAS_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'CDE (Y or N)', value: 'CDE' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Table Name', value: 'TABLE_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Column Name', value: 'COLUMN_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'If sourced Ultimate source system', value: 'ULT_SYSTEM_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Data SLA (Y/ N)', value: 'DATA_SLA' },
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

        state.all.param = {
            System: state.all.dspname,
            Filters: state.all.filters.left,
            Pagination: _.cloneDeep(state.all.left.pagination)
        }
        
        state.all.param.Pagination.rowsPerPage = -1;

        return dpoMyService.getDataelementsTable(state.all.param)
            .then(
                res => commit('getExportDataSuccess', res.Data),
                error => commit('getExportDataFailure', error)
            );
    },
    getLeftTable({ commit }, system) {
        commit('getLeftTableRequest');

        Object.keys(state.all.filters.left).map(function(key, index) {
            state.all.filters.left[key] = (typeof(state.all.filters.left[key]) == "object") ? state.all.filters.left[key] : (state.all.filters.left[key] ? state.all.filters.left[key].toString() : "");
        });

        state.all.param = {
            Filename: state.all.filename,
            Queryname: state.all.queryname,
            System: state.all.dspname,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return dpoMyService.getDataelementsTable(state.all.param)
            .then(
                res => commit('getLeftTableSuccess', res.Data),
                error => commit('getLeftTableFailure', error)
            );
    },
    getDetails({ commit }, param) {
        commit('getDetailsRequest');

        return dpoMyService.getDetails(param)
            .then(
                res => {
                    commit('getDetailsSuccess', res.Data)

                    return res;
                },
                error => commit('getDetailsFailure', error)
            );
    },
};

const mutations = {
    getExportDataRequest(state) {
        state.all.left.isLoading = true;
    },
    getExportDataSuccess(state, data) {
        state.all.exportDatas = data;

        state.all.left.isLoading = false;
    },
    getExportDataFailure(state, error) {
        state.all.left.isLoading = false;
        state.all.error = error;
    },
    getLeftTableRequest(state) {
        state.all.left.isLoading = true;
    },
    getLeftTableSuccess(state, data) {
        state.all.left.source = data;
        state.all.left.display = data;

        state.all.left.isLoading = false;
    },
    getLeftTableFailure(state, error) {
        state.all.left.isLoading = false;
        state.all.error = error;
    },
    getDetailsRequest(state) {
        state.all.detailsLoading = true;
    },
    getDetailsSuccess(state, data) {
        state.all.detailsSource = data.Detail;
        state.all.DDSource = data.DDSource;
        
        state.all.detailsLoading = false;
    },
    getDetailsFailure(state, error) {
        state.all.detailsLoading = false;
        state.all.error = error;
    },
};

export const dpodataelements = {
    namespaced: true,
    state,
    actions,
    mutations
};