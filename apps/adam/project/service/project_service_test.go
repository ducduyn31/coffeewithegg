package service

import (
	"coffeewithegg/apps/adam/graph/model"
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"time"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type Any struct{}

func (a Any) Match(_ driver.Value) bool {
	return true
}

var _ = Describe("ProjectService", func() {
	var projectService *ProjectService
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var db *sql.DB
		var err error

		db, mock, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())

		gdb, err := gorm.Open(postgres.New(
			postgres.Config{
				Conn: db,
			},
		))
		Expect(err).ShouldNot(HaveOccurred())

		projectService = &ProjectService{
			db: gdb,
			techService: &TechnologyService{
				db: gdb,
			},
		}
	})

	AfterEach(func() {
		err := mock.ExpectationsWereMet()
		Expect(err).ShouldNot(HaveOccurred())
	})

	When("GetProjects", func() {

		It("should list all projects and their technologies", func() {
			// Given
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "projects" WHERE "projects"."deleted_at" IS NULL LIMIT 25`)).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "key", "description"}).
						AddRow("1", "Test Project", "test", "").
						AddRow("2", "Test Project", "test", "").
						AddRow("3", "Test Project", "test", "").
						AddRow("4", "Test Project", "test", "").
						AddRow("5", "Test Project", "test", "").
						AddRow("6", "Test Project", "test", "").
						AddRow("7", "Test Project", "test", "").
						AddRow("8", "Test Project", "test", "").
						AddRow("9", "Test Project", "test", "").
						AddRow("10", "Test Project", "test", ""),
				)
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "project_technologies"`)).
				WithArgs(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "key", "description"}),
				)

			filter := model.ProjectFilter{}

			// When
			projects, err := projectService.GetProjects(nil, &filter)
			if err != nil {
				return
			}

			// Then
			Expect(err).ShouldNot(HaveOccurred())
			Expect(projects).Should(HaveLen(10))
		})

		It("should list maximum of 5 projects given filter count=5", func() {
			// Given
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "projects" WHERE "projects"."deleted_at" IS NULL LIMIT 5`)).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "key", "description"}).
						AddRow("1", "Test Project", "test", "").
						AddRow("2", "Test Project", "test", "").
						AddRow("3", "Test Project", "test", "").
						AddRow("4", "Test Project", "test", "").
						AddRow("5", "Test Project", "test", ""),
				)
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "project_technologies"`)).
				WithArgs(1, 2, 3, 4, 5).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "key", "description"}),
				)

			var count = 5
			filter := model.ProjectFilter{
				Count: &count,
			}

			// When
			projects, err := projectService.GetProjects(nil, &filter)
			if err != nil {
				return
			}

			// Then
			Expect(err).ShouldNot(HaveOccurred())
			Expect(projects).Should(HaveLen(5))
		})

		It("should list the nex 5 projects given filter count=5 and offset=5", func() {
			// Given
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "projects" WHERE "projects"."deleted_at" IS NULL LIMIT 5 OFFSET 5`)).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "key", "description"}).
						AddRow("6", "Test Project", "test", "").
						AddRow("7", "Test Project", "test", "").
						AddRow("8", "Test Project", "test", "").
						AddRow("9", "Test Project", "test", "").
						AddRow("10", "Test Project", "test", ""),
				)
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "project_technologies"`)).
				WithArgs(6, 7, 8, 9, 10).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "key", "description"}),
				)

			var count = 5
			var offset = 5
			filter := model.ProjectFilter{
				Count:  &count,
				Offset: &offset,
			}

			// When
			projects, err := projectService.GetProjects(nil, &filter)
			if err != nil {
				return
			}

			// Then
			Expect(err).ShouldNot(HaveOccurred())
			Expect(projects).Should(HaveLen(5))
		})
	})

	When("UpsertProjects", func() {

		It("should throw error if invalid input is provided", func() {
			// Given
			var input *model.ProjectInput

			// When
			project, err := projectService.UpsertProject(nil, input)

			// Then
			Expect(err).Should(HaveOccurred())
			Expect(project).To(BeNil())
		})

		It("should throw error if id is provided but project does not exist", func() {
			// Given
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "projects" WHERE "projects"."id" = $1 AND "projects"."deleted_at" IS NULL ORDER BY "projects"."id" LIMIT 1`)).
				WithArgs("1").
				WillReturnError(gorm.ErrRecordNotFound)
			var id = "1"
			input := &model.ProjectInput{
				ID: &id,
			}

			// When
			project, err := projectService.UpsertProject(nil, input)

			// Then
			Expect(err).Should(HaveOccurred())
			Expect(project).To(BeNil())
		})

		It("should create a new project if no existing project is found", func() {
			// Given
			mock.ExpectBegin()
			mock.
				ExpectQuery(regexp.QuoteMeta(`INSERT INTO "projects" ("created_at","updated_at","deleted_at","name","key","description") VALUES ($1,$2,$3,$4,$5,$6) ON CONFLICT ("id") DO UPDATE SET "updated_at"=$7,"deleted_at"="excluded"."deleted_at","name"="excluded"."name","key"="excluded"."key","description"="excluded"."description" RETURNING "id"`)).
				WithArgs(AnyTime{}, AnyTime{}, Any{}, "name", "key", "description", AnyTime{}).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "key", "description"}).
						AddRow("1", "name", "key", "description"),
				)
			mock.ExpectCommit()

			var name = "name"
			var key = "key"
			var description = "description"
			input := &model.ProjectInput{
				Name:        &name,
				Key:         &key,
				Description: &description,
			}

			// When
			project, err := projectService.UpsertProject(nil, input)

			// Then
			Expect(err).ShouldNot(HaveOccurred())
			Expect(project).To(HaveField("Key", "key"))
		})

		It("should not create project and rollback if insert fails", func() {
			// Given
			mock.ExpectBegin()
			mock.
				ExpectQuery(regexp.QuoteMeta(`INSERT INTO "projects" ("created_at","updated_at","deleted_at","name","key","description") VALUES ($1,$2,$3,$4,$5,$6) ON CONFLICT ("id") DO UPDATE SET "updated_at"=$7,"deleted_at"="excluded"."deleted_at","name"="excluded"."name","key"="excluded"."key","description"="excluded"."description" RETURNING "id"`)).
				WithArgs(AnyTime{}, AnyTime{}, Any{}, "name", "key", "description", AnyTime{}).
				WillReturnError(errors.New("failed to insert data"))
			mock.ExpectRollback()

			var name = "name"
			var key = "key"
			var description = "description"
			input := &model.ProjectInput{
				Name:        &name,
				Key:         &key,
				Description: &description,
			}

			// When
			project, err := projectService.UpsertProject(nil, input)

			// Then
			Expect(err).Should(HaveOccurred())
			Expect(project).To(BeNil())
		})
	})

	When("UpsertProjects with technologies", func() {
		var techInput *model.TechnologyInput
		var projInput *model.ProjectInput

		BeforeEach(func() {
			var projKey = "key"
			var projName = "name"
			var projDescription = "description"

			projInput = &model.ProjectInput{
				Key:          &projKey,
				Name:         &projName,
				Description:  &projDescription,
				Technologies: nil,
			}

			mock.ExpectBegin()

		})

		AfterEach(func() {
			techInput = nil
		})

		When("technologies are expected to exist", func() {
			BeforeEach(func() {
				var myUuid = "910f362e-4824-4c3d-a95d-e1d9e5f30f4a"
				techInput = &model.TechnologyInput{
					ID: &myUuid,
				}

				projInput.Technologies = []*model.TechnologyInput{techInput}
			})

			It("should fail if uuid is invalid", func() {
				// Given
				var project *model.Project
				var err error
				techInput.ID = new(string)

				// When
				defer func() {
					if r := recover(); r != nil {
						err = errors.New("test panic")

						Expect(project).To(BeNil())
						Expect(err).Should(HaveOccurred())
					}
				}()
				project, err = projectService.UpsertProject(nil, projInput)

			})

			It("should create a new technology if ID is not given", func() {
				// Given
				var project *model.Project
				var err error

				var techKey = "techKey"
				var techName = "tech name"
				var techDesc = "tech description"

				techInput.ID = nil
				techInput.Key = &techKey
				techInput.Name = &techName
				techInput.Description = &techDesc

				mock.
					ExpectExec("SAVEPOINT").
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.
					ExpectQuery(`INSERT INTO "technologies"`).
					WithArgs(AnyTime{}, AnyTime{}, Any{}, "tech name", "tech description", "techKey", AnyTime{}).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "name", "key", "description"}).
							AddRow("910f362e-4824-4c3d-a95d-e1d9e5f30f4a", "tech name", "techKey", "tech description"),
					)
				mock.
					ExpectQuery(regexp.QuoteMeta(`INSERT INTO "projects" ("created_at","updated_at","deleted_at","name","key","description") VALUES ($1,$2,$3,$4,$5,$6) ON CONFLICT ("id") DO UPDATE SET "updated_at"=$7,"deleted_at"="excluded"."deleted_at","name"="excluded"."name","key"="excluded"."key","description"="excluded"."description" RETURNING "id"`)).
					WithArgs(AnyTime{}, AnyTime{}, Any{}, "name", "key", "description", AnyTime{}).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "name", "key", "description"}).
							AddRow("1", "name", "key", "description"),
					)
				mock.ExpectCommit()

				// When
				project, err = projectService.UpsertProject(nil, projInput)

				// Then
				Expect(project).ToNot(BeNil())
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("should rollback technology if creating technology fails", func() {
				// Given
				var project *model.Project
				var err error

				var techKey = "techKey"
				var techName = "tech name"
				var techDesc = "tech description"

				techInput.ID = nil
				techInput.Key = &techKey
				techInput.Name = &techName
				techInput.Description = &techDesc

				mock.
					ExpectExec("SAVEPOINT").
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.
					ExpectQuery(`INSERT INTO "technologies"`).
					WithArgs(AnyTime{}, AnyTime{}, Any{}, "tech name", "tech description", "techKey", AnyTime{}).
					WillReturnError(errors.New("failed to insert data"))
				mock.ExpectExec("ROLLBACK TO SAVEPOINT").
					WillReturnResult(sqlmock.NewResult(1, 1))

				// When
				project, err = projectService.UpsertProject(nil, projInput)

				// Then
				Expect(project).To(BeNil())
				Expect(err).Should(HaveOccurred())
			})

		})
	})

})
