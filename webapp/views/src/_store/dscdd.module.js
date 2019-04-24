import { dscMyService } from '../_services/dscmy.service';
import { newTableObject } from '../_helpers/table-helper';

const state = {
    all: {
        tabName: '',
        filters: {
            left: {},
            right: {}
        },
        system: '',
        left: newTableObject(),
        exportDatas: [],
        leftHeaders: {
            technical : [
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Sytem Name', value: 'SYSTEM_NAME' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'ITAM ID', value: 'ITAM_ID' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Table Name', value: 'TABLE_NAME' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Column Name', value: 'COLUMN_NAME' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Alias Name', value: 'BUSINESS_ALIAS_NAME' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Alias Description', value: 'BUSINESS_ALIAS_DESCRIPTION' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'CDE (yes/no)', value: 'CDE_YES_NO' },
                // { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Status', value: 'STATUS' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Data Type', value: 'DATA_TYPE' },
                // { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Data Format', value: 'DATA_FORMAT' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Data Length', value: 'DATA_LENGTH' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Example', value: 'EXAMPLE' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Derived (Yes/No)', value: 'DERIVED_YES_NO' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Derivation logic', value: 'DERIVATION_LOGIC' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Sourced from Upstream (Yes/No)', value: 'SOURCED_FROM_UPSTREAM_YES_NO' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'System Checks', value: 'SYSTEM_CHECKS' },
            ],
            business : [
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Sytem Name', value: 'SYSTEM_NAME' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'ITAM ID', value: 'ITAM_ID' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Table Name', value: 'TABLE_NAME' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Column Name', value: 'COLUMN_NAME' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Alias Name', value: 'BUSINESS_ALIAS_NAME' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Alias Description', value: 'BUSINESS_ALIAS_DESCRIPTION' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'CDE (yes/no)', value: 'CDE_YES_NO' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Domain', value: 'DOMAIN' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Sub Domain', value: 'SUBDOMAIN' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Domain Owner', value: 'DOMAIN_OWNER' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Term', value: 'BUSINESS_TERM' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Term Description', value: 'BUSINESS_TERM_DESCRIPTION' },
            ],
            policy : [
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Sytem Name', value: 'SYSTEM_NAME' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'ITAM ID', value: 'ITAM_ID' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Table Name', value: 'TABLE_NAME' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Column Name', value: 'COLUMN_NAME' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Alias Name', value: 'BUSINESS_ALIAS_NAME' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Alias Description', value: 'BUSINESS_ALIAS_DESCRIPTION' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'CDE (yes/no)', value: 'CDE_YES_NO' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Information Asset Names', value: 'INFORMATION_ASSET_NAMES' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Information Asset Description', value: 'INFORMATION_ASSET_DESCRIPTION' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'C - Confidentiality', value: 'CONFIDENTIALITY' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'I - Integrity', value: 'INTEGRITY' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'A - Availability', value: 'AVAILABILITY' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Overall CIA Rating', value: 'OVERALL_CIA_RATING' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Record Categories', value: 'RECORD_CATEGORIES' },
                { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'PII Flag', value: 'PII_FLAG' },
            ],
            // interfaces : [
            //     { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Immediate Preceding System', value: 'IMM_PRECEEDING_SYSTEM' },
            //     { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Incoming Data Element', value: 'IMM_PREC_INCOMING' },
            //     { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Derived (Yes/No)', value: 'IMM_PREC_DERIVED' },
            //     { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Derivation Logic', value: 'IMM_PREC_DERIVATION_LOGIC' },
            //     { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Immediate Succeeding System', value: 'IMM_SUCCEEDING_SYSTEM' },
            //     { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Incoming Data Element', value: 'IMM_SUCC_INCOMING' },
            //     { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Derived (Yes/No)', value: 'IMM_SUCC_DERIVED' },
            //     { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Derivation Logic', value: 'IMM_SUCC_DERIVATION_LOGIC' },
            //     { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'DQ Standards | Threshold', value: 'THRESHOLD' },
            // ],
        },
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

        Object.keys(state.all.filters.left).map(function(key, index) {
            state.all.filters.left[key] = state.all.filters.left[key] ? state.all.filters.left[key].toString() : "";
        });

        var param = {
            System: state.all.system,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return dscMyService.getDdTable(param)
            .then(
                res => commit('getExportDataSuccess', res),
                error => commit('getExportDataFailure', error)
            );
    },
    getLeftTable({ commit }, system) {
        commit('getLeftTableRequest');

        Object.keys(state.all.filters.left).map(function(key, index) {
            state.all.filters.left[key] = state.all.filters.left[key] ? state.all.filters.left[key].toString() : "";
        });

        var param = {
            System: state.all.system,
            Filters: state.all.filters.left,
            Pagination: state.all.left.pagination
        }

        return dscMyService.getDdTable(param)
            .then(
                res => commit('getLeftTableSuccess', res),
                error => commit('getLeftTableFailure', error)
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
        state.all.left.totalItems = res.Data[0] ? res.Data[0].RESULT_COUNT : 0;

        state.all.left.isLoading = false;
    },
    getLeftTableFailure(state, error) {
        state.all.left.isLoading = false;
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

export const dscdd = {
    namespaced: true,
    state,
    actions,
    mutations
};