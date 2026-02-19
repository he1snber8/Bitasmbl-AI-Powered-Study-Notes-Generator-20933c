package domain

type NoteRequest struct{ Content string `json:"content"` }
type Keyword struct{ Term string `json:"term"` }
type NoteResponse struct{ Summary string `json:"summary"`; Keywords []Keyword `json:"keywords"` }