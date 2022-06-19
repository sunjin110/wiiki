// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type NoopInput struct {
	ClientMutationID *string `json:"clientMutationId"`
}

type NoopPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type TodoID struct {
	ID string `json:"id"`
}

type UpdateTodo struct {
	ID     string  `json:"id"`
	Text   *string `json:"text"`
	Done   *bool   `json:"done"`
	UserID *string `json:"userId"`
}

type UpdateUser struct {
	ID    string  `json:"id"`
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserID struct {
	ID string `json:"id"`
}
