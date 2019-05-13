import { userService } from '../_services/user.service';
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
        leftHeaders: [
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Username', value: 'USERNAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Email', value: 'EMAIL' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Name', value: 'NAME' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Role', value: 'ROLE' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Status', value: 'STATUS' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Created At', value: 'CREATEDAT' },
            { align: 'left', display: true, filterable: true, exportable: true, displayCount: false, sortable: true, text: 'Updated At', value: 'UPDATEDAT' },
            { align: 'left', display: true, filterable: false, exportable: false, displayCount: false, sortable: false, text: 'Actions', value: 'ACTIONS' },
        ],
        isRightTable: false,
        DDSource: [],
        detailsLoading: true,
        detailsSource: [],
        error: null
    }
};

const actions = {
    getAll({ commit }) {
        commit('getAllRequest');

        userService.getAll()
            .then(
                res => {
                    res.Data.map(v => {
                        v.CreatedAt = moment(v.CreatedAt.substring(0, 19)).format('DD MMM YYYY, hh:mm a');
                        v.UpdatedAt = moment(v.UpdatedAt.substring(0, 19)).format('DD MMM YYYY, hh:mm a');
                        return v; 
                    }),
                    commit('getAllSuccess', res.Data);
                }, 
                
                
                error => commit('getAllFailure', error)
            );
    },
    register({ commit }, user) {
        var tempUser = _.cloneDeep(user)
        tempUser.Username = parseInt(tempUser.Username);
        tempUser.Role = tempUser.Role.join();
        tempUser.Status = tempUser.Status == true ? 1 : 0;

        commit('registerRequest', tempUser);
    
        return userService.register(tempUser)
            .then(
                user => commit('registerSuccess', user),
                error => commit('registerFailure', error)
            );
    },
    update({ commit }, user) {
        user.Username = parseInt(user.Username);
        user.Role = user.Role.join();
        tempUser.Status = tempUser.Status == true ? 1 : 0;
        
        commit('updateRequest', user);
    
        return userService.update(user)
            .then(
                user => commit('updateSuccess', user),
                error => commit('updateFailure', error)
            );
    },
    delete({ commit }, id) {
        commit('deleteRequest', id);

        return userService.delete(id)
            .then(
                user => commit('deleteSuccess', user),
                error => commit('deleteFailure', { id, error: error.toString() })
            );
    }
};

const mutations = {
    getAllRequest(state) {
        state.all.isLoading = true;
    },
    getAllSuccess(state, users) {
        state.all.isLoading = false;
        state.all.items = users;
    },
    getAllFailure(state, error) {
        state.all.isLoading = false;
        state.all.error = error;
    },
    registerRequest(state) {
        state.all.isLoading = true;
    },
    registerSuccess(state) {
        state.all.isLoading = false;
    },
    registerFailure(state, error) {
        state.all.isLoading = false;
        state.all.error = error;
    },
    updateRequest(state) {
        state.all.isLoading = true;
    },
    updateSuccess(state) {
        state.all.isLoading = false;
    },
    updateFailure(state, error) {
        state.all.isLoading = false;
        state.all.error = error;
    },
    deleteRequest(state, id) {
        // add 'deleting:true' property to user being deleted
        state.all.isLoading = true;
        state.all.items = state.all.items.map(user =>
            user.id === id
                ? { ...user, deleting: true }
                : user
        )
    },
    deleteSuccess(state, id) {
        // remove deleted user from state
        state.all.isLoading = false;
        state.all.items = state.all.items.filter(user => user.id !== id)
    },
    deleteFailure(state, { id, error }) {
        // remove 'deleting:true' property and add 'deleteError:[error]' property to user 
        state.all.isLoading = false;
        state.all.items = state.items.map(user => {
            if (user.id === id) {
                // make copy of user without 'deleting:true' property
                const { deleting, ...userCopy } = user;
                // return copy of user with 'deleteError:[error]' property
                return { ...userCopy, deleteError: error };
            }

            return user;
        })
    }
};

export const users = {
    namespaced: true,
    state,
    actions,
    mutations
};