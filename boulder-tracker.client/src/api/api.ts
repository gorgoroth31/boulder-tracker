import { getAccessToken } from '../plugins/auth'
import axios, {AxiosInstance} from "axios";

const url = "http://localhost:8080/api";

let instance = axios.create({
    baseURL: url,
    timeout: 1000,
})

instance.interceptors.request.use(
    async (config) => {
        const accessToken = await getAccessToken()
        config.headers['Authorization'] = `Bearer ${accessToken}`
        return config
    },
    (error) => {
        return Promise.reject(error)
    }
)

export async function getHealthCheck() : Promise<void> {
    await instance.get(url + "/health").then((response) => {console.log(response)})
}
