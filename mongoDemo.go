package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type _user struct {
	Name  string  `json:"name" bson:"name"`
	Age   int     `json:"age" bson:"age"`
	Score float64 `json:"score" bson:"score"`
}

//新增单个
func addOne(c *mongo.Collection, object interface{}) error {
	result, err := c.InsertOne(context.TODO(), object)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("新增一个：", result.InsertedID)
	return err
}

//新增多个
func addMany(c *mongo.Collection, object []interface{}) error {
	result, err := c.InsertMany(context.TODO(), object)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("新增多个：", result.InsertedIDs)
	return err
}

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://admin:123456@localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	//获取collection
	collection := client.Database("test").Collection("user")

	var u _user
	//u1 := _user{"三儿", 13, 80}
	//u2 := _user{"六子", 16, 70}
	//u3 := _user{"柯比", 14, 77}

	//增one
	//addOne(collection, u1)

	//增Many
	//users := []interface{}{u2, u3}
	//addMany(collection, users)

	//查询单个
	filter := bson.M{
		"name": "六子",
	}
	err = collection.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", u)

	//查询多个:要查找多个文档，请使用collection.Find()。此方法返回一个游标。游标提供了一个文档流，你可以通过它一次迭代和解码一个文档。
	// 当游标用完之后，应该关闭游标。下面的示例将使用options包设置一个限制以便只返回两个文档。
	findOptions := options.Find()
	//设置获取的记录数
	//findOptions.SetLimit(2)
	//定义一个切片（List）来储存查询结果
	var results []*_user
	//把bson.D{{}}作为筛选条件传入，大括号里面没东西就是匹配所有
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	// 查找多个文档返回一个光标
	// 遍历游标允许我们一次解码一个文档
	for cursor.Next(context.TODO()) {
		var elem _user
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("get User %#v\n", elem)
		results = append(results, &elem)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	//修改
	update := bson.D{
		{"$set", bson.D{
			{"score", 89},
		}},
	}
	//更改匹配到的其中一个
	// 不等于 != ($ne)、 大于 > ($gt)、小于 < ($lt)、大于等于 >= ($gte)、小于等于 <= ($lte)、in($in)、是否包含这个键($exists)、正则匹配($regex)、$all查询所有、$or
	// 修改$set 、增加值$incr、向数组增加一个元素$push 、移除数组一个元素$pull https://my.oschina.net/u/2939474/blog/3201320
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	//匹配到的全更改
	//updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	////删除单个
	//delResult, err := collection.DeleteOne(context.TODO(), filter)
	////删除所有
	////deleteResult2, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Deleted %v documents in the trainers collection\n", delResult.DeletedCount)
}
