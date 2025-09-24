import {ref} from "vue";

const pageTitle = ref("Hallo")

const wrapperClass = ref('')

export default {
    css: wrapperClass,
    pageTitle: pageTitle
}