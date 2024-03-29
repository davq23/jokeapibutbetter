import '@mdi/font/css/materialdesignicons.css';
import '@fortawesome/fontawesome-free/css/all.css';
import 'flag-icons/css/flag-icons.min.css';
import App from './App.vue';
import router from './router';

import 'vuetify/styles';
import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';
import { aliases, fa } from 'vuetify/iconsets/fa';
import { mdi } from 'vuetify/iconsets/mdi';
import '@/assets/main.css';

const app = createApp(App);

app.use(createPinia());
app.use(router);

const vuetify = createVuetify({
    components,
    directives,
    icons: {
        defaultSet: 'mdi',
        aliases,
        sets: {
            fa,
            mdi,
        },
    },
});

app.use(vuetify);

app.mount('#app');
