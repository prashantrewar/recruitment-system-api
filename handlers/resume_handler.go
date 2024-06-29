package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"recruitment-system/middleware"
	"recruitment-system/models"
	"recruitment-system/utils"

	"gorm.io/gorm"
)

func UploadResumeHandler(db *gorm.DB, apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(middleware.UserContextKey).(*models.User)

		if user.UserType != "Applicant" {
			http.Error(w, "Only applicants can upload resumes", http.StatusForbidden)
			return
		}

		file, _, err := r.FormFile("resume")
		if err != nil {
			log.Printf("Failed to get form file: %v", err)
			http.Error(w, "Invalid file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		profile, err := utils.ParseResume(file, apiKey)
		if err != nil {
			log.Printf("Failed to parse resume: %v", err)
			http.Error(w, "Failed to parse resume", http.StatusInternalServerError)
			return
		}
		profile.UserID = user.ID

		if err := db.Create(&profile).Error; err != nil {
			log.Printf("Failed to save profile to database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(profile)
	}
}
