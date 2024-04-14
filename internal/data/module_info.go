package data

import "time"

type ModuleInfo struct {
	ID             int       `json:"id"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	ModuleName     string    `json:"module_name"`
	ModuleDuration int       `json:"module_duration"`
	ExamType       string    `json:"exam_type"`
	Version        int       `json:"version"`
}
