import { ATAMEKEN_API, getHeader } from '@/config';
// import { credentials } from 'amqplib';
// import { states } from '../service/composables/states';
// import { useMainStore } from '../service/composables/userstore';
// const nuxtApp = useNuxtApp();
// const useMainStore = nuxtApp.$mainstore();
const AtamekenService = {
    // const nuxtApp = useNuxtApp(),
    async askQuestion(prompt, conv_id) {
        const url = `${ATAMEKEN_API}/ask`;
        console.log('conv_id:', conv_id);

        const response = await useAsyncData(
            'question',
            async () =>
                await $fetch(url, {
                    method: 'POST',
                    body: {
                        prompt: prompt
                        // access_token: states.access_token()
                    },
                    params: {
                        conv_id: conv_id
                    },
                    headers: getHeader(),
                    credentials: 'include'
                    // headers: {
                    //     'X-CSRF-TOKEN': crsf
                    // }
                })
        );
        console.log('response:', response);
        return response;
    },
    async login(email, password) {
        const url = `${ATAMEKEN_API}/login`;
        const response = await useAsyncData(
            'question',
            async () =>
                await $fetch(url, {
                    method: 'POST',
                    body: {
                        email: email,
                        password: password
                    },
                    headers: getHeader(),
                    credentials: 'include'
                })
        );
        return response;
    },
    async getConversatinos() {
        const url = `${ATAMEKEN_API}/conversations`;
        const response = await useFetch(url, {
            method: 'GET',
            headers: getHeader(),
            credentials: 'include',
            async onResponseError({ request, response, options }) {
                // Handle the response errors
                if (response.status == 401) {
                    await AtamekenService.refreshToken();
                }
            }
        });
        return response;
    },
    async logout() {
        const url = `${ATAMEKEN_API}/logout`;
        await useFetch(url, { method: 'POST', headers: getHeader(), credentials: 'include' });
    },
    async deleteConv(conv_id) {
        const url = `${ATAMEKEN_API}/delete/conv`;
        const response = await useFetch(url, { method: 'DELETE', headers: getHeader(), body: { conv_id: conv_id }, credentials: 'include' });
    },
    async refreshToken() {
        const url = `${ATAMEKEN_API}/refresh`;
        await useFetch(url, { method: 'POST', headers: getHeader(), credentials: 'include' });
    },
    async get_messages(conv_id) {
        const url = `${ATAMEKEN_API}/get_messages`;
        return await useFetch(url, {
            headers: getHeader(),
            credentials: 'include',
            params: {
                conv_id: conv_id
            }
        });
    }
};
export default defineNuxtPlugin((nuxtApp) => {
    // Doing something with nuxtApp
    nuxtApp.provide('atamekenservice', () => AtamekenService);
});
