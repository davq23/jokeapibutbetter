import type { User } from '@/data/user';
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

    public save(user: User) {
        return this.sendRequest(
            `users`,
            'POST',
            new Headers(),
            JSON.stringify(user),
        );
    }

    public getProfilePicUploadLink() {
        return this.sendRequest(
            `users/profile/upload`,
            'GET',
            new Headers(),
            null,
        );
    }

    public getProfilePicDownloadLink() {
        return this.sendRequest(
            `users/profile/download`,
            'GET',
            new Headers(),
            null,
        );
    }
}
