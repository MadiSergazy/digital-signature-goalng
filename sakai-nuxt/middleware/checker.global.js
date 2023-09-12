export default defineNuxtRouteMiddleware((to, from) => {
    if (process.server) return;
    if (to.path !== '/auth/login' && useCookie('access_token')) {
        return navigateTo('/auth/login');
    }
});
