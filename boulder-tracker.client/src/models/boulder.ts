import type { Difficulty } from "./difficulty";
import type { Style } from "./style";

export class Boulder {
    id: string;
    screwedDifficulty: Difficulty;
    feltLikeDifficulty: Difficulty;
    attempts: number;
    sessionsTried: number;
    exhausting: boolean;
    style: Style[];
    like: boolean;
    sessionId: string
}
