package activity

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
	"time"

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
	timeNow  = time.Now().Unix()
	e        = pkg.LoadEcho()
	activity = entity.Activity{
		Title:     "satu",
		Email:     "satu@satu.satu",
		CreatedAt: int(timeNow),
		UpdatedAt: int(timeNow),
	}

	repo repository.ActivityRepository
	db   *sql.DB
	ctx  = context.Background()
	id   = 0
	err  error
)

func setUp() {
	config.ReadConfig("../../../.test.env")
	db = connector.LoadMysqlDatabase()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}
	repo = imysql.LoadActivityRepository(e.Logger)
	id, err = repo.Create(ctx, activity)
	if err != nil {
		log.Fatal(err)
	}
}

func tearDown() {
	_, err := db.Exec("DELETE FROM activity WHERE email IN ('satu@satu.satu', 'dua@dua.dua')")
	if err != nil {
		log.Println("failed deleting activity data")
	}
	db.Close()
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestActivityCreate(t *testing.T) {
	jsonData := `{
		"title": "baru",
		"email": "dua@dua.dua"
	}`
	req := httptest.NewRequest(http.MethodPost, "/activity-groups", strings.NewReader(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewActivityController(e)
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
		assert.Equal(t, "baru", saved.Title)
		log.Printf("Result: %s", rec.Body.String())
	}
}

func TestGetAllActivity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/activity-groups", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewActivityController(e)

	if assert.NoError(t, h.GetAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := model.BaseResponse{}
		json.Unmarshal(rec.Body.Bytes(), &body)
		data := reflect.ValueOf(body.Data)
		assert.Greater(t, data.Len(), 0)
		log.Printf("Result: %s", rec.Body.String())
	}
}

func TestGetActivity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/activity-groups", nil)
	rec := httptest.NewRecorder()
	h := NewActivityController(e)
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
		assert.Equal(t, "satu", saved.Title)
		assert.Equal(t, "satu@satu.satu", saved.Email)
		log.Printf("Result: %s", rec.Body.String())
	}
}

func TestUpdateActivity(t *testing.T) {
	jsonData := `{
		"title": "disalahkan"
	}`
	req := httptest.NewRequest(http.MethodPatch, "/activity-groups", strings.NewReader(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	h := NewActivityController(e)
	if assert.NoError(t, h.Update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		saved, _ := repo.Get(ctx, id)
		log.Printf("Saved entity: %s", beautify.JSONString(saved))
		assert.Equal(t, "disalahkan", saved.Title)
		log.Printf("Result: %s", rec.Body.String())
	}
}

func TestDeleteActivity(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/activity-groups", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	h := NewActivityController(e)
	if assert.NoError(t, h.Delete(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		saved, _ := repo.Get(ctx, id)
		assert.Nil(t, saved)
	}
}
