import { cryptoService } from '../_services/crypto.service';

const state = {
    all: {
        plainText: "",
        encrypted: "",
        isLoading: false,
        error: null
    }
};

const actions = {
    encrypt({ commit }) {
        commit('encryptRequest');

        var param = {
            PlainText: state.all.plainText
        }

        return cryptoService.encrypt(param)
            .then(
                res => commit('encryptSuccess', res.Data),
                error => commit('encryptFailure', error)
            );
    },
    decrypt({ commit }) {
        commit('decryptRequest');

        var param = {
            Encrypted: state.all.encrypted
        }

        return cryptoService.decrypt(param)
            .then(
                res => commit('decryptSuccess', res.Data),
                error => commit('decryptFailure', error)
            );
    }
};

const mutations = {
    encryptRequest(state) {
        state.all.isLoading = true;
    },
    encryptSuccess(state, data) {
        state.all.encrypted = data;

        state.all.isLoading = false;
    },
    encryptFailure(state, error) {
        state.all.isLoading = false;
        state.all.error = error;
    },
    decryptRequest(state) {
        state.all.isLoading = true;
    },
    decryptSuccess(state, data) {
        state.all.plainText = data;

        state.all.isLoading = false;
    },
    decryptFailure(state, error) {
        state.all.isLoading = false;
        state.all.error = error;
    },
};

export const crypto = {
    namespaced: true,
    state,
    actions,
    mutations
};