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
    roles: string[];
}

export const useUserStore = defineStore('user', {
    state: (): UserState => {
        return {
            id: null,
            username: null,
            email: null,
            authLoaded: false,
            roles: [],
        };
    },

    actions: {
        setCurrentUser(
            id: string | null,
            username: string | null,
            email: string | null,
            roles: string[],
        ) {
            this.$state.id = id;
            this.$state.username = username;
            this.$state.email = email;
            this.$state.roles = roles;
        },

        setAuthLoaded(authLoaded: boolean) {
            this.$state.authLoaded = authLoaded;
        },

        emptyCurrentUser() {
            this.setCurrentUser(null, null, null, []);
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
                        const { user_id, email, username, roles, token } =
                            jsonResponse.data as AuthResponse;

                        this.setCurrentUser(user_id, username, email, roles);

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

        async whoIAm() {
            const userService = new UserService(
                import.meta.env.VITE_JOKEAPI_URL ?? 'api',
                localStorage.getItem('token'),
            );

            try {
                const response = await userService.whoIAm();
                const jsonResponse = await response.json();
                if (jsonResponse.status === 200) {
                    const { id, email, username, roles } =
                        jsonResponse.data as User;

                    this.setCurrentUser(id, username, email, roles);

                    if (jsonResponse.token) {
                        localStorage.setItem('token', jsonResponse.token);
                    }
                } else {
                    this.emptyCurrentUser();
                }
            } finally {
                this.setAuthLoaded(true);
            }
        },
    },
});
