package redis

import (
	"encoding/json"
	"time"

	"github.com/drhax9908/Meme_Api/models"
)

// WritePostsToCache : Takes a List of Memes and writes it to Cache
func (r Redis) WritePostsToCache(sub string, memes []models.Meme) (ok bool) {
	memesBinary, err := json.Marshal(memes)

	if err != nil {
		return false
	}

	err = r.Client.Set(sub, memesBinary, 2*time.Hour).Err()

	if err != nil {
		return false
	}

	return true
}
