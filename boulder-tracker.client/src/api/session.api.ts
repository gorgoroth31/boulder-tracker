import {AxiosResponse} from "axios";
import instance from "./api";
import {Session} from "../models/session";

export async function getCurrentInProgressSession() : Promise<AxiosResponse<Session>> {
    return await instance.get<Session>("/session/currentLiveOrInProgress").then(response => {
        return response
    });
}