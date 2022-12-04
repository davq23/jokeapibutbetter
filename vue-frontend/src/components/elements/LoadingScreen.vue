<template>
    <div class="entire-screen">
        <v-progress-circular v-if="loading" indeterminate></v-progress-circular>
    </div>
</template>

<script>
export default {};
</script>

<style scoped>
.entire-screen {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100vw;
    height: 100vh;
}
</style>
