package consumer

import (
	log "github.com/sirupsen/logrus"
	"sync"
)
import "github.com/songxinjianqwe/go-practice/concurrent/channel/food"

func Consume(consumerName string, blockingQueue <-chan *food.Food, waitGroup *sync.WaitGroup) {
	for {
		if food, ok := <-blockingQueue; ok {
			log.Infof("【%s】consumed %s", consumerName, food)
		} else {
			break
		}
	}
	waitGroup.Done()
	log.Infof("【%s】stopped", consumerName)
}
