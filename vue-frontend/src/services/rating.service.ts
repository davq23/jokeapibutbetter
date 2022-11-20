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
}
