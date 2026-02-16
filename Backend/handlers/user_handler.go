package handlers

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"gatordash-backend/models"
	"gatordash-backend/store"
	"gatordash-backend/utils"

	"github.com/gin-gonic/gin"
)

// UserHandler handles all user-related HTTP requests
type UserHandler struct {
	userStore *store.UserStore
}

// NewUserHandler creates a new user handler
func NewUserHandler(userStore *store.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

// SignUp handles user registration
func (h *UserHandler) SignUp(c *gin.Context) {
	var req models.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, 400, "Invalid request body")
		return
	}

	// Validate required fields
	if req.Name == "" || req.Email == "" || req.Phone == "" || req.Password == "" {
		utils.SendError(c, 400, "All fields are required")
		return
	}

	// Check if user already exists
	if _, exists := h.userStore.GetUserByEmail(req.Email); exists {
		utils.SendError(c, 409, "User with this email already exists")
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendError(c, 500, "Failed to hash password")
		return
	}

	// Create user
	now := time.Now()
	user := &models.User{
		ID:        fmt.Sprintf("user_%d", now.UnixNano()),
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  string(hashedPassword),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := h.userStore.CreateUser(user); err != nil {
		utils.SendError(c, 500, err.Error())
		return
	}

	// Return user without password
	userResponse := *user
	userResponse.Password = ""

	utils.SendSuccess(c, "User registered successfully", userResponse)
}

// SignIn handles user authentication
func (h *UserHandler) SignIn(c *gin.Context) {
	var req models.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, 400, "Invalid request body")
		return
	}

	// Validate required fields
	if req.Email == "" || req.Password == "" {
		utils.SendError(c, 400, "Email and password are required")
		return
	}

	// Get user by email
	user, exists := h.userStore.GetUserByEmail(req.Email)
	if !exists {
		utils.SendError(c, 401, "Invalid email or password")
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.SendError(c, 401, "Invalid email or password")
		return
	}

	// Return user without password
	userResponse := *user
	userResponse.Password = ""

	utils.SendSuccess(c, "Sign in successful", userResponse)
}

// GetUser retrieves a user by ID
func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	if userID == "" {
		utils.SendError(c, 400, "User ID is required")
		return
	}

	user, exists := h.userStore.GetUserByID(userID)
	if !exists {
		utils.SendError(c, 404, "User not found")
		return
	}

	// Return user without password
	userResponse := *user
	userResponse.Password = ""

	utils.SendSuccess(c, "User retrieved successfully", userResponse)
}