package handlers

import (
	"encoding/json"
	"net/http"
	"gorm.io/gorm"
	"recruitment-system/models"
	"github.com/gorilla/mux"
	"strconv"
	"recruitment-system/middleware"
)

func ListApplicantsHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(middleware.UserContextKey).(*models.User)

		if user.UserType != "Admin" {
			http.Error(w, "Only admins can view applicants", http.StatusForbidden)
			return
		}

		var users []models.User
		if err := db.Where("user_type = ?", "Applicant").Find(&users).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}

func GetApplicantHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(middleware.UserContextKey).(*models.User)

		if user.UserType != "Admin" {
			http.Error(w, "Only admins can view applicant details", http.StatusForbidden)
			return
		}

		vars := mux.Vars(r)
		applicantID, err := strconv.Atoi(vars["applicant_id"])
		if err != nil {
			http.Error(w, "Invalid applicant ID", http.StatusBadRequest)
			return
		}

		var applicant models.User
		if err := db.First(&applicant, applicantID).Error; err != nil {
			http.Error(w, "Applicant not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(applicant)
	}
}
