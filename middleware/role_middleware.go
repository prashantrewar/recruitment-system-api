package middleware

import (
    "net/http"
    "recruitment-system/models"
)

func AdminMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user, ok := r.Context().Value(UserContextKey).(*models.User)
        if !ok || user.UserType != "Admin" {
            http.Error(w, "Only admins are allowed to access this resource", http.StatusForbidden)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func ApplicantMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user, ok := r.Context().Value(UserContextKey).(*models.User)
        if !ok || user.UserType != "Applicant" {
            http.Error(w, "Only applicants are allowed to access this resource", http.StatusForbidden)
            return
        }
        next.ServeHTTP(w, r)
    })
}
