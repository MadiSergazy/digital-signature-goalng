import { defineStore } from 'pinia';

// main is the name of the store. It is unique across your application
// and will appear in devtools
export const useConvStore = defineStore('conv', {
    // a function that returns a fresh state
    state: () => ({
        conv_id: null
    }),
    // optional getters
    getters: {
        // getters receive the state as first parameter
        get_conv_id: (state) => state.conv_id
    },
    // optional actions
    actions: {
        set_conv_id(conv_id: string) {
            this.conv_id = conv_id;
        }
    }
});
