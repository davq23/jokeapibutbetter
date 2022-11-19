import type { AuthRequest } from '@/libs/auth';
import { Service } from '@/services/service';

export default class UserService extends Service {
    public whoIAm(): Promise<Response> {
        return this.sendRequest(`users/whoiam`, 'GET', new Headers(), null);
    }

    public login(body: AuthRequest): Promise<Response> {
        return this.sendRequest(
            `auth/login`,
            'POST',
            new Headers(),
            JSON.stringify(body),
        );
    }
}
