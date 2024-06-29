package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"recruitment-system/models"
	"recruitment-system/middleware"
)

func CreateJobHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(middleware.UserContextKey).(*models.User)

		if user.UserType != "Admin" {
			http.Error(w, "Only admins can create job openings", http.StatusForbidden)
			return
		}

		var job models.Job
		if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		job.PostedBy = user.ID

		if err := db.Create(&job).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(job)
	}
}

func GetJobHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(middleware.UserContextKey).(*models.User)

		if user.UserType != "Admin" {
			http.Error(w, "Only admins can view job details", http.StatusForbidden)
			return
		}

		vars := mux.Vars(r)
		jobID, err := strconv.Atoi(vars["job_id"])
		if err != nil {
			http.Error(w, "Invalid job ID", http.StatusBadRequest)
			return
		}

		var job models.Job
		if err := db.First(&job, jobID).Error; err != nil {
			http.Error(w, "Job not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(job)
	}
}

func ListJobsHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var jobs []models.Job
		if err := db.Find(&jobs).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(jobs)
	}
}

func ApplyJobHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(middleware.UserContextKey).(*models.User)

		if user.UserType != "Applicant" {
			http.Error(w, "Only applicants can apply for jobs", http.StatusForbidden)
			return
		}

		jobID, err := strconv.Atoi(r.URL.Query().Get("job_id"))
		if err != nil {
			http.Error(w, "Invalid job ID", http.StatusBadRequest)
			return
		}

		application := models.Application{
			JobID:       uint(jobID),
			ApplicantID: user.ID,
		}

		if err := db.Create(&application).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(application)
	}
}
