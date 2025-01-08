package hello

type Hello struct {
	Say string
}

func GetHello1(s string) string {
	if s == "" {
		return "Hello1"
	} else if s != "DDD" {
		return "Not hello1 DDD.."
	}
	return "Hello1 " + s + "!!"
}

func (h *Hello) GetHello2() string {
	if h.Say == "" {
		return "Hello2"
	} else if h.Say != "DDD" {
		return "Not hello2 DDD.."
	}
	return "Hello2 " + h.Say + "!!"
}
