export enum AUTH_STATE {
    LOADING = "LOADING",
    CONNECTED = "CONNECTED", // require login
    DISCONNECTED = "DISCONNECTED", // successful login
    REQUIRE_VERIFY_EMAIL = "REQUIRE_VERIFY_EMAIL", // require email verification
}

export enum PAGE_TYPE {
    PUBLIC = "PUBLIC",
    PRIVATE = "PRIVATE",
}
