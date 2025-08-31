<script setup lang="ts">
import { ref } from "vue"
import { Icon } from "@iconify/vue";
import { generateRandomColor } from "../helpers/ColorHelpers"

interface Experience {
  companyName: string
  employmentType: string
  role: string
  url: string
  period: string
  description: string
}

const experiences = ref<Experience[]>([
  {
    companyName: "OG Consultancy Services",
    employmentType: "Internship",
    role: "Software Developer",
    url: "https://www.og-cs.hr",
    period: "Jul 2022 - Aug 2022",
    description:
      "During my internship, I learned how to develop an entire application from backend to frontend and deploy it with Docker Compose."
  },
  {
    companyName: "OG Consultancy Services",
    employmentType: "Full time",
    role: "Backend Developer",
    url: "https://www.og-cs.hr",
    period: "Jul 2023 - Present",
    description:
      "I focus on backend development and collaborate with other teams to deliver solutions, with occasional frontend work and involvement in PoCs for internal and client projects."
  }
])
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
                    <a :href="experience.url" class="bg-black text-white p-2 rounded-b-lg flex justify-between pr-4 transition ease-in-out hover:bg-opacity-70" target="_blank">
                        <h2>Go to the site</h2>
                        <Icon
                            icon="ci:external-link"
                            width="24" height="24"
                            :style="{ color: '#fff' }"
                            class="sm:text-6xl text-5xl"
                        />
                    </a>
                </div>
            </div>
        </div>
    </section>
</template>
