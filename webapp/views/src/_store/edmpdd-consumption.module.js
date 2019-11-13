import { header } from './header.module';
import { edmpService } from '../_services/edmp.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        filename: "edmp.sql", 
        queryname: "edmp-dd-consumption",
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
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Source System Name', value: 'EDM_SOURCE_SYSTEM_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Database Name', value: 'DATABASE_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Table Name', value: 'TABLE_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Column Name', value: 'COLUMN_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Consuming Application', value: 'CONSUMING_APPLICATION' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Consuming Application ITAM', value: 'CONSUMING_APPLICATION_ITAM' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Consuming Application Owner', value: 'CONSUMING_APPLICATION_OWNER' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Consumer Description', value: 'CONSUMER_DESCRIPTION' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Tech Contact', value: 'TECH_CONTACT' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Ownership', value: 'BUSINESS_OWNERSHIP' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Access Role', value: 'ACCESS_ROLE' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Role Description', value: 'ROLE_DESCRIPTION' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Consuming Tech Metadata', value: 'CONSUMING_TECH_METADATA' },
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

        return edmpService.getConsumptionTable(state.all.param)
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

        return edmpService.getConsumptionTable(state.all.param)
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

export const edmpddConsumption = {
    namespaced: true,
    state,
    actions,
    mutations
};