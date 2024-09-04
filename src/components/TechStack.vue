<script lang="ts">
import { Icon } from "@iconify/vue";

export default {
    data() {
        return {
            tech: [
                    {
                        name: "Java",
                        icon: "skill-icons:java-dark"
                    },
                    {
                        name: "SpringBoot",
                        icon: "skill-icons:spring-dark",
                    },
                    {
                        name: "Quarkus",
                        icon: "logos:quarkus-icon"
                    },
                    {
                        name: "Tailwind",
                        icon: "skill-icons:tailwindcss-dark",
                    },
                    {
                        name: "Angular",
                        icon: "skill-icons:angular-dark",
                    },
                    {
                        name: "SvelteKit",
                        icon: "skill-icons:svelte"
                    },
                    {
                        name: "Vue",
                        icon: "skill-icons:vuejs-dark",
                    },
                    {
                        name: "AWS",
                        icon: "skill-icons:aws-dark",
                    },
                    {
                        name: "Linode",
                        icon: "logos:linode",
                    },
                    {
                        name: "Docker",
                        icon: "skill-icons:docker",
                    },
                    {
                        name: "Git",
                        icon: "skill-icons:git",
                    },
                    {
                        name: "Python",
                        icon: "skill-icons:python-dark",
                    },
                    {
                        name: "Postman",
                        icon: "skill-icons:postman",
                    },
                    {
                        name: "Figma",
                        icon: "skill-icons:figma-dark",
                    },
                    {
                        name: "Photoshop",
                        icon: "skill-icons:photoshop",
                    },
                    {
                        name: "Unity",
                        icon: "skill-icons:unity-dark",
                    },
                    {
                        name: "Arduino",
                        icon: "skill-icons:arduino",
                    },
                    {
                        name: "OpenAI",
                        icon: "logos:openai-icon",
                    },
                ],
        };
    },

    components: {
        Icon,
    },
};
</script>

<template>
    <section class="sm:mt-60 mt-28 md:mx-16 mx-6">
        <h1 class="text-4xl font-semibold">Technologies I've used...</h1>
        <h2 class="text-xl mb-14">(going beyond main tech stack)</h2>

        <div
            class="mt-10">
            <div class="flex flex-wrap sm:gap-8 gap-8">
                <span
                    v-for="technology in tech"
                    :key="technology.name"
                    class="flex flex-col items-center w-[40px]">
                    <Icon
                        :icon="technology.icon"
                        class="sm:text-5xl text-4xl" />
                    <h1 class="text-sm">{{ technology.name }}</h1>
                </span>
            </div>
        </div>
    </section>
</template>
