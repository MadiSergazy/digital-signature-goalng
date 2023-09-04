import { useMainStore } from '../service/composables/userstore';
import { useConvStore } from '../service/composables/convstore';
export default defineNuxtRouteMiddleware((to, from) => {
    if (process.server) return;
    // if (localStorage.getItem('jwt_token') && localStorage.getItem('email')) {
    if (localStorage.getItem('email')) {
        // useMainStore().set_jwt(localStorage.getItem('jwt_token'));
        // console.log("localStorage.getItem('jwt_token'):", localStorage.getItem('jwt_token'));
        useMainStore().set_email(localStorage.getItem('email'));
    }
    useMainStore().set_istopbarvisible(false);
    if (!useMainStore().get_email && to.path !== '/auth/login') {
        return navigateTo('/auth/login');
    }
    if (from.path == '/') {
        useConvStore().set_conv_id(null);
    }
    // if (to.path !== '/conversation') {
    //     return navigateTo('/conversation');
    // }
});
