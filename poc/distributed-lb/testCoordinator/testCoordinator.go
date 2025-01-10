// server.go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"distributed-lb/coordinator"
	"distributed-lb/hash"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"fmt"
	"time"
)

func main() {
	k8sdata()
	// members := []hash.Member{}
	// size := 4
	// for i := 0; i < size; i++ {
	// 	member := hash.Member{Name: fmt.Sprintf("node%d", i)}
	// 	members = append(members, member)
	// }

	// b := coordinator.New(members)
	// names := []string{}
	// for i := size; i < size*2; i++ {
	// 	names = append(names, fmt.Sprintf("node%d", size+i))
	// }
	// for i := 0; ; i++ {
	// 	time.Sleep(1000 * time.Hour)
	// 	if i%2 == 1 {
	// 		m := b.GetMembers()
	// 		index := len(m) - 1
	// 		names = append(names, m[index].Name)
	// 		b.RemoveMember(m[index])
	// 	} else {
	// 		b.AddMember([]hash.Member{{
	// 			Name: names[0],
	// 		}})
	// 		names = names[1:]
	// 	}
	// 	// Wait for one second before sending the next message
	// }
}

func k8sdata() {
	certPath := "./client-cert.pem"
	keyPath := "./client-key.pem"

	url := "https://192.168.75.2:16443/api/v1/watch/namespaces/default/endpoints?fieldSelector=metadata.name=usage-engine-service"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		fmt.Println("Error loading client certificate and key:", err)
		return
	}

	caCert, err := ioutil.ReadFile(certPath)
	if err != nil {
		log.Fatalf("Error opening cert file %s, Error: %s", certPath, err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			RootCAs:            caCertPool,
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Timeout:   0,
		Transport: t,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", resp.StatusCode)
		return
	}
	var k Response
	members := []hash.Member{}
	err = json.NewDecoder(resp.Body).Decode(&k)
	if err != nil {
		panic(err)
	}

	for _, v := range k.Object.Subsets {
		for _, k := range v.Addresses {
			members = append(members, hash.Member{Name: k.IP})
		}
	}
	b := coordinator.New(members)
	decoder := json.NewDecoder(resp.Body)

	for {
		var data Response
		if err := decoder.Decode(&data); err == nil {
			fmt.Printf("Received event: %+v\n", data)
			m := []hash.Member{}
			for _, v := range data.Object.Subsets {
				for _, k := range v.Addresses {
					m = append(m, hash.Member{Name: k.IP})
				}
			}
			added, removed := diffMembers(members, m)
			b.AddMember(added)
			for _, j := range removed {
				b.RemoveMember(j)
			}
			members = m
		} else {
			break
		}
	}

}

func diffMembers(original []hash.Member, updated []hash.Member) ([]hash.Member, []hash.Member) {
	added := make([]hash.Member, 0)
	removed := make([]hash.Member, 0)

	// Create a map to efficiently check for existence
	originalMap := make(map[string]bool)
	for _, member := range original {
		originalMap[member.Name] = true
	}

	// Check for added and removed members
	for _, member := range updated {
		if _, exists := originalMap[member.Name]; exists {
			delete(originalMap, member.Name)
		} else {
			added = append(added, member)
		}
	}

	// Remaining members in originalMap are removed
	for name := range originalMap {
		removed = append(removed, hash.Member{Name: name})
	}

	return added, removed
}

type Response struct {
	Type   string `json:"type"`
	Object struct {
		Kind       string `json:"kind"`
		APIVersion string `json:"apiVersion"`
		Metadata   struct {
			Name              string    `json:"name"`
			Namespace         string    `json:"namespace"`
			UID               string    `json:"uid"`
			ResourceVersion   string    `json:"resourceVersion"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Annotations       struct {
				EndpointsKubernetesIoLastChangeTriggerTime time.Time `json:"endpoints.kubernetes.io/last-change-trigger-time"`
			} `json:"annotations"`
		} `json:"metadata"`
		Subsets []struct {
			Addresses []struct {
				IP        string `json:"ip"`
				NodeName  string `json:"nodeName"`
				TargetRef struct {
					Kind      string `json:"kind"`
					Namespace string `json:"namespace"`
					Name      string `json:"name"`
					UID       string `json:"uid"`
				} `json:"targetRef"`
			} `json:"addresses"`
			Ports []struct {
				Port     int    `json:"port"`
				Protocol string `json:"protocol"`
			} `json:"ports"`
		} `json:"subsets"`
	} `json:"object"`
}
