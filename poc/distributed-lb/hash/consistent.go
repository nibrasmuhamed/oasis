// Copyright (c) 2018-2022 Burak Sezer
// All rights reserved.
//
// This code is licensed under the MIT License.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files(the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and / or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions :
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package consistent provides a consistent hashing function with bounded loads. This implementation also adds
// partitioning logic on top of the original algorithm. For more information about the underlying algorithm,
// please take a look at https://research.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html
//
// Example Use:
//
//	cfg := consistent.Config{
//		PartitionCount:    71,
//		ReplicationFactor: 20,
//		Load:              1.25,
//		Hasher:            hasher{},
//	}
//
// Now you can create a new Consistent instance. This function can take a list of the members.
//
//	c := consistent.New(members, cfg)
//
// In the following sample, you add a new Member to the consistent hash ring. myMember is just a Go struct that
// implements the Member interface. You should know that modifying the consistent hash ring distributes partitions among
// members using the algorithm defined on Google Research Blog.
//
//	c.Add(myMember)
//
// Remove a member from the consistent hash ring:
//
//	c.Remove(member-name)
//
// LocateKey hashes the key and calculates partition ID with this modulo operation: MOD(hash result, partition count)
// The owner of the partition is already calculated by New/Add/Remove. LocateKey just returns the member that is responsible
// for the key.
//
//	key := []byte("my-key")
//	member := c.LocateKey(key)
package hash

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"sync"

	"github.com/cespare/xxhash"
)

const (
	DefaultPartitionCount    int     = 271
	DefaultReplicationFactor int     = 20
	DefaultLoad              float64 = 1.25
)

// ErrInsufficientMemberCount represents an error which means there are not enough members to complete the task.
var ErrInsufficientMemberCount = errors.New("insufficient member count")

// Hasher is responsible for generating unsigned, 64-bit hash of provided byte slice.
// Hasher should minimize collisions (generating same hash for different byte slice)
// and while performance is also important fast functions are preferable (i.e.
// you can use FarmHash family).
type Hasher func(data []byte) uint64

// Member interface represents a member in consistent hash ring.
type Member struct {
	Name string
}

func (m Member) String() string {
	return m.Name
}

type MemberList []Member

func (a MemberList) Len() int           { return len(a) }
func (a MemberList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MemberList) Less(i, j int) bool { return a[i].Name < a[j].Name }

// Config represents a structure to control consistent package.
type Config struct {
	// Hasher is responsible for generating unsigned, 64-bit hash of provided byte slice.
	Hasher Hasher

	// Keys are distributed among partitions. Prime numbers are good to
	// distribute keys uniformly. Select a big PartitionCount if you have
	// too many keys.
	PartitionCount int

	// Members are replicated on consistent hash ring. This number means that a member
	// how many times replicated on the ring.
	ReplicationFactor int

	// Load is used to calculate average load. See the code, the paper and Google's blog post to learn about it.
	Load float64
}

// Consistent holds the information about the members of the consistent hash circle.
type Consistent struct {
	mu sync.RWMutex

	config         Config
	hasher         Hasher
	sortedSet      []uint64
	partitionCount uint64
	loads          map[string]float64
	members        map[string]*Member
	partitions     map[int]*Member
	ring           map[uint64]*Member
}

// New creates and returns a new Consistent object.
func New(members []Member, config Config) *Consistent {
	if config.Hasher == nil {
		config.Hasher = xxhash.Sum64
	}
	if config.PartitionCount == 0 {
		config.PartitionCount = DefaultPartitionCount
	}
	if config.ReplicationFactor == 0 {
		config.ReplicationFactor = DefaultReplicationFactor
	}
	if config.Load == 0 {
		config.Load = DefaultLoad
	}

	c := &Consistent{
		config:         config,
		members:        make(map[string]*Member),
		partitionCount: uint64(config.PartitionCount),
		ring:           make(map[uint64]*Member),
	}

	c.hasher = config.Hasher
	for _, member := range members {
		c.add(member)
	}
	if members != nil {
		c.distributePartitions()
	}
	return c
}

// GetMembers returns a thread-safe copy of members. If there are no members, it returns an empty slice of Member.
func (c *Consistent) GetMembers() []Member {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Create a thread-safe copy of member list.
	members := make([]Member, 0, len(c.members))
	for _, member := range c.members {
		members = append(members, *member)
	}
	return members
}

// AverageLoad exposes the current average load.
func (c *Consistent) AverageLoad() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.averageLoad()
}

func (c *Consistent) averageLoad() float64 {
	if len(c.members) == 0 {
		return 0
	}

	avgLoad := float64(c.partitionCount/uint64(len(c.members))) * c.config.Load
	return math.Ceil(avgLoad)
}

func (c *Consistent) distributeWithLoad(partID, idx int, partitions map[int]*Member, loads map[string]float64) {
	avgLoad := c.averageLoad()
	var count int
	for {
		count++
		if count >= len(c.sortedSet) {
			// User needs to decrease partition count, increase member count or increase load factor.
			panic("not enough room to distribute partitions")
		}
		i := c.sortedSet[idx]
		// if _,ok  := c.ring[i]; !ok {
		// 	return
		// }
		member := *c.ring[i]
		load := loads[member.String()]
		if load+1 <= avgLoad {
			partitions[partID] = &member
			loads[member.String()]++
			return
		}
		idx++
		if idx >= len(c.sortedSet) {
			idx = 0
		}
	}
}

