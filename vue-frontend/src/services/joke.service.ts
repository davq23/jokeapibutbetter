import type Joke from '@/data/joke';
import { Service } from '@/services/service';

export class JokeService extends Service {
    public getJokes(
        offset: string | null,
        language: string | null,
        direction: number,
    ) {
        const params = new URLSearchParams();

        if (offset) {
            params.append('offset', offset);
        }

        if (language) {
            params.append('language', language);
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
        joke.id = '';

        return this.sendRequest(
            'jokes',
            'POST',
            new Headers(),
            JSON.stringify(joke),
        );
    }
}
