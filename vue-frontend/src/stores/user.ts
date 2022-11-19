import type { User } from '@/data/user';
import type { AuthResponse } from '@/libs/auth';
import type StandardResponse from '@/libs/standard';
import UserService from '@/services/user.service';
import { defineStore } from 'pinia';

export interface UserState {
    id: string | null;
    username: string | null;
    email: string | null;
}

export const useUserStore = defineStore('user', {
    state: (): UserState => {
        return {
            id: null,
            username: null,
            email: null,
        };
    },

    actions: {
        setUserState(user: UserState) {
            this.$state.id = user.id;
            this.$state.username = user.username;
            this.$state.email = user.email;
        },

        login(user: string, password: string) {
            const userService = new UserService(
                import.meta.env.VITE_JOKEAPI_URL ?? 'api',
                localStorage.getItem('token'),
            );

            userService
                .login({ user, password })
                .then((response) => {
                    return response.json();
                })
                .then((jsonResponse: StandardResponse) => {
                    if (jsonResponse.status === 200) {
                        const { user_id, email, username, token } =
                            jsonResponse.data as AuthResponse;

                        this.setUserState({ id: user_id, email, username });

                        localStorage.setItem('token', token);
                    }
                });
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
                    console.log(jsonResponse);
                    if (jsonResponse.status === 200) {
                        const { id, email, username } =
                            jsonResponse.data as User;

                        this.setUserState({ id, email, username });
                    } else {
                        this.setUserState({
                            id: null,
                            email: null,
                            username: null,
                        });
                    }
                });
        },
    },
});
