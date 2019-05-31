import { fetchWHeader } from '../_helpers/auth-header';

import moment from 'moment'

export const userService = {
    checkSession,
    login,
    logout,
    register,
    getAll,
    update,
    delete: _delete,
    saveLog,
    getUsageTable
};

function checkSession() {
    return fetchWHeader(`/users/checksession`);
}

function login(username, password) {
    return fetchWHeader(`/users/authenticate`, { username: username, password: password })
        .then(
            res => {
            
                // if (user.token) {
                //     // store user details and jwt token in local storage to keep user logged in between page refreshes
                //     localStorage.setItem('user', JSON.stringify(user));
                // }

                if (!res.Data) {
                    return Promise.reject(Error('I was created using a function call!'))
                }
                
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

function getAll(param) {
    return fetchWHeader(`/users/getall`, param).then(
        res => {
            res.Data = res.Data.map(v => {
                v.CREATEDAT = moment(v.CREATEDAT.substring(0, 19)).format('DD MMM YYYY, hh:mm a');
                v.UPDATEDAT = moment(v.UPDATEDAT.substring(0, 19)).format('DD MMM YYYY, hh:mm a');
                return v; 
            });

            return res;
        }
    );
}

function update(user) {
    return fetchWHeader(`/users/update`, user);
}

// prefixed function name with underscore because delete is a reserved word in javascript
function _delete(username) {
    return fetchWHeader(`/users/delete`, { Username: username});
}

function saveLog(param) {
    return fetchWHeader(`/users/saveusage`, param);
}

function getUsageTable(param) {
    return fetchWHeader(`/users/getusagetable`, param);
}