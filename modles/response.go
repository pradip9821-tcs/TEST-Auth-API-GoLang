package modles

type Data struct {
	Id    int
	Email string
	Token string
}

type Response struct {
	Message string
	Data    Data
	Status  int
}
