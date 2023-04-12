<script lang="ts">
import { Icon } from "@iconify/vue";

export default {
    data() {
        return {
            tech: {
                Web: [
                    {
                        name: "SpringBoot",
                        icon: "skill-icons:spring-dark",
                    },
                    {
                        name: "Java",
                        icon: "skill-icons:java-dark",
                    },
                    {
                        name: "HTML",
                        icon: "skill-icons:html",
                    },
                    {
                        name: "CSS",
                        icon: "skill-icons:css",
                    },
                    {
                        name: "Tailwind",
                        icon: "skill-icons:tailwindcss-dark",
                    },
                    {
                        name: "Javascript",
                        icon: "skill-icons:javascript",
                    },
                    {
                        name: "Typescript",
                        icon: "skill-icons:typescript",
                    },
                    {
                        name: "Angular",
                        icon: "skill-icons:angular-dark",
                    },
                    {
                        name: "Vue",
                        icon: "skill-icons:vuejs-dark",
                    },
                ],
                "Cloud and Containerization": [
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
                ],
                Other: [
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
            },
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
            v-for="(technologies, sectionName) in tech"
            :key="sectionName"
            class="mt-10">
            <h1 class="text-2xl font-semibold mb-5 relative">
                <span
                    class="border-b-2 border-black absolute bottom-0 left-0 right-0 sm:w-80 w-72"></span>
                {{ sectionName }}
            </h1>
            <div class="flex flex-wrap sm:gap-10 gap-3">
                <a
                    v-for="technology in technologies"
                    :key="technology.name"
                    class="flex flex-col items-center w-[80px]">
                    <Icon
                        :icon="technology.icon"
                        class="sm:text-7xl text-6xl" />
                    <h1>{{ technology.name }}</h1>
                </a>
            </div>
        </div>
    </section>
</template>
