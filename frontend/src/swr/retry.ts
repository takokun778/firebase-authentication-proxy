export const retryable = (err: any, retryCount: any) => {
    if (!err.status) {
        return false;
    }
    if (400 <= err.status && err.status < 500) {
        return false;
    }
    if (retryCount >= 4) {
        return false;
    }
    return true;
};
