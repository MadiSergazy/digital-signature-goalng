<template>
    <TabMenu :model="items" />
    <!-- <router-view /> -->
</template>

<script>
export default {
    data() {
        return {
            items: [
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
            ]
        };
    }
};
</script>
