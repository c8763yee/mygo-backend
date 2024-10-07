package data

import (
	"fmt"
	"strconv"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

const dsn = "ithome:ironman@tcp(localhost:3306)/mygo?charset=utf8&parseTime=True&loc=Local"

func CreateDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	if err := db.AutoMigrate(&SentenceItem{}); err != nil {
		panic(fmt.Sprintf("error migrating schema: %v", err))
		// return nil, fmt.Errorf("error migrating schema: %v", err)
	}
	return db
}

func getMaxConnections() (int, error) {
	var value string
	row := Database.Raw("SHOW VARIABLES LIKE 'max_connections'").Row()
	err := row.Scan(new(string), &value)
	if err != nil {
		return 0, err
	}
	maxConnections, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return maxConnections, nil
}

func coorInsert(wg *sync.WaitGroup, semaphore chan struct{}, item SentenceItem, db *gorm.DB, messageChan chan<- string) {
	defer wg.Done()
	semaphore <- struct{}{}
	defer func() { <-semaphore }()

	// Check if the record already exists
	var existingItem SentenceItem
	err := db.First(&existingItem, "segment_id = ?", item.SegmentId).Error
	if err == nil {
		messageChan <- fmt.Sprintf("Segment ID(%d) already exists", item.SegmentId)
		return
	} else if err != gorm.ErrRecordNotFound {
		messageChan <- fmt.Sprintf("Error checking Segment ID(%d): %v", item.SegmentId, err)
		return
	}

	// Insert the new record
	if err := db.Create(&item).Error; err != nil {
		messageChan <- fmt.Sprintf("Error inserting Segment ID(%d): %v", item.SegmentId, err)
		return
	}
	messageChan <- fmt.Sprintf("Successfully inserted Segment ID %d", item.SegmentId)
}

func insertOrUpdate(sentenceData *Sentence, Database *gorm.DB, messageChan chan string, wg *sync.WaitGroup) error {
	maxConnections, err := getMaxConnections()
	if err != nil {
		return fmt.Errorf("failed to get max_connections: %v", err)
	}
	reservedConnections := 20 // reserve 20 connections for the application
	maxConnections -= reservedConnections
	if maxConnections <= 0 {
		return fmt.Errorf("no available connections for the application")
	}

	sqlDB, err := Database.DB()
	if err != nil {
		return fmt.Errorf("failed to get DB instance: %v", err)
	}
	sqlDB.SetMaxOpenConns(maxConnections)
	sqlDB.SetMaxIdleConns(maxConnections)
	sqlDB.SetConnMaxLifetime(0)

	semaphore := make(chan struct{}, maxConnections) // setup a semaphore to limit concurrency
	for _, item := range sentenceData.Result {
		wg.Add(1)
		go coorInsert(wg, semaphore, item, Database, messageChan)
	}

	// wait until all insertItem goroutines are done, then close channel
	go func() {
		wg.Wait()
		close(messageChan)
	}()

	// dump all messages to stdout
	for message := range messageChan {
		fmt.Println(message)
	}
	return nil
}

func init() {
	messageChan := make(chan string)
	var wg sync.WaitGroup

	Database = CreateDB()

	sentenceData := GetDataFromFile()
	err := insertOrUpdate(&sentenceData, Database, messageChan, &wg)
	if err != nil {
		panic(err)
	}
}

func SearchBySegmentID(segmentID int) (*SentenceItem, error) {
	var item SentenceItem
	err := Database.First(&item, "segment_id = ?", segmentID).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func SearchByText(text, episode string, pagedBy, page int) ([]SentenceItem, int64, error) {
	var items []SentenceItem
	var count int64
	sentenceQuery := Database.Where("text LIKE ?", "%"+text+"%")
	countQuery := Database.Model(&SentenceItem{}).Where("text LIKE ?", "%"+text+"%")
	if episode != "" {
		sentenceQuery = sentenceQuery.Where("episode = ?", episode)
		countQuery = countQuery.Where("episode = ?", episode)
	}
	if err := sentenceQuery.Order("segment_id ASC").
		Offset((page - 1) * pagedBy).Limit(pagedBy).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return items, count, nil
}
