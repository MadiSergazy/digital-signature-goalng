<script setup>
import { ref, computed } from 'vue';
import AppConfig from '@/layouts/AppConfig.vue';
import { useMainStore } from '../../service/mainstore';
import { onMounted } from 'vue';
const loading = ref(false);
const loading2 = ref(false);
const logoUrl = computed(() => {
    return 'https://lift.kz/upload/CAllcorp3/a34/ajyqu8cvbfy1ktuc1nys5cxzsslndmtx/talgaatb2.png';
});

definePageMeta({
    layout: false
});
const router = useRouter();
const nuxtApp = useNuxtApp();
const link = ref(null);
const requirements = ref(null);
const iin = ref('');
const id = ref('');
const login = async () => {
    loading.value = true;
    var response = await nuxtApp.$liftservice().login();
    loading.value = false;
    console.log('response[data]:', response.data.value['link']);
    link.value = response.data.value['link'];
    requirements.value = response.data.value['requirements'];
    console.log('requirements:', requirements.value);
};
onMounted(() => {
    iin.value = window.localStorage.getItem('iin');
    if (!iin.value) {
        iin.value = '';
    }
});
const confirm = async () => {
    loading2.value = true;
    window.open(link.value, '_blank');
    var response = await nuxtApp.$liftservice().confirm(requirements.value);
    loading2.value = false;
    console.log('response:', response);
    const user = response.data.value['user'];
    iin.value = user['IIN'];
    console.log('user:', user);
    console.log("user['ID']:", user['ID']);
    console.log('iin:', iin.value);
    id.value = user['ID'];
    window.localStorage.setItem('iin', user['IIN']);
    const store = useMainStore();

    store.set_email(user['Email']);
    store.set_iin(user['IIN']);
    store.set_bin(user['BIN']);
    store.set_username(user['Username']);
    store.set_is_manager(false);
    // router.push('/test');
};
</script>

<template>
    <div class="surface-ground flex align-items-center justify-content-center min-h-screen min-w-screen overflow-hidden">
        <div class="flex flex-column align-items-center justify-content-center">
            <img :src="logoUrl" alt="Sakai logo" class="mb-5 w-6rem flex-shrink-0" />
            <p>iin: {{ iin.valueOf() }}</p>
            <p>id: {{ id.valueOf() }}</p>
            <div style="border-radius: 56px; padding: 0.3rem; background: linear-gradient(180deg, var(--primary-color) 10%, rgba(33, 150, 243, 0) 30%)">
                <div class="w-full surface-card py-8 px-5 sm:px-8" style="border-radius: 53px">
                    <div class="text-center mb-5">
                        <div class="text-900 text-3xl font-medium mb-3">Добро пожаловать!</div>
                        <span class="text-600 font-medium">Войдите через егов mobile что бы продолжить</span>
                    </div>
                    <div style="margin-bottom: 10px">
                        <nuxt-link v-if="link"><Button :loading="loading2" @click="confirm" label="Нажмите сюда что бы перейти в егов мобайл" class="w-full p-3 text-xl"></Button></nuxt-link>
                        <Button v-else :loading="loading" label="Войти" class="w-full p-3 text-xl" @click="login"></Button>
                    </div>
                    <div>
                        <nuxt-link v-if="link"><Button :loading="loading2" @click="confirm" label="Нажмите сюда что бы перейти в егов мобайл" class="w-full p-3 text-xl"></Button></nuxt-link>
                        <Button v-else :loading="loading" label="Войти как менеджер" class="w-full p-3 text-xl" @click="login"></Button>
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
