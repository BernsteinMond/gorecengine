package handler

import (
	"encoding/json"
	"fmt"
	"github.com/BernsteinMond/gorecengine/src/internal/inference/service"
	"github.com/google/uuid"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux, service service.Service) {
	mux.Handle("/recommendation/posts", GetPostsRecommendation(service))
}

func GetPostsRecommendation(service service.Service) http.HandlerFunc {
	type response struct {
		Recommendation []postDTO `json:"recommendation"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, fmt.Sprintf("method \"%s\" not allowed", r.Method), http.StatusMethodNotAllowed)
			return
		}

		userIDStr := r.URL.Query().Get("user_id")
		if userIDStr == "" {
			http.Error(w, "no \"user_id\" query parameter was provided", http.StatusBadRequest)
			return
		}
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			http.Error(w, "invalid \"user_id\" query parameter format", http.StatusBadRequest)
			return
		}

		recommendation, err := service.GetRecommendationByUserID(r.Context(), userID)
		if err != nil {
			http.Error(w, fmt.Sprintf("service: get recommendation by user id: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		resp := response{
			Recommendation: make([]postDTO, len(recommendation.Posts)),
		}

		for i, post := range recommendation.Posts {
			resp.Recommendation[i] = fromDomainToPostDTO(&post)
		}

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
