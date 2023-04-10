<script lang="ts">
import {
    IconLinkedinWithCircle,
    IconGithubWithCircle,
} from "@iconify-prerendered/vue-entypo-social";

export default {
    components: {
        IconLinkedinWithCircle,
        IconGithubWithCircle,
    },
};
</script>

<template>
    <div class="flex justify-between">
        <a href="#"><img src="../assets/logo.png" class="h-12 w-12" /></a>
        <div class="flex gap-2">
            <a href="https://www.linkedin.com/in/itomic7/" target="_blank"
                ><IconLinkedinWithCircle class="h-9 w-9 hover:animate-spin"
            /></a>
            <a href="https://github.com/ivantomic77" target="_blank"
                ><IconGithubWithCircle class="h-9 w-9 hover:animate-spin"
            /></a>
        </div>
    </div>
</template>
