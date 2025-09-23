import instance from "./api";
import {User} from "../models/user";
import {AxiosResponse} from "axios";

export async function createUserForClaim(user: User) : Promise<AxiosResponse> {
    return await instance.post("/user/login", user).then(response => {
        return response
    });
}