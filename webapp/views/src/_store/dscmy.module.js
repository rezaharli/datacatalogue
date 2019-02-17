import { dscMyService } from '../_services/dscmy.service';

const state = {
    all: {
        loading: true,
        systems: [],
        table: [],
        error: null
    }
};

const actions = {
    getAllSystem({ commit }) {
        commit('getAllSystemRequest');

        dscMyService.getAllSystem()
            .then(
                res => commit('getAllSystemSuccess', res.Data),
                error => commit('getAllSystemFailure', error)
            );
    },
    getTableName({ commit }, systemID) {
        commit('getAllTablenameRequest');

        dscMyService.getTableName(systemID)
            .then(
                res => commit('getAllTablenameSuccess', res.Data),
                error => commit('getAllTablenameFailure', error)
            );
    },
};

const mutations = {
    getAllSystemRequest(state) {
        state.all.loading = true;
    },
    getAllSystemSuccess(state, data) {
        state.all.loading = false;
        state.all.systems = data;
    },
    getAllSystemFailure(state, error) {
        state.all.loading = false;
        state.all.error = error;
    },
    getAllTablenameRequest(state) {
        state.all.loading = true;
    },
    getAllTablenameSuccess(state, data) {
        state.all.loading = false;
        state.all.table = data;
    },
    getAllTablenameFailure(state, error) {
        state.all.loading = false;
        state.all.error = error;
    },
};

export const dscmy = {
    namespaced: true,
    state,
    actions,
    mutations
};