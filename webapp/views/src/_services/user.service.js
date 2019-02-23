import { fetchWHeader } from '../_helpers/auth-header';

export const userService = {
    login,
    logout,
    register,
    getAll,
    update,
    delete: _delete
};

function login(username, password) {
    return fetchWHeader(`/users/authenticate`, { username: username, password: password })
        .then(
            res => {
            
                // if (user.token) {
                //     // store user details and jwt token in local storage to keep user logged in between page refreshes
                //     localStorage.setItem('user', JSON.stringify(user));
                // }

                localStorage.setItem('user', JSON.stringify(res.Data));
            
                return res.Data;
            }, 
            err => Promise.reject(err)
        );
}

function logout() {
    localStorage.removeItem('user');
}

function register(user) {
    return fetchWHeader(`/users/register`, user)
        .then(
            res => res, 
            err => Promise.reject(err)
        );
}

function getAll() {
    return fetchWHeader(`/users/getall`, {});
}

function update(user) {
    return fetchWHeader(`/users/update`, {});
}

// prefixed function name with underscore because delete is a reserved word in javascript
function _delete(username) {
    return fetchWHeader(`/users/delete`, {});
}