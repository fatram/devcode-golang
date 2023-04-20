package todo

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/fatram/devcode-golang/config"
	"github.com/fatram/devcode-golang/domain/entity"
	"github.com/fatram/devcode-golang/domain/model"
	"github.com/fatram/devcode-golang/domain/repository"
	imysql "github.com/fatram/devcode-golang/domain/repository/mysql"
	"github.com/fatram/devcode-golang/internal/connector"
	"github.com/fatram/devcode-golang/internal/pkg"
	"github.com/fatram/devcode-golang/internal/pkg/beautify"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	e    = pkg.LoadEcho()
	todo = entity.Todo{
		ActivityID: -9999,
		Title:      "satu",
		Priority:   "very-low",
	}

	repo repository.TodoRepository
	db   *sql.DB
	ctx  = context.Background()
	id   = 0
	err  error
)

func setUp() {
	config.ReadConfig("../../../.test.env")
	db = connector.LoadMysqlDatabase()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}
	repo = imysql.LoadTodoRepository(e.Logger)
	id, err = repo.Create(ctx, todo)
	if err != nil {
		log.Fatal(err)
	}
}

func tearDown() {
	_, err := db.Exec("DELETE FROM todo WHERE activity_group_id IN (-9999, -9)")
	if err != nil {
		log.Println("failed deleting todo data")
	}
	db.Close()
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestTodoCreate(t *testing.T) {
	jsonData := `{
		"activity_group_id": -9,
		"title": "minus sembilan",
		"priority": "sangat minus sembilan"
	}`
	req := httptest.NewRequest(http.MethodPost, "/todo-items", strings.NewReader(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewTodoController(e)
	idModel := model.BaseResponse{}
	if assert.NoError(t, h.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		json.Unmarshal(rec.Body.Bytes(), &idModel)
		jsonString, _ := json.Marshal(idModel.Data)
		data := model.Activity{}
		json.Unmarshal(jsonString, &data)
		id := data.ID
		assert.NotEmpty(t, id)
		saved, _ := repo.Get(ctx, id)
		log.Printf("Saved entity: %s", beautify.JSONString(saved))
		assert.NotEmpty(t, saved.ID)
		assert.Equal(t, -9, saved.ActivityID)
		assert.Equal(t, "minus sembilan", saved.Title)
		assert.Equal(t, "sangat minus sembilan", saved.Priority)
		log.Printf("Result: %s", rec.Body.String())
	}
}

func TestGetAllTodo(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/todo-items", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewTodoController(e)

	if assert.NoError(t, h.GetAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := model.BaseResponse{}
		json.Unmarshal(rec.Body.Bytes(), &body)
		data := reflect.ValueOf(body.Data)
		assert.Greater(t, data.Len(), 0)
		log.Printf("Result: %s", rec.Body.String())
	}
}

func TestGetTodo(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/todo-items", nil)
	rec := httptest.NewRecorder()
	h := NewTodoController(e)
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))

	if assert.NoError(t, h.Get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := model.BaseResponse{}
		json.Unmarshal(rec.Body.Bytes(), &body)
		jsonString, _ := json.Marshal(body.Data)
		data := model.Activity{}
		json.Unmarshal(jsonString, &data)
		id := data.ID
		assert.NotEmpty(t, id)
		saved, _ := repo.Get(ctx, id)
		log.Printf("Saved entity: %s", beautify.JSONString(saved))
		assert.NotEmpty(t, saved.ID)
		assert.Equal(t, -9999, saved.ActivityID)
		assert.Equal(t, "satu", saved.Title)
		assert.Equal(t, "very-low", saved.Priority)
		log.Printf("Result: %s", rec.Body.String())
	}
}

func TestUpdateTodo(t *testing.T) {
	jsonData := `{
		"is_active": false,
		"title": "disalahkan",
		"priority": "sangat-salah"
	}`
	req := httptest.NewRequest(http.MethodPut, "/todo-items", strings.NewReader(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	h := NewTodoController(e)
	if assert.NoError(t, h.Update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		saved, _ := repo.Get(ctx, id)
		log.Printf("Saved entity: %s", beautify.JSONString(saved))
		assert.Equal(t, false, saved.IsActive)
		assert.Equal(t, "disalahkan", saved.Title)
		assert.Equal(t, "sangat-salah", saved.Priority)
		log.Printf("Result: %s", rec.Body.String())
	}
}

func TestDeleteTodo(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/todo-items", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	h := NewTodoController(e)
	if assert.NoError(t, h.Delete(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		saved, _ := repo.Get(ctx, id)
		assert.Nil(t, saved)
	}
}
