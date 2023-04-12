<script lang="ts">
import {
    IconLinkedinWithCircle,
    IconGithubWithCircle,
} from "@iconify-prerendered/vue-entypo-social";
export default {
    mounted() {
        this.getYear();
    },
    data() {
        return {
            year: 0,
        };
    },
    methods: {
        getYear() {
            this.year = new Date().getFullYear();
        },
    },
    components: {
        IconLinkedinWithCircle,
        IconGithubWithCircle,
    },
};
</script>

<template>
    <div class="flex justify-center items-center flex-col gap-10 mt-40">
        <h1 class="font-bold text-4xl">Contact me</h1>
        <div class="flex flex-col gap-3">
            <a
                href="mailto: ivan.tomic2108@gmail.com"
                class="font-semibold text-2xl"
                >ivan.tomic2108@gmail.com</a
            >
            <div class="flex justify-center items-center gap-2">
                <a href="https://www.linkedin.com/in/itomic7/" target="_blank"
                    ><IconLinkedinWithCircle class="h-9 w-9 hover:animate-spin"
                /></a>
                <a href="https://github.com/ivantomic77" target="_blank"
                    ><IconGithubWithCircle class="h-9 w-9 hover:animate-spin"
                /></a>
            </div>
        </div>
        <div class="flex justify-center items-center flex-col gap-1 mt-20">
            <div class="w-80 h-2 bg-black"></div>
            <h1>&copy; {{ year }} itomic.tech</h1>
        </div>
    </div>
</template>
