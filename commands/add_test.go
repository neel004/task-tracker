package commands

import (
	originalModel "github.com/neel004/task-tracker/models"
	originalStorage "github.com/neel004/task-tracker/storage"
	"testing"
)

type MockStorage struct {
	ReadFunc   func() ([]originalStorage.TaskItem, error)
	ReadCalled bool

	UpdateCalled bool
	UpdateFunc   func([]originalStorage.TaskItem) error
	UpdateCalls  [][]originalStorage.TaskItem
}

func (m *MockStorage) Read() ([]originalStorage.TaskItem, error) {
	m.ReadCalled = true
	if m.ReadFunc != nil {
		return m.ReadFunc()
	}
	return []originalStorage.TaskItem{}, nil
}
func (m *MockStorage) Update(items []originalStorage.TaskItem) error {
	m.UpdateCalled = true
	m.UpdateCalls = append(m.UpdateCalls, items)
	if m.UpdateFunc != nil {
		return m.UpdateFunc(items)
	}
	return nil
}

func TestAddBase(t *testing.T) {
	mockStorage := &MockStorage{}
	args := []string{"something", "mockey"}
	err := Add(mockStorage, args...)
	if err != nil {
		t.Fatalf("expected Add to succeed, got error: %v", err)
	}
	if !mockStorage.ReadCalled {
		t.Fatalf("expected to call at least Read once, never called.")
	}

	if len(mockStorage.UpdateCalls) != 1 || !mockStorage.UpdateCalled {
		t.Fatalf("expected to call update once, got %d", len(mockStorage.UpdateCalls))
	}

	updatedItems := mockStorage.UpdateCalls[0]
	if len(updatedItems) != 1 {
		t.Fatalf("expected 1 item to be added, got %d", len(updatedItems))
	}

	task := updatedItems[0]
	if task.Description != "something mockey" {
		t.Errorf("expected description 'something mockey', got '%s'", task.Description)
	}
}

func TestAddToExisting(t *testing.T) {
	mockStorage := &MockStorage{
		ReadFunc: func() ([]originalStorage.TaskItem, error) {
			return []originalStorage.TaskItem{
				{
					Id:          1,
					Description: "existing note",
					Status:      originalModel.TODO,
				},
			}, nil
		},
	}
	args := []string{"something", "mockey"}
	err := Add(mockStorage, args...)
	if err != nil {
		t.Fatalf("expected Add to succeed, got error: %v", err)
	}
	if !mockStorage.ReadCalled {
		t.Fatalf("expected to call at least Read once, never called.")
	}

	if len(mockStorage.UpdateCalls) != 1 || !mockStorage.UpdateCalled {
		t.Fatalf("expected to call update once, got %d", len(mockStorage.UpdateCalls))
	}
	updatedItems := mockStorage.UpdateCalls[0]
	if len(updatedItems) != 2 {
		t.Fatalf("expected total 2 items, got %d", len(updatedItems))
	}

	task := updatedItems[0]
	if task.Description != "existing note" {
		t.Errorf("expected description 'existing note', got '%s'", task.Description)
	}

	task = updatedItems[1]
	if task.Description != "something mockey" {
		t.Errorf("expected description 'something mockey', got '%s'", task.Description)
	}

	tests := []struct {
		name                string
		readFunc            func() ([]originalStorage.TaskItem, error)
		readCalled          bool
		updateFunc          func([]originalStorage.TaskItem) error
		updateCalled        bool
		expectedTotal       int
		expectedId          int
		expectedDescription string
	}{
		{name: "Add to empty", readFunc: func() ([]originalStorage.TaskItem, error) { return []originalStorage.TaskItem{}, nil }, readCalled: true, updateFunc: nil, updateCalled: true, expectedTotal: 1, expectedId: 1, expectedDescription: "Add to empty"},
		{name: "Add to non empty", readFunc: func() ([]originalStorage.TaskItem, error) {
			return []originalStorage.TaskItem{{Id: 1, Description: "existing", Status: originalModel.TODO}}, nil
		}, readCalled: true, updateFunc: nil, updateCalled: true, expectedTotal: 2, expectedId: 2, expectedDescription: "Add to non empty"},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				mockStorage := &MockStorage{
					ReadFunc:     tt.readFunc,
					UpdateCalled: false,
					ReadCalled:   false,
					UpdateFunc:   tt.updateFunc,
				}
				args := []string{tt.name}

				err := Add(mockStorage, args...)
				if err != nil {
					t.Fatalf("expected Add to succeed, got error: %v", err)
				}
				if mockStorage.ReadCalled != tt.readCalled {
					t.Fatalf("expected to call at least Read once, never called.")
				}
			},
		)
	}

}
