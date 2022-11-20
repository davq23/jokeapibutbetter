export default interface Rating {
    id: string;
    user_id: string;
    joke_id: string;
    stars: number;
    comment: string | undefined;
}
