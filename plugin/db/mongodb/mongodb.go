package mongodb

import (
	"log"
)

var (
	globalS *mgo.Session
)

// Init 初始化
func Init(dialInfo *mgo.DialInfo) {
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalf("Create Session: %s\n", err.Error())
	}
	globalS = s
}

// connect 连接
func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	ms := globalS.Copy() //避免每次创建session
	c := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, c
}

// Insert 插入数据
func Insert(db, collection string, doc interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(doc)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).All(result)
}

func Update(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Update(selector, update)
}

// Upsert `upsert:true`
func Upsert(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	_, err := c.Upsert(selector, update)
	return err
}

func UpdateAll(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.UpdateAll(selector, update)
	return err
}

func Remove(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Remove(selector)
}

func RemoveAll(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.RemoveAll(selector)
	return err
}

func FindPage(db, collection string, page, limit int, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}

func IsEmpty(db, collection string) bool {
	ms, c := connect(db, collection)
	defer ms.Close()
	count, err := c.Count()
	if err != nil {
		log.Fatal(err)
	}
	return count == 0
}

func Count(db, collection string, query interface{}) (int, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Count()
}

// InsertAll 批量添加
func InsertAll(db, collection string, doc ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(doc)
}

// FindAllSort 根据字段排序
func FindAllSort(db, collection string, field []string, page, limit int, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).Sort(field...).Skip((page - 1) * limit).Limit(limit).All(result)
}
