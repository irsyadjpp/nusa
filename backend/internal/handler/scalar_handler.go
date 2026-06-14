package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ScalarHandler handles Scalar documentation UI requests
type ScalarHandler struct{}

// NewScalarHandler creates a new scalar handler
func NewScalarHandler() *ScalarHandler {
	return &ScalarHandler{}
}

// ServeScalar serves the Scalar API documentation UI
func (h *ScalarHandler) ServeScalar(c *gin.Context) {
	html := `<!DOCTYPE html>
<html>
  <head>
    <title>NUSA API Documentation</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
  </head>
  <body>
    <div id="app"></div>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
    <script>
      Scalar.createApiReference('#app', {
        url: window.location.origin + '/openapi.json',
        theme: 'default',
        darkMode: true,
        layout: 'modern',
        hideSidebar: false,
        showSidebar: true,
        defaultHttpClient: {
          targetKey: 'curl',
          client: 'fetch'
        }
      })
    </script>
  </body>
</html>`
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

// ServeSwaggerUI serves the Swagger UI as a fallback alternative to Scalar
func (h *ScalarHandler) ServeSwaggerUI(c *gin.Context) {
	html := `<!DOCTYPE html>
<html>
  <head>
    <title>NUSA API Documentation</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <link rel="stylesheet" type="text/css" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui.css" />
    <style>
      html { box-sizing: border-box; overflow: -moz-scrollbars-vertical; overflow-y: scroll; }
      *, *:before, *:after { box-sizing: inherit; }
      body { margin: 0; background: #fafafa; }
    </style>
  </head>
  <body>
    <div id="swagger-ui"></div>
    <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui-standalone-preset.js"></script>
    <script>
      window.onload = function() {
        const ui = SwaggerUIBundle({
          url: window.location.origin + '/openapi.json',
          dom_id: '#swagger-ui',
          deepLinking: true,
          presets: [
            SwaggerUIBundle.presets.apis,
            SwaggerUIBundle.StandalonePreset
          ],
          plugins: [
            SwaggerUIBundle.plugins.DownloadUrl
          ],
          layout: "StandaloneLayout",
          defaultModelsExpandDepth: 1,
          defaultModelExpandDepth: 1,
          docExpansion: "list"
        });
      }
    </script>
  </body>
</html>`
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

// ServeOpenAPISpec serves the OpenAPI specification
func (h *ScalarHandler) ServeOpenAPISpec(c *gin.Context) {
	// Generate comprehensive OpenAPI spec with all endpoints
	openAPISpec := map[string]interface{}{
		"openapi": "3.0.0",
		"info": map[string]interface{}{
			"title":       "NUSA Platform API",
			"description": "Education Management System for Kurikulum Merdeka 2026",
			"version":     "1.0.0",
			"contact": map[string]interface{}{
				"name":  "NUSA Platform Team",
				"email": "support@nusa.id",
			},
		},
		"servers": []map[string]interface{}{
			{
				"url":         "http://localhost:8081",
				"description": "Development server",
			},
		},
		"paths": map[string]interface{}{
			// System Endpoints
			"/health": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Health check endpoint",
					"description": "Check if the API is running",
					"tags":        []string{"System"},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Server is healthy",
						},
					},
				},
			},
			"/ready": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Readiness check endpoint",
					"description": "Check if the API is ready to handle requests",
					"tags":        []string{"System"},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Server is ready",
						},
					},
				},
			},
			"/version": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Version endpoint",
					"description": "Get API version information",
					"tags":        []string{"System"},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "API version information",
						},
					},
				},
			},

			// Public Auth Endpoints
			"/api/v1/public/auth/login": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "User login",
					"description": "Authenticate user and return JWT tokens",
					"tags":        []string{"Authentication"},
					"requestBody": map[string]interface{}{
						"required": true,
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type":     "object",
									"required": []string{"email", "password"},
									"properties": map[string]interface{}{
										"email": map[string]interface{}{
											"type":    "string",
											"format":  "email",
											"example": "admin@nusa.id",
										},
										"password": map[string]interface{}{
											"type":    "string",
											"format":  "password",
											"example": "admin123",
										},
									},
								},
							},
						},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Login successful",
						},
					},
				},
			},
			"/api/v1/public/auth/refresh": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Refresh access token",
					"description": "Refresh JWT access token using refresh token",
					"tags":        []string{"Authentication"},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Token refreshed successfully",
						},
					},
				},
			},

			// Protected Auth Endpoints
			"/api/v1/auth/logout": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "User logout",
					"description": "Logout user and invalidate refresh token",
					"tags":        []string{"Authentication"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Logout successful",
						},
					},
				},
			},
			"/api/v1/auth/me": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get current user",
					"description": "Get information about the currently authenticated user",
					"tags":        []string{"Authentication"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "User information",
						},
					},
				},
			},

			// Users Endpoints
			"/api/v1/users": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Create user",
					"description": "Create a new user in the system",
					"tags":        []string{"Users"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "User created successfully",
						},
					},
				},
				"get": map[string]interface{}{
					"summary":     "List users",
					"description": "Get all users in the system",
					"tags":        []string{"Users"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of users",
						},
					},
				},
			},
			"/api/v1/users/{id}": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get user by ID",
					"description": "Get user information by ID",
					"tags":        []string{"Users"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "User information",
						},
					},
				},
				"put": map[string]interface{}{
					"summary":     "Update user",
					"description": "Update user information",
					"tags":        []string{"Users"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "User updated successfully",
						},
					},
				},
				"patch": map[string]interface{}{
					"summary":     "Update user status",
					"description": "Update user status",
					"tags":        []string{"Users"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "User status updated successfully",
						},
					},
				},
			},

			// Schools Endpoints
			"/api/v1/schools": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Create school",
					"description": "Create a new school",
					"tags":        []string{"Schools"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "School created successfully",
						},
					},
				},
				"get": map[string]interface{}{
					"summary":     "List schools",
					"description": "Get all schools in the system",
					"tags":        []string{"Schools"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of schools",
						},
					},
				},
			},
			"/api/v1/schools/{id}": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get school by ID",
					"description": "Get school information by ID",
					"tags":        []string{"Schools"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "School information",
						},
					},
				},
				"put": map[string]interface{}{
					"summary":     "Update school",
					"description": "Update school information",
					"tags":        []string{"Schools"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "School updated successfully",
						},
					},
				},
				"patch": map[string]interface{}{
					"summary":     "Update school status",
					"description": "Update school status",
					"tags":        []string{"Schools"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "School status updated successfully",
						},
					},
				},
			},

			// Roles Endpoints
			"/api/v1/roles": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "List roles",
					"description": "Get all roles in the system",
					"tags":        []string{"Roles"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of roles",
						},
					},
				},
				"post": map[string]interface{}{
					"summary":     "Create role",
					"description": "Create a new role (SYSTEM_ADMIN only)",
					"tags":        []string{"Roles"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "Role created successfully",
						},
					},
				},
			},
			"/api/v1/roles/{id}": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get role by ID",
					"description": "Get role information by ID",
					"tags":        []string{"Roles"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Role information",
						},
					},
				},
				"put": map[string]interface{}{
					"summary":     "Update role",
					"description": "Update role information (SYSTEM_ADMIN only)",
					"tags":        []string{"Roles"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Role updated successfully",
						},
					},
				},
				"delete": map[string]interface{}{
					"summary":     "Delete role",
					"description": "Delete role (SYSTEM_ADMIN only)",
					"tags":        []string{"Roles"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"204": map[string]interface{}{
							"description": "Role deleted successfully",
						},
					},
				},
			},
			"/api/v1/roles/{id}/permissions": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get role permissions",
					"description": "Get permissions for a role",
					"tags":        []string{"Roles"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Role permissions",
						},
					},
				},
				"post": map[string]interface{}{
					"summary":     "Add permission to role",
					"description": "Add a permission to a role",
					"tags":        []string{"Roles"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Permission added successfully",
						},
					},
				},
				"delete": map[string]interface{}{
					"summary":     "Remove permission from role",
					"description": "Remove a permission from a role",
					"tags":        []string{"Roles"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Permission removed successfully",
						},
					},
				},
			},

			// Curriculum Endpoints
			"/api/v1/curriculum/subjects": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Create curriculum subject",
					"description": "Create a new curriculum subject",
					"tags":        []string{"Curriculum"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "Subject created successfully",
						},
					},
				},
				"get": map[string]interface{}{
					"summary":     "List curriculum subjects",
					"description": "Get all curriculum subjects",
					"tags":        []string{"Curriculum"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of subjects",
						},
					},
				},
			},
			"/api/v1/curriculum/subjects/{id}": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get curriculum subject",
					"description": "Get curriculum subject by ID",
					"tags":        []string{"Curriculum"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Subject information",
						},
					},
				},
				"put": map[string]interface{}{
					"summary":     "Update curriculum subject",
					"description": "Update curriculum subject",
					"tags":        []string{"Curriculum"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Subject updated successfully",
						},
					},
				},
				"delete": map[string]interface{}{
					"summary":     "Delete curriculum subject",
					"description": "Delete curriculum subject",
					"tags":        []string{"Curriculum"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"204": map[string]interface{}{
							"description": "Subject deleted successfully",
						},
					},
				},
			},

			"/api/v1/curriculum/phases": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Create curriculum phase",
					"description": "Create a new curriculum phase",
					"tags":        []string{"Curriculum"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "Phase created successfully",
						},
					},
				},
				"get": map[string]interface{}{
					"summary":     "List curriculum phases",
					"description": "Get all curriculum phases",
					"tags":        []string{"Curriculum"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of phases",
						},
					},
				},
			},
			"/api/v1/curriculum/phases/{id}": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get curriculum phase",
					"description": "Get curriculum phase by ID",
					"tags":        []string{"Curriculum"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Phase information",
						},
					},
				},
				"put": map[string]interface{}{
					"summary":     "Update curriculum phase",
					"description": "Update curriculum phase",
					"tags":        []string{"Curriculum"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Phase updated successfully",
						},
					},
				},
				"delete": map[string]interface{}{
					"summary":     "Delete curriculum phase",
					"description": "Delete curriculum phase",
					"tags":        []string{"Curriculum"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"204": map[string]interface{}{
							"description": "Phase deleted successfully",
						},
					},
				},
			},

			// Learning Planning Endpoints
			"/api/v1/learning-planning/tp-sets": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Create TP set",
					"description": "Create a new Tujuan Pembelajaran set",
					"tags":        []string{"Learning Planning"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "TP set created successfully",
						},
					},
				},
				"get": map[string]interface{}{
					"summary":     "List TP sets",
					"description": "Get all Tujuan Pembelajaran sets",
					"tags":        []string{"Learning Planning"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of TP sets",
						},
					},
				},
			},
			"/api/v1/learning-planning/tp-sets/{id}": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get TP set",
					"description": "Get Tujuan Pembelajaran set by ID",
					"tags":        []string{"Learning Planning"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "TP set information",
						},
					},
				},
				"put": map[string]interface{}{
					"summary":     "Update TP set",
					"description": "Update Tujuan Pembelajaran set",
					"tags":        []string{"Learning Planning"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "TP set updated successfully",
						},
					},
				},
			},
			"/api/v1/learning-planning/tp-sets/{id}/approve": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Approve TP set",
					"description": "Approve Tujuan Pembelajaran set",
					"tags":        []string{"Learning Planning"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "TP set approved successfully",
						},
					},
				},
			},
			"/api/v1/learning-planning/tp-sets/{id}/versions": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get TP set versions",
					"description": "Get version history for TP set",
					"tags":        []string{"Learning Planning"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "TP set versions",
						},
					},
				},
			},

			// Assessment Endpoints
			"/api/v1/assessment": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Create assessment",
					"description": "Create a new assessment",
					"tags":        []string{"Assessment"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "Assessment created successfully",
						},
					},
				},
				"get": map[string]interface{}{
					"summary":     "List assessments",
					"description": "Get all assessments",
					"tags":        []string{"Assessment"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of assessments",
						},
					},
				},
			},
			"/api/v1/assessment/{id}": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get assessment",
					"description": "Get assessment by ID",
					"tags":        []string{"Assessment"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Assessment information",
						},
					},
				},
				"put": map[string]interface{}{
					"summary":     "Update assessment",
					"description": "Update assessment",
					"tags":        []string{"Assessment"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Assessment updated successfully",
						},
					},
				},
			},
			"/api/v1/assessment/{id}/approve": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Approve assessment",
					"description": "Approve assessment",
					"tags":        []string{"Assessment"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Assessment approved successfully",
						},
					},
				},
			},
			"/api/v1/assessment/rubrics": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Create rubric",
					"description": "Create a new assessment rubric",
					"tags":        []string{"Assessment"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "Rubric created successfully",
						},
					},
				},
				"get": map[string]interface{}{
					"summary":     "List rubrics",
					"description": "Get all assessment rubrics",
					"tags":        []string{"Assessment"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of rubrics",
						},
					},
				},
			},
			"/api/v1/assessment/rubrics/{id}": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get rubric",
					"description": "Get rubric by ID",
					"tags":        []string{"Assessment"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Rubric information",
						},
					},
				},
				"put": map[string]interface{}{
					"summary":     "Update rubric",
					"description": "Update rubric",
					"tags":        []string{"Assessment"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Rubric updated successfully",
						},
					},
				},
				"delete": map[string]interface{}{
					"summary":     "Delete rubric",
					"description": "Delete rubric",
					"tags":        []string{"Assessment"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"204": map[string]interface{}{
							"description": "Rubric deleted successfully",
						},
					},
				},
			},

			// Reporting Endpoints
			"/api/v1/reporting/narrative-reports": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Create narrative report",
					"description": "Create a new narrative report",
					"tags":        []string{"Reporting"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "Narrative report created successfully",
						},
					},
				},
				"get": map[string]interface{}{
					"summary":     "List narrative reports",
					"description": "Get all narrative reports",
					"tags":        []string{"Reporting"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of narrative reports",
						},
					},
				},
			},
			"/api/v1/reporting/narrative-reports/{id}": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get narrative report",
					"description": "Get narrative report by ID",
					"tags":        []string{"Reporting"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Narrative report information",
						},
					},
				},
				"put": map[string]interface{}{
					"summary":     "Update narrative report",
					"description": "Update narrative report",
					"tags":        []string{"Reporting"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Narrative report updated successfully",
						},
					},
				},
				"delete": map[string]interface{}{
					"summary":     "Delete narrative report",
					"description": "Delete narrative report",
					"tags":        []string{"Reporting"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"204": map[string]interface{}{
							"description": "Narrative report deleted successfully",
						},
					},
				},
			},
			"/api/v1/reporting/narrative-reports/{id}/refresh-achievement": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Refresh report achievement",
					"description": "Refresh achievement data for narrative report",
					"tags":        []string{"Reporting"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Achievement data refreshed successfully",
						},
					},
				},
			},

			// Achievement Endpoints
			"/api/v1/students/{id}/achievement": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get student achievement",
					"description": "Get achievement data for a student",
					"tags":        []string{"Achievement"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Student achievement data",
						},
					},
				},
			},
			"/api/v1/students/{id}/progress": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get student progress",
					"description": "Get competency progress for a student",
					"tags":        []string{"Achievement"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Student progress data",
						},
					},
				},
			},
			"/api/v1/classes/{id}/achievement": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get class achievement",
					"description": "Get achievement data for a class",
					"tags":        []string{"Achievement"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Class achievement data",
						},
					},
				},
			},

			// Academic Foundation Endpoints
			"/api/v1/academic-years": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Create academic year",
					"description": "Create a new academic year",
					"tags":        []string{"Academic Foundation"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "Academic year created successfully",
						},
					},
				},
				"get": map[string]interface{}{
					"summary":     "List academic years",
					"description": "Get all academic years",
					"tags":        []string{"Academic Foundation"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of academic years",
						},
					},
				},
			},
			"/api/v1/academic-years/{id}": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get academic year",
					"description": "Get academic year by ID",
					"tags":        []string{"Academic Foundation"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Academic year information",
						},
					},
				},
				"put": map[string]interface{}{
					"summary":     "Update academic year",
					"description": "Update academic year",
					"tags":        []string{"Academic Foundation"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Academic year updated successfully",
						},
					},
				},
			},
			"/api/v1/academic-years/{id}/activate": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Activate academic year",
					"description": "Activate an academic year",
					"tags":        []string{"Academic Foundation"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Academic year activated successfully",
						},
					},
				},
			},
			"/api/v1/academic-years/{id}/archive": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Archive academic year",
					"description": "Archive an academic year",
					"tags":        []string{"Academic Foundation"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Academic year archived successfully",
						},
					},
				},
			},

			"/api/v1/semesters": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Create semester",
					"description": "Create a new semester",
					"tags":        []string{"Academic Foundation"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "Semester created successfully",
						},
					},
				},
				"get": map[string]interface{}{
					"summary":     "List semesters",
					"description": "Get all semesters",
					"tags":        []string{"Academic Foundation"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of semesters",
						},
					},
				},
			},
			"/api/v1/semesters/{id}": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get semester",
					"description": "Get semester by ID",
					"tags":        []string{"Academic Foundation"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Semester information",
						},
					},
				},
				"put": map[string]interface{}{
					"summary":     "Update semester",
					"description": "Update semester",
					"tags":        []string{"Academic Foundation"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Semester updated successfully",
						},
					},
				},
				"delete": map[string]interface{}{
					"summary":     "Delete semester",
					"description": "Delete semester",
					"tags":        []string{"Academic Foundation"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"204": map[string]interface{}{
							"description": "Semester deleted successfully",
						},
					},
				},
			},

			"/api/v1/subject-categories": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Create subject category",
					"description": "Create a new subject category",
					"tags":        []string{"Curriculum Governance"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "Subject category created successfully",
						},
					},
				},
				"get": map[string]interface{}{
					"summary":     "List subject categories",
					"description": "Get all subject categories",
					"tags":        []string{"Curriculum Governance"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of subject categories",
						},
					},
				},
			},
			"/api/v1/subject-categories/{id}": map[string]interface{}{
				"put": map[string]interface{}{
					"summary":     "Update subject category",
					"description": "Update subject category",
					"tags":        []string{"Curriculum Governance"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Subject category updated successfully",
						},
					},
				},
				"delete": map[string]interface{}{
					"summary":     "Delete subject category",
					"description": "Delete subject category",
					"tags":        []string{"Curriculum Governance"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"204": map[string]interface{}{
							"description": "Subject category deleted successfully",
						},
					},
				},
			},

			"/api/v1/graduate-profile-dimensions": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Create graduate profile dimension",
					"description": "Create a new graduate profile dimension",
					"tags":        []string{"Curriculum Governance"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "Graduate profile dimension created successfully",
						},
					},
				},
				"get": map[string]interface{}{
					"summary":     "List graduate profile dimensions",
					"description": "Get all graduate profile dimensions",
					"tags":        []string{"Curriculum Governance"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of graduate profile dimensions",
						},
					},
				},
			},
			"/api/v1/graduate-profile-dimensions/{id}": map[string]interface{}{
				"put": map[string]interface{}{
					"summary":     "Update graduate profile dimension",
					"description": "Update graduate profile dimension",
					"tags":        []string{"Curriculum Governance"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Graduate profile dimension updated successfully",
						},
					},
				},
				"delete": map[string]interface{}{
					"summary":     "Delete graduate profile dimension",
					"description": "Delete graduate profile dimension",
					"tags":        []string{"Curriculum Governance"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"204": map[string]interface{}{
							"description": "Graduate profile dimension deleted successfully",
						},
					},
				},
			},

			"/api/v1/cp-alignments": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Create CP alignment",
					"description": "Create a new CP alignment",
					"tags":        []string{"Curriculum Governance"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "CP alignment created successfully",
						},
					},
				},
				"get": map[string]interface{}{
					"summary":     "List CP alignments",
					"description": "Get all CP alignments",
					"tags":        []string{"Curriculum Governance"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of CP alignments",
						},
					},
				},
			},
			"/api/v1/cp-alignments/bulk": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Bulk create CP alignments",
					"description": "Create multiple CP alignments",
					"tags":        []string{"Curriculum Governance"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "CP alignments created successfully",
						},
					},
				},
			},
			"/api/v1/cp-alignments/report": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Generate CP alignment report",
					"description": "Generate CP alignment report",
					"tags":        []string{"Curriculum Governance"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "CP alignment report generated",
						},
					},
				},
			},
			"/api/v1/cp-alignments/{id}": map[string]interface{}{
				"put": map[string]interface{}{
					"summary":     "Update CP alignment",
					"description": "Update CP alignment",
					"tags":        []string{"Curriculum Governance"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "CP alignment updated successfully",
						},
					},
				},
				"delete": map[string]interface{}{
					"summary":     "Delete CP alignment",
					"description": "Delete CP alignment",
					"tags":        []string{"Curriculum Governance"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"204": map[string]interface{}{
							"description": "CP alignment deleted successfully",
						},
					},
				},
			},

			"/api/v1/system-configurations": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Create system configuration",
					"description": "Create a new system configuration",
					"tags":        []string{"System Configuration"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "System configuration created successfully",
						},
					},
				},
				"get": map[string]interface{}{
					"summary":     "List system configurations",
					"description": "Get all system configurations",
					"tags":        []string{"System Configuration"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of system configurations",
						},
					},
				},
			},
			"/api/v1/system-configurations/{id}": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get system configuration",
					"description": "Get system configuration by ID",
					"tags":        []string{"System Configuration"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "System configuration information",
						},
					},
				},
				"put": map[string]interface{}{
					"summary":     "Update system configuration",
					"description": "Update system configuration",
					"tags":        []string{"System Configuration"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "System configuration updated successfully",
						},
					},
				},
				"delete": map[string]interface{}{
					"summary":     "Delete system configuration",
					"description": "Delete system configuration",
					"tags":        []string{"System Configuration"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"204": map[string]interface{}{
							"description": "System configuration deleted successfully",
						},
					},
				},
			},
			"/api/v1/system-configurations/by-key/{key}": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Get system configuration by key",
					"description": "Get system configuration by key",
					"tags":        []string{"System Configuration"},
					"security":    []interface{}{map[string]interface{}{"bearerAuth": []interface{}{}}},
					"parameters": []map[string]interface{}{
						{"name": "key", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "System configuration information",
						},
					},
				},
			},
		},
		"components": map[string]interface{}{
			"schemas": map[string]interface{}{
				"Error": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"message": map[string]interface{}{"type": "string"},
						"error":   map[string]interface{}{"type": "string"},
					},
				},
			},
			"securitySchemes": map[string]interface{}{
				"bearerAuth": map[string]interface{}{
					"type":         "http",
					"scheme":       "bearer",
					"bearerFormat": "JWT",
				},
			},
		},
	}
	c.JSON(http.StatusOK, openAPISpec)
}
