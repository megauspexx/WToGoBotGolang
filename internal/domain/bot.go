package domain



type Question struct {
	Id    int
	Type  string
	Text  string
	Image string
}

type Motion struct {
	ActionId   int
	QuestionId int
	AnswerId   int
}

type ReciveResult struct {
	motion Motion
	Type   string
	Text   string
	Image  string
}