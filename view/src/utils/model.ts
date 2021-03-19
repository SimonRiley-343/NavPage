export enum reqCode {
    success = 0,
    paramErr = 101,
    loginFailed = 102,
    noPage = 201,
    unknownErr = 202
}

export enum reqMsg {
    success = 'Success',
    paramErr = 'Param error',
    loginFailed = 'Failed login',
    noPage = 'No page',
    unknownErr = 'Unknown error'
}

export enum cookies {
    navpage = 'navpageSession',
    sessionId = 'sessionId'
}
