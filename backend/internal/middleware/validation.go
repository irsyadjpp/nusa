package middleware

import (
	"net/http"
	"regexp"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

// InputValidation middleware validates input for common security issues
func InputValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate Content-Type for POST/PUT/PATCH requests
		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
			contentType := c.GetHeader("Content-Type")
			if contentType != "" && !strings.HasPrefix(contentType, "application/json") && !strings.HasPrefix(contentType, "multipart/form-data") && !strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
				c.JSON(http.StatusUnsupportedMediaType, gin.H{
					"error": "Unsupported Content-Type",
				})
				c.Abort()
				return
			}
		}

		// Validate query parameters for SQL injection patterns
		for _, values := range c.Request.URL.Query() {
			for _, value := range values {
				if containsSQLInjection(value) {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": "Invalid input detected",
					})
					c.Abort()
					return
				}
			}
		}

		// Validate path parameters
		for _, param := range c.Params {
			if containsSQLInjection(param.Value) {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid input detected",
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

// containsSQLInjection checks for common SQL injection patterns
func containsSQLInjection(input string) bool {
	// Common SQL injection patterns
	patterns := []string{
		"(?i)(union|select|insert|update|delete|drop|alter|create|truncate|exec|execute)",
		"(?i)(or|and)\\s+\\d+\\s*=\\s*\\d+",
		"(?i)(or|and)\\s+['\"]?\\w+['\"]?\\s*=\\s*['\"]?\\w+['\"]?",
		"(?i)(--|;|\\/\\*|\\*\\/)",
		"(?i)(xp_|sp_)",
		"(?i)(waitfor|delay|sleep)",
	}

	for _, pattern := range patterns {
		matched, _ := regexp.MatchString(pattern, input)
		if matched {
			return true
		}
	}

	return false
}

// ValidateEmail validates email format
func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePassword validates password strength
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return &ValidationError{Field: "password", Message: "Password must be at least 8 characters long"}
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return &ValidationError{Field: "password", Message: "Password must contain at least one uppercase letter"}
	}
	if !hasLower {
		return &ValidationError{Field: "password", Message: "Password must contain at least one lowercase letter"}
	}
	if !hasNumber {
		return &ValidationError{Field: "password", Message: "Password must contain at least one number"}
	}
	if !hasSpecial {
		return &ValidationError{Field: "password", Message: "Password must contain at least one special character"}
	}

	return nil
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface
func (ve *ValidationError) Error() string {
	return ve.Field + ": " + ve.Message
}

// SanitizeInput removes potentially dangerous characters from input
func SanitizeInput(input string) string {
	// Remove null bytes
	input = strings.ReplaceAll(input, "\x00", "")

	// Remove control characters except newline, tab, carriage return
	var sanitized strings.Builder
	for _, r := range input {
		if r == '\n' || r == '\t' || r == '\r' || !unicode.IsControl(r) {
			sanitized.WriteRune(r)
		}
	}

	return sanitized.String()
}
