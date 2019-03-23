import { dscMyService } from '../_services/dscmy.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        tabs: "dscmy",
        searchMain: '',
        filters: {
            left: {},
            right: {}
        },
        searchDropdown: {
            SystemName: '',
            ItamID: '',
            TableName: '',
            ColumnName: '',
        },
        left: newTableObject(),
        right: newTableObject(),
        DDSource: [],
        detailsLoading: true,
        detailsSource: [],
        error: null
    }
};

const actions = {
    getLeftTable({ commit }) {
        commit('getLeftTableRequest');

        var user = JSON.parse(localStorage.getItem("user"));

        Object.keys(state.all.filters.left).map(function(key, index) {
            state.all.filters.left[key] = state.all.filters.left[key].toString();
        });

        var param = {
            Tabs: state.all.tabs,
            LoggedInID: user.Username.toString(),
            Search: state.all.searchMain.toString(),
            SearchDD: state.all.searchDropdown,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return dscMyService.getLeftTable(param)
            .then(
                res => commit('getLeftTableSuccess', res.Data),
                error => commit('getLeftTableFailure', error)
            );
    },
    getRightTable({ commit }, systemID) {
        commit('getRightTableRequest');

        Object.keys(state.all.filters.right).map(function(key, index) {
            state.all.filters.right[key] = state.all.filters.right[key].toString();
        });

        var param = {
            Tabs: state.all.tabs,
            SystemID: systemID,
            Search: state.all.searchMain,
            SearchDD: state.all.searchDropdown,
            Filters: state.all.filters.right,
            Pagination: state.all.right.pagination
        }

        return dscMyService.getRightTable(param)
            .then(
                res => commit('getRightTableSuccess', res.Data),
                error => commit('getRightTableFailure', error)
            );
    },
    getDetails({ commit }, param) {
        commit('getDetailsRequest');

        return dscMyService.getDetails(param)
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
    getLeftTableRequest(state) {
        state.all.left.isLoading = true;
    },
    getLeftTableSuccess(state, data) {
        state.all.left.source = data;
        state.all.left.display = data;
        state.all.left.totalItems = data[0] ? data[0].RESULT_COUNT : 0;

        state.all.left.isLoading = false;
    },
    getLeftTableFailure(state, error) {
        state.all.left.isLoading = false;
        state.all.error = error;
    },
    getRightTableRequest(state) {
        state.all.right.isLoading = true;
    },
    getRightTableSuccess(state, data) {
        state.all.right.source = data;
        state.all.right.display = data;
        state.all.right.totalItems = data[0] ? data[0].RESULT_COUNT : 0;

        state.all.right.isLoading = false;
    },
    getRightTableFailure(state, error) {
        state.all.right.isLoading = false;
        state.all.error = error;
    },
    getDetailsRequest(state) {
        state.all.detailsLoading = true;
    },
    getDetailsSuccess(state, data) {
        state.all.detailsSource = data.Detail;
        state.all.DDSource = data.DDSource;
        state.all.detailsLoading = false;
    },
    getDetailsFailure(state, error) {
        state.all.detailsLoading = false;
        state.all.error = error;
    },
};

export const dscmy = {
    namespaced: true,
    state,
    actions,
    mutations
};