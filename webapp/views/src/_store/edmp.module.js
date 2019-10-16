import { edmpService } from '../_services/edmp.service';

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
        },
        DDSource: [],
        firstload: true,
        ddVal: {
            ddCountrySelected: [],
            ddBusinessSegmentSelected: [],
            ddSourceSystemSelected: [],
            ddClusterSelected: [],
            ddTierSelected: [],
            ddItamSelected: [],
        },
    }
};

const actions = {
    getCounts({ commit }, system) {
        commit('getCountsRequest');

        var param = {
            System: system
        }

        return edmpService.getHomepageCounts(param)
            .then(
                res => commit('getCountsSuccess', res.Data),
                error => commit('getCountsFailure', error)
            );
    },
    getDropdownOpts({ commit }, system) {
        commit('getDropdownOptsRequest');

        var param = {
            System: system
        }

        return edmpService.getDropdownOpts(param)
            .then(
                res => commit('getDropdownOptsSuccess', res.Data),
                error => commit('getDropdownOptsFailure', error)
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
    getDropdownOptsRequest(state) {
        state.all.firstload = true;
        state.all.isLoading = true;
    },
    getDropdownOptsSuccess(state, data) {
        state.all.firstload = true;
        state.all.DDSource = data.MappedDDSource;

        setTimeout(() => {
            state.all.firstload = false;
            state.all.isLoading = false;
        }, 100);
    },
    getDropdownOptsFailure(state, error) {
        state.all.isLoading = false;
        state.all.error = error;
    },
};

export const edmp = {
    namespaced: true,
    state,
    actions,
    mutations
};