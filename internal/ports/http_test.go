package ports

import (
	"github.com/purposeinplay/go-commons/http/router"
	"github.com/purposeinplay/go-starter/internal/adapter"
	"github.com/purposeinplay/go-starter/internal/domain"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"

	"gorm.io/gorm"

	"github.com/purposeinplay/go-starter/internal/app"
	"github.com/purposeinplay/go-starter/internal/app/command"
	"github.com/purposeinplay/go-starter/internal/app/query"
	"github.com/purposeinplay/go-starter/internal/repository"

	"github.com/purposeinplay/go-commons/logs"

	"github.com/purposeinplay/go-starter/internal/config"
)


type StarterTestSuite struct {
	suite.Suite
	db     *gorm.DB
	httpPort ServerInterface
	handler http.Handler
}

func (ts *StarterTestSuite) SetupTest() {
	ts.db.Exec("TRUNCATE TABLE users CASCADE")
}

func TestStarterTestSuite(t *testing.T) {
	db, httpPort, handler := CreateTestAPI(t)

	ts := &StarterTestSuite{
		db:     db,
		httpPort: httpPort,
		handler: handler,
	}

	suite.Run(t, ts)
}

func CreateTestAPI(t *testing.T) (*gorm.DB, ServerInterface, http.Handler) {
	cfg, err := config.LoadTestConfig("../../config.test.yaml")

	if err != nil {
		t.Fatal(err)
	}

	logger, err := logs.NewLogger()
	if err != nil {
		t.Fatal("could not create logger %+v", err)
	}
	defer logger.Sync()

	db, err := adapter.Connect(cfg)

	if err != nil {
		t.Fatal(err)
	}

	repo := repository.NewUserRepository(db)
	validator := domain.NewValidator()
	translation := domain.NewTranslator(validator)

	app := app.Application{
		Commands: app.Commands{
			CreateUser: command.NewCreateUserHandler(logger, repo, validator, translation),
		},
		Queries:  app.Queries{
			FindUsers: query.NewFindUsersHandler(logger, repo),
			UserByID: query.NewUserById(logger, repo),
			UserByEmail: query.NewUserByEmail(logger, repo),
		},
	}

	httpPort := NewHTTPPort(app, cfg, db, logger)
	r := router.NewDefaultRouter(logger)
	HandlerFromMux(httpPort, r)

	return db, httpPort, r
}
