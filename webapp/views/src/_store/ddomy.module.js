import { ddoMyService } from '../_services/ddomy.service';

const state = {
    all: {
        systemsLoading: true,
        systemsSource: [],
        systemsDisplay: [],
        tableLoading: true,
        tableSource: [],
        tableDisplay: [],
        detailsLoading: true,
        detailsSource: [],
        error: null
    }
};

const actions = {
    getAllSystem({ commit }) {
        commit('getAllSystemRequest');

        return ddoMyService.getAllSystem()
            .then(
                res => commit('getAllSystemSuccess', res.Data),
                error => commit('getAllSystemFailure', error)
            );
    },
    getTableName({ commit }, systemID) {
        commit('getAllTablenameRequest');

        ddoMyService.getTableName(systemID)
            .then(
                res => commit('getAllTablenameSuccess', res.Data),
                error => commit('getAllTablenameFailure', error)
            );
    },
    getDetails({ commit }, param) {
        commit('getDetailsRequest');

        ddoMyService.getDetails(param)
            .then(
                res => commit('getDetailsSuccess', res.Data),
                error => commit('getDetailsFailure', error)
            );
    },
};

const mutations = {
    getAllSystemRequest(state) {
        state.all.systemsLoading = true;
    },
    getAllSystemSuccess(state, data) {
        state.all.systemsSource = data;
        state.all.systemsDisplay = data;

        state.all.systemsLoading = false;
    },
    getAllSystemFailure(state, error) {
        state.all.systemsLoading = false;
        state.all.error = error;
    },
    getAllTablenameRequest(state) {
        state.all.tableLoading = true;
    },
    getAllTablenameSuccess(state, data) {
        state.all.tableSource = data;
        state.all.tableDisplay = data;

        state.all.tableLoading = false;
    },
    getAllTablenameFailure(state, error) {
        state.all.tableLoading = false;
        state.all.error = error;
    },
    getDetailsRequest(state) {
        state.all.detailsLoading = true;
    },
    getDetailsSuccess(state, data) {
        state.all.detailsSource = data;

        state.all.tableLoading = false;
    },
    getDetailsFailure(state, error) {
        state.all.detailsLoading = false;
        state.all.error = error;
    },
};

export const ddomy = {
    namespaced: true,
    state,
    actions,
    mutations
};