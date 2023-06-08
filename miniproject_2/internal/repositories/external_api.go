package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"miniproject_2/internal/models"
)

type ExternalAPIRepository struct {
	// Add any necessary configuration or dependencies here
}

func NewExternalAPIRepository() *ExternalAPIRepository {
	return &ExternalAPIRepository{}
}

func (r *ExternalAPIRepository) FetchUserData(page int) ([]*models.User, error) {
	url := fmt.Sprintf("https://reqres.in/api/users?page=%d", page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response struct {
		Data []*models.User `json:"data"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
