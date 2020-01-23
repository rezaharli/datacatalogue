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
        dd: {
            isNewPage: true,
            displayTable: true,
            DDSource: [],
            firstload: true,
            globalFilters: {},
            ddVal: {
                ddCountrySelected: [],
                ddBusinessSegmentSelected: [],
                ddSourceSystemSelected: [],
                ddClusterSelected: [],
                ddTierSelected: [],
                ddItamSelected: [],
            },
        },
        iarc: {
            isNewPage: true,
            displayTable: true,
            DDSource: [],
            firstload: true,
            globalFilters: {},
            ddVal: {
                ddCountrySelected: [],
                ddSourceSystemSelected: [],
                ddItamSelected: [],
            },
        },
    }
};

const getters = {
    isDdGlobalFilterEmpty: (state) => {
        var ddVal = state.all.dd.ddVal;
        return ddVal.ddCountrySelected.length == 0
            && ddVal.ddBusinessSegmentSelected.length == 0
            && ddVal.ddSourceSystemSelected.length == 0
            && ddVal.ddClusterSelected.length == 0
            && ddVal.ddTierSelected.length == 0
            && ddVal.ddItamSelected.length == 0;
    },
    isIarcGlobalFilterEmpty: (state) => {
        var ddVal = state.all.iarc.ddVal;
        return ddVal.ddCountrySelected.length == 0
            && ddVal.ddSourceSystemSelected.length == 0
            && ddVal.ddItamSelected.length == 0;
    },
}

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
    getDdDropdownOpts({ commit }, system) {
        commit('getDdDropdownOptsRequest');

        var param = {
            System: system
        }

        return edmpService.getDdDropdownOpts(param)
            .then(
                res => commit('getDdDropdownOptsSuccess', res.Data),
                error => commit('getDdDropdownOptsFailure', error)
            );
    },
    getIarcDropdownOpts({ commit }, system) {
        commit('getIarcDropdownOptsRequest');

        var param = {
            System: system
        }

        return edmpService.getIarcDropdownOpts(param)
            .then(
                res => commit('getIarcDropdownOptsSuccess', res.Data),
                error => commit('getIarcDropdownOptsFailure', error)
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
    getDdDropdownOptsRequest(state) {
        state.all.dd.firstload = true;
        state.all.isLoading = true;
    },
    getDdDropdownOptsSuccess(state, data) {
        state.all.dd.firstload = true;

        data.MappedDDSource = data.MappedDDSource.filter(v => {
            for (const key in v) {
                if(v[key] == "NA") {
                    return false;
                }
            }

            return true;
        });
        
        state.all.dd.DDSource = data.MappedDDSource;

        setTimeout(() => {
            state.all.dd.firstload = false;
            state.all.isLoading = false;
        }, 100);
    },
    getDdDropdownOptsFailure(state, error) {
        state.all.isLoading = false;
        state.all.error = error;
    },
    getIarcDropdownOptsRequest(state) {
        state.all.iarc.firstload = true;
        state.all.isLoading = true;
    },
    getIarcDropdownOptsSuccess(state, data) {
        state.all.iarc.firstload = true;
        state.all.iarc.DDSource = data.MappedDDSource;

        setTimeout(() => {
            state.all.iarc.firstload = false;
            state.all.isLoading = false;
        }, 100);
    },
    getIarcDropdownOptsFailure(state, error) {
        state.all.isLoading = false;
        state.all.error = error;
    },
};

export const edmp = {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
};