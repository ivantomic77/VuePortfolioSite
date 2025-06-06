<script lang="ts" setup>
import {
    IconLinkedinWithCircle,
    IconGithubWithCircle,
} from "@iconify-prerendered/vue-entypo-social";
import { ref, onMounted } from "vue";

const year = ref(0);

const getYear = () => {
    year.value = new Date().getFullYear();
};

onMounted(() => {
    getYear();
});
</script>

<template>
    <section class="flex justify-center items-center flex-col gap-10 mt-40">
        <h2 class="font-bold text-4xl">Contact me</h2>
        <div class="flex flex-col gap-3">
            <img class="w-80" alt="Email address" src="../assets/email.png" width="320" height="48" />
            <div class="flex justify-center items-center gap-2">
                <a href="https://www.linkedin.com/in/itomic7/" aria-label="Visit my LinkedIn page." target="_blank">
                    <IconLinkedinWithCircle class="h-9 w-9 hover:animate-spin" />
                </a>
                <a href="https://github.com/ivantomic77" aria-label="Visit my Github page." target="_blank">
                    <IconGithubWithCircle class="h-9 w-9 hover:animate-spin" />
                </a>
            </div>
        </div>
        <div class="flex justify-center items-center flex-col gap-1 mt-8">
            <div class="w-80 h-2 bg-black"></div>
            <p>&copy; {{ year }} tomic.dev</p>
        </div>
    </section>
</template>
