import { headerService } from '../_services/header.service';

const state = {
    isLoading: true,
    datas: [],
    error: null,
};

const actions = {
    getOpts({ commit }, param) {
        commit('getGetOptsRequest');

        return headerService.getOpts(param)
            .then(
                res => commit('getGetOptsSuccess', res),
                error => commit('getGetOptsFailure', error)
            );
    },
    getRowCount({ commit }, param) {
        return headerService.getRowCount(param)
    },
};

const mutations = {
    getGetOptsRequest(state) {
        state.isLoading = true;
    },
    getGetOptsSuccess(state, res) {
        state.datas = res.Data;

        state.isLoading = false;
    },
    getGetOptsFailure(state, error) {
        state.isLoading = false;
        state.error = error;
    },
};

export const header = {
    namespaced: true,
    state,
    actions,
    mutations
};