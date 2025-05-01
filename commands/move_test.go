package commands

import (
	commands "github.com/neel004/task-tracker/commands"
	storage "github.com/neel004/task-tracker/storage"
	"testing"
)

type MockStorage struct {
	ReadFunc   func() ([]storage.TaskItem, error)
	ReadCalled bool

	UpdateCalled bool
	UpdateFunc   func([]storage.TaskItem) error
	UpdateCalls  [][]storage.TaskItem
}

func (m *MockStorage) Read() ([]storage.TaskItem, error) {
	m.ReadCalled = true
	if m.ReadFunc != nil {
		return m.ReadFunc()
	}
	return []storage.TaskItem{}, nil
}
func (m *MockStorage) Update(items []storage.TaskItem) error {
	m.UpdateCalled = true
	m.UpdateCalls = append(m.UpdateCalls, items)
	if m.UpdateFunc != nil {
		return m.UpdateFunc(items)
	}
	return nil
}

func TestMove(t *testing.T) {
	/*
		# valid move commmand
		# move with un-proper args
		# move to valid status
		# move to invalid status
	*/
	mockStorage := &MockStorage{
		ReadFunc: func() ([]storage.TaskItem, error) {
			return []storage.TaskItem{
				{Id: 1, Description: "SampleMockTask1"},
			}, nil
		},
	}
	args := []string{"1", "inprogress"}
	err := commands.MoveTo(mockStorage, args...)
	if err != nil {
		t.Errorf("did not expected error %d", err)
	}
	if !mockStorage.UpdateCalled {
		t.Error("expected to Update be called.")
	}
	// Few other correct test checks...
}
