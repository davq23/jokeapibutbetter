import type { User } from '@/data/user';
import type { AuthResponse } from '@/libs/auth';
import type StandardResponse from '@/libs/standard';
import UserService from '@/services/user.service';
import { defineStore } from 'pinia';

export interface UserState {
    id: string | null;
    username: string | null;
    email: string | null;
    authLoaded: boolean;
}

export const useUserStore = defineStore('user', {
    state: (): UserState => {
        return {
            id: null,
            username: null,
            email: null,
            authLoaded: false,
        };
    },

    actions: {
        setCurrentUser(
            id: string | null,
            username: string | null,
            email: string | null,
        ) {
            this.$state.id = id;
            this.$state.username = username;
            this.$state.email = email;
        },

        setAuthLoaded(authLoaded: boolean) {
            this.$state.authLoaded = authLoaded;
        },

        emptyCurrentUser() {
            this.setCurrentUser(null, null, null);
        },

        async login(user: string, password: string): Promise<StandardResponse> {
            const userService = new UserService(
                import.meta.env.VITE_JOKEAPI_URL ?? 'api',
                localStorage.getItem('token'),
            );

            return userService
                .login({ user, password })
                .then((response) => {
                    return response.json();
                })
                .then((jsonResponse: StandardResponse) => {
                    if (jsonResponse.status === 200) {
                        const { user_id, email, username, token } =
                            jsonResponse.data as AuthResponse;

                        this.setCurrentUser(user_id, username, email);

                        localStorage.setItem('token', token);
                    }

                    return new Promise<StandardResponse>((resolve) => {
                        resolve(jsonResponse);
                    });
                });
        },

        async logout() {
            this.setAuthLoaded(false);

            localStorage.removeItem('token');

            return new Promise<null>((resolve) => {
                resolve(null);
            });
            // Send request to invalidate tokens
        },

        whoIAm() {
            const userService = new UserService(
                import.meta.env.VITE_JOKEAPI_URL ?? 'api',
                localStorage.getItem('token'),
            );

            userService
                .whoIAm()
                .then((response) => {
                    return response.json();
                })
                .then((jsonResponse: StandardResponse) => {
                    if (jsonResponse.status === 200) {
                        const { id, email, username } =
                            jsonResponse.data as User;

                        this.setCurrentUser(id, username, email);
                    } else {
                        this.emptyCurrentUser();
                    }
                })
                .finally(() => {
                    this.setAuthLoaded(true);
                });
        },
    },
});
