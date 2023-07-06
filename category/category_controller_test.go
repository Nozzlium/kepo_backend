package category

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/tools"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCategory(t *testing.T) {

	mockCall := mockCategories()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:2637/category", nil)
	ctx := tools.GetMockClaimContext(request.Context())
	recorder := httptest.NewRecorder()

	categoryController.Get(recorder, request.WithContext(ctx), nil)

	resp := response.CategoryWebResponse{}
	res := recorder.Result()
	dec := json.NewDecoder(res.Body)
	dec.Decode(&resp)

	assert.Equal(t, http.StatusOK, resp.Code)
	for i, category := range resp.Data {
		entity := expectedCategories[i]
		assert.Equal(t, entity.ID, category.ID)
	}

	mockCall.Unset()
}
