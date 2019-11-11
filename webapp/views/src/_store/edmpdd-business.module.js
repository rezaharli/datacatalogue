import { header } from './header.module';
import { edmpService } from '../_services/edmp.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        filename: "edmp.sql", 
        queryname: "edmp-dd-business",
        param: {},
        tabName: '',
        filters: {
            left: {},
            right: {}
        },
        system: '',
        left: newTableObject(),
        exportDatas: [],
        selected: [],
        leftHeaders: [
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'ITAM ID', value: 'ITAM' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'EDM Source System Name', value: 'EDM_SOURCE_SYSTEM_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Database Name', value: 'DATABASE_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Table Name', value: 'TABLE_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Table Description', value: 'TABLE_DESCRIPTION' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Column Name', value: 'COLUMN_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Column Description', value: 'COLUMN_DESCRIPTION' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Term', value: 'BUSINESS_TERM' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Term Description', value: 'BUSINESS_DESCRIPTION' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Determines Client Location (Yes/No)', value: 'DETERMINES_CLIENT_LOCATION' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Determines Account / Deal Location (Yes/No)', value: 'DETERMINES_ACCOUNT' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Segment', value: 'BUSINESS_SEGMENT' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Product Category', value: 'PRODUCT_CATEGORY' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'CDE (Yes/No)', value: 'CDE' },
        ],
        isRightTable: false,
        DDSource: [],
        detailsLoading: true,
        detailsSource: [],
        error: null
    }
};

const actions = {
    exportData({ commit }) {
        commit('getExportDataRequest');

        Object.keys(state.all.filters.left).map(function(key) {
            state.all.filters.left[key] = (typeof(state.all.filters.left[key]) == "object") ? state.all.filters.left[key] : (state.all.filters.left[key] ? state.all.filters.left[key].toString() : "");
        });

        state.all.param = {
            System: state.all.system,
            Filters: state.all.filters.left,
            Pagination: _.cloneDeep(state.all.left.pagination)
        }

        state.all.param.Pagination.rowsPerPage = -1;

        return edmpService.getBusinessTable(state.all.param)
            .then(
                res => commit('getExportDataSuccess', res),
                error => commit('getExportDataFailure', error)
            );
    },
    getLeftTable({ commit }) {
        commit('getLeftTableRequest');

        Object.keys(state.all.filters.left).map(function(key) {
            state.all.filters.left[key] = (typeof(state.all.filters.left[key]) == "object") ? state.all.filters.left[key] : (state.all.filters.left[key] ? state.all.filters.left[key].toString() : "");
        });

        state.all.param = {
            Filename: state.all.filename,
            Queryname: state.all.queryname,
            System: state.all.system,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return edmpService.getBusinessTable(state.all.param)
            .then(
                res => {
                    commit('getLeftTableSuccess', res)
                    
                    header.actions.getRowCount(state.all.param).then(v => {
                        state.all.left.totalItems = v.Data;
                    });
                },
                error => commit('getLeftTableFailure', error)
            );
    },
};

const mutations = {
    getExportDataRequest(state) {
        state.all.left.isLoading = true;
    },
    getExportDataSuccess(state, res) {
        state.all.exportDatas = res.DataFlat;

        state.all.left.isLoading = false;
    },
    getExportDataFailure(state, error) {
        state.all.left.isLoading = false;
        state.all.error = error;
    },
    getLeftTableRequest(state) {
        state.all.left.isLoading = true;
    },
    getLeftTableSuccess(state, res) {
        state.all.left.source = res.DataFlat;
        state.all.left.display = res.Data;

        state.all.left.isLoading = false;
    },
    getLeftTableFailure(state, error) {
        state.all.left.isLoading = false;
        state.all.error = error;
    },
};

export const edmpddBusiness = {
    namespaced: true,
    state,
    actions,
    mutations
};