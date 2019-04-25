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
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'CDE', value: 'CDE' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'Description', value: 'DESCRIPTION' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'Table Name', value: 'TABLE_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'Column Name', value: 'ColumnsVal.COLUMN_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'Downstream Process', value: 'ColumnsVal.DSP_NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'Downstream Process Owner', value: 'ColumnsVal.PROCESS_OWNER' },
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
            state.all.filters.left[key] = state.all.filters.left[key] ? state.all.filters.left[key].toString() : "";
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