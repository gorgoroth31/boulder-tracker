import type { Boulder } from "./boulder";

export interface Session {
    id: string;
    startTime: Date,
    endTime: Date,
    routesSolved: Boulder[];
    isDeleted: boolean;
    userId: string
    sessionState: SessionState;
}