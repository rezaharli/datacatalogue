import { rfoMyService } from '../_services/rfomy.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        filename: "rfo.sql", 
        queryname: "rfo-view",
        param: {},
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
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Principal Risk Types', value: 'PRINCIPAL_RISK_TYPES' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Risk Sub Type', value: 'RISK_SUB_TYPE' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Risk Framework Owner', value: 'RISK_FRAMEWORK_OWNER' },
            { align: 'left', display: false, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Risk Reporting Lead', value: 'RISK_REPORTING_LEAD' },
            { align: 'left', display: false, filterable: false, exportable: true, displayCount: false, sortable: true, text: 'Unique Count of Priority Reports', value: 'PR_COUNT' },
            { align: 'left', display: false, filterable: false, exportable: true, displayCount: false, sortable: true, text: 'Unique Count of Critical Risk Measures', value: 'CRM_COUNT' },
            { align: 'left', display: false, filterable: false, exportable: true, displayCount: false, sortable: true, text: 'Unique Count of Critical Data Elements', value: 'CDE_COUNT' },
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

        state.all.param = {
            LoggedInID: "",
            Tabs: state.all.tabName,
            Search: state.all.searchMain,
            SearchDD: state.all.searchDropdown,
            Filters: state.all.filters.left,
            Pagination: _.cloneDeep(state.all.left.pagination)
        }

        state.all.param.Pagination.rowsPerPage = -1;

        return rfoMyService.getAllRisk(state.all.param)
            .then(
                res => commit('getExportDataSuccess', res),
                error => commit('getExportDataFailure', error)
            );
    },
    getLeftTable({ commit }) {
        commit('getLeftTableRequest');

        Object.keys(state.all.filters.left).map(function(key, index) {
            state.all.filters.left[key] = (typeof(state.all.filters.left[key]) == "object") ? state.all.filters.left[key] : (state.all.filters.left[key] ? state.all.filters.left[key].toString() : "");
        });

        state.all.param = {
            LoggedInID: "",
            Tabs: state.all.tabName,
            Search: state.all.searchMain,
            SearchDD: state.all.searchDropdown,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return rfoMyService.getAllRisk(state.all.param)
            .then(
                res => commit('getLeftTableSuccess', res),
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

export const rfoall = {
    namespaced: true,
    state,
    actions,
    mutations
};