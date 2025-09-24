import instance from "./api";
import {User} from "../models/user";
import {AxiosResponse} from "axios";

export async function createUserForClaim(user: User) : Promise<AxiosResponse> {
    return await instance.post("/user/login", user).then(response => {
        return response
    });
}

export async function getCurrentLoggedInUser() : Promise<AxiosResponse> {
    return await instance.get("/user/login").then(response => {
        return response
    });
}