export class SitomatError extends Error {
    constructor(message, status = 500) {
        super(message);
        this.status = status;
    }
}

export function checkResponse(message, res) {
    if (!res.ok) {
        throw new SitomatError(message, res.status);
    }

    return res;
}

export function asSitomatError(error) {
    if (error instanceof SitomatError) {
        return error;
    }

    return SitomatError(error.message);
}