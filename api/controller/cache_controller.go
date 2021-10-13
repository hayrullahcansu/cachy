package controller

import (
	"net/http"
	"time"

	"github.com/hayrullahcansu/cachy/core/services"
	"github.com/hayrullahcansu/cachy/data/request"
	"github.com/hayrullahcansu/cachy/data/view"
)

type CacheController struct {
	*BaseController
	cacheService services.CacheService
}

func NewCacheController(w http.ResponseWriter, r *http.Request) *CacheController {
	return &CacheController{
		BaseController: NewBaseController(w, r),
		cacheService:   *services.NewCacheService(),
	}
}

func (c *CacheController) GetItem(key string) {
	item := c.cacheService.GetCacheEntry(key)
	if item == nil {
		c.NotFound()
		return
	}
	itemView := view.NewCacheItemView(item.Key, item.Value, item.Dead)
	c.OkWithBody(*itemView)
}
func (c *CacheController) SetItem(key string, model *request.CreateCacheItemRequest) {
	data := model.Data
	timespan := time.Second * time.Duration(model.TimeSpan)
	dead := time.Now().Add(timespan)
	// var str = dead.Format("02-01-2006 15:04:05")
	// var str2 = time.Now().Local().Format("02-01-2006 15:04:05")
	// logging.Info(str)
	// logging.Info(str2)
	item := c.cacheService.SetCacheEntry(key, data, dead.UnixNano())
	if item == nil {
		c.InternalServerError()
		return
	}
	itemView := view.NewCacheItemView(item.Key, item.Value, item.Dead)
	c.OkWithBody(*itemView)
}
func (c *CacheController) DeleteItem(key string) {
	item := c.cacheService.DeleteCacheEntry(key)
	if item == nil {
		c.NotFound()
		return
	}
	itemView := view.NewCacheItemView(item.Key, item.Value, item.Dead)
	c.OkWithBody(*itemView)
}
func (c *CacheController) Flush() {
	c.cacheService.FlushEntries()
	c.Ok()
}
func (c *CacheController) ListItems() {
	entities := c.cacheService.GetCacheEntries()
	var itemViewList []*view.CacheItemView

	for _, entity := range entities {
		itemViewList = append(itemViewList, view.NewCacheItemView(entity.Key, entity.Value, entity.Dead))
	}
	c.OkWithBody(itemViewList)
}
