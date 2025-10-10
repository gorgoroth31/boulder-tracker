import {AxiosResponse} from "axios";
import instance from "./api";
import {Style} from "node:util";

export async function getAllStyles() : Promise<AxiosResponse<Style[]>> {
    return await instance.get<Style[]>("/style").then(response => {
        return response
    });
}