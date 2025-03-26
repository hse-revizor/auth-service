package di

import (
	trmgorm "github.com/avito-tech/go-transaction-manager/gorm"
	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/hse-revizor/auth-service/internal/pkg/service/auth"
	"github.com/hse-revizor/auth-service/internal/pkg/storage/sql"
	"github.com/hse-revizor/auth-service/internal/utils/config"
	"gorm.io/gorm"
)

type Container struct {
	cfg *config.Config

	storage *sql.Storage

	db                 *gorm.DB
	transactionManager trm.Manager
	authService        *auth.Service
}

func New(cfg *config.Config) *Container {
	return &Container{cfg: cfg}
}

func (c *Container) GetDB() *gorm.DB {
	return get(&c.db, func() *gorm.DB {
		return sql.MustNewPostgresDB(c.cfg)
	})
}

func (c *Container) GetSQLStorage() *sql.Storage {
	return get(&c.storage, func() *sql.Storage {
		return sql.New(c.GetDB(), trmgorm.DefaultCtxGetter)
	})
}

func (c *Container) GetTransactionManager() trm.Manager {
	return get(&c.transactionManager, func() trm.Manager {
		return manager.Must(trmgorm.NewDefaultFactory(c.GetDB()))
	})
}

func (c *Container) GetAuthService() *auth.Service {
	return get(&c.authService, func() *auth.Service {
		return auth.New(c.GetSQLStorage())
	})
}

func get[T comparable](obj *T, builder func() T) T {
	if *obj != *new(T) {
		return *obj
	}

	*obj = builder()
	return *obj
}
