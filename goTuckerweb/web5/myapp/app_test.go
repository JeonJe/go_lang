package myapp

import (
	"encoding/json"
	"fmt"
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
	assert.Contains(string(data), "No Users")
}

func TestUsers_WithUsersData(t *testing.T){
	assert := assert.New(t)
	//테스트 용도 핸들러가 나옴 
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()
	
	resp, err := http.Post(ts.URL + "/users", "application/json",
	strings.NewReader(`{"first_name":"whssodi","last_name":"last","email":"whssodi@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	resp, err = http.Post(ts.URL + "/users", "application/json",
	strings.NewReader(`{"first_name":"tucker","last_name":"last2","email":"tucker@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	
}

func TestGetUserInfo(t *testing.T){
	assert := assert.New(t)
	//테스트 용도 핸들러가 나옴 
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode) 
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID:89")
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

func TestDeleteUser(t *testing.T){
	assert := assert.New(t)
	//테스트 용도 핸들러가 나옴 
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	//delte는 기본적으로 제공하는 메소드 아님 
	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID:")
}

func TestUpdateUser(t *testing.T){
	assert := assert.New(t)
	//테스트 용도 핸들러가 나옴 
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest("PUT", ts.URL+"/users",
	strings.NewReader(`{"id":1, "first_name":"update"}`))

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	//유저 생성
	resp, err = http.Post(ts.URL + "/users", "application/json",
	strings.NewReader(`{"first_name":"whssodi","last_name":"last","email":"whssodi@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	updateStr := fmt.Sprintf(`{"id":%d, "first_name":"update"}`,user.ID)
	//PUT는 기본적으로 제공하는 메소드 아님 
	req, _ = http.NewRequest("PUT", ts.URL+"/users",
	strings.NewReader(updateStr))

	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	updateUser := new(User)
	err = json.NewDecoder(resp.Body).Decode(updateUser)
	assert.NoError(err)
	assert.Equal(user.ID, updateUser.ID)
	assert.Equal(updateUser.FirstName,"update")
	assert.Equal(user.LastName, updateUser.LastName)
	assert.Equal(user.Email, updateUser.Email)
}