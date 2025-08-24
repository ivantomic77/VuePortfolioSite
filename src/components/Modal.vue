<template>
    <div
      v-if="isOpen"
      class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-start pt-20 z-50"
      @click.self="close"
    >
      <div class="bg-white rounded-lg max-w-3xl w-full p-6 overflow-y-auto max-h-[80vh]">
        <h3 class="text-sm">{{ project.role }}</h3>
        <h2 class="text-2xl font-semibold">{{ project.name }}</h2>
        <div class="flex flex-wrap text-sm mb-4">
          <template v-for="(tech, index) in project.technologies" :key="tech">
            <span class="flex items-center">
              <span>{{ tech }}</span>
              <span v-if="index < project.technologies.length - 1" class="mx-2">•</span>
            </span>
          </template>
        </div>
  
        <div class="space-y-4 article-content">
            <div v-for="(section, index) in project.articleSections" :key="index" v-html="section"></div>
        </div>

  
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
  
        <button @click="close" class="mt-6 bg-black text-white p-2 rounded w-full">Close</button>
      </div>
    </div>
</template>
  
<script setup lang="ts">
  import { defineProps, defineEmits } from "vue";
  import { Icon } from "@iconify/vue";
  
  defineProps<{
    isOpen: boolean;
    project: {
      name: string;
      role: string;
      technologies: string[];
      articleSections: string[];
      linksToProject?: { displayName: string; link: string }[];
    };
  }>();
  
  const emit = defineEmits<{
    (e: "close"): void;
  }>();
  
  function close() {
    emit("close");
  }
</script>

<style scoped>
.article-content >>> a {
  font-style: italic;
}
</style>
