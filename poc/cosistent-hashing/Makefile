add:
	curl http://127.0.0.1:9002/add
	curl http://127.0.0.1:9001/stats
	curl http://127.0.0.1:9002/stats
	curl http://127.0.0.1:9003/stats

get:
	curl http://127.0.0.1:9002/
	curl http://127.0.0.1:9001/stats
	curl http://127.0.0.1:9002/stats
	curl http://127.0.0.1:9003/stats

removeNode:
	curl http://127.0.0.1:9002/removeNode/1
	curl http://127.0.0.1:9003/removeNode/1

addNode:
	curl http://127.0.0.1:9002/addNode/1
	curl http://127.0.0.1:9003/addNode/1

start:
	go run groupCacheServer.go -port $(port)