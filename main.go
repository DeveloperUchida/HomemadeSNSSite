package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Post データベース内の投稿を示す構造体
type Post struct {
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

func main() {
	// MongoDBへ接続
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Ginルーターを作成
	router := gin.Default()

	// ポストの一覧を取得するエンドポイント
	router.GET("/posts", func(c *gin.Context) {
		var posts []Post
		err := session.DB(dbName).C(collection).Find(nil).All(&posts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, posts)
	})

	// ポストを作成するエンドポイント
	router.POST("/posts", func(c *gin.Context) {
		var post Post
		if err := c.BindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		post.ID = bson.NewObjectId()
		post.Timestamp = nowUnix()
		if err := session.DB(dbName).C(collection).Insert(post); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, post)
	})

	// サーバーを起動
	router.Run(":8080")
}

// nowUnix 現在のUNIXタイムスタンプを返すヘルパー関数
func nowUnix() int64 {
	return time.Now().Unix()
}
