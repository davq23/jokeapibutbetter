<template>
    <login-form
        @redirect="(route) => $router.replace({ name: route })"></login-form>
</template>

<script lang="ts">
import LoginForm from '@/components/auth/LoginForm.vue';

export default {
    components: {
        LoginForm,
    },
};
</script>
