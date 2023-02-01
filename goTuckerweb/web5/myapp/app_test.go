package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T){
	assert := assert.New(t)
	//테스트 용도 핸들러가 나옴 
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode) 
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello, World!", string(data))
}

func TestUsers(t *testing.T){
	assert := assert.New(t)
	//테스트 용도 핸들러가 나옴 
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode) 
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "Get UserInfo")
}

func TestGetUserInfo(t *testing.T){
	assert := assert.New(t)
	//테스트 용도 핸들러가 나옴 
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/30")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode) 
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "User ID:30")
}

func TestCreateUser(t *testing.T){
	assert := assert.New(t)
	//테스트 용도 핸들러가 나옴 
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL + "/users", "application/json",
	strings.NewReader(`{"first_name":"whssodi","last_name":"last","email":"whssodi@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	id := user.ID 
	resp, err = http.Get(ts.URL + "/users/"+ strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode) 
	user2 := new(User)
	err = json.NewDecoder(resp.Body).Decode(user2)
	assert.NoError(err)

	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.FirstName, user2.FirstName)
	assert.Equal(user.LastName, user2.LastName)
	assert.Equal(user.Email, user2.Email)

}