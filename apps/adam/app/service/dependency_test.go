package service

import (
	"coffeewithegg/apps/adam/project/service"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golobby/container/v3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ = Describe("App Dependencies", func() {
	var mock sqlmock.Sqlmock
	var gdb *gorm.DB

	BeforeEach(func() {
		var db *sql.DB
		var err error

		db, mock, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())

		gdb, err = gorm.Open(postgres.New(
			postgres.Config{
				Conn: db,
			},
		))
		Expect(err).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		err := mock.ExpectationsWereMet()
		Expect(err).ShouldNot(HaveOccurred())
	})

	When("DB is initialized", func() {
		BeforeEach(func() {
			err := container.Singleton(func() *gorm.DB {
				return gdb
			})
			Expect(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			container.Reset()
		})

		It("should implement project module", func() {
			err := InitServicesContainer()

			Expect(err).ShouldNot(HaveOccurred())

			By("Implementing Technology Service")
			var techService *service.TechnologyService
			err = container.Resolve(&techService)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(techService).ToNot(BeNil())

			By("Implementing Project Service")
			var projectService *service.ProjectService
			err = container.Resolve(&projectService)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(projectService).ToNot(BeNil())
		})
	})
})
