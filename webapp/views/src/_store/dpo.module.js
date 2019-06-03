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
        firstload: true,
        detailsLoading: true,

        DDSourceImmediatePrecedingSystem: [],
        selectedDetailsImmediatePrecedingSystem: null,
        ddValImmediatePrecedingSystem: {},

        DDSourceUltimateSourceSystem: [],
        selectedDetailsUltimateSourceSystem: null,
        ddValUltimateSourceSystem: {},

        DDSourceDomainView: [],
        selectedDetailsDomainView: null,
        ddValDomainView: {},

        DDSourceDataStandards: [],
        selectedDetailsDataStandards: null,
        ddValDataStandards: {},

        error: null
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
        state.all.selectedDetailsImmediatePrecedingSystem = data.SelectedDetailImmediatePrecedingSystem;
        state.all.DDSourceImmediatePrecedingSystem = data.DDSourceImmediatePrecedingSystem;

        state.all.selectedDetailsUltimateSourceSystem = data.SelectedDetailUltimateSourceSystem;
        state.all.DDSourceUltimateSourceSystem = data.DDSourceUltimateSourceSystem;

        state.all.selectedDetailsDomainView = data.SelectedDetailDomainView;
        state.all.DDSourceDomainView = data.DDSourceDomainView;

        state.all.selectedDetailsDataStandards = data.SelectedDetailDataStandards;
        state.all.DDSourceDataStandards = data.DDSourceDataStandards;
        
        setTimeout(() => {
            state.all.firstload = true;
            state.all.ddValImmediatePrecedingSystem = data.DDValImmediatePrecedingSystem;
            state.all.ddValUltimateSourceSystem = data.DDValUltimateSourceSystem;
            state.all.ddValDomainView = data.DDValDomainView;
            state.all.ddValDataStandards = data.DDValDataStandards;

            setTimeout(() => {
                state.all.firstload = false;
                state.all.detailsLoading = false;
            }, 100);
        }, 100);
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