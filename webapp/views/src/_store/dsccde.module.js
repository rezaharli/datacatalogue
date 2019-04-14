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
        leftHeaders: [
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'CDE', value: 'CDE' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Description', value: 'DESCRIPTION' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Table Name', value: 'TABLE_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Column Name', value: 'ColumnsVal.COLUMN_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Downstream Process', value: 'ColumnsVal.DSP_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: true, sortable: false, text: 'Downstream Process Owner', value: 'ColumnsVal.PROCESS_OWNER' },
        ],
        isRightTable: false,
        DDSource: [],
        detailsLoading: true,
        detailsSource: [],
        error: null
    }
};

const actions = {
    getLeftTable({ commit }, system) {
        commit('getLeftTableRequest');

        Object.keys(state.all.filters.left).map(function(key, index) {
            state.all.filters.left[key] = state.all.filters.left[key].toString();
        });

        var param = {
            System: state.all.system,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return dscMyService.getCdeTable(param)
            .then(
                res => commit('getLeftTableSuccess', res.Data),
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

export const dsccde = {
    namespaced: true,
    state,
    actions,
    mutations
};