package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pointfive/internal/import_job/domain"
)

type ImportJobHandler struct {
	importJobService *domain.ServiceImpl
}

func NewHandler(importJobService *domain.ServiceImpl) *ImportJobHandler {
	h := ImportJobHandler{importJobService: importJobService}
	return &h
}

func (c *ImportJobHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(100 << 20)
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Missing file field", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileInfo, err := c.importJobService.ImportFile(ctx, file)

	if err != nil {
		fmt.Printf("error: %s\n", err)
		http.Error(w, "no data", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fileInfo)
	w.WriteHeader(http.StatusOK)

	return

}
