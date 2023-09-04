<template>
    <div style="display: flex; flex-direction: column; width: 40%; margin-left: 50dvh">
        <div style="display: flex; flex-direction: row; justify-content: space-between; width: 100%; margin-bottom: 1dvh">
            <span style="font-weight: bold; font-size: 2.5dvh">Сұраныстар</span>
            <div style="display: flex; flex-direction: row">
                <Button icon="pi pi-pencil" aria-label="Filter" style="margin-right: 20%; width: 5dvh" />
                <Button icon="pi pi-trash" aria-label="Filter" style="width: 5dvh; color: black" severity="danger" />
            </div>
        </div>
        <div v-if="!isloading" v-for="conversation in conversations">
            <div style="display: flex; flex-direction: row; justify-content: space-between; width: 100%; margin-top: 5%">
                <Button @click="pushtoconversation(conversation.id)" text>
                    <div style="display: flex; flex-direction: column; align-items: flex-start">
                        <span style="color: white">{{ conversation.title }}</span>
                        <span style="color: white">{{ formattedDate(conversation.update_time) }}</span>
                    </div>
                </Button>
                <Button icon="pi pi-trash" aria-label="Filter" style="width: 4.5dvh; height: 4.5dvh; color: black" severity="danger" @click="(visible = true), (selectedConv = conversation)" />
            </div>
            <hr aria-orientation="horizontal" style="margin-top: 3dvh" />
        </div>
        <div v-else>
            <Skeleton height="4rem" style="margin-top: 3%"></Skeleton>
            <Skeleton height="4rem" style="margin-top: 3%"></Skeleton>
            <Skeleton height="4rem" style="margin-top: 3%"></Skeleton>
            <Skeleton height="4rem" style="margin-top: 3%"></Skeleton>
            <Skeleton height="4rem" style="margin-top: 3%"></Skeleton>
        </div>
    </div>
    <Dialog v-model:visible="visible" modal header="Сенімдісіз бе?" :style="{ width: '50vw' }">
        <p>Егер келіссеңіз "{{ selectedConv.title }}" сұраныс парақшасын қайтара алмайсыз</p>
        <template #footer>
            <Button label="Жоқ" icon="pi pi-times" @click="visible = false" text />
            <Button label="Иә" icon="pi pi-check" @click="deleteConv" autofocus />
        </template>
    </Dialog>
</template>
<script>
// import { useMainStore } from '../service/composables/userstore';
import { useConvStore } from '../service/composables/convstore';
import moment from 'moment';
export default {
    setup() {
        definePageMeta({
            middleware: ['history-middleware']
        });
        const nuxtApp = useNuxtApp();
        return { nuxtApp };
    },
    data() {
        return {
            conversations: null,
            inputFormat: 'YYYY-MM-DDTHH:mm:ss.SSSSSSZ',
            outputFormat: 'DD.MM.YYYY HH:mm',
            isloading: true,
            visible: false,
            selectedConv: null
        };
    },
    mounted() {
        this.init();
    },
    methods: {
        formattedDate(inputDateString) {
            const parsedDate = moment(inputDateString, this.inputFormat);
            return parsedDate.format(this.outputFormat);
        },
        async deleteConv() {
            this.visible = false;
            this.nuxtApp.$atamekenservice().deleteConv(this.selectedConv.id);
            this.$toast.add({ severity: 'success', summary: 'Өшірілді', detail: this.selectedConv.title, life: 3000 });
            await this.init();
        },
        async init() {
            this.isloading = true;
            const response = await this.nuxtApp.$atamekenservice().getConversatinos();
            this.isloading = false;
            this.conversations = response.data;
        },
        async pushtoconversation(conv_id) {
            useConvStore().set_conv_id(conv_id);
            this.$router.push('/conversation/' + conv_id);
        }
    }
};
</script>
<style></style>
