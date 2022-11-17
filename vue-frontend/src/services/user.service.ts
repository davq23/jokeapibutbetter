import { Service } from '@/services/service';

export default class UserService extends Service {
    public whoIAm(): Promise<Response> {
        return this.sendRequest(`users/whoiam`, 'GET', new Headers(), null);
    }
}
