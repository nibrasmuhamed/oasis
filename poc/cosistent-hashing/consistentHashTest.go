package main

import (
	"crypto/rand"
	"fmt"
	"hash/fnv"
	"math"

	"github.com/buraksezer/consistent"
	// "github.com/cespare/xxhash"
)

type Member string

func (m Member) String() string {
	return string(m)
}

const (
	// offset64 FNVa offset basis. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	offset64 = 14695981039346656037
	// prime64 FNVa prime value. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	prime64 = 1099511628211
)

// Sum64 gets the string and returns its uint64 hash value.


type hasher struct{}

// func (h hasher) Sum64(data []byte) uint64 {
// 	return xxhash.Sum64(data)
// }

func (hs hasher) Sum64(key []byte) uint64 {
	h := fnv.New64()
	h.Write(key)
	return h.Sum64()
}

func init() {
	//rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	members := []consistent.Member{}
	for i := 0; i < 8; i++ {
		member := Member(fmt.Sprintf("node%d", i))
		members = append(members, member)
	}
	cfg := consistent.Config{
		PartitionCount:    16384,
		ReplicationFactor: 100,
		Load:              1.2,
		Hasher:            hasher{},
	}
	c := consistent.New(members, cfg)

	keyCount := 1000000
	load := (c.AverageLoad() * float64(keyCount)) / float64(cfg.PartitionCount)
	fmt.Println("Maximum key count for a member should be around this: ", math.Ceil(load))
	distribution := make(map[string]int)
	keys := make([][]byte, keyCount)
	for i := 0; i < keyCount; i++ {
		key := make([]byte, 4)
		_, err:=rand.Read(key)
		if (err != nil ) {
			fmt.Println("Key generation failed")
			return
		}
		//fmt.Println(key)
		member := c.LocateKey(key)
		//closestMember, _ := c.GetClosestN(key, 2)
		distribution[member.String()]++
		//fmt.Printf("\n%v %s %v", key, member.String(), closestMember)
		keys[i] = key
	}
	fmt.Println()
	
	for member, count := range distribution {
		fmt.Printf("member: %s, key count: %d\n", member, count)
	}
	c.Remove("node2")
	distribution = map[string]int{}
	for i := 0; i < keyCount; i++ {
		member := c.LocateKey(keys[i])
		//closestMember, _ := c.GetClosestN(keys[i], 2)
		distribution[member.String()]++
		//fmt.Printf("\n%v %s %v", keys[i], member.String(), closestMember)
	}
	
	fmt.Println()
	for member, count := range distribution {
		fmt.Printf("member: %s, key count: %d\n", member, count)
	}
}