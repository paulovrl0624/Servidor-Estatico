
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"sync"
)

var (
	cache      = make(map[string][]byte)
	cacheMutex = sync.RWMutex{}
	publicDir  = "./public"
)

func getMimeType(path string) string {
	ext := filepath.Ext(path)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		return "application/octet-stream"
	}
	return mimeType
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "/index.html"
	}
	filePath := filepath.Join(publicDir, filepath.Clean(path))

	cacheMutex.RLock()
	data, cached := cache[filePath]
	cacheMutex.RUnlock()

	if !cached {
		fileData, err := ioutil.ReadFile(filePath)
		if err != nil {
			http.Error(w, "404 Not Found", http.StatusNotFound)
			return
		}
		cacheMutex.Lock()
		cache[filePath] = fileData
		cacheMutex.Unlock()
		data = fileData
	}

	mimeType := getMimeType(filePath)
	w.Header().Set("Content-Type", mimeType)
	w.Write(data)
}

func main() {
	fmt.Println("Servidor rodando em http://localhost:8080")
	http.HandleFunc("/", fileHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
