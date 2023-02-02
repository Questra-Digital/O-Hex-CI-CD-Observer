package main

// "context"
// "log"

const (
	RedisAddr     string = "localhost : 6379"
	RedisPassword string = "saad1221"
	RedisDB       int    = 0
)

func main() {

	// ctx := context.Background()

	// c := cache.New(RedisAddr, RedisPassword, RedisDB)
	// if err := c.Ping(ctx); err != nil {
	// 	log.Panic("Failed to connect to redis")
	// }
}
