package entities

type Job struct {
	Uuid string `json:"uuid"`
	Title string `json:"title"`
	NormalizedJobTitle string `json:"normalized_job_title"`
	ParentUuid string `json:"parent_uuid"`
}

type TypeNumberFilter struct {
	Even string
	Odd  string
}