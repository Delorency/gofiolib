package schema

type Filter struct {
	ID         uint   `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Surname    string `json:"surname,omitempty"`
	Patronymic string `json:"patronymic,omitempty"`
	Age        uint8  `json:"age,omitempty"`
	Gender     string `json:"gender,omitempty"`
	Nat        string `json:"nat,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	DeletedAt  string `json:"deleted_at,omitempty"`
}
