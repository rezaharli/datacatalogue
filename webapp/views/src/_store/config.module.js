import { configService } from '../_services/config.service';

const state = {
    all: {
        config: {},
        isLoading: false,
        error: null
    }
};

const actions = {
    getConfig({ commit }, param) {
        commit('getConfigRequest');

        param = { Attr: param }

        configService.getConfig(param).then(
                res => commit('getConfigSuccess', res.Data),
                error => commit('getConfigFailure', error)
            );
    }
};

const mutations = {
    getConfigRequest(state) {
        state.all.isLoading = true;
    },
    getConfigSuccess(state, data) {
        state.all.config = data;

        state.all.isLoading = false;
    },
    getConfigFailure(state, error) {
        state.all.isLoading = false;
        state.all.error = error;
    },
};

export const config = {
    namespaced: true,
    state,
    actions,
    mutations
};