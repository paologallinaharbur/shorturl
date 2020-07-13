package storage

import (
	"errors"
	"github.com/boltdb/bolt"
	"log"
	"sync"
)

//Storage interface is the contract that a struct should meet in order to be used as storage for the API
type Storage interface {
	Write(short string, long string) error
	Delete(short string) error
	Read(short string) (string, error)
}

//StorageDB implements Storage interface and it is used to save URL in a real key/value db
type StorageDB struct {
	db    *bolt.DB
	mutex sync.RWMutex
}

//NewStorageDB create a new instance of the db
func NewStorageDB(dbName string) *StorageDB {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return &StorageDB{
		db:    db,
		mutex: sync.RWMutex{},
	}
}

//Close cleanUp resources
func (db StorageDB) Close() error {
	return db.db.Close()
}

//Delete deletes a short URL from the db
func (db StorageDB) Delete(short string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	err := db.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("urlBucket"))
		if b == nil {
			return nil
		}
		err := b.Delete([]byte(short))
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

//Delete save a [shortURL-longURL] in the db
func (db StorageDB) Write(short string, long string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	err := db.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("urlBucket"))
		if err != nil {
			return err
		}
		tmp := b.Get([]byte(short))

		if len(tmp) != 0 {
			return errors.New("shortUrlAlreadyAssigned")
		}

		err = b.Put([]byte(short), []byte(long))
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

//Read retrieves a longURL from a shortURL from the db
func (db StorageDB) Read(short string) (string, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	var long []byte

	err := db.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("urlBucket"))
		if b == nil {
			return errors.New("urlNotFoundInDB")
		}
		long = b.Get([]byte(short))

		if len(long) == 0 {
			return errors.New("urlNotFoundInDB")
		}
		return nil
	})

	return string(long), err
}
