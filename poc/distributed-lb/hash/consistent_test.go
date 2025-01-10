package hash

import (
	"crypto/rand"
	"fmt"
	"hash/crc32"
	"hash/fnv"
	"math"
	"sort"
	"testing"

	"github.com/cespare/xxhash"
	"github.com/montanaflynn/stats"
	"github.com/segmentio/fasthash/fnv1"
)

var members MemberList

func init() {
	for i := 0; i < 8; i++ {
		member := Member{Name: fmt.Sprintf("node%d", i)}
		members = append(members, member)
	}
}

func TestConsistentXXhash(t *testing.T) {

	cfg := Config{
		PartitionCount:    16384,
		ReplicationFactor: 100,
		Load:              1.2,
		Hasher:            xxhash.Sum64,
	}
	c := New(members, cfg)
	loadKeys(c)
}

func checksum(data []byte) uint64 {
	return uint64(crc32.ChecksumIEEE(data))
}

func TestConsistentCrc(t *testing.T) {

	cfg := Config{
		PartitionCount:    16384,
		ReplicationFactor: 100,
		Load:              1.2,
		Hasher:            checksum,
	}
	c := New(members, cfg)
	loadKeys(c)
}

func fnvHash(data []byte) uint64 {
	h := fnv.New64()
	h.Write(data)
	return h.Sum64()
}
func TestConsistentfnv(t *testing.T) {

	cfg := Config{
		PartitionCount:    16384,
		ReplicationFactor: 100,
		Load:              1.2,
		Hasher:            fnvHash,
	}
	c := New(members, cfg)
	loadKeys(c)
}

func TestConsistentfnv1(t *testing.T) {

	cfg := Config{
		PartitionCount:    16384,
		ReplicationFactor: 100,
		Load:              1.2,
		Hasher:            fnv1.HashBytes64,
	}
	c := New(members, cfg)
	loadKeys(c)
}

func loadKeys(c *Consistent) {
	keyCount := 1000000
	load := (c.AverageLoad() * float64(keyCount)) / float64(c.partitionCount)
	fmt.Println("Maximum key count for a member should be around this: ", math.Ceil(load))

	keys := make([][]byte, keyCount)
	for i := 0; i < keyCount; i++ {
		key := make([]byte, 4)
		_, err := rand.Read(key)

		if err != nil {
			fmt.Println("Key generation failed")
			return
		}
		//fmt.Println(key)
		keys[i] = key
	}

	std(keys, c, load)
	c.Remove("node2")
	std(keys, c, load)
}

func std(keys [][]byte, c *Consistent, load float64) {
	distribution := map[string]int{}
	for _, key := range keys {

		member := c.LocateKey(key)
		//closestMember, _ := c.GetClosestN(key, 2)
		distribution[member.String()]++
		//fmt.Printf("\n%v %s %v", key, member.String(), closestMember)

	}
	fmt.Println()
	var list []float64
	for member, count := range distribution {
		fmt.Printf("member: %s, key count: %d\n", member, count)
		list = append(list, float64(count))
	}
	r, _ := stats.StandardDeviation(list)
	fmt.Println("Std Deviation: ", r)
	fmt.Println("%Dev: ",r*100/load)
}

func TestSort(t *testing.T) {
	fmt.Println("Original List: ", members)
	cfg := Config{
		PartitionCount:    16384,
		ReplicationFactor: 10000,
		Load:              1.2,
		Hasher:          fnvHash,
	}
	c := New(members, cfg)
	p := c.GetPartitionList()
	
	sort.Sort(sort.Reverse(members))
	c = New(members, cfg)
	reversep := c.GetPartitionList()
	fmt.Println("Reverse List: ", members)

	for i,m := range p {
		if reversep[i].Name != m.Name {
			fmt.Println("Not Identical: ", i)
		}
	}

}
