
package users
import "gopkg.in/mgo.v2/bson"
//User represents a member of the app
type User struct {
	ID     bson.ObjectId `bson:"_id"`
	FirstName  string        `json:"firstname"`
	LastName string        `json:"lastname"`
}
//Users is an array of User
type Users []User
