import { dscMyService } from '../_services/dscmy.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        tabName: '',
        searchMain: '',
        searchDropdown: {
            SystemName: '',
            ItamID: '',
            TableName: '',
            ColumnName: '',
        },
        filters: {
            left: {},
            right: {}
        },
        left: newTableObject(),
        right: newTableObject(),
        exportDatas: [],
        leftHeaders: [
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'System Name', value: 'SYSTEM_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'ITAM ID', value: 'Custodians.ITAM_ID' },
            { align: 'left', display: false, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'Dataset Custodian', value: 'Custodians.DATASET_CUSTODIAN' },
            { align: 'left', display: false, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'Bank ID', value: 'Custodians.BANK_ID' }
        ],
        // rightHeaders: [
        //   { text: 'Table Name', align: 'left', sortable: false, value: 'TABLE_NAME', displayCount: true, width: "25%" },
        //   { text: 'Column Name', align: 'left', sortable: false, value: 'Columns.COLUMN_NAME', displayCount: true, width: "25%" },
        //   { text: 'Business Alias Name', align: 'left', sortable: false, value: 'Columns.BUSINESS_ALIAS_NAME', displayCount: false, width: "25%" },
        //   { text: 'CDE (Yes/No)', align: 'left', sortable: false, value: 'Columns.CDE_YES_NO', displayCount: true, width: "25%" }
        // ],
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
            Tabs: state.all.tabName,
            Search: state.all.searchMain,
            SearchDD: state.all.searchDropdown,
            Filters: state.all.filters.left,
            Pagination: _.cloneDeep(state.all.left.pagination)
        }

        param.Pagination.rowsPerPage = -1;

        return dscMyService.getLeftTable(param)
            .then(
                res => commit('getExportDataSuccess', res.Data),
                error => commit('getExportDataFailure', error)
            );
    },
    getLeftTable({ commit }) {
        commit('getLeftTableRequest');

        Object.keys(state.all.filters.left).map(function(key, index) {
            state.all.filters.left[key] = (typeof(state.all.filters.left[key]) == "object") ? state.all.filters.left[key] : (state.all.filters.left[key] ? state.all.filters.left[key].toString() : "");
        });

        var param = {
            Tabs: state.all.tabName,
            Search: state.all.searchMain,
            SearchDD: state.all.searchDropdown,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return dscMyService.getLeftTable(param)
            .then(
                res => commit('getLeftTableSuccess', res.Data),
                error => commit('getLeftTableFailure', error)
            );
    },
    getRightTable({ commit }, systemID) {
        commit('getRightTableRequest');

        Object.keys(state.all.filters.right).map(function(key, index) {
            state.all.filters.right[key] = state.all.filters.right[key] ? state.all.filters.right[key].toString() : "";
        });

        var param = {
            Tabs: state.all.tabName,
            SystemID: systemID,
            Search: state.all.searchMain,
            SearchDD: state.all.searchDropdown,
            Filters: state.all.filters.right,
            Pagination: state.all.right.pagination
        }

        return dscMyService.getRightTable(param)
            .then(
                res => commit('getRightTableSuccess', res.Data),
                error => commit('getRightTableFailure', error)
            );
    },
    getDetails({ commit }, param) {
        commit('getDetailsRequest');

        param.Tabs = state.all.tabName;
        
        return dscMyService.getDetails(param)
            .then(
                res => commit('getDetailsSuccess', res.Data),
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
        state.all.left.totalItems = data[0] ? data[0].RESULT_COUNT : 0;

        state.all.left.isLoading = false;
    },
    getLeftTableFailure(state, error) {
        state.all.left.isLoading = false;
        state.all.error = error;
    },
    getRightTableRequest(state) {
        state.all.right.isLoading = true;
    },
    getRightTableSuccess(state, data) {
        state.all.right.source = data;
        state.all.right.display = data;
        state.all.right.totalItems = data[0] ? data[0].RESULT_COUNT : 0;

        state.all.right.isLoading = false;
    },
    getRightTableFailure(state, error) {
        state.all.right.isLoading = false;
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

export const dscall = {
    namespaced: true,
    state,
    actions,
    mutations
};