# bazel-redis-unittest
Experimentation into Bazel-built Redis client with a unittest driving a test-containers turnup of redis.

This is a demonstration of turning up a container for unittesting, specifically exercises with a
Redis container.


# Building

This is Bazel, so...

    bazel build //...:all

# Testing

This is Bazel, so ...

    bazel test //...:all

# What does it do?

This project is a minimalist example of turning up a redis container (with clean slate of state)
for unittests, turning it down when complete, and allowing stateful transitions of data GET/SET in
that container.  This example is intented to be guidance for turning up live redis for unittesting
rather than a complete project.

This project has a single unittest that encloses unittest actions with:

```
    req := testcontainers.ContainerRequest{
        Image:        "redis:3-alpine",
        ExposedPorts: []string{"6379/tcp"},
    }

    redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
        ContainerRequest: req,
        Started:          true,
    })
    if err != nil {
        t.Fatal(err)
    }
    defer redisC.Terminate(ctx)
```

Essentially, this brackets the unittests with the orchestration of the redis K/V store,
automatically requesting that it turnup before unittests, and down when complete.

In a larger environment, this might be in the setup() for a group/array/collection of tests,
perhaps separating groups into different DBs.  Such would amortize the turnup/turndown cost over
all the tests.

... and the tests PASS :)
