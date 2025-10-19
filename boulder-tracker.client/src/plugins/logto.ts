import { LogtoConfig } from "@logto/vue";

const logtoClient: LogtoConfig = {
    endpoint: import.meta.env.VITE_LOGTO_ENDPOINT,
    appId: import.meta.env.VITE_LOGTO_APPID
}

export default logtoClient;
