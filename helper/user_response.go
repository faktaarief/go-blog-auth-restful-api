package helper

import "github.com/faktaarief/go-blog-auth-restful-api/model/domain"

type userResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserResponses(users []domain.User) []userResponse {
	var responses []userResponse

	for _, user := range users {
		response := userResponse{
			ID:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}

		responses = append(responses, response)
	}

	return responses
}

func UserResponse(user domain.User) userResponse {
	response := userResponse{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	return response
}
