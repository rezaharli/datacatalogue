import { dscMyService } from '../_services/dscmy.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        filename: "dsc.sql", 
        queryname: "dsc-view-dd",
        param: {},
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
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Table Name', value: 'TABLE_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Column Name', value: 'COLUMN_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Alias Name', value: 'BUSINESS_ALIAS_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Alias Description', value: 'BUSINESS_ALIAS_DESCRIPTION' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'CDE (Yes/No)', value: 'CDE_YES_NO' },
            // { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Status', value: 'STATUS' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Data Type', value: 'DATA_TYPE' },
            // { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Data Format', value: 'DATA_FORMAT' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Data Length', value: 'DATA_LENGTH' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Example', value: 'EXAMPLE' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Derived (Yes/No)', value: 'DERIVED_YES_NO' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Derivation logic', value: 'DERIVATION_LOGIC' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Sourced from Upstream (Yes/No)', value: 'SOURCED_FROM_UPSTREAM_YES_NO' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'System Checks', value: 'SYSTEM_CHECKS' },
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
            System: state.all.system,
            Filters: state.all.filters.left,
            Pagination: _.cloneDeep(state.all.left.pagination)
        }

        state.all.param.Pagination.rowsPerPage = -1;

        return dscMyService.getDdTable(state.all.param)
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

        state.all.param = {
            Filename: state.all.filename,
            Queryname: state.all.queryname,
            System: state.all.system,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return dscMyService.getDdTable(state.all.param)
            .then(
                res => {
                    commit('getLeftTableSuccess', res)
                },
                error => commit('getLeftTableFailure', error)
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

        state.all.left.isLoading = false;
    },
    getLeftTableFailure(state, error) {
        state.all.left.isLoading = false;
        state.all.error = error;
    },
};

export const dscddTechnical = {
    namespaced: true,
    state,
    actions,
    mutations
};