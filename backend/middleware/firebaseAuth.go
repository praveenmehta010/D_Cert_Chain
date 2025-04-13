package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/praveenmehta010/Decentralized-Cert-Verification/database"
)

// FirebaseAuthMiddleware verifies Firebase ID token
func FirebaseAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		idToken := strings.TrimPrefix(authHeader, "Bearer ")

		authClient, err := database.FirebaseApp.Auth(context.Background())
		if err != nil {
			http.Error(w, "Failed to init Firebase Auth client", http.StatusInternalServerError)
			return
		}

		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// You can access token.Claims["email"] or token.UID if needed
		// Optionally store it in context for the next handler
		// Basically to see the UID and email of the auth-user
		ctx := context.WithValue(r.Context(), "uid", token.UID)
		ctx = context.WithValue(ctx, "email", token.Claims["email"])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
