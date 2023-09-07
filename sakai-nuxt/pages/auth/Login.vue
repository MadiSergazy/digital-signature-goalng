<script setup>
import { useLayout } from '@/layouts/composables/layout';
import { ref, computed } from 'vue';
import AppConfig from '@/layouts/AppConfig.vue';
const { layoutConfig } = useLayout();
const email = ref('');
const password = ref('');
const checked = ref(false);
const loading = ref(false);
const logoUrl = computed(() => {
    return 'https://lift.kz/upload/CAllcorp3/a34/ajyqu8cvbfy1ktuc1nys5cxzsslndmtx/talgaatb2.png';
});

definePageMeta({
    layout: false
});
const nuxtApp = useNuxtApp();
const link = ref(null);
const login = async () => {
    loading.value = true;
    var response = await nuxtApp.$liftservice().login();
    loading.value = false;
    console.log('response[data]:', response.data.value['link']);
    link.value = response.data.value['link'];
    return response;
};
</script>

<template>
    <div class="surface-ground flex align-items-center justify-content-center min-h-screen min-w-screen overflow-hidden">
        <div class="flex flex-column align-items-center justify-content-center">
            <img :src="logoUrl" alt="Sakai logo" class="mb-5 w-6rem flex-shrink-0" />
            <div style="border-radius: 56px; padding: 0.3rem; background: linear-gradient(180deg, var(--primary-color) 10%, rgba(33, 150, 243, 0) 30%)">
                <div class="w-full surface-card py-8 px-5 sm:px-8" style="border-radius: 53px">
                    <div class="text-center mb-5">
                        <!-- <img src="/demo/images/login/avatar.png" alt="Image" height="50" class="mb-3" /> -->
                        <div class="text-900 text-3xl font-medium mb-3">Добро пожаловать!</div>
                        <span class="text-600 font-medium">Войдите через егов mobile что бы продолжить</span>
                    </div>

                    <div>
                        <!-- <label for="email1" class="block text-900 text-xl font-medium mb-2">ЭЦП</label>
                        <InputText id="email1" v-model="email" type="text" placeholder="введите ЭЦП ключь" class="w-full mb-3" style="padding: 1rem" /> -->
                        <nuxt-link v-if="link" :to="link"><Button :loading="loading" label="Нажмите сюда что бы перейти в егов мобайл" class="w-full p-3 text-xl"></Button></nuxt-link>
                        <Button v-if="!link" :loading="loading" label="Войти" class="w-full p-3 text-xl" @click="login"></Button>
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
