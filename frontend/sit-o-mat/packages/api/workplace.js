import { BACKEND_URL } from './config.js';
import './error.js';
import { checkResponse } from './error.js';

export function getWorkplaces() {
    return fetch(
        `${BACKEND_URL}/workplace`,
        {
            method: 'GET',
        }
    ).then(res => {
        checkResponse("Arbeitsplätze konnten nicht geladen werden", res);
        return res.json();
    });
}