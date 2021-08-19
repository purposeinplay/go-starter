package ports

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pborman/uuid"
	"github.com/purposeinplay/go-starter/internal/domain"
	"github.com/purposeinplay/go-starter/internal/domain/user"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func (ts *StarterTestSuite) TestStarter_UserList() {
	req, _ := http.NewRequest("GET", "/users", nil)
	req.Header.Set("Content-Type", "application/json")

	u := &user.User{
		Email: "my@email.com",
	}
	createUser := ts.db.Create(u)
	require.NoError(ts.T(), createUser.Error)

	response := httptest.NewRecorder()

	ts.handler.ServeHTTP(response, req)

	expected := fmt.Sprintf(`{
		"users": [{
			"id": "%v",
			"email": "%v"
		}]
	}`, u.ID, u.Email)

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body), expected)

	ts.Assert().Equal(response.Code, http.StatusOK)
	ts.Assert().JSONEq(string(body), expected)
}

func (ts *StarterTestSuite) TestStarter_FindUser() {
	req, _ := http.NewRequest("GET", "/user", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI5NzU5MDMwMS1iNWJjLTRkM2YtOWRmNS0zNjdhOWMzYjVjMmQiLCJuYW1lIjoiSm9obiBEb2UiLCJpYXQiOjE1MTYyMzkwMjJ9._IqeaDNs7Px9j6SRsLP0mEleObv5lL_zeQammoEj72g"))

	u := &user.User{
		Base: domain.Base{
			ID: uuid.Parse("97590301-b5bc-4d3f-9df5-367a9c3b5c2d"),
		},
		Email: "my@email.com",
	}
	createUser := ts.db.Create(u)
	require.NoError(ts.T(), createUser.Error)

	response := httptest.NewRecorder()

	ts.handler.ServeHTTP(response, req)

	expected := fmt.Sprintf(`{
		"user": {
			"id": "%v",
			"email": "%v"
		}
	}`, u.ID, u.Email)

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))

	ts.Assert().Equal(response.Code, http.StatusOK)
	ts.Assert().JSONEq(string(body), expected)
}

func (ts *StarterTestSuite) TestStarter_FindUserNotAuthenticated() {
	req, _ := http.NewRequest("GET", "/user", nil)
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	ts.handler.ServeHTTP(response, req)

	ts.Assert().Equal(response.Code, http.StatusUnauthorized)
}

func (ts *StarterTestSuite) TestStarterAPI_CreateUser() {
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(map[string]interface{}{
		"email": "test@example.com",
	})
	ts.NoError(err)

	req, _ := http.NewRequest("POST", "/users", &b)
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	ts.handler.ServeHTTP(response, req)

	body, _ := ioutil.ReadAll(response.Body)

	var createUserRes CreateUserRes
	err = json.Unmarshal(body, &createUserRes)
	ts.NoError(err)

	ts.Assert().Equal(response.Code, http.StatusOK)
	ts.Assert().Equal("test@example.com", createUserRes.User.Email)
}


func (ts *StarterTestSuite) TestStarterAPI_CreateUserWithValidationError() {
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(map[string]interface{}{
		"email": "",
	})
	ts.NoError(err)

	req, _ := http.NewRequest("POST", "/users", &b)
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	ts.handler.ServeHTTP(response, req)

	body, _ := ioutil.ReadAll(response.Body)

	var createUserRes CreateUserRes
	err = json.Unmarshal(body, &createUserRes)
	ts.NoError(err)

	ts.Assert().Equal(http.StatusUnprocessableEntity, response.Code)
}
