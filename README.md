# Tavern

All in one Caching library built in go to be used by go applications. The purpose is to expose an interface that is easy to be used and provides abstraction with the underlying caching techniques.

Supports following Caching engines:
- Inmemory Caching
- Redis
- Redis Cluster
- AWS Elasticache

## How to Install:

Simple run, below command.

```bash
go get -u github.com/sanksons/tavern
```
## How to use:

Detailed examples are kept in examples directory. But for a quick view: 

#### Initialization

Initilaize InMemory adapter:
```go
cacheAdapter := local.Initialize(local.LocalAdapterConfig{})
```
Initialize Redis adapter:
```go
cacheAdapter := redis.InitializeRedisSimple(redis.RedisSimpleConfig{
    Addr: "localhost:6379",
})
```
#### Usage

Set a Key
```go
//set a key
cacheAdapter.Set(utils.CacheItem{
    Key: "A", Value: []byte("I am A")
})
```
Get a key
```go
data, err := cacheAdapter.Get("A")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%s\n", data)
```