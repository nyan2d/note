package util

import "go.etcd.io/bbolt"

func ReadBoltNodesCount(bolt *bbolt.DB, bucketName string) int {
	result := 0
	bolt.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return nil
		}
		result = bucket.Stats().KeyN
		return nil
	})

	return result
}
