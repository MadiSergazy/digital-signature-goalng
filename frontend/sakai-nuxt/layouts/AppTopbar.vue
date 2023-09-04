<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue';
import { useLayout } from './composables/layout';
import { useRouter } from 'vue-router';
import { useMainStore } from '../service/composables/userstore';
import { useConfirm } from 'primevue/useconfirm';
import { useToast } from 'primevue/usetoast';

const { layoutConfig, onMenuToggle } = useLayout();
const outsideClickListener = ref(null);
const topbarMenuActive = ref(false);
const router = useRouter();
const confirm = useConfirm();
const toast = useToast();
const showConfirmDialog = () => {
    confirm.require({
        target: event.currentTarget,
        message: 'Аккаунтыңыздан шығуға сенімдісіз бе?',
        icon: 'pi pi-exclamation-triangle',
        accept: () => {
            localStorage.clear();
            try {
                nuxtApp.$atamekenservice().logout();
                useMainStore().set_email(null);
                toast.add({ severity: 'info', summary: 'Қабылданды', detail: 'Сіз аккаунтыңыздан шықтыңыз', life: 3000 });
            } catch {
                toast.add({ severity: 'error', summary: 'Ақау', detail: 'Бұл ақау жөнделуде', life: 3000 });
            }
        },
        acceptLabel: 'Иә',
        rejectLabel: 'Жоқ'
    });
};
const nuxtApp = useNuxtApp();
const items = [
    {
        label: 'Home',
        icon: 'pi pi-fw pi-home',
        to: '/'
    },
    {
        label: 'Documentation',
        icon: 'pi pi-fw pi-file',
        to: '/documentation'
    },
    {
        label: 'Settings',
        icon: 'pi pi-fw pi-cog',
        to: '/settings'
    }
];
const width = ref(0);

onMounted(() => {
    width.value = window.innerWidth;
    window.addEventListener('resize', resizer);
    bindOutsideClickListener();
});
onBeforeUnmount(() => {
    unbindOutsideClickListener();
});
const resizer = () => {
    width.value = window.innerWidth;
};
const logoUrl = computed(() => {
    return `/layout/images/${layoutConfig.darkTheme.value ? 'logo-white' : 'logo-dark'}.svg`;
});

const onTopBarMenuButton = () => {
    topbarMenuActive.value = !topbarMenuActive.value;
};

const onSettingsClick = () => {
    topbarMenuActive.value = false;
    router.push('/utilities/documentation');
};

const topbarMenuClasses = computed(() => {
    return {
        'layout-topbar-menu-mobile-active': topbarMenuActive.value
    };
});

const bindOutsideClickListener = () => {
    if (!outsideClickListener.value) {
        outsideClickListener.value = (event) => {
            if (isOutsideClicked(event)) {
                topbarMenuActive.value = false;
            }
        };

        document.addEventListener('click', outsideClickListener.value);
    }
};

const unbindOutsideClickListener = () => {
    if (outsideClickListener.value) {
        document.removeEventListener('click', outsideClickListener);
        outsideClickListener.value = null;
    }
};

const isOutsideClicked = (event) => {
    if (!topbarMenuActive.value) return;
    const sidebarEl = document.querySelector('.layout-topbar-menu');
    const topbarEl = document.querySelector('.layout-topbar-menu-button');

    return !(sidebarEl.isSameNode(event.target) || sidebarEl.contains(event.target) || topbarEl.isSameNode(event.target) || topbarEl.contains(event.target));
};
</script>

<template>
    <Toast />
    <ConfirmPopup></ConfirmPopup>
    <div v-if="width.valueOf() >= 768" class="layout-topbar">
        <div class="layout-topbar-left-bottom-menu" style="margin-left: 30px"><TabMenu :model="items" /></div>
        <NuxtLink v-if="!useMainStore().get_email" to="auth/login" style="margin-left: auto"><Button label="Log-in" outlined /></NuxtLink>
        <Button v-else @click="showConfirmDialog($event)" style="margin-left: auto; display: flex; flex-direction: row; align-items: center" text>
            <span style="margin-right: 5%; color: white">{{ useMainStore().get_email }}</span>
            <Avatar :label="useMainStore().get_email[0]" style="background-color: #2196f3; font-size: large; color: #ffffff; min-height: 5dvh; min-width: 15%; margin-right: 5dvh" />
        </Button>
        <!-- <Button label="Log-out" outlined style="margin-left: 2%" /> -->
        <!-- <TopMenu class="layout-topbar-left-bottom-menu" style="margin-left: 30px"></TopMenu> -->
        <!-- <router-link to="/" class="layout-topbar-logo">
            <img :src="logoUrl" alt="logo" />
            <span>SAKAI</span>
        </router-link>

        <button class="p-link layout-menu-button layout-topbar-button" @click="onMenuToggle()">
            <i class="pi pi-bars"></i>
        </button>

        <button class="p-link layout-topbar-menu-button layout-topbar-button" @click="onTopBarMenuButton()">
            <i class="pi pi-ellipsis-v"></i>
        </button>

        <div class="layout-topbar-menu" :class="topbarMenuClasses">
            <button @click="onTopBarMenuButton()" class="p-link layout-topbar-button">
                <i class="pi pi-calendar"></i>
                <span>Calendar</span>
            </button>
            <button @click="onTopBarMenuButton()" class="p-link layout-topbar-button">
                <i class="pi pi-user"></i>
                <span>Profile</span>
            </button>
            <button @click="onSettingsClick()" class="p-link layout-topbar-button">
                <i class="pi pi-cog"></i>
                <span>Settings</span>
            </button>
        </div> -->
    </div>
    <div v-else class="layout-topbar">
        <button class="p-link layout-menu-button layout-topbar-button" @click="onMenuToggle()">
            <i class="pi pi-bars"></i>
        </button>
        <Button v-if="useMainStore().get_email" @click="showConfirmDialog($event)" style="margin-right: auto; display: flex; flex-direction: row; align-items: center" text>
            <span style="margin-right: 5%; color: white">{{ useMainStore().get_email }}</span>
            <Avatar :label="useMainStore().get_email[0]" style="background-color: #2196f3; font-size: large; color: #ffffff; min-height: 5dvh; min-width: 15%; margin-right: 5dvh" />
        </Button>
    </div>
</template>

<style lang="scss" scoped>
.layout-topbar {
    position: relative;
    padding-bottom: 0px;
    margin-bottom: 0;
}

.layout-topbar-content {
    display: flex;
    justify-content: flex-end; /* Align to the right-center */
    align-items: flex-end; /* Align to the bottom */
}

.layout-topbar-menu {
    position: absolute;
    bottom: 0;
    right: 0;
}

.layout-topbar-left-bottom-menu {
    position: absolute;
    bottom: 0;
    left: 0;
}
</style>
