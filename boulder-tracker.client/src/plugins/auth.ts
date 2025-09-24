import type {Auth0Plugin} from '@auth0/auth0-vue'
import {createAuth0} from '@auth0/auth0-vue'
import authConfig from '../../auth.config.json'

const auth0Scopes = "oidc profile email "

const client: Auth0Plugin = createAuth0({
    domain: authConfig.domain,
    clientId: authConfig.clientId,
    authorizationParams: {
        redirect_uri: window.location.origin,
        audience: authConfig.audience,
        scope: auth0Scopes + authConfig.scopes,
    }
})

export default client

export const getAccessToken = async () => {
    await client.checkSession()
    return await client.getAccessTokenSilently()
}

export const isUserAuthenticated = async () => {
    await client.checkSession();
    return client.isAuthenticated.value;
}