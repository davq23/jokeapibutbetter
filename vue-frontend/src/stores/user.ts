import type { User } from '@/data/user';
import type StandardResponse from '@/libs/standard';
import UserService from '@/services/user.service';
import { defineStore } from 'pinia';

interface UserState {
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
        login(user: UserState) {
            this.$state.id = user.id;
            this.$state.username = user.username;
            this.$state.email = user.email;
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

                        this.login({ id, email, username });
                    } else {
                        this.login({ id: null, email: null, username: null });
                    }
                });
        },
    },
});
