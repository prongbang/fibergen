package common

func Map(items []string, fn func(item string) string) string {
	data := ""
	for _, it := range items {
		data += fn(it)
	}
	return data
}
