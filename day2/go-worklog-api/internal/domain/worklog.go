package domain

type WorklogStatus string

const (
	StatusDraft     WorklogStatus = "DRAFT"
	StatusSubmitted WorklogStatus = "SUBMITTED"
	StatusReviewed  WorklogStatus = "REVIEWED"
)

type Worklog struct {
	ID     int64
	Week   int
	Status WorklogStatus
	Rating int
}
