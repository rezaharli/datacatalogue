export function saveUsage() {
    const requestOptions = {
        method: 'POST',
        headers: authHeader(),
        body: JSON.stringify({ ProcessID: parseInt(processID) })
    };

    return fetch(`/dpo/getlefttable`, requestOptions).then(handleResponse);
}