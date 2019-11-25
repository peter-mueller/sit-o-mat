export const BACKEND_URL = `https://sit-o-mat.appspot.com`;
//export const BACKEND_URL = `http://localhost:8080`;

export function authHeaders(username, password) {
    var headers = new Headers()
    headers.set('Authorization', 'Basic ' + btoa(username + ":" + password));
    return headers;
}
