package middleware

import (
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
)

var SessionStore = sessions.NewCookieStore([]byte("my_super_secret_key"))

const SessionName = "react-go-spa-session"

func init() {
	SessionStore.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,  // Prevent JavaScript access
		Secure:   false, // Set to true in production (requires HTTPS)
		MaxAge:   3600,  // 1 hour
	}
}

func CorsMiddleware() func(http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	}).Handler
}

func SessionMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := SessionStore.Get(r, SessionName)
			if err != nil {
				http.Error(w, "Failed to get session", http.StatusInternalServerError)
				return
			}

			// Example: Set a default value in the session if it doesn't exist
			if session.Values["authenticated"] == nil {
				session.Values["authenticated"] = false
			}

			// Save the session
			if err := session.Save(r, w); err != nil {
				http.Error(w, "Failed to save session", http.StatusInternalServerError)
				return
			}

			// Proceed to the next handler
			next.ServeHTTP(w, r)
		})
	}
}

func LoggingMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			log.Info().
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Msg("Incoming request")

			next.ServeHTTP(w, r)

			log.Info().
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Dur("duration", time.Since(start)).
				Msg("Request processed")
		})
	}
}
