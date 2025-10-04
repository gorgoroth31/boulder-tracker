import type { Session } from "./session";

export interface User {
    id: string;
    userName: string;
    email: string;
    isDeleted: boolean;
    sessions: Session[]
}