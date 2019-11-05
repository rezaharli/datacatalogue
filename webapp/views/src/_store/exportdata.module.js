import { exportService } from '../_services/export.service';

const state = {
    isLoading: true,
    filename: "",
    error: null,
};

const actions = {
    doExport({ commit }, param) {
        commit('getGetOptsRequest');

        return exportService.doExport(param)
            .then(
                res => commit('getGetOptsSuccess', res),
                error => commit('getGetOptsFailure', error)
            );
    },
};

const mutations = {
    getGetOptsRequest(state) {
        state.isLoading = true;
    },
    getGetOptsSuccess(state, res) {
        state.filename = res.Data;

        state.isLoading = false;
    },
    getGetOptsFailure(state, error) {
        state.isLoading = false;
        state.error = error;
    },
};

export const exportData = {
    namespaced: true,
    state,
    actions,
    mutations
};