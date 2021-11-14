package helper

import (
	"time"

	"github.com/faktaarief/go-blog-auth-restful-api/model/domain"
)

type userResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UserResponses(users []domain.User) []userResponse {
	var responses []userResponse

	for _, user := range users {
		response := userResponse{
			ID:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		responses = append(responses, response)
	}

	return responses
}

func UserResponse(user domain.User) userResponse {
	response := userResponse{
		ID:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response
}
