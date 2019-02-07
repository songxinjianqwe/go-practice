package producer

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/songxinjianqwe/go-practice/concurrent/channel/food"
	"sync"
)

func Produce(producerName string, blockingQueue chan<- *food.Food, waitProducerStop *sync.WaitGroup) {
	for i := 0; i < 20; i++ {
		food := food.NewFood(fmt.Sprintf("【Food-%d-From-%s】", i, producerName))
		blockingQueue <- food
		log.Infof("【%s】produced %#v", producerName, food)
	}
	waitProducerStop.Done()
	log.Infof("【%s】stopped", producerName)
}
