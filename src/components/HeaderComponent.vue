<script setup lang="ts">
import {
    IconLinkedinWithCircle,
    IconGithubWithCircle,
} from "@iconify-prerendered/vue-entypo-social";
</script>

<template>
    <div class="flex justify-between">
        <router-link to="/" aria-label="Go to the Homepage.">
            <img alt="Ivan Tomic logo" src="../assets/logo.png" class="h-12 w-12" />
        </router-link>
        <div class="flex gap-2">
            <a href="https://www.linkedin.com/in/itomic7/" aria-label="Visit my LinkedIn page." target="_blank"
                ><IconLinkedinWithCircle class="h-9 w-9 hover:animate-spin"
            /></a>
            <a href="https://github.com/ivantomic77" aria-label="Visit my Github page." target="_blank"
                ><IconGithubWithCircle class="h-9 w-9 hover:animate-spin"
            /></a>
        </div>
    </div>
</template>
