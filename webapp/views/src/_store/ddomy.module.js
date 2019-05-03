import { ddoMyService } from '../_services/ddomy.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        tabName: '',
        searchMain: '',
        searchDropdown: {
            SubDataDomain: '',
            DataDomain: '',
            SubDataDomainOwner: '',
            BusinessTerm: '',
        },
        filters: {
            left: {},
            right: {}
        },
        left: newTableObject(),
        right: newTableObject(),
        exportDatas: [],
        leftHeaders: [
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'Data Domain', value: 'DATA_DOMAIN' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'Sub Domains', value: 'SUB_DOMAINS' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: true, sortable: true, text: 'Domain / Sub domain owner', value: 'SUB_DOMAIN_OWNER' },
        ],
        isRightTable: false,
        detailsSource: [],
        DDSource: [],

        DetailsBusinessMetadata: [],
        DDSourceBusinessMetadata: [],
        DetailsDownstreamUsage: [],
        DDSourceDownstreamUsage: [],
        DetailsBTResiding: [],
        DDSourceBTResiding: [],
        detailsLoading: true,
        error: null
    }
};

const actions = {
    getLeftTable({ commit }) {
        commit('getLeftTableRequest');

        var user = JSON.parse(localStorage.getItem("user"));

        Object.keys(state.all.filters.left).map(function(key, index) {
            state.all.filters.left[key] = state.all.filters.left[key] ? state.all.filters.left[key].toString() : "";
        });

        var param = {
            Tabs: state.all.tabName,
            LoggedInID: user.Username.toString(),
            Search: state.all.searchMain.toString(),
            SearchDD: state.all.searchDropdown,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return ddoMyService.getLeftTable(param)
            .then(
                res => commit('getLeftTableSuccess', res.Data),
                error => commit('getLeftTableFailure', error)
            );
    },
    getRightTable({ commit }, systemID) {
        commit('getRightTableRequest');

        Object.keys(state.all.filters.right).map(function(key, index) {
            state.all.filters.right[key] = state.all.filters.right[key] ? state.all.filters.right[key].toString() : "";
        });

        var param = {
            Tabs: state.all.tabName,
            SystemID: systemID,
            Search: state.all.searchMain,
            SearchDD: state.all.searchDropdown,
            Filters: state.all.filters.right,
            Pagination: state.all.right.pagination
        }

        return ddoMyService.getRightTable(param)
            .then(
                res => commit('getRightTableSuccess', res.Data),
                error => commit('getRightTableFailure', error)
            );
    },
    getDetails({ commit }, param) {
        commit('getDetailsRequest');

        return ddoMyService.getDetails(param)
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

        state.all.DetailsBusinessMetadata = data.DetailsBusinessMetadata;
        state.all.DDSourceBusinessMetadata = data.DDSourceBusinessMetadata;
        state.all.DetailsDownstreamUsage = data.DetailsDownstreamUsage;
        state.all.DDSourceDownstreamUsage = data.DDSourceDownstreamUsage;
        state.all.DetailsBTResiding = data.DetailsBTResiding;
        state.all.DDSourceBTResiding = data.DDSourceBTResiding;
        
        state.all.detailsLoading = false;
    },
    getDetailsFailure(state, error) {
        state.all.detailsLoading = false;
        state.all.error = error;
    },
};

export const ddomy = {
    namespaced: true,
    state,
    actions,
    mutations
};