package redis_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sanksons/tavern/common/entity"
	"github.com/sanksons/tavern/common/errors"
)

var _ = Describe("Redis Cluster AdapterC", func() {
	Describe("Checking Single Key Behaviour", func() {

		//
		// Set key behaviour check
		//
		Context("SET KEY", func() {
			It("should succeed", func() {
				err := AdapterC.Set(entity.CacheItem{
					Key:        entity.CacheKey("key1"),
					Value:      []byte("I am Key1"),
					Expiration: 0,
				})
				Expect(err).To(BeNil())

			})
		})

		//
		// Get key behaviour check
		//
		Context("GET KEY", func() {
			It("key1 should match `I am Key1`", func() {
				dataBytes, err := AdapterC.Get(entity.CacheKey("key1"))
				Expect(err).To(BeNil())
				Expect(dataBytes).To(Equal([]byte("I am Key1")))

			})
			It("key2 should match `not found`", func() {
				_, err := AdapterC.Get(entity.CacheKey("key2"))
				Expect(err).To(Equal(errors.KeyNotExists))
			})
		})

		//
		// Delete key behaviour check
		//
		Context("DELETE KEY", func() {
			It("should succeed", func() {
				m, err := AdapterC.Destroy(entity.CacheKey("key1"))

				Expect(err).To(BeNil())
				Expect(m).To(Equal(map[entity.CacheKey]bool{"key1": true}))
			})
			It("should actually be deleted", func() {
				_, err := AdapterC.Get(entity.CacheKey("key1"))
				Expect(err).To(Equal(errors.KeyNotExists))
			})

		})

	})

	Describe("Checking MULTI Key Behaviour", func() {

		//
		// Set multi key behaviour check
		//
		Context("SET Multi KEY", func() {
			It("should succeed", func() {

				Items := []entity.CacheItem{
					entity.CacheItem{
						Key:        entity.CacheKey("mkey1"),
						Value:      []byte("I am mkey1"),
						Expiration: time.Duration(340) * time.Second,
					},
					entity.CacheItem{
						Key:        entity.CacheKey("mkey2"),
						Value:      []byte("I am mkey2"),
						Expiration: time.Duration(340) * time.Second,
					},
					entity.CacheItem{
						Key:        entity.CacheKey("mkey3"),
						Value:      []byte("I am mkey3"),
						Expiration: time.Duration(340) * time.Second,
					},
				}

				result, err := AdapterC.MSet(Items...)
				Expect(err).To(BeNil())
				Expect(result).To(Equal(map[entity.CacheKey]bool{"mkey1": true, "mkey2": true, "mkey3": true}))

			})
		})

		//
		// Get key behaviour check
		//
		Context("GET Multi KEY", func() {
			It("keys should match their content", func() {
				dataBytes, err := AdapterC.MGet(
					entity.CacheKey("mkey1"), entity.CacheKey("mkey2"), entity.CacheKey("mkey3"),
				)
				Expect(err).To(BeNil())
				for k, v := range dataBytes {
					if k == "mkey1" {
						Expect(v).To(Equal([]byte("I am mkey1")))
					}
					if k == "mkey2" {
						Expect(v).To(Equal([]byte("I am mkey2")))
					}
					if k == "mkey3" {
						Expect(v).To(Equal([]byte("I am mkey3")))
					}
				}

			})

		})

		//
		// Delete key behaviour check
		//
		Context("DELETE Multi KEY", func() {
			It("should succeed", func() {
				m, err := AdapterC.Destroy(entity.CacheKey("mkey1"), entity.CacheKey("mkey2"))

				Expect(err).To(BeNil())
				Expect(m).To(Equal(map[entity.CacheKey]bool{"mkey1": true, "mkey2": true}))
			})
			It("should actually be deleted", func() {
				_, err1 := AdapterC.Get(entity.CacheKey("mkey1"))
				Expect(err1).To(Equal(errors.KeyNotExists))

				_, err2 := AdapterC.Get(entity.CacheKey("mkey2"))
				Expect(err2).To(Equal(errors.KeyNotExists))
			})

		})

	})

})
