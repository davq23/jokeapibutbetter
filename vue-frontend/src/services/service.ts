import { useUserStore } from '@/stores/user';

export abstract class Service {
    protected baseUrl: string;
    protected apiKey: string | null;

    public constructor(baseUrl: string, apiKey: string | null) {
        this.apiKey = apiKey;
        this.baseUrl = baseUrl;
    }

    protected setAuthorization(headers: Headers) {
        if (this.apiKey !== null) {
            headers.append('Authorization', `Bearer ${this.apiKey}`);
        }
    }

    protected async sendRequest(
        path: string,
        method: string,
        headers: Headers,
        body: BodyInit | null | undefined,
    ): Promise<Response> {
        const user = useUserStore();

        this.setAuthorization(headers);

        const response = await fetch(`${this.baseUrl}/${path}`, {
            method,
            headers,
            body,
            credentials: 'include',
        });
        if (response.status === 403) {
            user.emptyCurrentUser();
        }
        return response;
    }
}
