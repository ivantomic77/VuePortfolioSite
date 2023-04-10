<script lang="ts">
import HeaderComponent from "./HeaderComponent.vue";
import Hero from "./Hero.vue";
export default {
    components: {
        HeaderComponent,
        Hero,
    },
};
</script>

<template>
    <div
        class="bg-white w-screen min-h-screen sm:min-h-[70vh] sm:w-[80vw] rounded-lg p-4 max-w-6xl shadow-lg"
    >
        <header-component />
        <hero />
    </div>
</template>
