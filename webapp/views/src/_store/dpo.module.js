import { dpoMyService } from '../_services/dpomy.service';

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
        DetailsImmediatePrecedingSystem: [],
        DDSourceImmediatePrecedingSystem: [],
        DetailsUltimateSourceSystem: [],
        DDSourceUltimateSourceSystem: [],
        DetailsDomainView: [],
        DDSourceDomainView: [],
        DetailsDataStandards: [],
        DDSourceDataStandards: [],
    }
};

const actions = {
    getCounts({ commit }, system) {
        commit('getCountsRequest');

        var param = {
            System: system
        }

        return dpoMyService.getHomepageCounts(param)
            .then(
                res => commit('getCountsSuccess', res.Data),
                error => commit('getCountsFailure', error)
            );
    },
    getDetails({ commit }, param) {
        commit('getDetailsRequest');

        return dpoMyService.getDetails(param)
            .then(
                res => {
                    commit('getDetailsSuccess', res.Data)

                    return res;
                },
                error => commit('getDetailsFailure', error)
            );
    },
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
    getDetailsRequest(state) {
        state.all.detailsLoading = true;
    },
    getDetailsSuccess(state, data) {
        state.all.DetailsImmediatePrecedingSystem = data.DetailsImmediatePrecedingSystem;
        state.all.DDSourceImmediatePrecedingSystem = data.DDSourceImmediatePrecedingSystem;
        state.all.DetailsUltimateSourceSystem = data.DetailsUltimateSourceSystem;
        state.all.DDSourceUltimateSourceSystem = data.DDSourceUltimateSourceSystem;
        state.all.DetailsDomainView = data.DetailsDomainView;
        state.all.DDSourceDomainView = data.DDSourceDomainView;
        state.all.DetailsDataStandards = data.DetailsDataStandards;
        state.all.DDSourceDataStandards = data.DDSourceDataStandards;
        
        state.all.detailsLoading = false;
    },
    getDetailsFailure(state, error) {
        state.all.detailsLoading = false;
        state.all.error = error;
    },
};

export const dpo = {
    namespaced: true,
    state,
    actions,
    mutations
};