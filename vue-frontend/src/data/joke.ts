export default interface Joke {
    id: string;
    author_id: string;
    description: string | undefined;
    text: string;
    added_at: string;
    lang: string;
}
