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
                        name: "PostgreSQL",
                        icon: "skill-icons:postgresql-dark"
                    },
                    {
                        name: "Redis",
                        icon: "logos:redis",
                    },
                    {
                        name: "GitLab CI",
                        icon: "skill-icons:gitlab-dark"
                    },
                    {
                        name: "Tailwind",
                        icon: "skill-icons:tailwindcss-dark",
                    },
                    {
                        name: "React",
                        icon: "skill-icons:react-dark",
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
                        name: "Kubernetes",
                        icon: "skill-icons:kubernetes",
                    },
                    {
                        name: "Git",
                        icon: "skill-icons:git",
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
        <h2 class="text-4xl font-semibold">Technologies I've used...</h2>
        <p class="text-xl mb-14">(going beyond main tech stack)</p>

        <div class="mt-10">
            <div class="grid gap-10 grid-cols-[repeat(auto-fit,minmax(55px,1fr))]">
                <span
                    v-for="technology in tech"
                    :key="technology.name"
                    class="flex flex-col items-center">
                    <Icon
                        :icon="technology.icon"
                        class="sm:text-6xl text-5xl block" />
                    <p class="text-sm whitespace-nowrap">{{ technology.name }}</p>
                </span>
            </div>
        </div>
    </section>
</template>
