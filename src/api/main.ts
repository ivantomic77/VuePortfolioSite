import type { Post } from "./types";

const apiBaseUrl = import.meta.env.VITE_API_URL;

export const getPreviews = async (): Promise<Post[]> => {
    try {
        const response = await fetch(`${apiBaseUrl}/api/previews`);
        
        if (!response.ok) {
            throw new Error(`Failed to fetch previews: ${response.statusText}`);
        }

        const previews = await response.json();
        return previews;
    } catch (error) {
        console.error("Error fetching previews:", error);
        throw error;
    }
};

export const getPost = async (slug: string): Promise<string> => {
    try {
        const response = await fetch(`${apiBaseUrl}/api/posts/${slug}`);

        if (!response.ok) {
            throw new Error(`Failed to fetch post: ${response.statusText}`);
        }

        return await response.text();
    } catch (error) {
        console.error("Error fetching post:", error);
        throw error;
    }
};

