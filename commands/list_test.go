package commands

import (
	"github.com/neel004/task-tracker/commands"
	"github.com/neel004/task-tracker/storage"
	"strconv"
	"strings"
	"testing"
)

type MockStorage struct {
	ReadFunc     func() ([]storage.TaskItem, error)
	UpdateFunc   func([]storage.TaskItem) error
	ReadCalled   bool
	UpdateCalled bool
	UpdateArgs   [][]storage.TaskItem
}

func (m *MockStorage) Read() ([]storage.TaskItem, error) {
	m.ReadCalled = true
	if m.ReadFunc != nil {
		return m.ReadFunc()
	}
	return nil, nil
}

func (m *MockStorage) Update(items []storage.TaskItem) error {
	m.UpdateCalled = true
	m.UpdateArgs = append(m.UpdateArgs, items)
	if m.UpdateFunc != nil {
		return m.UpdateFunc(items)
	}
	return nil
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name          string
		readFunc      func() ([]storage.TaskItem, error)
		updateFunc    func([]storage.TaskItem) error
		args          []string
		wantErr       bool
		errContains   string
		finalItemsLen int // only checked on success
	}{
		{
			name:        "No args",
			readFunc:    func() ([]storage.TaskItem, error) { return []storage.TaskItem{}, nil },
			args:        []string{},
			wantErr:     true,
			errContains: "id needs to be passed",
		},
		{
			name:        "Invalid ID parse",
			readFunc:    func() ([]storage.TaskItem, error) { return []storage.TaskItem{}, nil },
			args:        []string{"abc"},
			wantErr:     true,
			errContains: "converting input to valid type",
		},
		{
			name:        "ID not found",
			readFunc:    func() ([]storage.TaskItem, error) { return []storage.TaskItem{{Id: 2}}, nil },
			args:        []string{"1"},
			wantErr:     true,
			errContains: "item with queried id is not present",
		},
		{
			name: "Successful delete",
			readFunc: func() ([]storage.TaskItem, error) {
				return []storage.TaskItem{
					{Id: 1, Description: "foo"},
					{Id: 2, Description: "bar"},
				}, nil
			},
			args:          []string{"1"},
			wantErr:       false,
			finalItemsLen: 1,
		},
	}

	for _, tt := range tests {
		tt := tt // capture
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockStorage{
				ReadFunc:   tt.readFunc,
				UpdateFunc: tt.updateFunc,
			}

			err := commands.Delete(mock, tt.args...)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if !contains(err.Error(), tt.errContains) {
					t.Errorf("error %q does not contain %q", err.Error(), tt.errContains)
				}
				return
			}

			// success path
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if !mock.ReadCalled {
				t.Error("expected Read to be called")
			}
			if !mock.UpdateCalled {
				t.Error("expected Update to be called")
			}
			if len(mock.UpdateArgs) != 1 {
				t.Fatalf("expected one Update call, got %d", len(mock.UpdateArgs))
			}
			final := mock.UpdateArgs[0]
			if len(final) != tt.finalItemsLen {
				t.Errorf("expected %d items after delete, got %d", tt.finalItemsLen, len(final))
			}
			// ensure the remaining item's ID is not the deleted one
			for _, item := range final {
				id, _ := strconv.ParseUint(tt.args[0], 10, 16)
				if item.Id == uint16(id) {
					t.Errorf("found deleted ID %d in final items", id)
				}
			}
		})
	}
}

// helper for substring check
func contains(s, substr string) bool {
	return substr == "" || (s != "" && strings.Contains(s, substr))
}
