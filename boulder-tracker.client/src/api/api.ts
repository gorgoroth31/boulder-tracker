import axios from "axios";
import client, {getAccessToken} from "../plugins/auth";

const url = "http://192.168.178.95:8080/api";

let instance = axios.create({
    baseURL: url,
    timeout: 1000,
})

instance.interceptors.request.use(
    async (config) => {
        const token = await getAccessToken()
        config.headers['Authorization'] = `Bearer ${token}`
        return config
    },
    (error) => {
        return Promise.reject(error)
    }
)

export default instance

export async function getHealthCheck() : Promise<string> {
    return await instance.get(url + "/health").then((response) => {
        return response.data;
    })
}

export async function existsUserByClaims() : Promise<boolean> {
    await client.checkSession();
    return await instance.get(url + "/user/exists").then(response => response.data);
}