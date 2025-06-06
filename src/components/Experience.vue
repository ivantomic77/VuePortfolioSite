<script lang="ts">
export default {
    data() {
        return {
            experiences: [
                {
                    companyName: "OG Consultancy Services",
                    employmentType: "Internship",
                    role: "Software Developer",
                    url: "https://www.og-cs.hr",
                    period: "Jul 2022 - Aug 2022",
                    description: "Worked on a full-stack web application using Spring Boot for backend APIs and Angular for the frontend. Learned how to containerize apps with Docker and got familiar with dev workflows."
                },
                {
                    companyName: "OG Consultancy Services",
                    employmentType: "Full time",
                    role: "Backend Developer",
                    url: "https://www.og-cs.hr",
                    period: "Jul 2023 - Present",
                    description: "I develop backend systems using Java and Spring Boot, design solutions, build REST APIs, and manage data with PostgreSQL and other databases. I work closely with frontend developers and DevOps team to automate testing and deployment. Since the projects vary, I’m always learning new tools and technologies as needed."
                }
            ],
            colorCache: {}
        }
    },
    methods: {
        generateRandomColor() {
            const letters = "0123456789ABCDEF";
            let color = "#";

            for (let i = 0; i < 6; i++) {
                color += letters[Math.floor(Math.random() * 16)];
            }

            return color;
        }
    }
}
</script>

<template>
    <section class="sm:mt-36 mt-32 md:mx-16 mx-6">
        <h2 class="text-4xl font-semibold">Experience</h2>
        <p class="text-xl mb-14">(a tale of bugs, bytes, and brilliant breakthroughs)</p>

        <div class="mt-10">
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-8">
                <div v-for="experience in experiences"
                    class="flex flex-col justify-between shadow-md border-t-4 rounded-lg w-full"
                    :style="{ borderColor: generateRandomColor() }">
                    <div class="px-4 py-5">
                        <p class="text-sm">{{ experience.period }}</p>
                        <h3 class="text-lg font-semibold">{{ experience.companyName }}</h3>
                        <div class="flex gap-2 text-sm">
                            <p>{{ experience.employmentType }}</p>
                            <span>&#8226;</span>
                            <p>{{ experience.role }}</p>
                        </div>
                        <div class="mt-4">{{ experience.description }}</div>
                    </div>
                    <a :href="experience.url"
                        class="bg-black text-white p-2 rounded-b-lg flex justify-between pr-4 transition ease-in-out hover:bg-opacity-70"
                        target="_blank">
                        <span>Go to the site</span>
                        <span class="font-bold">></span>
                    </a>
                </div>
            </div>
        </div>
    </section>
</template>
