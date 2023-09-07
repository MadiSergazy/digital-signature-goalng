import { API_URL, getHeader } from '@/config';

const LiftService = {
    async login() {
        const url = `${API_URL}/api/users/login`;
        const response = await useAsyncData('login', async () => await $fetch(url, { method: 'POST', headers: getHeader() }));
        
        return response;
    }
};

export default defineNuxtPlugin((nuxtApp) => {
    // Doing something with nuxtApp
    nuxtApp.provide('liftservice', () => LiftService);
});
