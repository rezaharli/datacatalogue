const state = {
    type: null,
    message: null
};

const actions = {
    success({ commit }, message) { commit('success', message); },
    warning({ commit }, message) { commit('warning', message); },
    error({ commit }, message) { commit('error', message); },
    clear({ commit }, message) { commit('success', message); }
};

const mutations = {
    success(state, message) {
        state.type = 'success';
        state.message = message;
    },
    warning(state, message) {
        state.type = 'warning';
        state.message = message;
    },
    error(state, message) {
        state.type = 'danger';
        state.message = message;
    },
    clear(state) {
        state.type = null;
        state.message = null;
    }
};

export const alert = {
    namespaced: true,
    state,
    actions,
    mutations
};