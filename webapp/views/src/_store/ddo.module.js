import { ddoMyService } from '../_services/ddomy.service';

const state = {
    all: {
        isLoading: true,
        error: null,
        counts: null,
        drawer: null,
        drawerContent: {
            systemName: "",
            itamID: "",
            owners: []
        }
    }
};

const actions = {
    getCounts({ commit }, system) {
        commit('getCountsRequest');

        var param = {
            System: system
        }

        return ddoMyService.getHomepageCounts(param)
            .then(
                res => commit('getCountsSuccess', res.Data),
                error => commit('getCountsFailure', error)
            );
    }
};

const mutations = {
    getCountsRequest(state) {
        state.all.isLoading = true;
    },
    getCountsSuccess(state, data) {
        state.all.counts = data;

        state.all.isLoading = false;
    },
    getCountsFailure(state, error) {
        state.all.isLoading = false;
        state.all.error = error;
    },
};

export const ddo = {
    namespaced: true,
    state,
    actions,
    mutations
};