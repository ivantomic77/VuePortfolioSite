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
                    description: "As a high school intern, I gained skills in full-stack development with Spring Boot and Angular and explored Docker containerization"
                },
                {
                    companyName: "OG Consultancy Services",
                    employmentType: "Full time",
                    role: "Backend Developer",
                    url: "https://www.og-cs.hr",
                    period: "Jul 2023 - Present",
                    description: "I collaborate with clients to create innovative solutions, primarily focusing on Spring Boot while continuously expanding my tech horizons."
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
        <h1 class="text-4xl font-semibold">Experience</h1>
        <h2 class="text-xl mb-14">(a tale of bugs, bytes, and brilliant breakthroughs)</h2>

        <div
            class="mt-10">
            <div class="flex flex-wrap sm:gap-8 gap-8">
                <div v-for="experience in experiences" class="flex flex-col justify-between shadow-md border-t-4 rounded-lg max-w-sm" :style="{ borderColor: generateRandomColor() }">
                    <div class="px-4 py-5">
                        <h2 class="text-sm">{{ experience.period }}</h2>
                        <h2 class="text-lg font-semibold">{{ experience.companyName }}</h2>
                        <div class="flex gap-2 text-sm">
                            <h2>{{ experience.employmentType }}</h2>
                            <h2>&#8226;</h2>
                            <h2>{{ experience.role }}</h2>
                        </div>
                        <div class="mt-4">{{ experience.description }}</div>
                    </div>
                    <a :href="experience.url" class="bg-black text-white p-2 rounded-b-lg flex justify-between pr-4 transition ease-in-out hover:bg-opacity-70" target="_blank">
                        <h2>Go to the site</h2>
                        <h2 class="font-bold">></h2>
                    </a>
                </div>
            </div>
        </div>
    </section>
</template>