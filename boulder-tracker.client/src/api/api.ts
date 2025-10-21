import { useAuthStore } from "@/stores/authStore";
import axios from "axios";
import {parse, stringify} from 'qs'

const url = import.meta.env.VITE_SERVER_URL;

let instance = axios.create({
    baseURL: url,
    timeout: 1000,
    paramsSerializer: {
        encode: parse,
        serialize: stringify,
    },
})

instance.interceptors.request.use(
    async (config) => {
        const { getToken } = useAuthStore()
        const token = await getToken()
        config.headers['Authorization'] = `Bearer ${token}`
        return config
    },
    (error) => {
        return error
    }
)

instance.interceptors.response.use((response) => {
        return response;
    }, (error) => {
        return error
    }
)

export default instance

export async function getHealthCheck(): Promise<string> {
    return await instance.get(url + "/health").then((response) => {
        return response.data;
    })
}

export async function existsUserByClaims(): Promise<boolean> {
    //await client.checkSession();
    return await instance.get(url + "/user/exists").then(response => response.data);
}