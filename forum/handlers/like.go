package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anamivale/01talent-projects.git/middlewares"
)

type Like struct {
	PostID string
	UserID string
}

func LikeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		PostID string `json:"postID"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil || requestData.PostID == "" {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}
	cookie, err := middlewares.GetCookies(r, "session-cookies")
	if err != nil {
		return
	}
	UserId := cookie.Value
	// Example: UserID is hardcoded for now; use your session or auth system here
	if middlewares.CheckIfAlreadyLiked(requestData.PostID, UserId) {

		err = middlewares.RemoveLike(requestData.PostID, UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {

		err = middlewares.AddLike(requestData.PostID, UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Get updated like count
	likeCount, err := middlewares.GetLikeCount(requestData.PostID)

	middlewares.UpdateLikeCount(likeCount, requestData.PostID)

	if err != nil {
		http.Error(w, "Unable to retrieve like count", http.StatusInternalServerError)
		return
	}

	response := struct {
		Success   bool `json:"success"`
		LikeCount int  `json:"likeCount"`
	}{
		Success:   true,
		LikeCount: likeCount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
