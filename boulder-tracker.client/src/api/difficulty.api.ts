import {AxiosResponse} from "axios";
import instance from "./api";
import {Difficulty} from "../models/difficulty";

export async function getAllForUser() : Promise<AxiosResponse<Difficulty[]>> {
    return await instance.get<Difficulty[]>("/difficulty/forUser").then(response => {
        return response
    });
}