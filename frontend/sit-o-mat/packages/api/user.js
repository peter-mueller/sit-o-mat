import { BACKEND_URL } from './config.js';

export function getUser(username, password) {
    return fetch(`${BACKEND_URL}/user`,
        {
            method: 'GET',
        }
    ).catch()
}