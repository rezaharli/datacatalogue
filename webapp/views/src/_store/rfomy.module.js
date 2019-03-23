import { ddoMyService } from '../_services/rfomy.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        searchMain: '',
        left: newTableObject(),
        right: newTableObject(),
        detailsLoading: true,
        detailsSource: [],
        error: null
    }
};

const actions = {
    getLeftTable({ commit }) {
        commit('getLeftTableRequest');

        var param = {
            Search: state.all.searchMain,
            Pagination: state.all.left.pagination
        }

        return ddoMyService.getLeftTable(param)
            .then(
                res => commit('getLeftTableSuccess', res.Data),
                error => commit('getLeftTableFailure', error)
            );
    },
    getRightTable({ commit }, systemID) {
        commit('getRightTableRequest');

        var param = {
            SystemID: systemID,
            Search: state.all.searchMain,
            Pagination: state.all.right.pagination
        }

        ddoMyService.getRightTable(param)
            .then(
                res => commit('getRightTableSuccess', res.Data),
                error => commit('getRightTableFailure', error)
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
};

export const rfomy = {
    namespaced: true,
    state,
    actions,
    mutations
};