package klnethttp

import "fmt"

func MapInterfaceToMapString(ms map[string]interface{}) map[string]string {

	mapString := make(map[string]string)
	for key, value := range ms {
		strValue := fmt.Sprintf("%v", value)

		mapString[key] = strValue
	}

	return mapString
}
