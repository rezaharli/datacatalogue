import Vue from "vue";

const state = {
    loading: false,
    items: [],
    pagination: {
        descending: true,
        page: 1,
        rowsPerPage: 5,
        totalItems: 0,
    },
    filters: {},
};

const actions = {
    setData ({commit}, res) {
        console.log("--------", res);
        var items = res.Data;
        var totalItems = res.Data.length;

        commit('setItems', { items, totalItems });
    }
};

const getters = {
    allState(state){
        return state
    },
    loading (state) {
      return state.loading;
    },
    items (state) {
      return state.items;
    },
    pagination (state) {
      return state.pagination;
    },
    filters (state) {
      return state.filters;
    },
}

const mutations = {
    setItems (state, { items, totalItems }) {
        state.items = items;
        Vue.set(state.pagination, 'totalItems', totalItems);
    },
    setPagination (state, payload) {
        state.pagination = payload;
    },
    setFilters (state, payload) {
        state.filters = payload;
    },
};

export const tableModule = {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
};