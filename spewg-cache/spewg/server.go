package spewg

import (
	"encoding/json"
	"net/http"
	"time"
)

type CacheServer struct {
	cache *Cache
}

func NewCacheServer() *CacheServer {
	return &CacheServer{
		cache: NewCache(),
	}
}

func (cs *CacheServer) SetHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cs.cache.Set(req.Key, req.Value, 1*time.Hour)
	w.WriteHeader(http.StatusOK)
}

func (cs *CacheServer) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, found := cs.cache.Get(key)
	if !found {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"value": value})
}


