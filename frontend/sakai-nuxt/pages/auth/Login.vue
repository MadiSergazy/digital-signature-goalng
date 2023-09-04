<script setup>
import { useLayout } from '@/layouts/composables/layout';
import { ref, computed } from 'vue';
import AppConfig from '@/layouts/AppConfig.vue';
import { useToast } from 'primevue/usetoast';
// import { states } from '../../service/composables/states';
import { useMainStore } from '../../service/composables/userstore';
// const useMainStore = nuxtApp.$mainstore;
const toast = useToast();
const nuxtApp = useNuxtApp();
const { layoutConfig } = useLayout();
const email = ref('');
const password = ref('');
const checked = ref(false);
const loading = ref(false);
const logoUrl = computed(() => {
    return `/layout/images/${layoutConfig.darkTheme.value ? 'logo-white' : 'logo-dark'}.svg`;
});

definePageMeta({
    layout: false
});
const register = async () => {
    try {
        loading.value = true;
        var response = await nuxtApp.$atamekenservice().login(email.value, password.value);
        loading.value = false;
        // states.setJWT(response['data'].value['jwt_token']);
        // states.set_email(email.value);
        useCookie();
        console.log('response:', response);
        console.log("response['data'].value['jwt_token']:", response['data'].value['jwt_token']);
        console.log('email:', email.value);
        // useMainStore().set_jwt(response['data'].value['jwt_token']);
        useMainStore().set_email(email.value);
        // localStorage.setItem('jwt_token', response['data'].value['jwt_token']);
        localStorage.setItem('email', email.value);
        // console.log("useCookie('access_token_cookie'):", useCookie('access_token_cookie'));
        useRouter().push('/');
    } catch (e) {
        console.log('ERROR:', e);
        toast.add({ severity: 'error', summary: 'Info', detail: 'Қате пароль немесе пароль', life: 3000 });
    }
};
</script>

<template>
    <Toast />
    <div class="surface-ground flex align-items-center justify-content-center min-h-screen min-w-screen overflow-hidden">
        <div class="flex flex-column align-items-center justify-content-center">
            <img :src="logoUrl" alt="Sakai logo" class="mb-5 w-6rem flex-shrink-0" />
            <div style="border-radius: 56px; padding: 0.3rem; background: linear-gradient(180deg, var(--primary-color) 10%, rgba(33, 150, 243, 0) 30%)">
                <div class="w-full surface-card py-8 px-5 sm:px-8" style="border-radius: 53px">
                    <div class="text-center mb-5">
                        <div style="display: flex; flex-direction: row; margin-left: 13%; margin-bottom: 5%">
                            <span style="font-size: 115%">Алдымен </span>
                            <a
                                target="_blank"
                                href="https://auth0.openai.com/u/signup/identifier?state=hKFo2SB0SWMtNXJRVmNnYkVlNENqcmxmdGJ5UExPTFN2bGNTM6Fur3VuaXZlcnNhbC1sb2dpbqN0aWTZIC1mc3dNV0J3T1U4dlRyUk1fOWoyQTV4SFAwVE5xLWpmo2NpZNkgVGRKSWNiZTE2V29USHROOTVueXl3aDVFNHlPbzZJdEc"
                                style="margin-left: 1%; margin-right: 1%; font-weight: bold; font-size: 115%"
                                class="text-blue-600"
                                >chatGPT</a
                            >
                            <span style="font-size: 115%"> сайтында тіркелу міндетті !!!</span>
                        </div>

                        <div class="text-900 text-3xl font-medium mb-3">Қош келдіңіз!</div>
                        <span class="text-600 font-medium">Бірінші рет кіріп тұрсаңыз келесі сұраныстарды орындаңыз</span>
                    </div>

                    <div>
                        <label for="email1" class="block text-900 text-xl font-medium mb-2">Email</label>
                        <InputText id="email1" v-model="email" type="text" placeholder="Email address" class="w-full md:w-30rem mb-5" style="padding: 1rem" />

                        <label for="password1" class="block text-900 font-medium text-xl mb-2">Password</label>
                        <Password id="password1" v-model="password" placeholder="Password" :toggleMask="true" class="w-full mb-3" inputClass="w-full" :inputStyle="{ padding: '1rem' }"></Password>

                        <div class="flex align-items-center justify-content-between mb-5 gap-5">
                            <div class="flex align-items-center">
                                <Checkbox id="rememberme1" v-model="checked" binary class="mr-2"></Checkbox>
                                <!-- <label for="rememberme1">Remember me</label> -->
                            </div>
                        </div>
                        <Button label="Тіркелу" class="w-full p-3 text-xl" :loading="loading" @click="register"></Button>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <AppConfig simple />
</template>

<style scoped>
.pi-eye {
    transform: scale(1.6);
    margin-right: 1rem;
}
.pi-eye-slash {
    transform: scale(1.6);
    margin-right: 1rem;
}
</style>
