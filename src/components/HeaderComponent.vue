<script setup lang="ts">
import {
    IconLinkedinWithCircle,
    IconGithubWithCircle,
} from "@iconify-prerendered/vue-entypo-social";
import { inject } from "vue";


const saEvent = inject<any>("saEvent");

const visitGithub = () => {
    saEvent('github_visit'); 
};

const visitLinkedIn = () => {
    saEvent('linkedin_visit');
};
</script>

<template>
    <div class="flex justify-between">
        <a href="#" aria-label="Go to the top."><img alt="Ivan Tomic logo" src="../assets/logo.png" class="h-12 w-12" /></a>
        <div class="flex gap-2">
            <a href="https://www.linkedin.com/in/itomic7/" v-on:click="visitLinkedIn" aria-label="Visit my LinkedIn page." target="_blank"
                ><IconLinkedinWithCircle class="h-9 w-9 hover:animate-spin"
            /></a>
            <a href="https://github.com/ivantomic77" v-on:click="visitGithub" aria-label="Visit my Github page." target="_blank"
                ><IconGithubWithCircle class="h-9 w-9 hover:animate-spin"
            /></a>
        </div>
    </div>
</template>
