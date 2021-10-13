package routing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/hayrullahcansu/cachy/api/controller"
	"github.com/hayrullahcansu/cachy/data/request"
	"github.com/hayrullahcansu/cachy/framework/logging"
	"github.com/hayrullahcansu/cachy/utility"
)

var pattern = `api\/v1\/cache\/{0,1}([a-zA-Z_\{\}\{\[\]\(\)\\\-@$0-9]+)\/{0,1}`

func CacheRouteHandler(w http.ResponseWriter, r *http.Request) {
	controller := controller.NewCacheController(w, r)
	reg := regexp.MustCompile(pattern)
	url := r.URL.Path[1:]
	queries := r.URL.Query()
	requestPayload := utility.FormatRequest(r)
	logging.Serverf(requestPayload)
	logging.Infof("Cache Controller , %s!", url)
	logging.Infof("Cache Controller queries, %v!", queries)
	if regexGroup := reg.FindStringSubmatch(url); regexGroup != nil {
		lengthRegexGroup := len(regexGroup)
		if lengthRegexGroup == 2 {
			key := regexGroup[1]
			switch method := r.Method; method {
			//Check GET & Show
			case "GET":
				controller.GetItem(key)
			//Check POST & Create
			//Check PUT & Update
			case "POST", "PUT":

				var model request.CreateCacheItemRequest
				err := r.ParseForm()
				if err != nil || len(r.Form) == 0 {
					//JSON
					body, err := utility.ReadBodyAsBytes(r)
					if err != nil {
						controller.InternalServerErrorWithBody(err.Error())
						return
					}
					var msg interface{}
					model = request.CreateCacheItemRequest{
						Data: &msg,
					}
					if err := json.Unmarshal(body, &model); err != nil {
						controller.InternalServerErrorWithBody(err.Error())
					}

				} else {
					//FORM DATA
					timespan_str := r.Form.Get("time_span")
					data_str := r.Form.Get("data")
					// data := r.Form.Get("data")
					timespan, err := strconv.Atoi(timespan_str)
					if err != nil {
						timespan = 0
					}
					var data interface{}
					err = json.Unmarshal([]byte(data_str), &data)
					if err != nil {
						data = data_str
					}
					model = request.CreateCacheItemRequest{
						TimeSpan: timespan,
						Data:     data,
					}
				}
				controller.SetItem(key, &model)
				return
			//Check DELETE & Delete
			case "DELETE":
				if strings.Contains(url, "cache/flush") {
					controller.Flush()
				} else {
					controller.DeleteItem(key)
				}
				return
			default:
			}

		} else {
			fmt.Fprintf(w, "Invalid Route, %s!", url)
		}
	} else {
		//LİSTING
		controller.ListItems()
		return
	}
}
