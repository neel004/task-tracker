package commands

import (
	"testing"

	originalStorage "github.com/neel004/task-tracker/storage"
)

type MockStorage struct{
	ReadFunc func() ([]originalStorage.TaskItem, error)
	ReadCalled bool
	
	UpdateCalled bool
	UpdateFunc func([]originalStorage.TaskItem) (error)
	UpdateCalls [][]originalStorage.TaskItem
}
func (m *MockStorage) Read() ([]originalStorage.TaskItem, error){
	m.ReadCalled = true
	if m.ReadFunc != nil {
		return m.ReadFunc()
	}
	return []originalStorage.TaskItem{}, nil
}
func (m *MockStorage) Update(items []originalStorage.TaskItem) error{
	m.UpdateCalled = true
	m.UpdateCalls = append(m.UpdateCalls, items)
	if m.UpdateFunc != nil {
		return m.UpdateFunc(items)
	}
	return nil
}

func TestAddBase(t *testing.T){
	mockStorage := MockStorage{}
	args := []string{"something", "mockey"}
	err := Add(&mockStorage, args...)
	if err != nil{
		t.Fatalf("expected Add to succeed, got error: %v", err)
	}
	if !mockStorage.ReadCalled{
		t.Fatalf("expected to call at least Read once, never called.")	
	}

	if len(mockStorage.UpdateCalls) != 1 || !mockStorage.UpdateCalled{
		t.Fatalf("expected to call update once, got %d", len(mockStorage.UpdateCalls))	
	}

	updatedItems := mockStorage.UpdateCalls[0]
	if len(updatedItems) != 1{
		t.Fatalf("expected 1 item to be added, got %d", len(updatedItems))
	}

	task := updatedItems[0]
	if task.Description != "something mockey"{
		t.Errorf("expected description 'buy groceries', got '%s'", task.Description)
	}
}
