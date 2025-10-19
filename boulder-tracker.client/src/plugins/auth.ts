import { useLogto } from "@logto/vue";
import logtoClient from "./logto";

export async function isAuthenticated() : Promise<boolean> {
    //const isAuthenticated = await logtoClient.isAuthenticated()
    return true
}

export async function getToken(): Promise<string | undefined> {
    const { getIdToken } = useLogto()
    const claims = await getIdToken()
    return claims!
}