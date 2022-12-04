import { defineStore } from 'pinia';

export interface AlertState {
    message: string;
    messageType: 'error' | 'info' | 'warning' | 'success' | undefined;
}
export const useAlertStore = defineStore('alert', {
    state: (): AlertState => {
        return {
            message: '',
            messageType: 'info',
        };
    },

    actions: {
        showAlert(alertState: AlertState) {
            this.$state.message = alertState.message;
            this.$state.messageType = alertState.messageType;
        },
    },
});
