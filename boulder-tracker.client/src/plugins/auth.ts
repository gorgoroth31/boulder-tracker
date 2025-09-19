import type {Auth0Plugin} from '@auth0/auth0-vue'
import {createAuth0} from '@auth0/auth0-vue'
import authConfig from '../../auth.config.json'

const client: Auth0Plugin = createAuth0({
    domain: authConfig.domain,
    clientId: authConfig.clientId,
    authorizationParams: {
        redirect_uri: window.location.origin,
    }
})

export default client

export const getAccessToken = async () => {
    return await client.getAccessTokenSilently()
}