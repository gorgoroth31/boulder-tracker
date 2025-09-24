import {ref} from "vue";

const pageTitle = ref("Hallo")

const appCss = ref('')
const containerCss = ref('')

export default {
    appCss: appCss,
    pageTitle: pageTitle,
    containerCss: containerCss,
}