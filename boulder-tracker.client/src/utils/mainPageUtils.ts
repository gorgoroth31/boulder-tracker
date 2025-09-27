import {Ref, ref} from "vue";

const pageTitle = ref("Hallo")

const appCss = ref('')
const containerCss = ref('')

const isLoggedIn: Ref<boolean> = ref(false);

export default {
    appCss: appCss,
    pageTitle: pageTitle,
    containerCss: containerCss,
    isLoggedIn: isLoggedIn,
}