import type { Session } from "./session";

export class User {
    id: string;
    userName: string;
    email: string;
    isDeleted: boolean;
    sessions: Session[]
}