import { authHeader } from '../_helpers/auth-header';
import router from '../routes';

const config = {
    apiUrl: 'http://0.0.0.0:9001'
}

export const userService = {
    login,
    logout,
    register,
    getAll,
    getById,
    update,
    delete: _delete
};

function login(username, password) {
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username: username, password: password })
    };

    return fetch(`/users/authenticate`, requestOptions)
        .then(handleResponse)
        .then(res => {
            
            // if (user.token) {
            //     // store user details and jwt token in local storage to keep user logged in between page refreshes
            //     localStorage.setItem('user', JSON.stringify(user));
            // }

            localStorage.setItem('user', JSON.stringify(res.Data));
        
            return res.Data;
        }, err => {
            return Promise.reject(err);
        });
}

function logout() {
    localStorage.removeItem('user');
}

function register(user) {
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(user)
    };

    return fetch(`/users/register`, requestOptions).then(handleResponse).then(res => {
        return res;
    }, err => {
        return Promise.reject(err);
    });;
}

function getAll() {
    const requestOptions = {
        method: 'POST',
        headers: authHeader()
    };

    return fetch(`/users/getall`, requestOptions).then(handleResponse);
}


function getById(id) {
    const requestOptions = {
        method: 'GET',
        headers: authHeader()
    };

    return fetch(`${config.apiUrl}/users/${id}`, requestOptions).then(handleResponse);
}

function update(user) {
    const requestOptions = {
        method: 'PUT',
        headers: { ...authHeader(), 'Content-Type': 'application/json' },
        body: JSON.stringify(user)
    };

    return fetch(`/users/update`, requestOptions).then(handleResponse);
}

// prefixed function name with underscore because delete is a reserved word in javascript
function _delete(username) {
    const requestOptions = {
        method: 'POST',
        headers: authHeader(),
        body: JSON.stringify({ Username: username })
    };

    return fetch(`/users/delete`, requestOptions).then(handleResponse);
}

function handleResponse(response) {
    return response.text().then(text => {
        const data = text && JSON.parse(text);
        
        if (!response.ok) {
            if (response.status === 401) {
                // auto logout if 401 response returned from api
                logout();
                router.push("/");
            }
            
            const error = (data && data.Message) || response.statusText;
            return Promise.reject(error);
        } else {
            if (data.Status == "NOK"){
                const error = (data && data.Message) || response.statusText;
                return Promise.reject(error);
            }
        }

        return data;
    });
}