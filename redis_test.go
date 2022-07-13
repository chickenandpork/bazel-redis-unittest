// Straightforward container turnup similar to (MIT) https://github.com/testcontainers/
// testcontainers-go/blob/master/README.md driving a few test steps similar to (BSD-2)
// https://github.com/go-redis/redis#quickstart

package main

import (
	"context"
	"net"
	"strings"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	//"github.com/testcontainers/testcontainers-go/wait"

	"github.com/go-redis/redis/v8"
)

func TestRedisLatestReturn(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:3-alpine",
		ExposedPorts: []string{"6379/tcp"},
		//WaitingFor:   wait.ForHTTP("/").WithPort("6379/tcp"),
	}
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer redisC.Terminate(ctx)
	t.Log(" -> configured")

	ip, err := redisC.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(" -> got a host: %+v\n", ip)

	// port is a nat.Port, which is a string with a new type.  string(port) is trivial yet a
	// typechange.
	port, err := redisC.MappedPort(ctx, "6379")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(" -> got a port: %+v\n", port)

	// the port looks like "12345/tcp" but we cannot use that directly to connect; scrub off
	// the protocol, and then connect, and we silently ignore that other protocols exist in
	// this world.  Nah, they're not really there.  Crazysauce to think so.  totally.

	ports := strings.Split(string(port), "/")
	fixedport := ports[0]
	t.Logf(" -> cleaned a port: %+v\n", fixedport)

	rdb := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(ip, fixedport),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	t.Logf(" -> got a client: %+v\n", rdb)

	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		t.Fatal(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		t.Log("key2 does not exist, as expected")
	} else if err != nil {
		t.Fatal(err)
	} else {
		t.Log("key2", val2)
	}
}
