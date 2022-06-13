//go:generate mockgen -package=store -source=store.go -destination=./store_mock.go Store
package store

import (
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const ErrDuplicateKey = "23505"

type Storer interface {
	CreateArticleIfNotExists(request NewsArticle) error
	GetArticleByURL(url string) (NewsArticle, error)
	GetRecordsAfterID(ID string, numberOfRecords int, filters Filters) ([]NewsArticle, error)
}

type Store struct {
	db *gorm.DB
}

type NewsArticle struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Link        string `gorm:"unique_index:idx_link"`
	Category    string
	Thumbnail   string
	CreatedAt   time.Time
}

type Filters struct {
	Title       string
	Description string
	Link        string `gorm:"unique_index:idx_link"`
	Category    string
	Thumbnail   string
	CreatedAt   time.Time
}

func NewStore() (*Store, error) {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Println(err)
	}
	return &Store{
		db: db,
	}, nil
}

func (s *Store) CreateArticleIfNotExists(request NewsArticle) error {
	// there's probably a way to leverage FirstOrCreate instead of using db schemas to do this
	log.Println("store request", &request.Title)
	log.Println("creating record as not found")
	result := s.db.Create(&request)
	if result.Error != nil {
		var pgErr *pgconn.PgError
		// bit of a hack because of the above comment and this error isn't available from gorm as standard
		if errors.As(result.Error, &pgErr) && pgErr.Code == ErrDuplicateKey {
			log.Println("record already exists")
			return nil
		}
		return fmt.Errorf("unable to create record, %w", result.Error)
	}
	return nil
}

func (s *Store) GetRecordsAfterID(ID string, numberOfRecords int, filters Filters) ([]NewsArticle, error) {
	log.Println("get store request", ID, numberOfRecords)
	var FindResult []NewsArticle
	resp := s.db.Where("ID > ?", ID).Where(NewsArticle{
		Title:       filters.Title,
		Description: filters.Description,
		Link:        filters.Link,
		Category:    filters.Category,
		Thumbnail:   filters.Thumbnail,
	}).Order("ID asc").Limit(numberOfRecords).Find(&FindResult)
	if resp.Error != nil {
		return []NewsArticle{}, fmt.Errorf("failed to get record %e", resp.Error)
	}
	log.Println("find result", FindResult)
	return FindResult, nil
}

func (s *Store) GetArticleByURL(url string) (NewsArticle, error) {
	log.Println("get store request by url", url)
	var FindResult NewsArticle
	resp := s.db.Where(NewsArticle{Link: url}).First(&FindResult)
	if resp.Error != nil {
		return NewsArticle{}, fmt.Errorf("failed to get record %e", resp.Error)
	}
	log.Println("find result", FindResult)
	return FindResult, nil
}
