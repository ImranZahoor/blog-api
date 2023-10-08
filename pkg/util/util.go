package util

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ImranZahoor/blog-api/internal/models"
)

func JsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
func ToUUID(id string) (models.Uuid, error) {

	intId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return models.Uuid(intId), nil
}
