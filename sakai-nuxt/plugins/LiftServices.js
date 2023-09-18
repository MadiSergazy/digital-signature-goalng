import { API_URL, getHeader } from '@/config';

const LiftService = {
    async login() {
        const url = `${API_URL}/api/users/login`;
        const response = await useAsyncData('login', async () => await $fetch(url, { method: 'POST', headers: getHeader() }));

        return response;
    },
    async confirm(requirements) {
        const url = `${API_URL}/api/users/confirm`;
        const response = await useAsyncData('confirm', async () => await $fetch(url, { method: 'POST', headers: getHeader(), body: requirements }));
        return response;
    },
    async get_survey(iin) {
        const url = `${API_URL}/api/survey/get/` + iin;
        const response = await useAsyncData('getsurvey', async () => await $fetch(url, { method: 'GET', headers: getHeader() }));
        return response;
    },
    async post_survey(requirements) {
        const url = `${API_URL}/api/survey/create`;
        const response = await useAsyncData('createproduct', async () => await $fetch(url, { method: 'POST', headers: getHeader(), body: requirements }));
        return response;
    }
};
export default defineNuxtPlugin((nuxtApp) => {
    // Doing something with nuxtApp
    nuxtApp.provide('liftservice', () => LiftService);
});
