package user

import (
	"fmt"
	"github.com/oakeshq/go-starter/config"
	"github.com/oakeshq/go-starter/internal/user/storage"
	"github.com/oakeshq/go-starter/pkg/router"
	"github.com/oakeshq/go-starter/pkg/storage/dialer"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
	db      *gorm.DB
	router      *router.Router
}

func TestCollection(t *testing.T) {
	cfg, err := config.LoadTestConfig("../../config/config.test.yaml")

	if err != nil {
		logrus.Fatalf("Unable to read config %v", err)
	}

	db, err := dialer.Connect(cfg)
	require.NoError(t, err)

	r := router.NewRouter()

	RegisterHandlers(r, db, cfg)

	ts := &UserTestSuite{
		db:      db,
		router:      r,
	}

	suite.Run(t, ts)
}

func (ts *UserTestSuite) SetupTest() {
	ts.db.Exec("TRUNCATE TABLE users")
}

func (ts *UserTestSuite) TestCollection_UserList() {
	req, _ := http.NewRequest("GET", "/users", nil)
	req.Header.Set("Content-Type", "application/json")

	user := &storage.User{
		Email: "my@email.com",
	}

	createUser := ts.db.Create(user)
	require.NoError(ts.T(), createUser.Error)

	response := httptest.NewRecorder()

	ts.router.ServeHTTP(response, req)

	expected := fmt.Sprintf(`[
		{
			"email": "my@email.com"
		}
	]`)

	body, _ := ioutil.ReadAll(response.Body)

	ts.Assert().Equal(response.Code, http.StatusOK)
	ts.Assert().JSONEq(string(body), expected)
}

func (ts *UserTestSuite) TestCollection_UserListEmpty() {
	req, _ := http.NewRequest("GET", "/users", nil)
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	ts.router.ServeHTTP(response, req)

	expected := `[]`

	body, _ := ioutil.ReadAll(response.Body)

	ts.Assert().Equal(response.Code, http.StatusOK)
	ts.Assert().JSONEq(string(body), expected)
}
