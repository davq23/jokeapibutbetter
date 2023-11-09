import { fileURLToPath, URL } from 'node:url';

import { defineConfig } from 'vite';
import { loadEnv } from 'vite';
import vue from '@vitejs/plugin-vue';



export default ({ mode }: any) => {
    process.env = {...process.env, ...loadEnv(mode, process.cwd())};

    const FRONTEND_ASSETS_URL = process.env.VITE_FRONTEND_ASSETS_URL || '';


    // import.meta.env.VITE_NAME available here with: process.env.VITE_NAME
    // import.meta.env.VITE_PORT available here with: process.env.VITE_PORT

    // https://vitejs.dev/config/
    return defineConfig({
        plugins: [vue()],
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url)),
            },
        },
        base: FRONTEND_ASSETS_URL,
    });
}


