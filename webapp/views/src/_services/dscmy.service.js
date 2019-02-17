import { authHeader } from '../_helpers/auth-header';
import router from '../routes';

export const dscMyService = {
    getAllSystem,
    getTableName,
};

function logout() {
    localStorage.removeItem('user');
}

function getAllSystem() {
    const requestOptions = {
        method: 'POST',
        headers: authHeader()
    };

    return fetch(`/dsc/getallsystems`, requestOptions).then(handleResponse);
}

function getTableName(system) {
    const requestOptions = {
        method: 'POST',
        headers: authHeader(),
        body: JSON.stringify({ SystemID: parseInt(system) })
    };

    return fetch(`/dsc/gettablename`, requestOptions).then(handleResponse);
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