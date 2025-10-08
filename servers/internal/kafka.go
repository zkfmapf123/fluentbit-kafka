package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/IBM/sarama"
)

type Pubsub struct {
	producer sarama.SyncProducer
	logger   ILogger
}

func MustNewPubsub(logger ILogger) *Pubsub {
	config := sarama.NewConfig()

	// 추후 Optional Parameter로 처리
	config.Producer.Return.Successes = true          // 성공 여부 반환
	config.Producer.Return.Errors = true             // 에러 반환
	config.Producer.RequiredAcks = sarama.WaitForAll // 모든 파티션에 데이터 전송 여부
	/*
		NoResponse -> 0 ACK 안 기다림 (빠름 , 유실)
		WaitForLocal -> 1 리더만 ACK
		WaitForAll -> -1 모든 ISR 복제본 ACK
	*/
	config.Producer.Partitioner = sarama.NewHashPartitioner
	/*
		NewRandomPartitioner -> 랜덤 파티션 할당
		NewRoundRobinPartitioner -> 순환 파티션 할당
		NewHashPartitioner -> 해시 파티션 할당
		NewManualPartitioner -> 수동 파티션 할당
	*/
	config.Producer.Retry = struct {
		Max         int
		Backoff     time.Duration
		BackoffFunc func(retries int, maxRetries int) time.Duration
	}{
		Max:     5,
		Backoff: time.Second * 1,
		BackoffFunc: func(retries int, maxRetries int) time.Duration {
			return time.Second * time.Duration(retries)
		},
	}

	// config.Producer.Idempotent = true // 중복방지

	// producer
	if os.Getenv("KAFKA_BROKERS") == "" {
		panic("KAFKA_BROKERS is not set")
	}

	producer, err := sarama.NewSyncProducer(strings.Split(os.Getenv("KAFKA_BROKERS"), ","), config)
	if err != nil {
		panic(err)
	}

	return &Pubsub{
		producer: producer,
		logger:   logger,
	}
}

func (p *Pubsub) Producer(topic string, value map[string]any) {

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1, //-1 은 랜덤 파티션 할당
		Value:     getStringEncoder(value),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
		p.logger.ErrorLogger("home", spreadLog(value, map[string]any{
			"error":     err.Error(),
			"partition": partition,
			"offset":    offset,
			"Status":    "Failed",
		}))
	} else {
		p.logger.InfoLogger("home", spreadLog(value, map[string]any{
			"partition": partition,
			"offset":    offset,
			"Status":    "Success",
		}))
	}
}

func (p *Pubsub) Consumer() {

}

func getStringEncoder(value map[string]any) sarama.Encoder {
	b, err := json.Marshal(value)
	if err != nil {
		return nil
	}
	return sarama.StringEncoder(b)
}

func spreadLog(value map[string]any, addValue map[string]any) map[string]any {
	m := map[string]any{}

	for k, v := range value {
		m[k] = v
	}

	for k, v := range addValue {
		m[k] = v
	}

	return m
}
