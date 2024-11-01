package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var (
	basePath             string
	fileCache            = make(map[string][]byte)
	cacheCleanupInterval time.Duration
	cacheMutex           sync.RWMutex
	slugRegex            = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
)

type Preview struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Summary     string `json:"summary"`
	PublishedAt string `json:"publishedAt"`
	Author      string `json:"author"`
}

func init() {
	basePath = os.Getenv("POSTS_BASE_DIR")
	intervalStr := os.Getenv("CACHE_EVICTION_INTERVAL")

	if intervalMinutes, err := strconv.Atoi(intervalStr); err == nil {
		cacheCleanupInterval = time.Duration(intervalMinutes) * time.Minute
	} else {
		cacheCleanupInterval = 60 * time.Minute
	}

	go startCacheCleanup()
}

func startCacheCleanup() {
	for {
		time.Sleep(cacheCleanupInterval)
		cacheMutex.Lock()
		fileCache = make(map[string][]byte)
		cacheMutex.Unlock()
		fmt.Println("Cache cleared")
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getPost(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	slug := r.PathValue("slug")

	if !slugRegex.MatchString(slug) {
		http.Error(w, "Invalid slug format", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join(basePath, slug, "content.md")

	cacheMutex.RLock()
	content, cached := fileCache[filePath]
	cacheMutex.RUnlock()

	if !cached {
		var err error
		content, err = os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "File not found or unable to read file", http.StatusNotFound)
			return
		}
		cacheMutex.Lock()
		fileCache[filePath] = content
		cacheMutex.Unlock()
	}

	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write(content)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Received slug: %s\n", slug)
}

func getPreviews(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var previews []Preview

	entries, err := os.ReadDir(basePath)
	if err != nil {
		http.Error(w, "Failed to read base directory", http.StatusInternalServerError)
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		previewFilePath := filepath.Join(basePath, entry.Name(), "preview.json")

		file, err := os.Open(previewFilePath)
		if err != nil {
			fmt.Printf("Failed to open preview for %s: %v\n", entry.Name(), err)
			continue
		}
		defer file.Close()

		var preview Preview
		if err := json.NewDecoder(file).Decode(&preview); err != nil {
			fmt.Printf("Failed to parse preview for %s: %v\n", entry.Name(), err)
			file.Close()
			continue
		}

		previews = append(previews, preview)
		file.Close()
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(previews); err != nil {
		http.Error(w, "Failed to encode previews", http.StatusInternalServerError)
	}
}
