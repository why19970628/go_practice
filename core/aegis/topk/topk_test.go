package topk

import (
	"github.com/go-kratos/aegis/topk"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
	"math"
	"strconv"
	"testing"
	"time"
)

func TestTopkList(t *testing.T) {
	// zipfan distribution NewHeavyKeeper
	zipf := rand.NewZipf(rand.New(rand.NewSource(uint64(time.Now().Unix()))), 3, 2, 1000)
	topk := topk.NewHeavyKeeper(10, 10000, 5, 0.925, 0)
	dataMap := make(map[string]int)
	for i := 0; i < 10000; i++ {
		key := strconv.FormatUint(zipf.Uint64(), 10)
		dataMap[key] = dataMap[key] + 1
		topk.Add(key, 1)
	}
	var rate float64
	for _, node := range topk.List() {
		rate += math.Abs(float64(node.Count)-float64(dataMap[node.Key])) / float64(dataMap[node.Key])
		t.Logf("item %s, count %d, expect %d", node.Key, node.Count, dataMap[node.Key])
	}
	t.Logf("err rate avg:%f", rate)
	for i, node := range topk.List() {
		assert.Equal(t, strconv.FormatInt(int64(i), 10), node.Key)
		t.Logf("%s: %d", node.Key, node.Count)
	}
}
