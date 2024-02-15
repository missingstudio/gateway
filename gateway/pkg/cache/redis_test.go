package cache

import (
	"context"
	"fmt"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

type RedisTestSuite struct {
	suite.Suite
	pool        *dockertest.Pool
	resource    *dockertest.Resource
	redisConfig *RedisConfig
	redisClient ICache
}

func (suite *RedisTestSuite) SetupSuite() {
	pool, err := dockertest.NewPool("")
	if err != nil {
		suite.FailNow("Could not connect to Docker", err)
	}

	// Define the Redis container options
	redisVersion := "7.2-alpine"
	redisOptions := &dockertest.RunOptions{
		Repository: "redis",
		Tag:        redisVersion,
	}

	resource, err := pool.RunWithOptions(redisOptions)
	if err != nil {
		suite.FailNow("Could not start Redis container", err)
	}

	suite.pool = pool
	suite.resource = resource

	if err := suite.pool.Retry(func() error {
		client := redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("localhost:%s", resource.GetPort("6379/tcp")),
		})
		return client.Ping(context.Background()).Err()
	}); err != nil {
		suite.FailNow("Could not connect to Redis container", err)
	}

	// Set up RedisConfig for the test
	suite.redisConfig = &RedisConfig{
		Host:     "localhost",
		Port:     resource.GetPort("6379/tcp"),
		Database: 0,
		Password: "",
	}
}

func (suite *RedisTestSuite) TearDownSuite() {
	if err := suite.pool.Purge(suite.resource); err != nil {
		suite.Fail("Could not purge resource", err)
	}
}

func (suite *RedisTestSuite) SetupTest() {
	// Create a new Redis client before each test
	redisClient, err := NewRedisClient(suite.redisConfig)
	if err != nil {
		suite.Fail("Error creating Redis client", err)
	}
	suite.redisClient = redisClient
}

func (suite *RedisTestSuite) TearDownTest() {}

func TestRedisTestSuite(t *testing.T) {
	suite.Run(t, new(RedisTestSuite))
}
