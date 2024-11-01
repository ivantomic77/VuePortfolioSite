<template>
    <section class="md:mx-16 mx-6">
        <h1 class="text-4xl font-semibold">Articles</h1>
        <h2 class="text-xl mb-14">(Insights, updates, and reflections)</h2>

        <div class="mt-10">
            <div class="flex flex-wrap sm:gap-8 gap-8">
                <div v-for="preview in previews" class="flex flex-col justify-between shadow-md border-t-4 rounded-lg max-w-sm" :style="{ borderColor: generateRandomColor() }">
                    <div class="px-4 py-5">
                        <h2 class="text-sm">{{ preview.author }} <span>&#8226;</span> {{ preview.publishedAt }}</h2>
                        <h2 class="text-lg font-semibold">{{ preview.title }}</h2>
                        <div class="mt-2">{{ preview.summary }}</div>
                    </div>
                    <button @click="goToArticle(preview.slug)" class="bg-black text-white p-2 rounded-b-lg flex justify-between pr-4 transition ease-in-out hover:bg-opacity-70">
                        <h2>Read full article</h2>
                        <h2 class="font-bold">></h2>
                    </button>
                </div>
            </div>
        </div>
    </section>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { getPreviews } from '../api/main';
import type { Post } from '../api/types';
import router from '../router';

const previews = ref<Post[]>([]);

const fetchPreviews = async () => {
    try {
        previews.value = await getPreviews();
    } catch (error) {
        console.error("Failed to load previews:", error);
    }
};

const goToArticle = (slug: string) => {
    router.push({ 
        path: `/posts/${slug}`
    })
};

const generateRandomColor = () => {
    const letters = "0123456789ABCDEF";
    let color = "#";
    for (let i = 0; i < 6; i++) {
        color += letters[Math.floor(Math.random() * 16)];
    }
    return color;
};

onMounted(fetchPreviews);
</script>
