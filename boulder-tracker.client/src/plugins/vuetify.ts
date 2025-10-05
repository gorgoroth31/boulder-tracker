import { createVuetify } from 'vuetify'
import * as vuetifyComponents from 'vuetify/components'
import * as directives from 'vuetify/directives'
import * as labsComponents from 'vuetify/labs/components'

const myCustomDarkTheme = {
    dark: true,
    colors: {
        background: '#181818',
        surface: '#282828',
        primary: '#7ab41d',
        secondary: '#48A9A6',
        error: '#B00020',
        info: '#2196F3',
        success: '#4CAF50',
        warning: '#FB8C00',
    },
}

export default createVuetify({
    components: {
        ...vuetifyComponents,
        ...labsComponents,
    },
    directives,
    theme: {
        defaultTheme: 'myCustomDarkTheme',
        themes: {
            myCustomDarkTheme
        }
    },
})
