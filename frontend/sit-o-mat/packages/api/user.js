import { BACKEND_URL, authHeaders } from './config.js';
import './error.js';
import { checkResponse } from './error.js';

export function getUser(username, password) {
    return fetch(
        `${BACKEND_URL}/user/${username}`,
        {
            method: 'GET',
            headers: authHeaders(username, password),
        }
    ).then(res => {
        checkResponse("Benutzer kann nicht geladen werden", res);
        return res.json();
    });
}

export function patchWeeklyRequests(weeklyRequests) {
    return fetch(
        `${BACKEND_URL}/user/${authentication.username}`,
        {
            method: 'PATCH',
            headers: authHeaders(authentication.username, authentication.password),
            body: JSON.stringify({
                op: "replace",
                path: "/WeeklyRequests",
                value: weeklyRequests
            })
        }
    ).then(res => {
        checkResponse("Anwesenheiten konnten nicht geändert werden", res);
        return res.json();
    });
}

export class LocalAuthentication {
    get username() {
        return window.localStorage.getItem('username');
    }
    get password() {
        return window.localStorage.getItem('password');
    }

    set username(username) {
        return window.localStorage.setItem('username', username);
    }
    set password(password) {
        return window.localStorage.setItem('password', password);
    }

    clear() {
        window.localStorage.clear('username');
        window.localStorage.clear('password');
    }
}

export var authentication = new LocalAuthentication();
