# Tavern
<img align="right" width="159px" src="http://www.creativeuncut.com/gallery-08/art/f2-sketch-tavern.jpg">


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
Set multiple keys
```go
//Set multiple items in cache
items := prepareCacheItems()
fmt.Println("\nSet multiple items:")
result, err := cacheAdapter.MSet(items...)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Result: \n%+v\n", result)

func prepareCacheItems() []utils.CacheItem {
    data := map[string]string{
        "A": "I am A",
        "B": "I am A",
        "C": "I am C",
    }
    cacheItems := make([]utils.CacheItem, 0)
    for k, v := range data {
        item := utils.CacheItem{
            Key:   utils.CacheKey(k),
            Value: []byte(v),
        }
        cacheItems = append(cacheItems, item)
    }
    return cacheItems
}
```
Get multiple keys
```go
//Get multiple items from cache
fmt.Println("\n get multiple Items:")
resultget, err := cacheAdapter.MGet("A", "B", "C", "D")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Result: \n%+v\n", resultget)
```
Delete keys
```go
resultdelete, err := cacheAdapter.Destroy("A", "B", "C")
fmt.Printf("Result: \n%+v\n", resultdelete)
```
