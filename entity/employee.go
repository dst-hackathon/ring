package entity

type Employee struct {
	ID        string `bson:"_id,omitempty"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Phone     string
}
