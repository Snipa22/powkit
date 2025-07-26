//go:generate ../.bin/gen-lookup -package octopus -mixBytes 256 -cacheInit 16777216 -cacheGrowth 65536 -datasetInit 4294967296 -datasetGrowth 16777216

package octopus

import (
	"fmt"
	"runtime"

	"github.com/snipa22/powkit/support/common"
	"github.com/snipa22/powkit/support/dag"
)

type Client struct {
	data *dag.DAG
}

func New(cfg dag.Config) *Client {
	client := &Client{
		data: dag.New(cfg),
	}

	return client
}

func NewConflux() *Client {
	var cfg = dag.Config{
		Name:       "CFX",
		Revision:   1,
		StorageDir: common.DefaultDir(".powcache"),

		DatasetInitBytes:   2 * (1 << 31),
		DatasetGrowthBytes: 1 << 24,
		CacheInitBytes:     2 * (1 << 23),
		CacheGrowthBytes:   1 << 16,

		CacheSizes:   dag.NewLookupTable(cacheSizes, 2048),
		DatasetSizes: dag.NewLookupTable(datasetSizes, 2048),

		MixBytes:        256,
		DatasetParents:  256,
		EpochLength:     1 << 19,
		SeedEpochLength: 1 << 19,

		CacheRounds:    3,
		CachesCount:    3,
		CachesLockMmap: false,

		L1Enabled: false,
	}

	return New(cfg)
}

func (c *Client) Compute(hash []byte, height, nonce uint64) ([]byte, error) {
	if len(hash) != 32 {
		return nil, fmt.Errorf("hash must be 32 bytes")
	}

	epoch := c.data.CalcEpoch(height)
	size := c.data.DatasetSize(epoch)
	cache := c.data.GetCache(epoch)
	lookup := c.data.NewLookupFunc512(cache, epoch)

	digest := octopus(hash, nonce, size, lookup)
	runtime.KeepAlive(cache)

	return digest, nil
}
