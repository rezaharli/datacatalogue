import { userService } from '../_services/user.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        filename: "users.sql", 
        queryname: "users-usage",
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
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Username', value: 'USERNAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Fullname', value: 'FULLNAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Role', value: 'ROLE' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Module', value: 'MODULE' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Description', value: 'DESCRIPTION' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Date', value: 'THEDATE' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Time', value: 'TIME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Resource URL', value: 'RESOURCEURL' },
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

        return userService.getUsageTable(state.all.param)
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
            System: state.all.system,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return userService.getUsageTable(state.all.param)
            .then(
                res => commit('getLeftTableSuccess', res.Data),
                error => commit('getLeftTableFailure', error)
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
};

export const usersusage = {
    namespaced: true,
    state,
    actions,
    mutations
};