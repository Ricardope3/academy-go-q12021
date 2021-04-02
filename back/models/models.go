package models

type Pokemon struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Todo struct {
	Userid    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
type TodoArray struct {
	Collection []Todo `json:"todos"`
}
