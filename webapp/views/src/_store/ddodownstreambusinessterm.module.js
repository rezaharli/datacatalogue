import { ddoMyService } from '../_services/ddomy.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        filename: "ddo.sql", 
        queryname: "ddo-downstream-businessterm",
        param: {},
        tabName: '',
        filters: {
            left: {},
            right: {}
        },
        subdomain: '',
        system: '',
        left: newTableObject(),
        exportDatas: [],
        leftHeaders: [
            { align: 'left', display: true, filterable: false, exportable: false, displayCount: false, sortable: false, text: 'Details', value: 'Details' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Business Term', value: 'BT_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Business Alias Name', value: 'ALIAS_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'System Consumed from', value: 'SYSTEM_CONSUMED' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'System Table Name', value: 'TABLE_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'System Column Name', value: 'COLUMN_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Golden source (Yes / No)', value: 'GOLDEN_SOURCE' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Golden Source System name', value: 'GS_SYSTEM_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Golden Source Table Name', value: 'GS_TABLE_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Golden Source Column Name', value: 'GS_COLUMN_NAME' },
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
            Subdomain: state.all.subdomain,
            DspName: state.all.subdomain,
            System: state.all.system,
            Filters: state.all.filters.left,
            Pagination: _.cloneDeep(state.all.left.pagination)
        }
        
        state.all.param.Pagination.rowsPerPage = -1;

        return ddoMyService.getDownstreamBusinesstermTable(state.all.param)
            .then(
                res => commit('getExportDataSuccess', res.DataFlat),
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
            Subdomain: state.all.subdomain,
            DspName: state.all.subdomain,
            System: state.all.system,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return ddoMyService.getDownstreamBusinesstermTable(state.all.param)
            .then(
                res => commit('getLeftTableSuccess', res.Data),
                error => commit('getLeftTableFailure', error)
            );
    },
    getDetails({ commit }, param) {
        commit('getDetailsRequest');

        return ddoMyService.getDetails(param)
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

export const ddodownstreambusinessterm = {
    namespaced: true,
    state,
    actions,
    mutations
};