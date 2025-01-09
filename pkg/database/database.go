package database

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/c8763yee/mygo-backend/internal/config"
	"github.com/c8763yee/mygo-backend/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ImportDataFromFile(db *gorm.DB, filename string) error {
	ext := filepath.Ext(filename)
	switch ext {
	case ".json":
		return importDataFromJSON(db, filename)
	case ".csv":
		return importDataFromCSV(db, filename)
	default:
		return fmt.Errorf("unsupported file type: %s", ext)
	}
}

func importDataFromJSON(db *gorm.DB, filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %v", err)
	}

	var sentenceData struct {
		Result []models.SentenceItem `json:"result"`
	}

	if err := json.Unmarshal(file, &sentenceData); err != nil {
		return fmt.Errorf("failed to parse JSON data: %v", err)
	}

	return importItems(db, sentenceData.Result)
}

func importDataFromCSV(db *gorm.DB, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Skip header
	if _, err := reader.Read(); err != nil {
		return fmt.Errorf("failed to read CSV header: %v", err)
	}

	var items []models.SentenceItem

	for {
		record, err := reader.Read()
		if err != nil {
			break // EOF or error
		}
		if len(record) != 6 {
			return fmt.Errorf("invalid CSV format: expected 6 columns, but got %d\ndata: %+v", len(record), record)
		}
		// column: text,episode,video_name,segment_id,frame_start,frame_end
		segmentID, _ := strconv.ParseUint(record[3], 10, 32)
		frameStart, _ := strconv.ParseUint(record[4], 10, 32)
		frameEnd, _ := strconv.ParseUint(record[5], 10, 32)
		item := models.SentenceItem{
			Text:       record[0],
			Episode:    record[1],
			VideoName:  record[2],
			SegmentId:  uint(segmentID),
			FrameStart: uint(frameStart),
			FrameEnd:   uint(frameEnd),
		}

		items = append(items, item)
	}

	return importItems(db, items)
}

func importItems(db *gorm.DB, items []models.SentenceItem) error {
	var maxConnections int
	row := db.Raw("SHOW VARIABLES LIKE 'max_connections'").Row()
	if err := row.Scan(new(string), &maxConnections); err != nil {
		return fmt.Errorf("failed to get max_connections: %v", err)
	}

	workerCount := maxConnections - 10
	if workerCount <= 0 {
		return fmt.Errorf("insufficient available connections")
	}

	semaphore := make(chan struct{}, workerCount)
	var wg sync.WaitGroup

	for _, item := range items {
		wg.Add(1)
		go func(item models.SentenceItem) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			var existingItem models.SentenceItem
			err := db.First(&existingItem, "segment_id = ?", item.SegmentId).Error
			if err == nil {
				fmt.Printf("Segment ID(%d) already exists\n", item.SegmentId)
				return
			} else if err != gorm.ErrRecordNotFound {
				fmt.Printf("Error checking Segment ID(%d): %v\n", item.SegmentId, err)
				return
			}

			if err := db.Create(&item).Error; err != nil {
				fmt.Printf("Error inserting Segment ID(%d): %v\n", item.SegmentId, err)
				return
			}
			fmt.Printf("Successfully inserted Segment ID %d\n", item.SegmentId)
		}(item)
	}

	wg.Wait()
	return nil
}

func init() {
	config.LoadConfig()
	db, err := gorm.Open(mysql.Open(config.AppConfig.DBConnectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	if err := db.AutoMigrate(&models.SentenceItem{}); err != nil {
		panic(fmt.Sprintf("error migrating schema: %v", err))
	}
	DB = db
}
