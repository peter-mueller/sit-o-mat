export const LEVEL_INFO = 'INFO';
export const LEVEL_WARNING = 'WARNING';
export const LEVEL_ERROR = 'ERROR';


export function notify(message, level = LEVEL_ERROR) {
    const event = new CustomEvent('sitomat-notify', {
        detail: { message: message, level: level },
    })

    window.dispatchEvent(event);
}

export function info(message) {
    notify(message, LEVEL_INFO);
}

export function warning(message) {
    notify(message, LEVEL_WARNING);
}

export function error(message) {
    notify(message, LEVEL_ERROR);
}