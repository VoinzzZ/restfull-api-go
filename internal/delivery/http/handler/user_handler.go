package handler

import (
	"encoding/json"
	"net/http"
	"restfull-api-go/internal/domain"
	"restfull-api-go/internal/usecase"
	"restfull-api-go/pkg/helper"
	"strconv"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: uc,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		helper.ResponseError(w, http.StatusBadRequest, "All fields are required")
		return
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.userUsecase.CreateUser(user); err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	user.Password = ""
	helper.ResponseJSON(w, http.StatusCreated, "User created successfully", user)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	user, err := h.userUsecase.GetUserByID(id)
	if err != nil {
		helper.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}
	helper.ResponseJSON(w, http.StatusOK, "User found", user)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.ResponseJSON(w, http.StatusOK, "Users retrieved successfully", users)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	updateData := &domain.User{
		Name:  req.Name,
		Email: req.Email,
	}
	if err := h.userUsecase.UpdateUser(id, updateData); err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	helper.ResponseJSON(w, http.StatusOK, "User updated successfully", nil)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	if err := h.userUsecase.DeleteUser(id); err != nil {
		helper.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}
	helper.ResponseJSON(w, http.StatusOK, "User deleted successfully", nil)
}
