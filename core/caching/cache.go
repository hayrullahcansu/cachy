package caching

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/hayrullahcansu/cachy/cross"
	"github.com/hayrullahcansu/cachy/data/constants"
	"github.com/hayrullahcansu/cachy/framework/config"
	"github.com/hayrullahcansu/cachy/framework/logging"
	"github.com/hayrullahcansu/cachy/utility"
)

var _instance *Cache
var once sync.Once

type Cache struct {
	lock           *sync.Mutex
	deleteInterval time.Duration
	backupInterval time.Duration
	backupFilePath string
	stop           chan time.Time
	items          map[string]CacheItem
}

// Instance return new or existing instance of Cache
func Instance() *Cache {
	once.Do(func() {
		backupInterval := constants.DefaultBackupInterval
		if config.Instance().BackupInterval > 0 {
			backupInterval = time.Second * time.Duration(config.Instance().BackupInterval)
		}
		filePath := config.Instance().BackupFilePath
		if utility.IsNullOrEmpty(filePath) {
			filePath = cross.DefaultBackupFilePath
		}
		_instance = newCacheInstance(backupInterval, filePath)
		_instance.loadBackupIfExists()
		go _instance.run()

	})
	return _instance
}

func (c *Cache) Set(key string, value interface{}) *CacheItem {
	var dead time.Time
	c.lock.Lock()
	defer c.lock.Unlock()
	newCacheItem := NewCacheItem(key, value, dead.UnixNano())
	c.set(key, newCacheItem)
	return newCacheItem
}
func (c *Cache) SetWithTimeSpan(key string, value interface{}, dead int64) *CacheItem {
	c.lock.Lock()
	defer c.lock.Unlock()
	newCacheItem := NewCacheItem(key, value, dead)
	c.set(key, newCacheItem)
	return newCacheItem
}

func (c *Cache) Get(key string) *CacheItem {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.get(key)
}

func (c *Cache) Delete(key string) *CacheItem {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.checkKeyExists(key) {
		item := c.get(key)
		c.delete(key)
		return item
	}
	return nil
}
func (c *Cache) Flush() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.flush()
}
func (c *Cache) List() []*CacheItem {
	c.lock.Lock()
	defer c.lock.Unlock()

	arrayOfItems := make([]*CacheItem, 0, len(c.items))
	for _, val := range c.items {
		arrayOfItems = append(arrayOfItems, &val)
	}
	return arrayOfItems
}

// PRIVATE METHODS

func newCacheInstance(backupInterval time.Duration, backupFilePath string) *Cache {
	deleteInterval := time.Second * 2

	instance := &Cache{
		lock:           &sync.Mutex{},
		deleteInterval: deleteInterval,
		backupInterval: backupInterval,
		backupFilePath: backupFilePath,
		stop:           make(chan time.Time),
		items:          make(map[string]CacheItem),
	}

	return instance
}

func (c *Cache) run() {
	expirationTicker := time.NewTicker(c.deleteInterval)
	backupIntervalTicker := time.NewTicker(c.backupInterval)
	for {
		select {
		case <-expirationTicker.C:
			c.deleteExpiredItems()
		case <-backupIntervalTicker.C:
			c.writeToFile()
		case <-c.stop:
			expirationTicker.Stop()
			return
		}
	}
}

func (c *Cache) loadBackupIfExists() {
	c.lock.Lock()
	defer c.lock.Unlock()

	f, err := os.OpenFile(c.backupFilePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		logging.Errorf("open file error: %v", err)
		return
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var cacheItem = &CacheItem{}
		err = json.Unmarshal([]byte(line), cacheItem)
		if err != nil {
			logging.Errorf("error while reading backup file%s", err.Error())
		}
		if !utility.IsNullOrEmpty(cacheItem.Key) {
			c.set(cacheItem.Key, cacheItem)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	f.Close()
}

func (c *Cache) deleteExpiredItems() {
	c.lock.Lock()
	defer c.lock.Unlock()
	now := time.Now().UnixNano()
	for key, cacheItem := range c.items {
		if cacheItem.Dead > 0 && now > cacheItem.Dead {
			c.delete(key)
		}
	}
}

func (c *Cache) writeToFile() {
	c.lock.Lock()
	defer c.lock.Unlock()

	file, err := os.OpenFile(c.backupFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		logging.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)
	for _, item := range c.items {
		jsonString := utility.ToJson(item)
		_, _ = datawriter.WriteString(jsonString + cross.NewLine)
	}

	datawriter.Flush()
	file.Close()
}

func (c *Cache) delete(key string) {
	delete(c.items, key)
}
func (c *Cache) get(key string) *CacheItem {
	item, check := c.items[key]
	if check {
		return &item
	} else {
		return nil
	}
}
func (c *Cache) set(key string, item *CacheItem) {
	c.items[key] = *item
}
func (c *Cache) flush() {
	c.items = map[string]CacheItem{}
}
func (c *Cache) checkKeyExists(key string) bool {
	_, check := c.items[key]
	return check
}
