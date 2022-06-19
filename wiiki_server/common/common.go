package common

func ExcludeNilFromMap(m map[string]interface{}) map[string]interface{} {
	excludeMap := map[string]interface{}{}
	for key, value := range m {

		if value == nil {
			continue
		}
		excludeMap[key] = value
	}
	return excludeMap
}
