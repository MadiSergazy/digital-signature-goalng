import { useMainStore } from '../service/composables/userstore';

export default defineNuxtRouteMiddleware((to) => {
    // skip middleware on server
    if (process.client) {
        useMainStore().set_email(localStorage.getItem('email'));
        // useMainStore().set_jwt(localStorage.getItem('jwt_token'));
        useMainStore().set_istopbarvisible(true);
        return;
    }
    // skip middleware on client side entirely
    // if (process.client) return
    // or only skip middleware on initial client load
});
