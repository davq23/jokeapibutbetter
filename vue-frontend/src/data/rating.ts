import type { User } from "./user";

export default interface Rating {
    id: string;
    user_id: string;
    joke_id: string;
    user: User | undefined;
    stars: number;
    comment: string | undefined;
}
