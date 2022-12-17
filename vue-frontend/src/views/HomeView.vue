<template>
    <main>
        <welcome-panel />
    </main>
</template>

<script lang="ts">
import WelcomePanel from '@/components/WelcomePanel.vue';
import { defineComponent } from 'vue';

export default defineComponent({
    components: {
        WelcomePanel,
    },
});
</script>
