package models
import (
	"digileaps-ojt/conn"
	"time"
	"gopkg.in/mgo.v2/bson"
)
// User structure
type Berita struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Address   string        `bson:"address"`
	Age       int           `bson:"age"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}
// Users list
type Beritas []Berita
// UserInfo model function
func BeritaInfo(id bson.ObjectId, beritaCollection string) (Berita, error) {
	// Get DB from Mongo Config
	db := conn.GetMongoDB()
	berita := Berita{}
	err := db.C(beritaCollection).Find(bson.M{"_id": &id}).One(&berita)
	return berita, err
}
