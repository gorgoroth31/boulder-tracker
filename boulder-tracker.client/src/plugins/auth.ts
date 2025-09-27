import type {Auth0Plugin} from '@auth0/auth0-vue'
import {createAuth0} from '@auth0/auth0-vue'
const auth0Scopes = "oidc profile email "

const client: Auth0Plugin = createAuth0({
    domain: import.meta.env.VITE_AUTH_DOMAIN,
    clientId: import.meta.env.VITE_AUTH_CLIENT_ID,
    authorizationParams: {
        redirect_uri: window.location.origin,
        audience: import.meta.env.VITE_AUDIENCE,
        scope: auth0Scopes + import.meta.env.VITE_SCOPES,
    },
    useRefreshTokens: true,
    useRefreshTokensFallback: true,
    cacheLocation: 'localstorage',
})

// only acceptable use to access client directly is from main.ts
// every other call should use a wrapper function
export default client

export const getAccessToken = async () => {
    await client.checkSession()
    return await client.getAccessTokenSilently()
}

export const isUserAuthenticated = async () => {
    await client.checkSession();
    return client.isAuthenticated.value;
}

export const loginWithRedirect = async() => {
    await client.loginWithRedirect()
}