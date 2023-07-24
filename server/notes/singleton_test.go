package notes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	manager := New()
	assert.NotNil(t, manager)
	assert.Empty(t, manager.notesByOrder)

	manager.AddNoteToOrder("order1", "note1")
	assert.NotEmpty(t, manager.notesByOrder)

	notesOrder1 := manager.GetNotesFromOrder("order1")
	assert.NotEmpty(t, notesOrder1)
	assert.Equal(t, 1, len(notesOrder1))
	assert.Equal(t, "note1", notesOrder1[0].Text)
	assert.NotEmpty(t, notesOrder1[0].Date)

	manager2 := New()
	notesOrder1fromOther := manager2.GetNotesFromOrder("order1")
	assert.Equal(t, notesOrder1, notesOrder1fromOther)
}
