package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"social/api/routes"
	"testing"
)

// Test for the post /users request
func TestPostUser(t *testing.T){
	var jsonStr = []byte(`{"_id":"569ed8269353e9f4c51617aa","name":"test","email":"test@gmail.com","password":"testPassword"}`)
	req,err := http.NewRequest("POST","/users",bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type","application/json")
	rr:= httptest.NewRecorder()
	handler := http.HandlerFunc(routes.UserHandler);
	handler.ServeHTTP(rr,req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"message":"User Created"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	t.Log("PASS")
}

// Test for the post /posts request
func TestPostPosts(t *testing.T){
	var jsonStr = []byte(`{"caption":"test","imageurl":"http://test/test.jpeg","user_id":"4313144fa3"}`)
	req,err := http.NewRequest("POST","/posts",bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type","application/json")
	rr:= httptest.NewRecorder()
	handler := http.HandlerFunc(routes.PostHandler);
	handler.ServeHTTP(rr,req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"message":"Post Created"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

// Test for the get /posts request
func TestGetPosts(t *testing.T){
	id := "5d8f8f8f8f8f8f8f8f8f8f8"
	url:= "/posts"+id
	req,err := http.NewRequest("GET",url,nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type","application/json")
	rr:= httptest.NewRecorder()
	handler := http.HandlerFunc(routes.GetPostHandler);
	handler.ServeHTTP(rr,req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"_id":"5d8f8f8f8f8f8f8f8f8f8f8","caption":"test","imageurl":"https://test/test.jpeg"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// Test for the get /users request
func TestGetUser(t *testing.T){
	id := "5d8f8f8f8f8f8f3f8f8f8f8"
	url:= "/users"+id
	req,err := http.NewRequest("GET",url,nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type","application/json")
	rr:= httptest.NewRecorder()
	handler := http.HandlerFunc(routes.GetUserHandler);
	handler.ServeHTTP(rr,req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"_id":"5d8f8f8f8f8f8f3f8f8f8f8","name":"test","email":"test@gmail.com","password":""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	t.Log("PASS")
}

// Test for the get /posts/users request
func TestGetUserPosts(t *testing.T){
	id := "61616530c1471ce279c6fdd6"
	url:= "/posts/users/"+id
	body:= []byte(`{"Page":4,"Limit":3}`)
	req,err := http.NewRequest("GET",url,bytes.NewBuffer(body));
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type","application/json")
	rr:= httptest.NewRecorder()
	handler := http.HandlerFunc(routes.GetUserPostsHandler);
	handler.ServeHTTP(rr,req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	
	if rr.Body.Len()>3{
		t.Errorf("handler returned More then the page limit")
	}
	t.Log("PASS")
}