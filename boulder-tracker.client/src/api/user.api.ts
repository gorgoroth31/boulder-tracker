import instance from "./api";
import {User} from "../models/user";

export async function createUserForClaim(user: User) : Promise<void> {
    let config = {
        data: user,
    }
    return await instance.post("/user/login", config).then(response => console.log(response.data));
}