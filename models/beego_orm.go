package models
type User struct {
	Id        int64   `json:"id" `
	Name      string  `json:"name,omitempty" orm:"size(50)"`
	Passwords string  `json:"passwords" orm:"size(32)"`
	Baby      []*Baby `json:"baby" orm:"reverse(many)"`
}
type Baby struct {
	Id int64
	Name string `json:"name" orm:"size(50)"`
	User *User `json:"user" orm:"rel(fk);index"`

}



