<template>
    <div class="text-center">
        <div class="text-center" v-if="preview">
            <h1 class="text-4xl font-bold">{{ preview.title }}</h1>
            <h2 class="text-sm">
                {{ preview.author }} <span>&#8226;</span> {{ preview.publishedAt }}
            </h2>
            <div class="mt-4">{{ preview.summary }}</div>
        </div>
        <div v-if="loading" class="flex justify-center items-center mt-10">
            <svg class="animate-spin h-10 w-10 text-blue-500" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 1 1 16 0A8 8 0 0 1 4 12z"></path>
            </svg>
        </div>
        <div v-else-if="error" class="mt-4 text-red-500">
            <h1 class="text-2xl">404 - Post Not Found</h1>
        </div>
        <div v-else v-html="postContent" class="mt-4 prose mx-auto p-4 border border-gray-300 rounded-lg shadow-md"></div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { getPost, getPreviews } from '../api/main';
import type { Post } from '../api/types';

const postContent = ref<string>('');
const preview = ref<Post | null>(null);

const loading = ref<boolean>(true);
const error = ref<boolean>(false);
const route = useRoute();

const fetchPost = async () => {
    const slug = route.params.slug as string;

    try {
        postContent.value = await getPost(slug);
    } catch (err) {
        console.error("Failed to load post:", err);
        error.value = true;
    } finally {
        loading.value = false;
    }
};

const fetchPreviews = async () => {
    try {
        const previews = await getPreviews();
        const slug = route.params.slug as string;
        preview.value = previews.find((p: Post) => p.slug === slug) || null;
    } catch (error) {
        console.error("Failed to load previews:", error);
    }
};

onMounted(fetchPost);
onMounted(fetchPreviews);
</script>
