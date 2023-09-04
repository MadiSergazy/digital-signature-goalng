import { defineStore } from 'pinia';

// main is the name of the store. It is unique across your application
// and will appear in devtools
const useMainStore = defineStore('main', {
    // a function that returns a fresh state
    state: () => ({
        jwtToken: localStorage.getItem('jwtToken'),
        email: localStorage.getItem('email')
    }),
    // optional getters
    getters: {
        // getters receive the state as first parameter
        get_email: (state) => state.email,
        // get_jwt: (state) => state.jwtToken
    },
    // optional actions
    actions: {
        // set_jwt(new_jwt) {
        //     this.jwtToken = new_jwt;
        // },
        set_email(new_email) {
            this.email = new_email;
        }
    }
});

export default defineNuxtPlugin((nuxtApp) => {
    // Doing something with nuxtApp
    nuxtApp.provide('mainstore', () => useMainStore);
});
