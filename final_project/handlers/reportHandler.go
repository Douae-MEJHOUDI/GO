package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ReportHandler struct {
	*Handler
	outputDir string
}

type reportExtractDate struct {
	GeneratedAt time.Time `json:"generated_at"`
}

func NewReportHandler(handler *Handler, outputDir string) *ReportHandler {
	return &ReportHandler{
		Handler:   handler,
		outputDir: outputDir,
	}
}

func (h *ReportHandler) ListReports(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	var startTime, endTime time.Time
	var err error

	if startDate != "" {
		startTime, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			h.JsonWriteResponse(w, r, http.StatusBadRequest, "Invalid start_date format. Use YYYY-MM-DD")
			return
		}
	}

	if endDate != "" {
		endTime, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			h.JsonWriteResponse(w, r, http.StatusBadRequest, "Invalid end_date format. Use YYYY-MM-DD")
			return
		}
	}

	files, err := os.ReadDir(h.outputDir)
	if err != nil {
		h.JsonWriteResponse(w, r, http.StatusInternalServerError, "Failed to read reports directory")
		return
	}

	var reports []json.RawMessage
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		data, err := os.ReadFile(filepath.Join(h.outputDir, file.Name()))
		if err != nil {
			continue
		}

		if startDate != "" || endDate != "" {
			var report reportExtractDate
			if err := json.Unmarshal(data, &report); err != nil {
				continue
			}
			if startDate != "" && report.GeneratedAt.Before(startTime) {
				continue
			}
			if endDate != "" && report.GeneratedAt.After(endTime) {
				continue
			}
		}

		reports = append(reports, json.RawMessage(data))
	}

	h.JsonWriteResponse(w, r, http.StatusOK, reports)
}

func (h *ReportHandler) GetReport(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		h.JsonWriteResponse(w, r, http.StatusBadRequest, "Invalid URL")
		return
	}

	reportID := paths[2]
	filename := fmt.Sprintf("report_%s.json", reportID)

	data, err := os.ReadFile(filepath.Join(h.outputDir, filename))
	if err != nil {
		if os.IsNotExist(err) {
			h.JsonWriteResponse(w, r, http.StatusNotFound, "Report not found")
		} else {
			h.JsonWriteResponse(w, r, http.StatusInternalServerError, "Failed to read report")
		}
		return
	}

	h.JsonWriteResponse(w, r, http.StatusOK, json.RawMessage(data))
}

func (h *ReportHandler) ReportsRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.withTimeout(10*time.Second, h.ListReports)(w, r)
	default:
		h.JsonWriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *ReportHandler) ReportRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.withTimeout(10*time.Second, h.GetReport)(w, r)
	default:
		h.JsonWriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
