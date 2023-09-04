<template>
    <div :class="['flex-row', alignmentClass]">
        <slot></slot>
    </div>
</template>

<script>
export default {
    props: {
        mainAxisAlignment: {
            type: String,
            default: 'flex-start' // Default alignment
        },
        crossAxisAlignment: {
            type: String,
            default: 'stretch' // Default alignment
        }
    },
    computed: {
        alignmentClass() {
            return `justify-${this.mainAxisAlignment} align-${this.crossAxisAlignment}`;
        }
    }
};
</script>

<style>
.flex-row {
    display: flex;
    flex-direction: row;
}

.justify-flex-start {
    justify-content: flex-start;
}
.justify-center {
    justify-content: center;
}
.justify-flex-end {
    justify-content: flex-end;
}
/* Add more justifyContent classes as needed */

.align-stretch {
    align-items: stretch;
}
.align-center {
    align-items: center;
}
.align-flex-start {
    align-items: flex-start;
}
.align-flex-end {
    align-items: flex-end;
}
/* Add more alignItems classes as needed */
</style>
