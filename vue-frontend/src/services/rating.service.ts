import type Rating from '@/data/rating';
import { Service } from '@/services/service';

export default class RatingService extends Service {
    public rate(rating: Rating): Promise<Response> {
        return this.sendRequest(
            'ratings',
            'POST',
            new Headers(),
            JSON.stringify(rating),
        );
    }

    public getAllByUserID(userID: string): Promise<Response> {
        return this.sendRequest(
            `users/${userID}/ratings`,
            'GET',
            new Headers(),
            null,
        );
    }

    public getAllByJokeID(jokeID: string): Promise<Response> {
        return this.sendRequest(
            `jokes/${jokeID}/ratings`,
            'GET',
            new Headers(),
            null,
        );
    }
}
