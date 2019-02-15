package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/songxinjianqwe/go-practice/concurrent/channel/consumer"
	"github.com/songxinjianqwe/go-practice/concurrent/channel/food"
	"github.com/songxinjianqwe/go-practice/concurrent/channel/producer"
	"os"
	"sync"
)

/**
全局有效
*/
func init() {
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	log.SetFormatter(&log.TextFormatter{})
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	log.SetOutput(os.Stdout)
	//设置最低loglevel
	log.SetLevel(log.InfoLevel)
}

/**
这个示例是一个多生产者-多消费者模式的示例；
1）启动了三个生产者goroutine，分别生产20个食物，当结束时，将waitProducersStop的计数值--
2）启动了两个消费者goroutine，无限消费食物，直至通道被关闭，当结束时，将waitConsumersStop的计数值--
3）构造了大小为6的BlockingQueue（channel），传入生产者和消费者的goroutine
4）main函数在等待生产者均生产完毕时，会关闭通道；
5）消费者在消费时如果发现通道被关闭，则会从无限循环中退出
*/
func main() {
	log.Infoln("【MAIN】started")
	blockingQueue := make(chan *food.Food, 6)
	waitProducersStop := sync.WaitGroup{}
	waitConsumersStop := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		producerName := fmt.Sprintf("producer-%d", i)
		log.Infof("launching %s", producerName)
		go producer.Produce(producerName, blockingQueue, &waitProducersStop)
		waitProducersStop.Add(1)
	}
	for i := 0; i < 2; i++ {
		consumerName := fmt.Sprintf("consumer-%d", i)
		log.Infof("launching %s", consumerName)
		go consumer.Consume(consumerName, blockingQueue, &waitConsumersStop)
		waitConsumersStop.Add(1)
	}
	log.Infoln("wait all producers stop")
	waitProducersStop.Wait()

	log.Infoln("all producers stopped")
	log.Infoln("closing channel")
	close(blockingQueue)

	log.Infoln("wait all consumers stop")
	waitConsumersStop.Wait()
	log.Infoln("all consumers stopped")
	log.Infoln("【MAIN】finished")
}
