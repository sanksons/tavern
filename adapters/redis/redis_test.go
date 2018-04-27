package redis_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sanksons/tavern/utils"
)

var _ = Describe("Redis Simple Adapter", func() {
	Describe("Checking Single Key Behaviour", func() {

		//
		// Set key behaviour check
		//
		Context("SET KEY", func() {
			It("should succeed", func() {
				err := Adapter.Set(utils.CacheItem{
					Key:        utils.CacheKey("key1"),
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
				dataBytes, err := Adapter.Get(utils.CacheKey("key1"))
				Expect(err).To(BeNil())
				Expect(dataBytes).To(Equal([]byte("I am Key1")))

			})
			It("key2 should match `not found`", func() {
				_, err := Adapter.Get(utils.CacheKey("key2"))
				Expect(err).To(Equal(utils.KeyNotExists))
			})
		})

		//
		// Delete key behaviour check
		//
		Context("DELETE KEY", func() {
			It("should succeed", func() {
				m, err := Adapter.Destroy(utils.CacheKey("key1"))

				Expect(err).To(BeNil())
				Expect(m).To(Equal(map[utils.CacheKey]bool{"key1": true}))
			})
			It("should actually be deleted", func() {
				_, err := Adapter.Get(utils.CacheKey("key1"))
				Expect(err).To(Equal(utils.KeyNotExists))
			})

		})

	})

	Describe("Checking MULTI Key Behaviour", func() {

		//
		// Set multi key behaviour check
		//
		Context("SET Multi KEY", func() {
			It("should succeed", func() {

				Items := []utils.CacheItem{
					utils.CacheItem{
						Key:        utils.CacheKey("mkey1"),
						Value:      []byte("I am mkey1"),
						Expiration: time.Duration(340) * time.Second,
					},
					utils.CacheItem{
						Key:        utils.CacheKey("mkey2"),
						Value:      []byte("I am mkey2"),
						Expiration: time.Duration(340) * time.Second,
					},
					utils.CacheItem{
						Key:        utils.CacheKey("mkey3"),
						Value:      []byte("I am mkey3"),
						Expiration: time.Duration(340) * time.Second,
					},
				}

				result, err := Adapter.MSet(Items...)
				Expect(err).To(BeNil())
				Expect(result).To(Equal(map[utils.CacheKey]bool{"mkey1": true, "mkey2": true, "mkey3": true}))

			})
		})

		//
		// Get key behaviour check
		//
		Context("GET Multi KEY", func() {
			It("keys should match their content", func() {
				dataBytes, err := Adapter.MGet(
					utils.CacheKey("mkey1"), utils.CacheKey("mkey2"), utils.CacheKey("mkey3"),
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
				m, err := Adapter.Destroy(utils.CacheKey("mkey1"), utils.CacheKey("mkey2"))

				Expect(err).To(BeNil())
				Expect(m).To(Equal(map[utils.CacheKey]bool{"mkey1": true, "mkey2": true}))
			})
			It("should actually be deleted", func() {
				_, err1 := Adapter.Get(utils.CacheKey("mkey1"))
				Expect(err1).To(Equal(utils.KeyNotExists))

				_, err2 := Adapter.Get(utils.CacheKey("mkey2"))
				Expect(err2).To(Equal(utils.KeyNotExists))
			})

		})

	})

})
