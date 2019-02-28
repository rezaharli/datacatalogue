import { dscMyService } from '../_services/dscmy.service';

const state = {
    all: {
        systemsLoading: true,
        systemsSource: [],
        systemsDisplay: [],
        tableLoading: true,
        tableSource: [],
        tableDisplay: [],
        error: null
    }
};

const actions = {
    getAllSystem({ commit }) {
        commit('getAllSystemRequest');

        return dscMyService.getAllSystem()
            .then(
                res => commit('getAllSystemSuccess', res.Data),
                error => commit('getAllSystemFailure', error)
            );
    },
    getTableName({ commit }, systemID) {
        commit('getAllTablenameRequest');

        dscMyService.getInterfacesRightTable(systemID)
            .then(
                res => commit('getAllTablenameSuccess', res.Data),
                error => commit('getAllTablenameFailure', error)
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
};

export const dscinterfaces = {
    namespaced: true,
    state,
    actions,
    mutations
};