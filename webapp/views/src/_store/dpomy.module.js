import { dpoMyService } from '../_services/dpomy.service';
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

        return dpoMyService.getLeftTable(param)
            .then(
                res => commit('getLeftTableSuccess', res.Data),
                error => commit('getLeftTableFailure', error)
            );
    },
    getRightTable({ commit }, ProcessID) {
        commit('getRightTableRequest');

        var param = {
            ProcessID: ProcessID,
            Search: state.all.searchMain,
            Pagination: state.all.right.pagination
        }

        dpoMyService.getRightTable(param)
            .then(
                res => commit('getRightTableSuccess', res.Data),
                error => commit('getRightTableFailure', error)
            );
    },
    getDetails({ commit }, param) {
        commit('getDetailsRequest');

        dpoMyService.getDetails(param)
            .then(
                res => commit('getDetailsSuccess', res.Data),
                error => commit('getDetailsFailure', error)
            );
    },
};

const mutations = {
    getLeftTableRequest(state) {
        state.all.left.loading = true;
    },
    getLeftTableSuccess(state, data) {
        state.all.left.source = data;
        state.all.left.display = data;
        state.all.left.totalItems = data[0] ? data[0].RESULT_COUNT : 0;

        state.all.left.loading = false;
    },
    getLeftTableFailure(state, error) {
        state.all.left.loading = false;
        state.all.error = error;
    },
    getRightTableRequest(state) {
        state.all.right.loading = true;
    },
    getRightTableSuccess(state, data) {
        state.all.right.source = data;
        state.all.right.display = data;
        state.all.right.totalItems = data[0] ? data[0].RESULT_COUNT : 0;

        state.all.right.loading = false;
    },
    getRightTableFailure(state, error) {
        state.all.right.loading = false;
        state.all.error = error;
    },
    getDetailsRequest(state) {
        state.all.detailsLoading = true;
    },
    getDetailsSuccess(state, data) {
        state.all.detailsSource = data;

        state.all.right.loading = false;
    },
    getDetailsFailure(state, error) {
        state.all.detailsLoading = false;
        state.all.error = error;
    },
};

export const dpomy = {
    namespaced: true,
    state,
    actions,
    mutations
};