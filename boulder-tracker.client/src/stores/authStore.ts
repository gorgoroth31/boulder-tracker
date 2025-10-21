import { useLogto } from "@logto/vue";
import { defineStore } from "pinia";

export const useAuthStore = defineStore('auth',() => {
    const logto = useLogto()

    function isAuthenticated() {
        return logto.isAuthenticated.value
    }

    async function getToken() {
        return await logto.getIdToken()
    }

    return { isAuthenticated, getToken }
})