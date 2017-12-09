package render

import (
	"fmt"
)

type ModelLibrary struct {
	models    []Model
	modelsMap map[string]int
}

func NewModelLibrary(models []Model) ModelLibrary {
	var modelsMap = make(map[string]int)

	for i, model := range models {
		_, ok := modelsMap[model.name]
		if ok {
			fmt.Println(model.name, " already exists")
		} else {
			modelsMap[model.name] = i
		}
	}

	return ModelLibrary{
		models:    models,
		modelsMap: modelsMap,
	}
}

func (m ModelLibrary) GetModelID(modelName string) int {
	return m.modelsMap[modelName]
}
