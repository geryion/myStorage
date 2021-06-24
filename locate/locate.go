package locate

import (
	"myStorage/rabbitmq"
	"os"
)

func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate()  {
	mq := rabbitmq.RNew(os.Getenv("RABBITMQ_SERVER"))
	defer mq.RClose()


}
