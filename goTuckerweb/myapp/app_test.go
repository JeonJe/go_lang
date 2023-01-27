package myapp

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T){
	assert := assert.New(t)
 	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	indexHandler(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello, world!",string(data))
}

func TestBarPathHandler_withoutName(t *testing.T){
	assert := assert.New(t)
 	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	//request마다 path를 라우팅해주도록 mux 사용 
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello, World!", string(data))
}

func TestBarPathHandler_withName(t *testing.T){
	assert := assert.New(t)
 	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=whssodi", nil)

	//request마다 path를 라우팅해주도록 mux 사용 
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello, whssodi!", string(data))
}
func TestFooPathHandler_without_JSON(t *testing.T){
	assert := assert.New(t)
 	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	//request마다 path를 라우팅해주도록 mux 사용 
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}
func TestFooPathHandler_with_JSON(t *testing.T){
	assert := assert.New(t)
 	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo",strings.NewReader(`{"first_name":"whssodi","last_name":"lee","email":"whssodi@gmail.com"}`))

	//request마다 path를 라우팅해주도록 mux 사용 
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("whssodi", user.FristName)
	assert.Equal("lee", user.LastName)
}
