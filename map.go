package log

func Map(val map[string]string) Fields {
	tmp := make(map[string]interface{})
	for k, v := range val {
		tmp[k] = v
	}
	return Fields(tmp)
}
