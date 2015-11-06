package entity

type Employee struct {
	ID           string     `bson:"_id,omitempty" jsonapi:"primary,employees"`
	FirstName    string     `bson:"first_name" jsonapi:"attr,first-name"`
	LastName     string     `bson:"last_name" jsonapi:"attr,last-name"`
	Phone        string     `bson:"phone" jsonapi:"attr,phone"`
	Subordinates []Employee `bson:"subordinates" jsonapi:"attr,subordinates"`
}
