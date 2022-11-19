export interface AuthRequest {
    user: string;
    password: string;
}

export interface AuthResponse {
    user_id: string;
    username: string;
    email: string;
    token: string;
}
