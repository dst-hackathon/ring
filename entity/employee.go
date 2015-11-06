package entity

type Employee struct {
	ID           string     `bson:"_id,omitempty" json:"id"`
	FirstName    string     `bson:"first_name" json:"first-name"`
	LastName     string     `bson:"last_name" json:"last-name"`
	Phone        string     `bson:"phone" json:"phone"`
	Subordinates []Employee `bson:"subordinates" json:"subordinates"`
}
