package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	store := NewStore()
	mux := http.NewServeMux()

	// Handles the put function given the key in the URL path
	// Overwrites any existing value, replies with 204 No Content
	mux.HandleFunc("PUT /keys/{key}", func(w http.ResponseWriter, r *http.Request) {
		key := r.PathValue("key")
		body, _ := io.ReadAll(r.Body)

		store.Put(key, string(body))
		w.WriteHeader(http.StatusNoContent)
	})

	// Handles the GET request given the key in the URL path
	// Replies with 404 if key is not found, else writes the value as the response body
	mux.HandleFunc("GET /keys/{key}", func(w http.ResponseWriter, r *http.Request) {
		key := r.PathValue("key")

		value, ok := store.Get(key)
		if !ok {
			http.NotFound(w, r)
		} else {
			fmt.Fprint(w, value)
		}
	})

	// Handles the DELETE request given the key in the URL path
	// If the key doesn't exist, replies with 404 error, else deletes the value at the key
	mux.HandleFunc("DELETE /keys/{key}", func(w http.ResponseWriter, r *http.Request) {
		key := r.PathValue("key")

		_, ok := store.Get(key)
		if !ok {
			http.NotFound(w, r)
		} else {
			store.Delete(key)
			w.WriteHeader(http.StatusNoContent)
		}
	})

	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", mux)
}
