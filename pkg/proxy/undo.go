package proxy

import (
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	"reflect"
	"strings"
)

// Undo insert param
func UndoInsert(param *InsertParam) (params []*RemoveParam) {
	if param == nil {
		log.Error("received nil input")
		return
	}
	for _, doc := range param.Docs {
		var removeFilter = bson.M{}
		objType := reflect.TypeOf(doc)
		objVal := reflect.ValueOf(doc)
		for i := 0; i < objVal.NumField(); i++ {
			if objVal.Field(i).CanInterface() {
				if objVal.Field(i).Kind() == reflect.Struct {
					// TODO support nested struct
					log.Warning("by pass nested struct")
				} else {
					lowerName := strings.ToLower(objType.Field(i).Name)
					removeFilter[lowerName] = objVal.Field(i).Interface()
				}
			}
		}

		params = append(params, &RemoveParam{
			Filter: removeFilter,
			Amp: param.Amp,
		})
	}
	return
}
