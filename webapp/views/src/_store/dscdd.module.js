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
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Table Name', value: 'CDE' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Column Name', value: 'DESCRIPTION' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Business Alias Name*', value: 'TABLE_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Business Alias Description', value: 'BUSINESS_ALIAS_DESCRIPTION' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'CDE (yes/no)', value: 'CDE_YES_NO' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Status*', value: 'STATUS' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Data Type', value: 'DATA_TYPE' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Data Format', value: 'DATA_FORMAT' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Data Length', value: 'DATA_LENGTH' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Example', value: 'EXAMPLE' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Derived (Yes/No)*', value: 'DERIVED_YES_NO' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Derivation logic*', value: 'DERIVATION_LOGIC' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Sourced from Upstream (Yes/No)*', value: 'SOURCED_FROM_UPSTREAM_YES_NO' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'System Checks*', value: 'SYSTEM_CHECKS' },
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
            state.all.filters.left[key] = state.all.filters.left[key] ? state.all.filters.left[key].toString() : "";
        });

        var param = {
            System: state.all.system,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return dscMyService.getCdeTable(param)
            .then(
                res => commit('getExportDataSuccess', res),
                error => commit('getExportDataFailure', error)
            );
    },
    getLeftTable({ commit }, system) {
        commit('getLeftTableRequest');

        Object.keys(state.all.filters.left).map(function(key, index) {
            state.all.filters.left[key] = state.all.filters.left[key] ? state.all.filters.left[key].toString() : "";
        });

        var param = {
            System: state.all.system,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return dscMyService.getCdeTable(param)
            .then(
                res => commit('getLeftTableSuccess', res),
                error => commit('getLeftTableFailure', error)
            );
    },
    getDetails({ commit }, param) {
        commit('getDetailsRequest');

        return dscMyService.getDetails(param)
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
    getExportDataSuccess(state, res) {
        state.all.exportDatas = res.DataFlat;

        state.all.left.isLoading = false;
    },
    getExportDataFailure(state, error) {
        state.all.left.isLoading = false;
        state.all.error = error;
    },
    getLeftTableRequest(state) {
        state.all.left.isLoading = true;
    },
    getLeftTableSuccess(state, res) {
        state.all.left.source = res.DataFlat;
        state.all.left.display = res.Data;
        state.all.left.totalItems = res.Data[0] ? res.Data[0].RESULT_COUNT : 0;

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

export const dscdd = {
    namespaced: true,
    state,
    actions,
    mutations
};