func (c *Consistent) distributePartitions() {
	loads := make(map[string]float64)
	partitions := make(map[int]*Member)

	bs := make([]byte, 8)
	for partID := uint64(0); partID < c.partitionCount; partID++ {
		binary.LittleEndian.PutUint64(bs, partID)
		key := c.hasher(bs)
		idx := sort.Search(len(c.sortedSet), func(i int) bool {
			return c.sortedSet[i] >= key
		})
		if idx >= len(c.sortedSet) {
			idx = 0
		}
		c.distributeWithLoad(int(partID), idx, partitions, loads)
	}
	c.partitions = partitions
	c.loads = loads
}

func (c *Consistent) add(member Member) {
	dup := 0
	for i := 0; i < c.config.ReplicationFactor; i++ {
		h := c.getMemberHash(member.Name, i)
		if _, ok := c.ring[h]; ok {
			dup++
		}
		c.ring[h] = &member
		c.sortedSet = append(c.sortedSet, h)
	}
	if dup > 0 {
		fmt.Println("DUPLICATE HASH FOUND: ", dup)
		fmt.Printf("Sorted:%d, Ring: %d\n", len(c.sortedSet), len(c.ring))
	}
	// sort hashes ascendingly
	sort.Slice(c.sortedSet, func(i int, j int) bool {
		return c.sortedSet[i] < c.sortedSet[j]
	})
	// Storing member at this map is useful to find backup members of a partition.
	c.members[member.String()] = &member
}

// Add adds a new member to the consistent hash circle.
func (c *Consistent) Add(member Member) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.members[member.String()]; ok {
		// We already have this member. Quit immediately.
		return
	}
	c.add(member)
	c.distributePartitions()
}

func (c *Consistent) delSlice(val uint64) {
	for i := 0; i < len(c.sortedSet); i++ {
		if c.sortedSet[i] == val {
			c.sortedSet = append(c.sortedSet[:i], c.sortedSet[i+1:]...)
			//break - potentially there could be duplicate hash b'cos of collision
		}
	}
}

// Remove removes a member from the consistent hash circle.
func (c *Consistent) Remove(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.members[name]; !ok {
		// There is no member with that name. Quit immediately.
		return
	}

	for i := 0; i < c.config.ReplicationFactor; i++ {
		h := c.getMemberHash(name, i)
		delete(c.ring, h)
		c.delSlice(h)
	}
	delete(c.members, name)
	if len(c.members) == 0 {
		// consistent hash ring is empty now. Reset the partition table.
		c.partitions = make(map[int]*Member)
		return
	}
	c.distributePartitions()
}

// LoadDistribution exposes load distribution of members.
func (c *Consistent) LoadDistribution() map[string]float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Create a thread-safe copy
	res := make(map[string]float64)
	for member, load := range c.loads {
		res[member] = load
	}
	return res
}

// FindPartitionID returns partition id for given key.
func (c *Consistent) FindPartitionID(key []byte) int {
	hkey := c.hasher(key)
	return int(hkey % c.partitionCount)
}

// GetPartitionOwner returns the owner of the given partition.
func (c *Consistent) GetPartitionOwner(partID int) Member {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.getPartitionOwner(partID)
}

// getPartitionOwner returns the owner of the given partition. It's not thread-safe.
func (c *Consistent) getPartitionOwner(partID int) Member {
	member, ok := c.partitions[partID]
	if !ok {
		return Member{}
	}
	// Create a thread-safe copy of member and return it.
	return *member
}

// LocateKey finds a home for given key
func (c *Consistent) LocateKey(key []byte) Member {
	partID := c.FindPartitionID(key)
	return c.GetPartitionOwner(partID)
}

func (c *Consistent) getClosestN(partID, count int) ([]Member, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var res []Member
	if count > len(c.members) {
		return res, ErrInsufficientMemberCount
	}

	var ownerKey uint64
	owner := c.getPartitionOwner(partID)
	// Hash and sort all the names.
	var keys []uint64
	kmems := make(map[uint64]*Member)
	for name, member := range c.members {
		key := c.hasher([]byte(name))
		if name == owner.String() {
			ownerKey = key
		}
		keys = append(keys, key)
		kmems[key] = member
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// Find the key owner
	idx := 0
	for idx < len(keys) {
		if keys[idx] == ownerKey {
			key := keys[idx]
			res = append(res, *kmems[key])
			break
		}
		idx++
	}

	// Find the closest(replica owners) members.
	for len(res) < count {
		idx++
		if idx >= len(keys) {
			idx = 0
		}
		key := keys[idx]
		res = append(res, *kmems[key])
	}
	return res, nil
}

// GetClosestN returns the closest N member to a key in the hash ring.
// This may be useful to find members for replication.
func (c *Consistent) GetClosestN(key []byte, count int) ([]Member, error) {
	partID := c.FindPartitionID(key)
	return c.getClosestN(partID, count)
}

// GetClosestNForPartition returns the closest N member for given partition.
// This may be useful to find members for replication.
func (c *Consistent) GetClosestNForPartition(partID, count int) ([]Member, error) {
	return c.getClosestN(partID, count)
}

func (c *Consistent) getMemberHash(name string, i int) uint64 {
	//key := fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(i)+name)))
	key := strconv.Itoa(i) + name
	return c.hasher([]byte(key))
}

func (c *Consistent) GetPartitionList() map[int]*Member {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.partitions
}

func (c *Consistent) MemberExists(name string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.members[name]
	return exists
}
