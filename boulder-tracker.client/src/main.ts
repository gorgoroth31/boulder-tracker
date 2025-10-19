import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import 'vuetify/styles'

import '@mdi/font/css/materialdesignicons.css'
import vuetify from "./plugins/vuetify";
import { createLogto, LogtoConfig } from '@logto/vue'
import logtoClient from './plugins/logto'

const app = createApp(App)

app.use(router)

app.use(createLogto, logtoClient)

app.use(vuetify)

app.mount('#app')
