import { defineStore } from 'pinia';

// main is the name of the store. It is unique across your application
// and will appear in devtools
export const useMainStore = defineStore('main', {
    // a function that returns a fresh state
    state: () => ({
        email: null,
        istopbarvisible: true
    }),
    // optional getters
    getters: {
        // getters receive the state as first parameter
        get_email: (state) => state.email,
        // get_jwt: (state) => state.jwtToken,
        get_istopbarvisible: (state) => state.istopbarvisible
    },
    // optional actions
    actions: {
        // set_jwt(new_jwt: string) {
        //     this.jwtToken = new_jwt;
        // },
        set_email(new_email: string) {
            this.email = new_email;
        },
        set_istopbarvisible(new_state: boolean) {
            this.istopbarvisible = new_state;
        }
    }
});
