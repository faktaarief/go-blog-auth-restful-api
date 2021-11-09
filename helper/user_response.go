package helper

import "github.com/faktaarief/go-blog-auth-restful-api/model/domain"

type userResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserResponses(users []domain.User) []userResponse {
	var responses []userResponse

	for _, user := range users {
		response := userResponse{
			ID:       user.Id,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}

		responses = append(responses, response)
	}

	return responses
}
