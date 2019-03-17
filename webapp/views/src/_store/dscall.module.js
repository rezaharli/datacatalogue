import { dscMyService } from '../_services/dscmy.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        tabs: "dscall",
        searchMain: '',
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

        var param = {
            Tabs: state.all.tabs,
            Search: state.all.searchMain,
            SearchDD: state.all.searchDropdown,
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

        var param = {
            Tabs: state.all.tabs,
            SystemID: systemID,
            Search: state.all.searchMain,
            SearchDD: state.all.searchDropdown,
            Pagination: state.all.right.pagination
        }

        dscMyService.getRightTable(param)
            .then(
                res => commit('getRightTableSuccess', res.Data),
                error => commit('getRightTableFailure', error)
            );
    },
    getDetails({ commit }, param) {
        commit('getDetailsRequest');

        param.Tabs = state.all.tabs;
        
        return dscMyService.getDetails(param)
            .then(
                res => commit('getDetailsSuccess', res.Data),
                error => commit('getDetailsFailure', error)
            );
    },
};

const mutations = {
    getLeftTableRequest(state) {
        state.all.left.loading = true;
    },
    getLeftTableSuccess(state, data) {
        state.all.left.source = data;
        state.all.left.display = data;
        state.all.left.totalItems = data[0] ? data[0].RESULT_COUNT : 0;

        state.all.left.loading = false;
    },
    getLeftTableFailure(state, error) {
        state.all.left.loading = false;
        state.all.error = error;
    },
    getRightTableRequest(state) {
        state.all.right.loading = true;
    },
    getRightTableSuccess(state, data) {
        state.all.right.source = data;
        state.all.right.display = data;
        state.all.right.totalItems = data[0] ? data[0].RESULT_COUNT : 0;

        state.all.right.loading = false;
    },
    getRightTableFailure(state, error) {
        state.all.right.loading = false;
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

export const dscall = {
    namespaced: true,
    state,
    actions,
    mutations
};