import {AxiosResponse} from "axios";
import instance from "./api";
import {Difficulty} from "../models/difficulty";

export async function getAllDifficulties() : Promise<AxiosResponse<Difficulty[]>> {
    return await instance.get<Difficulty[]>("/difficulty").then(response => {
        return response
    });
}