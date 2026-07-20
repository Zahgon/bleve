package util

import (
	"os"

	bolt "go.etcd.io/bbolt"
)

type RootBoltImpl struct {
	*bolt.DB
}

type BoltTxImpl struct {
	*bolt.Tx
}

type BoltBucketImpl struct {
	*bolt.Bucket

	name string
}

func OpenBolt(path string, mode os.FileMode, options *bolt.Options) (*RootBoltImpl, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RootBoltImpl) Begin(writable bool) (*BoltTxImpl, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RootBoltImpl) View(fn func(*BoltTxImpl) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RootBoltImpl) Update(fn func(*BoltTxImpl) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *BoltTxImpl) CreateBucketIfNotExists(name []byte) (*BoltBucketImpl, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *BoltTxImpl) Bucket(name []byte) *BoltBucketImpl { _ = "STUB: not implemented"; return nil }

func (b *BoltBucketImpl) GetBucket(name []byte) *BoltBucketImpl {
	_ = "STUB: not implemented"
	return nil
}

func (b *BoltBucketImpl) CreateBucketIfNotExists(name []byte) (*BoltBucketImpl, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BoltBucketImpl) ForEach(fn func(key []byte, value []byte) error, reader FileReader) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *BoltBucketImpl) Put(key []byte, value []byte, writer FileWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *BoltBucketImpl) Get(key []byte, reader FileReader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
