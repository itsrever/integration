package notes

import "time"

type Note struct {
	Text string
	Date time.Time
}

type Manager interface {
	// AddNoteToOrder adds a note to an order
	AddNoteToOrder(orderID, note string) error
	// GetNotesFromOrder returns all notes from an order
	GetNotesFromOrder(orderID string) ([]Note, error)
}

type manager struct {
	notesByOrder map[string][]Note
}

var singleton *manager = nil

// New creates a new instance of the notes manager
func New() *manager {
	if singleton == nil {
		singleton = &manager{
			notesByOrder: make(map[string][]Note),
		}
	}
	return singleton
}

func (m *manager) AddNoteToOrder(orderID, note string) {
	curr := m.GetNotesFromOrder(orderID)
	curr = append(curr, Note{
		Text: note,
		Date: time.Now(),
	})
	m.notesByOrder[orderID] = curr
}

func (m *manager) GetNotesFromOrder(orderID string) []Note {
	if notes, ok := m.notesByOrder[orderID]; ok {
		return notes
	}
	return []Note{}
}
