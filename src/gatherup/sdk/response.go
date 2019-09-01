package sdk

type Response struct {
	Data map[string]interface{}
}

func (r Response) GetCode() int {
	if ok, val := r.Get("errorCode"); ok {
		return int(val.(float64))
	}
	return -1
}

func (r Response) GetMessage() string {
	if ok, val := r.Get("errorMessage"); ok {
		return val.(string)
	}
	return "Unknown"
}

func (r Response) IsSuccess() bool {
	return r.GetCode() == 0
}

func (r Response) Get(key string) (bool, interface{}) {
	if val, ok := r.Data[key]; ok {
		return true, val
	}
	return false, nil
}
