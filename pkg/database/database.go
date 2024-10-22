package database

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/c8763yee/mygo-backend/internal/config"
	"github.com/c8763yee/mygo-backend/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
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

	return db
}

func ImportDataFromJSON(db *gorm.DB) error {
	// 讀取 data.json 檔案
	file, err := os.ReadFile("data.json")
	if err != nil {
		return fmt.Errorf("無法讀取 data.json 檔案: %v", err)
	}

	var sentenceData struct {
		Result []models.SentenceItem `json:"result"`
	}

	err = json.Unmarshal(file, &sentenceData)
	if err != nil {
		return fmt.Errorf("無法解析 JSON 資料: %v", err)
	}

	// 獲取 MySQL 的最大連接數
	var maxConnections int
	row := db.Raw("SHOW VARIABLES LIKE 'max_connections'").Row()
	err = row.Scan(new(string), &maxConnections)
	if err != nil {
		return fmt.Errorf("無法獲取 max_connections: %v", err)
	}

	workerCount := maxConnections - 10
	if workerCount <= 0 {
		return fmt.Errorf("可用的連接數不足")
	}

	// 創建一個帶緩衝的通道來限制並發的 goroutine 數量
	semaphore := make(chan struct{}, workerCount)
	var wg sync.WaitGroup

	for _, item := range sentenceData.Result {
		wg.Add(1)
		go func(item models.SentenceItem) {
			defer wg.Done()
			semaphore <- struct{}{}        // 獲取信號量
			defer func() { <-semaphore }() // 釋放信號量

			// 檢查記錄是否已存在
			var existingItem models.SentenceItem
			err := db.First(&existingItem, "segment_id = ?", item.SegmentId).Error
			if err == nil {
				fmt.Printf("Segment ID(%d) 已存在\n", item.SegmentId)
				return
			} else if err != gorm.ErrRecordNotFound {
				fmt.Printf("檢查 Segment ID(%d) 時發生錯誤: %v\n", item.SegmentId, err)
				return
			}

			// 插入新記錄
			if err := db.Create(&item).Error; err != nil {
				fmt.Printf("插入 Segment ID(%d) 時發生錯誤: %v\n", item.SegmentId, err)
				return
			}
			fmt.Printf("成功插入 Segment ID %d\n", item.SegmentId)
		}(item)
	}

	wg.Wait()
	return nil
}

func init() {
	DB = InitDB()
	if err := ImportDataFromJSON(DB); err != nil {
		panic(fmt.Sprintf("導入資料時發生錯誤: %v", err))
	}
}
