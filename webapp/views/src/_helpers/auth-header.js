import router from '../router/routes';

export function authHeader(contentType) {
    let user = JSON.parse(localStorage.getItem('user'));

    var header = {
        'Content-Type': contentType ? contentType : 'application/json'
    }

    if (user && user.token) {
        header['Authorization'] = 'Bearer ' + user.token;
    }

    return header
}

export function fetchWHeader(url, body) {
    const requestOptions = {
        method: 'POST',
        headers: authHeader(),
        body: JSON.stringify(body)
    };

    return fetch(url, requestOptions).then(handleResponse);
}

export function fetchFile(url, body) {
    const requestOptions = {
        method: 'POST',
        headers: authHeader('text/csv'),
        body: JSON.stringify(body)
    };

    return fetch(url, requestOptions).then(handleResponse);
}

export function handleResponse(response) {
    return response.text().then(text => {
        const data = text && JSON.parse(text);
        
        if (!response.ok) {
            if (response.status === 401) {
                // auto logout if 401 response returned from api
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