import { header } from './header.module';
import { dscMyService } from '../_services/dscmy.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        filename: "dsc.sql", 
        queryname: "dsc-view",
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
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'System Name', value: 'SYSTEM_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'ITAM ID', value: 'ITAM_ID' },
            { align: 'left', display: false, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'Dataset Custodian', value: 'Custodians.DATASET_CUSTODIAN' },
            { align: 'left', display: false, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'Bank ID', value: 'Custodians.BANK_ID' }
        ],
        isRightTable: false,
        firstload: true,
        detailsLoading: true,
        DDSource: [],
        selectedDetails: null,
        ddVal: {},
        DDSourceLeftPanel: [],
        selectedDetailsLeftPanel: null,
        ddValLeftPanel: {},
        error: null,
    }
};

const actions = {
    exportData({ commit }) {
        commit('getExportDataRequest');

        var user = JSON.parse(localStorage.getItem("user"));

        Object.keys(state.all.filters.left).map(function(key, index) {
            state.all.filters.left[key] = (typeof(state.all.filters.left[key]) == "object") ? state.all.filters.left[key] : (state.all.filters.left[key] ? state.all.filters.left[key].toString() : "");
        });

        state.all.param = {
            Tabs: state.all.tabName,
            LoggedInID: user.USERNAME.toString(),
            Search: state.all.searchMain.toString(),
            SearchDD: state.all.searchDropdown,
            Filters: state.all.filters.left,
            Pagination: _.cloneDeep(state.all.left.pagination)
        }

        state.all.param.Pagination.rowsPerPage = -1;

        return dscMyService.getLeftTable(state.all.param)
            .then(
                res => commit('getExportDataSuccess', res.Data),
                error => commit('getExportDataFailure', error)
            );
    },
    getLeftTable({ commit }) {
        commit('getLeftTableRequest');

        var user = JSON.parse(localStorage.getItem("user"));

        Object.keys(state.all.filters.left).map(function(key, index) {
            state.all.filters.left[key] = state.all.filters.left[key].toString();
        });

        state.all.param = {
            Filename: state.all.filename,
            Queryname: state.all.queryname,
            Tabs: state.all.tabName,
            LoggedInID: user.USERNAME.toString(),
            Search: state.all.searchMain.toString(),
            SearchDD: state.all.searchDropdown,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return dscMyService.getLeftTable(state.all.param)
            .then(
                res => {
                    commit('getLeftTableSuccess', res.Data)
                    
                    header.actions.getRowCount(state.all.param).then(v => {
                        state.all.left.totalItems = v.Data;
                    });
                },
                error => commit('getLeftTableFailure', error)
            );
    },
    getRightTable({ commit }, systemID) {
        commit('getRightTableRequest');

        Object.keys(state.all.filters.right).map(function(key, index) {
            state.all.filters.right[key] = state.all.filters.right[key].toString();
        });

        state.all.param = {
            Tabs: state.all.tabName,
            SystemID: systemID,
            Search: state.all.searchMain,
            SearchDD: state.all.searchDropdown,
            Filters: state.all.filters.right,
            Pagination: state.all.right.pagination
        }

        return dscMyService.getRightTable(state.all.param)
            .then(
                res => commit('getRightTableSuccess', res.Data),
                error => commit('getRightTableFailure', error)
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
        state.all.firstload = true;
        state.all.detailsLoading = true;
    },
    getDetailsSuccess(state, data) {
        state.all.selectedDetails = data.SelectedDetail;
        state.all.DDSource = data.DDSource;
        state.all.selectedDetailsLeftPanel = data.SelectedDetailLeftPanel;
        state.all.DDSourceLeftPanel = data.DDSourceLeftPanel;
        
        setTimeout(() => {
            state.all.firstload = true;
            state.all.ddVal = data.DDVal;
            state.all.ddValLeftPanel = data.DDValLeftPanel;

            setTimeout(() => {
                state.all.firstload = false;
                state.all.detailsLoading = false;
            }, 100);
        }, 100);
    },
    getDetailsFailure(state, error) {
        state.all.detailsLoading = false;
        state.all.error = error;
    },
};

export const dscmy = {
    namespaced: true,
    state,
    actions,
    mutations
};