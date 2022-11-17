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

    protected sendRequest(
        path: string,
        method: string,
        headers: Headers,
        body: BodyInit | null | undefined,
    ): Promise<Response> {
        this.setAuthorization(headers);
        return fetch(`${this.baseUrl}/${path}`, {
            method,
            headers,
            body,
        });
    }
}
