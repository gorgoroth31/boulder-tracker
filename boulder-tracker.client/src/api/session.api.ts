import {AxiosResponse} from "axios";
import instance from "./api";
import {Session} from "../models/session";

export async function getCurrentInProgressSession() : Promise<AxiosResponse<Session>> {
    return await instance.get<Session>("/session/currentLiveOrInProgress").then(response => {
        return response
    });
}

export async function updateSession(session : Session) : Promise<AxiosResponse<Session>> {
    return await instance.put("/session", session).then(response => {
        return response
    })
}

export async function submitCurrentSession() : Promise<AxiosResponse> {
    return await instance.put("/session/submitCurrent").then(response => {
        return response
    })
}