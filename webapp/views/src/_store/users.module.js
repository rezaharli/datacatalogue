import { userService } from '../_services/user.service';

const state = {
    all: {
        loading: true,
        items: [],
        error: null
    }
};

const actions = {
    getAll({ commit }) {
        commit('getAllRequest');

        userService.getAll()
            .then(
                res => commit('getAllSuccess', res.Data),
                error => commit('getAllFailure', error)
            );
    },
    register({ commit }, user) {
        commit('registerRequest', user);
    
        return userService.register(user)
            .then(
                user => commit('registerSuccess', user),
                error => commit('registerFailure', error)
            );
    },
    update({ commit }, user) {
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
        state.all.loading = true;
    },
    getAllSuccess(state, users) {
        state.all.loading = false;
        state.all.items = users;
    },
    getAllFailure(state, error) {
        state.all.loading = false;
        state.all.error = error;
    },
    registerRequest(state) {
        state.all.loading = true;
    },
    registerSuccess(state) {
        state.all.loading = false;
    },
    registerFailure(state, error) {
        state.all.loading = false;
        state.all.error = error;
    },
    updateRequest(state) {
        state.all.loading = true;
    },
    updateSuccess(state) {
        state.all.loading = false;
    },
    updateFailure(state, error) {
        state.all.loading = false;
        state.all.error = error;
    },
    deleteRequest(state, id) {
        // add 'deleting:true' property to user being deleted
        state.all.loading = true;
        state.all.items = state.all.items.map(user =>
            user.id === id
                ? { ...user, deleting: true }
                : user
        )
    },
    deleteSuccess(state, id) {
        // remove deleted user from state
        state.all.loading = false;
        state.all.items = state.all.items.filter(user => user.id !== id)
    },
    deleteFailure(state, { id, error }) {
        // remove 'deleting:true' property and add 'deleteError:[error]' property to user 
        state.all.loading = false;
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