import { API_URL, getHeader } from '@/config';

const LiftService = {
    async login(email, password) {
        const url = `${API_URL}/users/login`;
        const response = await useAsyncData('login', async () => await $fetch(url, { method: 'POST', body: { email: email, password: password }, headers: getHeader(), credentials: 'include' }));
        return response;
    }
};

export default defineNuxtPlugin((nuxtApp) => {
    // Doing something with nuxtApp
    nuxtApp.provide('liftservice', () => LiftService);
});
