import { dscMyService } from '../_services/dscmy.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        tabs: "dscinterfaces",
        searchMain: '',
        searchDropdown: {
            SystemName: '',
            ItamID: '',
            TableName: '',
            ColumnName: '',
        },
        left: newTableObject(),
        right: newTableObject(),
        detailsLoading: true,
        detailsSource: [],
        error: null
    }
};

const actions = {
    getLeftTable({ commit }) {
        commit('getAllSystemRequest');

        var user = JSON.parse(localStorage.getItem("user"));

        var param = {
            Tabs: state.all.tabs,
            // LoggedInID: user.Username,
            Search: state.all.searchMain,
            SearchDD: state.all.searchDropdown,
            Pagination: state.all.left.pagination
        }

        return dscMyService.getLeftTable(param)
            .then(
                res => commit('getAllSystemSuccess', res.Data),
                error => commit('getAllSystemFailure', error)
            );
    },
    getRightTable({ commit }, systemID) {
        commit('getAllTablenameRequest');

        var param = {
            Tabs: state.all.tabs,
            SystemID: systemID,
            Search: state.all.searchMain,
            SearchDD: state.all.searchDropdown,
            Pagination: state.all.right.pagination
        }

        dscMyService.getInterfacesRightTable(param)
            .then(
                res => commit('getAllTablenameSuccess', res.Data),
                error => commit('getAllTablenameFailure', error)
            );
    },
};

const mutations = {
    getAllSystemRequest(state) {
        state.all.left.loading = true;
    },
    getAllSystemSuccess(state, data) {
        state.all.left.source = data;
        state.all.left.display = data;
        state.all.left.totalItems = data[0] ? data[0].RESULT_COUNT : 0;

        state.all.left.loading = false;
    },
    getAllSystemFailure(state, error) {
        state.all.left.loading = false;
        state.all.error = error;
    },
    getAllTablenameRequest(state) {
        state.all.right.loading = true;
    },
    getAllTablenameSuccess(state, data) {
        state.all.right.source = data;
        state.all.right.display = data;
        state.all.right.totalItems = data[0] ? data[0].RESULT_COUNT : 0;

        state.all.right.loading = false;
    },
    getAllTablenameFailure(state, error) {
        state.all.right.loading = false;
        state.all.error = error;
    },
};

export const dscinterfaces = {
    namespaced: true,
    state,
    actions,
    mutations
};