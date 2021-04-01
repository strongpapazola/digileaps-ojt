package berita
import (
        "errors"
        "net/http"
        "time"
        "digileaps-ojt/conn"
        berita "digileaps-ojt/models"
        "github.com/gin-gonic/gin"
        "gopkg.in/mgo.v2/bson"
)
// BeritaCollection statically declared
const BeritaCollection = "berita"
var (
       errNotExist        = errors.New("Berita are not exist")
       errInvalidID       = errors.New("Invalid ID")
       errInvalidBody     = errors.New("Invalid request body")
       errInsertionFailed = errors.New("Error in the berita insertion")
       errUpdationFailed  = errors.New("Error in the berita updation")
       errDeletionFailed  = errors.New("Error in the berita deletion")
)
// GetAllUser Endpoint
func GetAllBerita(c *gin.Context) {
       // Get DB from Mongo Config
       db := conn.GetMongoDB()
       beritas := berita.Beritas{}
       err := db.C(BeritaCollection).Find(bson.M{}).All(&beritas)
       if err != nil {
                 c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
                 return
       }
       c.JSON(http.StatusOK, gin.H{"status": "success", "berita": &beritas})
}
// GetUser Endpoint
func GetBerita(c *gin.Context) {
       var id bson.ObjectId = bson.ObjectIdHex(c.Param("id")) // Get Param
       berita, err := berita.BeritaInfo(id, BeritaCollection)
       if err != nil {
                  c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidID.Error()})
                  return
       }
       c.JSON(http.StatusOK, gin.H{"status": "success", "berita": &berita})
}
// CreateUser Endpoint
func CreateBerita(c *gin.Context) {
       // Get DB from Mongo Config
       db := conn.GetMongoDB()
       berita := berita.Berita{}
       err := c.Bind(&berita)
       if err != nil {
                  c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
                  return
       }
       berita.ID = bson.NewObjectId()
       berita.CreatedAt = time.Now()
       berita.UpdatedAt = time.Now()
       err = db.C(BeritaCollection).Insert(berita)
       if err != nil {
                  c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error()})
                  return
       }
       c.JSON(http.StatusOK, gin.H{"status": "success", "berita": &berita})
}
// UpdateUser Endpoint
func UpdateBerita(c *gin.Context) {
       // Get DB from Mongo Config
       db := conn.GetMongoDB()
       var id bson.ObjectId = bson.ObjectIdHex(c.Param("id")) // Get Param
       existingBerita, err := berita.BeritaInfo(id, BeritaCollection)
       if err != nil {
                  c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidID.Error()})
                  return
       }
       // user := user.User{}
       err = c.Bind(&existingBerita)
       if err != nil {
                 c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
                  return
       }
       existingBerita.ID = id
       existingBerita.UpdatedAt = time.Now()
       err = db.C(BeritaCollection).Update(bson.M{"_id": &id}, existingBerita)
       if err != nil {
                  c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errUpdationFailed.Error()})
                  return
       }
       c.JSON(http.StatusOK, gin.H{"status": "success", "berita": &existingBerita})
}
// DeleteUser Endpoint
func DeleteBerita(c *gin.Context) {
       // Get DB from Mongo Config
       db := conn.GetMongoDB()
       var id bson.ObjectId = bson.ObjectIdHex(c.Param("id")) // Get Param
       err := db.C(BeritaCollection).Remove(bson.M{"_id": &id})
       if err != nil {
                  c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errDeletionFailed.Error()})
                  return
       }
       c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Berita deleted successfully"})
}