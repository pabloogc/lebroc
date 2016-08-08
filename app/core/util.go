package core

func ToStringSlice(data interface{}) []string {
	s := data.([]interface{})
	out := make([]string, len(s))
	for i, v := range s {
		out[i] = v.(string)
	}
	return out
}
