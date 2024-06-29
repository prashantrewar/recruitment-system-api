package main

import (
	"fmt"
	"log"
	"net/http"
	"recruitment-system/config"
	"recruitment-system/db"
	"recruitment-system/handlers"
	"recruitment-system/middleware"
	"recruitment-system/models"
	"recruitment-system/utils"

	"github.com/gorilla/mux"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("could not load config: %v", err)
    }

    // Initialize database
    database, err := db.InitDB(cfg.DB.DSN)
    if err != nil {
        log.Fatalf("could not connect to the database: %v", err)
    }

	adminUser := models.User{
		ID:       1,
		UserType: "Admin",
		Email:    "admin@example.com",
	}

	secretKey := "your_secret_key"
	token, err := utils.GenerateJWT(adminUser, secretKey)
	if err != nil {
		fmt.Printf("Error generating token: %v\n", err)
		return
	}

	fmt.Printf("Generated Token: %s\n", token)
	
    // Set up routes and middleware
    r := mux.NewRouter()

    // Public routes
    r.HandleFunc("/signup", handlers.SignUpHandler(database)).Methods("POST")
    r.HandleFunc("/login", handlers.LoginHandler(database, cfg.JWT.SecretKey)).Methods("POST")

    // Protected routes
    authRouter := r.PathPrefix("/").Subrouter()
    authRouter.Use(middleware.AuthMiddleware(cfg.JWT.SecretKey))

    applicantRouter := authRouter.PathPrefix("/").Subrouter()
    applicantRouter.Use(middleware.ApplicantMiddleware)
    applicantRouter.HandleFunc("/uploadResume", handlers.UploadResumeHandler(database, cfg.API.ResumeParserAPIKey)).Methods("POST")
    applicantRouter.HandleFunc("/jobs/apply", handlers.ApplyJobHandler(database)).Methods("POST")

    adminRouter := authRouter.PathPrefix("/admin").Subrouter()
    adminRouter.Use(middleware.AdminMiddleware)
    adminRouter.HandleFunc("/job", handlers.CreateJobHandler(database)).Methods("POST")
    adminRouter.HandleFunc("/job/{job_id}", handlers.GetJobHandler(database)).Methods("GET")
    adminRouter.HandleFunc("/applicants", handlers.ListApplicantsHandler(database)).Methods("GET")
    adminRouter.HandleFunc("/applicant/{applicant_id}", handlers.GetApplicantHandler(database)).Methods("GET")

    authRouter.HandleFunc("/jobs", handlers.ListJobsHandler(database)).Methods("GET")

    // Start server
    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
