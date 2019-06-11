import { rfoMyService } from '../_services/rfomy.service';
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
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Priority Reports', value: 'PRIORITY_REPORT' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Priority Report Rationale', value: 'PRIORITY_REPORT_RATIONALE' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'CRM Name', value: 'CRM_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'CRM Rationale', value: 'CRM_RATIONALE' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'CDE Name', value: 'CDE_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'CDE Rationale', value: 'CDE_RATIONALE' },
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

        return rfoMyService.getPriorityTable(param)
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

        return rfoMyService.getPriorityTable(param)
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
        state.all.left.totalItems = res.Data[0] ? res.Data[0].COUNT_CDE : 0;
        
        state.all.left.isLoading = false;
    },
    getLeftTableFailure(state, error) {
        state.all.left.isLoading = false;
        state.all.error = error;
    },
};

export const rfopriority = {
    namespaced: true,
    state,
    actions,
    mutations
};