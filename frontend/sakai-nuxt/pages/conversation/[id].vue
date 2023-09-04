<template>
    <div class="page-container">
        <Toast />
        <FlexColumn>
            <FlexColumn class="card content" style="height: 85dvh">
                <div style="display: flex; justify-content: space-between; align-items: center">
                    <div style="display: flex; align-items: center">
                        <Button icon="pi pi-angle-left" style="margin-right: 10px; min-width: 2%" @click="router.push('/')" />
                        <span class="text-900">Артқа қайту</span>
                    </div>
                    <Dropdown v-model="selectedLanguage" :options="languages" optionLabel="language" placeholder="🇰🇿 Қазақ тілі" style="margin-left: 10px" />
                </div>
                <span v-if="isConversationEmpty && !loading" class="small-span" style="align-self: center">KAZAI</span>
                <Animation v-if="isConversationEmpty && !loading" :animationData="animationData" style="height: 33%; width: 33%; margin-left: 32%"></Animation>
                <span v-if="isConversationEmpty && !loading" class="responsive-text">Сұрақ қойып әңгіме бастаңыз...</span>
                <div style="display: flex; flex-direction: column; margin-top: 4dvh" v-if="loading">
                    <Skeleton class="mb-2"></Skeleton>
                    <Skeleton width="10rem" class="mb-2"></Skeleton>
                    <Skeleton width="5rem" class="mb-2"></Skeleton>
                    <Skeleton height="2rem" class="mb-2"></Skeleton>
                    <Skeleton width="10rem" height="4rem"></Skeleton>
                </div>
                <FlexColumn style="margin-top: 30px">
                    <div v-for="(message, index) in messages" :key="index">
                        <FlexRow style="margin-bottom: 10px">
                            <Avatar v-if="message.isgpt" label="G" class="mr-2" style="background-color: green; color: #ffffff; min-width: 34px; min-height: 34px" />
                            <Avatar v-else label="M" class="mr-2" style="background-color: #2196f3; color: #ffffff; min-width: 34px; min-height: 34px" />
                            <div class="card" style="flex-direction: column; display: flex; padding-top: 2px; padding-right: 10px; padding-left: 15px; margin-right: 5px; min-width: 100px; padding-bottom: 15px">
                                <Button icon="pi pi-copy" severity="secondary" text style="margin-left: auto; margin-top: 0" @click="copyToClipboard(message.message)" />
                                <span style="margin-top: 5px; white-space: normal">{{ message.message }}</span>
                            </div>
                        </FlexRow>
                    </div>
                </FlexColumn>
            </FlexColumn>

            <FlexRow class="bottom-row" style="margin-bottom: env(safe-area-inset-bottom)">
                <FlexColumn class="bottom-column">
                    <InputText :class="inputClass" id="username" v-model="propmtToServer" style="font-size: large" placeholder="Сұрақ қойыңыз" />
                    <small v-if="isEmptySubmit" class="p-error" id="text-error">Бірнәрсе жазу міндетті</small>
                </FlexColumn>
                <Button icon="pi pi-send" type="submit" aria-label="Filter" style="height: 5vh; width: 5vh; margin-left: 2vw" @click="sendPrompt" />
            </FlexRow>
        </FlexColumn>
    </div>
</template>

<script>
import { useRouter } from 'vue-router';
import animationData from '~/assets/animations/animation_lljapsvn.json';
// import { useMainStore } from '../service/composables/userstore';
import { useConvStore } from '../../service/composables/convstore';
export default {
    setup() {
        definePageMeta({
            middleware: ['conversation-middleware']
        });
        const nuxtApp = useNuxtApp();

        return { nuxtApp };
    },
    mounted() {
        // const route = useRoute();
        // const id = route.params.id;
        // this.conv_id = null;
        const id = useConvStore().get_conv_id;
        console.log('id:', id);
        this.init(id);
    },
    data() {
        return {
            animationData: animationData,
            isEmptySubmit: false,
            inputClass: 'input-text',
            propmtToServer: '',
            messages: [],
            router: useRouter(),
            selectedLanguage: { language: 'Қазақ тілі' },
            languages: [{ language: 'Қазақ тілі' }, { language: 'English' }],
            conv_id: null,
            loading: false
        };
    },
    methods: {
        copyToClipboard(text) {
            const tempInput = document.createElement('textarea');
            tempInput.value = text;
            document.body.appendChild(tempInput);
            tempInput.select();
            document.execCommand('copy');
            document.body.removeChild(tempInput);
            this.$toast.add({ severity: 'success', summary: 'Сәтті', detail: 'Сақталды', life: 3000 });
        },
        async sendPrompt() {
            if (this.propmtToServer.length == 0) {
                this.isEmptySubmit = true;
                this.inputClass = 'p-invalid';
            } else {
                this.messages.push({ message: this.propmtToServer, isgpt: false });
                var prompt = this.propmtToServer;
                this.propmtToServer = '';
                this.isEmptySubmit = false;
                this.inputClass = 'input-text';
                try {
                    var response = await this.nuxtApp.$atamekenservice().askQuestion(prompt, this.conv_id);
                } catch (e) {
                    console.log('Error:', e);
                    this.$toast.add({ severity: 'info', summary: 'Info', detail: e.response.data['detail'], life: 3000 });
                    return;
                }

                console.log('response[data]:', response['data'].value);
                this.conv_id = response['data'].value['conv_id'];
                this.messages.push({ message: response['data'].value['answer'], isgpt: true });
            }
        },
        async init(id) {
            console.log(id != null);
            if (id != 'newchat') {
                console.log(typeof id);
                this.conv_id = id;
                this.loading = true;
                const response = await this.nuxtApp.$atamekenservice().get_messages(this.conv_id);
                this.loading = false;
                console.log(response.data.value);
                for (let index = 0, j = 1; index < response.data.value.length; index += 2, j += 2) {
                    this.messages.push({ message: response.data.value[index], isgpt: false });
                    this.messages.push({ message: response.data.value[j], isgpt: true });
                }
                return;
            }
            // this.messages = [];
        }
    },

    computed: {
        isConversationEmpty: function () {
            return this.messages.length == 0;
        }
    }
};
</script>
<style scoped>
/* Apply flex layout to the immediate parent of the card column */
body {
    -webkit-overflow-scrolling: auto;
}
.card-container {
    display: flex;
    flex-direction: column;
    flex-grow: 1; /* Ensure the container grows to fill the available space */
}
.content {
    overflow-x: auto;
    white-space: nowrap;
}
.small-span {
    font-size: 300%; /* Adjust the font size as needed */
    max-width: 500%; /* Adjust the maximum width as needed */
    overflow: hidden; /* Hide content that exceeds the width */
    white-space: nowrap; /* Prevent line breaks */
    opacity: 0.3;
}
/* Default styles for larger screens */
.bottom-row {
    margin-left: 25%;
}

.bottom-column {
    width: 60%;
}
.responsive-text {
    margin-left: 40%;
    font-weight: 500;
    font-size: 2dvh; /* Default font size */
}
/* Styles for screens with max-width of 768px (e.g., mobile) */
@media screen and (max-width: 768px) {
    .bottom-row {
        margin-left: 0;
        width: 100%;
    }
    .bottom-column {
        margin-left: 0;
        width: 90%;
    }
    body {
        padding-bottom: env(safe-area-inset-bottom);
    }
    .responsive-text {
        font-size: small;
        margin-left: 20%;
    }
}
</style>
