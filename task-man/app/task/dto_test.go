package task

import (
	"github.com/google/uuid"
	"testing"
)

func TestChangeTasksColumnDTOAddTaskId(t *testing.T) {
	changeTasksColumnDTO := ChangeTasksColumnDTO{}
	if len(changeTasksColumnDTO.TasksIDs) > 0 {
		t.Errorf("ChangeTasksColumnDTO.AddTaskId: expected DTO with empty tasks ids. Got with len %v", len(changeTasksColumnDTO.TasksIDs))
	}
	firstTaskId := uuid.New()
	secondTaskId := uuid.New()
	changeTasksColumnDTO.AddTaskId(firstTaskId)
	if len(changeTasksColumnDTO.TasksIDs) != 1 {
		t.Errorf("ChangeTasksColumnDTO.AddTaskId: expected DTO with len 1. Got with len %v", len(changeTasksColumnDTO.TasksIDs))
	}
	changeTasksColumnDTO.AddTaskId(firstTaskId)
	changeTasksColumnDTO.AddTaskId(secondTaskId)
	if len(changeTasksColumnDTO.TasksIDs) != 3 {
		t.Errorf("ChangeTasksColumnDTO.AddTaskId: expected DTO with len 3. Got with len %v", len(changeTasksColumnDTO.TasksIDs))
	}
	lastTaskId := changeTasksColumnDTO.TasksIDs[len(changeTasksColumnDTO.TasksIDs)-1]
	if lastTaskId != secondTaskId {
		t.Errorf("ChangeTasksColumnDTO.AddTaskId: expected last task id in slace with value %s. Got with value %s", secondTaskId, lastTaskId)
	}
}

func TestDeleteTasksDTOAddTaskId(t *testing.T) {
	deleteTasksDTO := DeleteTasksDTO{}
	if len(deleteTasksDTO.TasksIDs) > 0 {
		t.Errorf("DeleteTasksDTO.AddTaskId: expected DTO with empty tasks ids. Got with len %v", len(deleteTasksDTO.TasksIDs))
	}
	firstTaskId := uuid.New()
	secondTaskId := uuid.New()
	deleteTasksDTO.AddTaskId(firstTaskId)
	deleteTasksDTO.AddTaskId(secondTaskId)
	if len(deleteTasksDTO.TasksIDs) != 2 {
		t.Errorf("DeleteTasksDTO.AddTaskId: expected DTO with len 2. Got with len %v", len(deleteTasksDTO.TasksIDs))
	}
	deleteTasksDTO.AddTaskId(firstTaskId)
	deleteTasksDTO.AddTaskId(firstTaskId)
	deleteTasksDTO.AddTaskId(secondTaskId)
	deleteTasksDTO.AddTaskId(firstTaskId)
	if len(deleteTasksDTO.TasksIDs) != 6 {
		t.Errorf("DeleteTasksDTO.AddTaskId: expected DTO with len 6. Got with len %v", len(deleteTasksDTO.TasksIDs))
	}
	lastTaskId := deleteTasksDTO.TasksIDs[len(deleteTasksDTO.TasksIDs)-1]
	if lastTaskId != firstTaskId {
		t.Errorf("DeleteTasksDTO.AddTaskId: expected last task id in slace with value %s. Got with value %s", firstTaskId, lastTaskId)
	}
}
