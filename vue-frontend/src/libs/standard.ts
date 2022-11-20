interface StandardResponse {
    status: number;
    data: unknown;
    message: string | undefined;
    token: string | undefined;
}

export default StandardResponse;
