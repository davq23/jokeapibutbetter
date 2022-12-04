import type Joke from '@/data/joke';
import { Service } from '@/services/service';

export class JokeService extends Service {
    public getJokeByID(jokeID: string) {
        return this.sendRequest(`jokes/${jokeID}`, 'GET', new Headers(), null);
    }

    public getJokes(
        offset: string | null | undefined,
        language: string | null | undefined,
        direction: number,
        authorID: string | null | undefined,
    ) {
        const params = new URLSearchParams();

        if (offset) {
            params.append('offset', offset);
        }
        if (language) {
            params.append('language', language);
        }
        if (authorID) {
            params.append('author_id', authorID);
        }

        params.append('direction', direction.toString());

        return this.sendRequest(
            `jokes?${params.toString()}`,
            'GET',
            new Headers(),
            null,
        );
    }

    public save(joke: Joke) {
        if (!joke.id) {
            joke.id = '';
        }

        return this.sendRequest(
            'jokes',
            'POST',
            new Headers(),
            JSON.stringify(joke),
        );
    }
}
