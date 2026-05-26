package httpserver

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/linkedin/Burrow/core/protocol"

	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	consumerTotalLagGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "burrow_kafka_consumer_lag_total",
			Help: "The sum of all partition current lag values for the group",
		},
		[]string{"cluster", "consumer_group"},
	)

	consumerStatusGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "burrow_kafka_consumer_status",
			Help: "The status of the consumer group. It is calculated from the highest status for the individual partitions. Statuses are an index list from NOTFOUND, OK, WARN, or ERR",
		},
		[]string{"cluster", "consumer_group"},
	)

	partitionStatusGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "burrow_kafka_topic_partition_status",
			Help: "The status of topic partition. It is calculated from the highest status for the individual partitions. Statuses are an index list from OK, WARN, STOP, STALL, REWIND",
		},
		[]string{"cluster", "consumer_group", "topic", "partition"},
	)

	consumerPartitionCurrentOffset = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "burrow_kafka_consumer_current_offset",
			Help: "Latest offset that Burrow is storing for this partition",
		},
		[]string{"cluster", "consumer_group", "topic", "partition"},
	)

	consumerPartitionLagGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "burrow_kafka_consumer_partition_lag",
			Help: "Number of messages the consumer group is behind by for a partition as reported by Burrow",
		},
		[]string{"cluster", "consumer_group", "topic", "partition"},
	)

	topicPartitionOffsetGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "burrow_kafka_topic_partition_offset",
			Help: "Latest offset the topic that Burrow is storing for this partition",
		},
		[]string{"cluster", "topic", "partition"},
	)
)

// DeleteConsumerMetrics deletes all metrics that are labeled with a consumer group
func DeleteConsumerMetrics(cluster, consumer string) { _ = "STUB: not implemented"; return }

// DeleteTopicMetrics deletes all metrics that are labeled with a topic
func DeleteTopicMetrics(cluster, topic string) { _ = "STUB: not implemented"; return }

// If a topic is deleted there cannot be any consumers, so delete all consumer metrics too
// Not strictly necessary as Kafka will delete the consumer groups, which will eventually trigger DeleteConsumerMetrics

// DeleteConsumerTopicMetrics deletes all metrics that are labeled with the provided consumer group AND topic
func DeleteConsumerTopicMetrics(cluster, consumer, topic string) { _ = "STUB: not implemented"; return }

func (hc *Coordinator) handlePrometheusMetrics() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// Topics

func listClusters(app *protocol.ApplicationContext) []string { _ = "STUB: not implemented"; return nil }

func listConsumers(app *protocol.ApplicationContext, cluster string) []string {
	_ = "STUB: not implemented"
	return nil
}

func getFullConsumerStatus(app *protocol.ApplicationContext, cluster, consumer string) *protocol.ConsumerGroupStatus {
	_ = "STUB: not implemented"
	return nil
}

func listTopics(app *protocol.ApplicationContext, cluster string) []string {
	_ = "STUB: not implemented"
	return nil
}

func getTopicDetail(app *protocol.ApplicationContext, cluster, topic string) []int64 {
	_ = "STUB: not implemented"
	return nil
}
