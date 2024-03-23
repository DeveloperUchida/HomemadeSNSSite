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
	//MongoDBへ接続
	session, err := mgo.Dial("mongodb://localhost")
	if err != nul{
		panic(err)
	}
	defer session.Close()

	//コレクションを取得
	collection := gin.Default()

	//ポストの一覧を取得するエンドポイント
	router.GET("/posts",func (c *gin.Content)  {
		var posts []Post
		err := collection.Find(nil).All(&posts)
		if  err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, posts)
	})
	//ポストを作成するエンドポイント
	router.POST("/posts", func (c *gin.Content)  {
		if err := c.BindJSON(&post); err != nil{
			c.JSON(http.StatusOK)
		}
	})
}