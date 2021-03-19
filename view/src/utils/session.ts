export function removeSession(context: any, sessionKey: string) {
    if (context.$cookies.isKey(sessionKey)) {
        context.$cookies.remove(sessionKey);
    }
}

export function isSessionExist(context: any, sessionKey: string) {
    return context.$cookies.isKey(sessionKey);
}
