import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import 'vuetify/styles'

import '@mdi/font/css/materialdesignicons.css'
import vuetify from "./plugins/vuetify";
import { createLogto, LogtoConfig } from '@logto/vue'
import { createPinia } from 'pinia'

const client: LogtoConfig = {
    endpoint: import.meta.env.VITE_LOGTO_ENDPOINT,
    appId: import.meta.env.VITE_LOGTO_APPID
}

const pinia = createPinia()

const app = createApp(App)

app.use(router)

app.use(createLogto, client)

app.use(vuetify)

app.use(pinia)

app.mount('#app')
