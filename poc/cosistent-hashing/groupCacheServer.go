package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"crypto/rand"
	"io"
	"math"
	"math/big"
	"strings"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/mailgun/groupcache/v2"
)

var group *groupcache.Group
var max = big.NewInt(10000)
var pool *groupcache.HTTPPool

func main() {

	port := flag.String("port", "1", "port number")
	flag.Parse()
	// NOTE: It is important to pass the same peer `http://192.168.1.1:8080` to `NewHTTPPoolOpts`
	// which is provided to `pool.Set()` so the pool can identify which of the peers is our instance.
	// The pool will not operate correctly if it can't identify which peer is our instance.

	// Pool keeps track of peers in our cluster and identifies which peer owns a key.
	ip := "127.0.0.1:808" + *port
	fmt.Println("Ip: " + ip)
	pool = groupcache.NewHTTPPoolOpts("http://"+ip, &groupcache.HTTPPoolOptions{Replicas: 1000})

	// Add more peers to the cluster You MUST Ensure our instance is included in this list else
	// determining who owns the key accross the cluster will not be consistent, and the pool won't
	// be able to determine if our instance owns the key.
	pool.Set("http://127.0.0.1:8081", "http://127.0.0.1:8082", "http://127.0.0.1:8083")
	//pool.Set("http://127.0.0.1:8082")
	server := http.Server{
		Addr:    ip,
		Handler: pool,
	}

	// Start a HTTP server to listen for peer requests from the groupcache
	go func() {
		log.Printf("Serving....\n")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	defer server.Shutdown(context.Background())

	// Create a new group cache with a max cache size of 3MB
	group = groupcache.NewGroup("users", 3000000, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {

			// Returns a protobuf struct `User`
			//user, err := fetchUserFromMongo(ctx, id)
			// if err != nil {
			//     return err
			// }
			fmt.Println(">>>>>>>>>>>>\nGoing to DB for Key: " + id)
			user := User{
				Id:      id,
				Name:    "John Doe",
				Age:     40,
				IsSuper: true,
			}

			// Set the user in the groupcache to expire after 5 minutes
			return dest.SetProto(&user, time.Now().Add(time.Minute*5))
		},
	))
	mux := http.NewServeMux()
	mux.HandleFunc("/", getUser)
	mux.HandleFunc("/add", addUser)
	mux.HandleFunc("/stats", stats)
	mux.HandleFunc("/removeNode/", removeNode)
	mux.HandleFunc("/addNode/", addNode)

	err := http.ListenAndServe("127.0.0.1:900"+*port, mux)
	if err != nil {
		fmt.Println("Error starting the http server:" + err.Error())
	}
}

func stats(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n %+v", group.CacheStats(groupcache.MainCache))
}

func removeNode(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/removeNode/")
	ip := "http://127.0.0.1:808" + id
	var ips []string
	peers := pool.GetAll()
	for _,p := range peers {
		url:= strings.TrimSuffix(p.GetURL(), "/_groupcache/")
		if url != ip {
			ips = append(ips, url)
		} 
	}
	pool.Set(ips[:]...)
	fmt.Println("Added: "+strings.Join(ips,","))

}

func addNode(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/addNode/")
	ip := "http://127.0.0.1:808" + id
 	ips := []string{ip}
	peers := pool.GetAll()
	for _,p := range peers {
		url:= strings.TrimSuffix(p.GetURL(), "/_groupcache/")
		if url != ip {
			ips = append(ips, url)
		} 
	}
	pool.Set(ips[:]...)
	fmt.Println("Added: "+strings.Join(ips,","))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	var user User

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	randomNumber, _ := rand.Int(rand.Reader, max)
	fmt.Println("Getting Key:", randomNumber)
	if err := group.Get(ctx, randomNumber.String(), groupcache.ProtoSink(&user)); err != nil {
		fmt.Println("Key Not found " + randomNumber.String() + " : " + err.Error())
		io.WriteString(w, "FAILED!\n")
		return
	}
	fmt.Printf("-- Got User --\n")
	fmt.Printf("Id: %s\n", user.Id)
	io.WriteString(w, "SUCCESS!\n")
	// fmt.Printf("Name: %s\n", user.Name)
	// fmt.Printf("Age: %d\n", user.Age)
	// fmt.Printf("IsSuper: %t\n", user.IsSuper)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	for i := 0; i < int(max.Int64()); i++ {
		//randomNumber, err := rand.Int(rand.Reader, max)
		// if err != nil {
		// 	io.WriteString(w, "FAILED!\n")
		// 	return
		// }
		user := User{
			Id:      strconv.Itoa(i),
			Name:    "John Doe",
			Age:     40,
			IsSuper: true,
		}
		//fmt.Println("Adding Key:", user.Id)
		val, _ := proto.Marshal(&user)
		err := group.Set(ctx,
			user.Id,
			val,
			time.Now().Add(24*time.Hour),
			false)
		if err != nil {
			fmt.Println("Error setting key: ", err)
			io.WriteString(w, "FAILED!\n")
			return
		}
	}
	io.WriteString(w, "SUCESS!\n")
	return
}

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Age     int64  `protobuf:"varint,3,opt,name=age" json:"age,omitempty"`
	IsSuper bool   `protobuf:"varint,4,opt,name=is_super,json=isSuper" json:"is_super,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetIsSuper() bool {
	if m != nil {
		return m.IsSuper
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "groupcachepb.User")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x49, 0x2f, 0xca, 0x2f, 0x2d,
	0x48, 0x4e, 0x4c, 0xce, 0x48, 0x2d, 0x48, 0x52, 0x0a, 0xe7, 0x62, 0x09, 0x2d, 0x4e, 0x2d, 0x12,
	0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11,
	0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0x8b, 0x80, 0xd9, 0x42, 0x02, 0x5c,
	0xcc, 0x89, 0xe9, 0xa9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x20, 0xa6, 0x90, 0x24, 0x17,
	0x47, 0x66, 0x71, 0x7c, 0x71, 0x69, 0x41, 0x6a, 0x91, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x47, 0x10,
	0x7b, 0x66, 0x71, 0x30, 0x88, 0xeb, 0x24, 0x18, 0xc5, 0x8f, 0xb0, 0x28, 0xbe, 0x24, 0xb5, 0xb8,
	0x24, 0x89, 0x0d, 0xec, 0x00, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x26, 0x2e, 0x5f, 0x1a,
	0x91, 0x00, 0x00, 0x00,
}
