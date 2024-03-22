package main
import(
	"fmt"
	"net/http"
	"github.com/gin-going/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Post データベース内の投稿を示す構造体
type Post struct{
	ID        bson.ObjectId `bson:"_id,omitempty"`
    Content   string        `bson:"content"`
    Author    string        `bson:"author"`
    Timestamp int64         `bson:"timestamp"`
}
var (
	session    *mgo.Session
    dbName     = "somedb"
    collection = "posts"
)
func main(){
	
}