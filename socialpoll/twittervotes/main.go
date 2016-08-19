package main

import (
	"log"

	"github.com/bitly/go-nsq"

	"labix.org/v2/mgo"
)

var db *mgo.Session

func dialDB() error {
	var err error
	log.Println("MongoDBにダイヤル中: localhost")
	db, err = mgo.Dial("localhost")
	return err
}

func closedb() {
	db.Close()
	log.Println("データベース接続が閉じられました")
}

type poll struct {
	Options []string
}

func loadOptions() ([]string, error) {
	var options []string
	iter := db.DB().C().Find().Iter()
	var p poll
	for iter.Next(&p) {
		options = append(options, p.Options...)
	}
	iter.Close()
	return options, iter.Err()
}

func publushVotes() <-chan struct{} {
	stopchan := make(chan struct{}, 1)
	pub, _ := nsq.NewProducer("localhost:4150", nsq.NewConfig())
    go func() {
        for vote := range votes {
            pub.Publish("votes", []byte(vote))
        }
        log.Println("Publisher: 停止中です")
        pub.Stop()
        log.Println("Publisher: 停止しました")
        stopchan <- struct{}{}
    }()
    return stopchan
}

func main() {
    var stoplock sync.Mutex
    stop := false
    stopChan := make(chan struct{}, 1)
    signalChan := make(chan os.Signal, 1)
    go func ()  {
        <-signalChan // <- はチャネルからの受信を試みる演算子
        stoplock.Lock()
        stop = true
        stoplock.Unlock()
        log.Println("停止します...")
        stopChan <- struct{}{}
        closeConn()
    }
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

    if err := dialDB(); err != nil {
        log.Println("MongoDBへのダイヤルに失敗しました", err)
    }
    defer closedb()
}
