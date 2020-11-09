package models

type People struct {
	//ID 		primitive.ObjectID 	`json:"id,omitempty"bson:"_id,omitempty"`
	Name 	string 				`json:"name,omitempty"`
	Age 	int 				`json:"age,omitempty"`
	Gender 	string 				`json:"gender,omitempty"`
	Country string 				`json:"country,omitempty"`
}
