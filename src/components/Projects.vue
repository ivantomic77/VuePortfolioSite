<template>
    <section class="sm:mt-36 mt-32 md:mx-16 mx-6">
      <h1 class="text-4xl font-semibold">Notable Personal Projects</h1>
      <h2 class="text-xl mb-14">(building, testing, and being our own clients)</h2>
  
      <div class="mt-10">
        <div class="flex flex-wrap sm:gap-8 gap-8">
          <div
            v-for="project in projects"
            :key="project.name"
            class="flex flex-col justify-between shadow-md border-t-4 rounded-lg max-w-sm"
            :style="{ borderColor: generateRandomColor() }"
          >
            <div class="px-4 py-5">
              <h2 class="text-sm">{{ project.role }}</h2>
              <h2 class="text-lg font-semibold">{{ project.name }}</h2>
              <div class="flex flex-wrap text-sm">
                <template v-for="(tech, index) in project.technologies" :key="tech">
                  <span class="flex items-center">
                    <span>{{ tech }}</span>
                    <span v-if="index < project.technologies.length - 1" class="mx-2">•</span>
                  </span>
                </template>
              </div>
  
              <div class="mt-4 break-words">{{ project.description }}</div>
  
              <div class="mt-4 flex gap-4 flex-wrap">
                <a
                  v-for="link in project.linksToProject || []"
                  :key="link.link"
                  :href="link.link"
                  target="_blank"
                  class="hover:underline text-sm flex items-center"
                >
                  {{ link.displayName }}
                  <Icon icon="ci:external-link" width="18" height="18" :style="{ color: '#000' }" />
                </a>
              </div>
            </div>
  
            <button
              v-if="project.articleSections && !project.linkToArticle"
              @click="openModal(project)"
              class="bg-black text-white p-2 rounded-b-lg flex justify-between pr-4 transition ease-in-out hover:bg-opacity-70 w-full"
            >
              <h2>Read the article</h2>
              <Icon
                icon="ci:external-link"
                width="24"
                height="24"
                :style="{ color: '#fff' }"
                class="sm:text-6xl text-5xl"
              />
            </button>
  
            <a
              v-else-if="project.linkToArticle"
              :href="project.linkToArticle"
              target="_blank"
              class="bg-black text-white p-2 rounded-b-lg flex justify-between pr-4 transition ease-in-out hover:bg-opacity-70 w-full"
            >
              <h2>Read the article</h2>
              <Icon
                icon="ci:external-link"
                width="24"
                height="24"
                :style="{ color: '#fff' }"
                class="sm:text-6xl text-5xl"
              />
            </a>
          </div>
        </div>
      </div>

      <Modal :isOpen="modalOpen" :project="selectedProject" @close="closeModal" />
    </section>
  </template>
  
<script setup lang="ts">
  import { ref } from "vue";
  import { generateRandomColor } from "../helpers/ColorHelpers";
  import { Icon } from "@iconify/vue";
  import Modal from "./Modal.vue";
  
  const modalOpen = ref(false);
  const selectedProject = ref<any>(null);
  
  function openModal(project: any) {
    selectedProject.value = project;
    modalOpen.value = true;
  }
  
  function closeModal() {
    modalOpen.value = false;
    selectedProject.value = null;
  }
  
  const projects = [
    {
      name: "UsporediMe",
      description:
        "UsporediMe was developed together with two colleagues and uses daily online supermarket price lists in Croatia to help users compare prices and find the lowest-cost options quickly.",
      role: "Backend Developer, DevOps",
      technologies: ["Spring Boot", "PostgreSQL", "Docker Swarm"],
      linksToProject: [
        { displayName: "Play Store", link: "https://play.google.com/store/apps/details?id=com.kgt.usporedime" },
        { displayName: "Website", link: "https://www.usporedime.com/" },
      ],
      linkToArticle:
        "https://www.linkedin.com/posts/itomic7_usporedimecom-usporedite-cijene-namirnica-activity-7331197636181397504-Gvtj",
    },
    {
      name: "Linode → Oracle Kubernetes Migration",
      description:
        "Migrated multiple personal services (websites, Node-RED, Uptime Kuma, and more) from a Linode server to a (free) k3s Kubernetes cluster on Oracle Cloud. This project helped me learn Kubernetes, container orchestration, and DevOps practices, enabling me to create and manage my own deployments.",
      role: "DevOps Engineer, Cloud Enthusiast",
      technologies: ["Kubernetes", "Docker", "Oracle Cloud", "CI/CD"],
      articleSections: [
        // tbd
      ]
    }
  ];
</script>
  