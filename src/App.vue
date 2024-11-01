<script setup lang="ts">
import HeaderComponent from "./components/HeaderComponent.vue";
import Footer from "./components/Footer.vue";
import CustomCursor from "./components/CustomCursor.vue";
</script>

<template>
    <div
        class="bg-white w-screen sm:w-[80vw] rounded-lg p-4 max-w-6xl shadow-lg flex flex-col gap-24">
        <header-component />
        <router-view />
        <Footer />
    </div>
    <CustomCursor />
</template>

<style scoped></style>
