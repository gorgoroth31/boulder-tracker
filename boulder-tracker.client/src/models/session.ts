import type { Boulder } from "./boulder";
import type { DateRange } from "./daterange";

export class Session {
    id: string;
    visitTime: DateRange;
    boulderedSolo: boolean;
    routesSolved: Boulder[];
    isDeleted: boolean;
    userId: string
}