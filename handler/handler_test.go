package handler

import 
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sample/model"
	"sample/store"
	"testing"
	"gofr.dev/pkg/gofr/request"
	"gofr.dev/pkg/gofr"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// MockStore is a mock of store.Employee interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

func TestCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := store.NewMockEmployee(mockCtrl)
	h := New(mockStore)

	// Create a sample employee to use as request body
	emp := model.Employee{
		Name: "JohnDoe",
		Dept: "Engineering",
	}

	empJSON, _ := json.Marshal(emp)
	r := httptest.NewRequest(http.MethodPost, "/employee", bytes.NewBuffer(empJSON))
	req := request.NewHTTPRequest(r)
	ctx := gofr.NewContext(nil, req, nil)

	// Set up the expected behavior of the store
	mockStore.EXPECT().Create(ctx, &emp).Return(&emp, nil)

	// Call the Create method
	resp, err := h.Create(ctx)

	assert.Nil(t, err)
	assert.Equal(t, &emp, resp)
}

func TestGetByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := store.NewMockEmployee(mockCtrl)
	h := New(mockStore)

	// Set up a sample employee ID
	employeeID := 1
	emp := model.Employee{
		ID:   1,
		Name: "John Doe",
		Dept: "Engineering",
	}
	
	r := httptest.NewRequest(http.MethodGet, "/employee/1", nil)
	req := request.NewHTTPRequest(r)
	ctx := gofr.NewContext(nil, req, nil)
	

	// ctx.PathParam("0")

	// Set up the expected behavior of the store
	mockStore.EXPECT().GetByID(ctx, employeeID).Return(&emp, nil)

	// Call the GetByID method
	resp, err := h.GetByID(ctx, employeeID)

	assert.Nil(t, err)
	assert.Equal(t, &emp, resp)
}

func TestUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := store.NewMockEmployee(mockCtrl)
	h := New(mockStore)

	// Set up a sample employee ID and data
	employeeID := 1
	emp := model.Employee{
		ID:   employeeID,
		Name: "Jane Doe",
		Dept: "HR",
	}

	// Marshal the employee struct to JSON for the request body


	empJSON, _ := json.Marshal(emp)
	r := httptest.NewRequest(http.MethodPut, "/employee", bytes.NewBuffer(empJSON))
	req := request.NewHTTPRequest(r)
	ctx := gofr.NewContext(nil, req, nil)

	// Set up the expected behavior of the store
	mockStore.EXPECT().Update(ctx, &emp).Return(&emp, nil)

	// Call the Update method
	resp, err := h.Update(ctx, employeeID)

	assert.Nil(t, err)
	assert.Equal(t, &emp, resp)
}

func TestDelete(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    mockStore := store.NewMockEmployee(mockCtrl)
    h := New(mockStore)

    // Set up a sample employee ID
    employeeID := 1

    // Mock HTTP request
    r := httptest.NewRequest(http.MethodDelete, "/employee/1", nil)
    req := request.NewHTTPRequest(r)
    ctx := gofr.NewContext(nil, req, nil)

    // Set up the expected behavior of the store
    mockStore.EXPECT().Delete(ctx, employeeID).Return(nil).Times(1)

    // Call the Delete method and capture both return values
    _, err := h.Delete(ctx, employeeID)

    assert.Nil(t, err)
}