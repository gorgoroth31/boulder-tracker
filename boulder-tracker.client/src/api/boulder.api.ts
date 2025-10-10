import {AxiosResponse} from "axios";
import instance from "./api";

export async function deleteBoulder(boulderId: string) : Promise<AxiosResponse> {
    return await instance.delete("/boulder/" + boulderId).then(response => {
        return response
    });
}