package util

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/store"
)

func GetFirstValueInRange[T any](storeObj store.KVStore, keyStart []byte, keyEnd []byte, reverseIterate bool, parseValue func([]byte) (T, error)) (T, error) {
	// Get the last value if it exist because iterator not catch the end key
	bz := storeObj.Get(keyEnd)
	if bz != nil {
		return parseValue(bz)
	}

	iterator := makeIterator(storeObj, keyStart, keyEnd, reverseIterate)
	defer iterator.Close()

	if !iterator.Valid() {
		var blankValue T
		return blankValue, errors.New("no values in range")
	}

	return parseValue(iterator.Value())
}

func GetValueInRange[T any](storeObj store.KVStore, keyStart []byte, keyEnd []byte, reverseIterate bool, parseValue func([]byte) (T, error)) ([]T, error) {
	var entryList []T

	iterator := makeIterator(storeObj, keyStart, keyEnd, reverseIterate)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		entry, err := parseValue(iterator.Value())
		if err != nil {
			return nil, err
		}
		entryList = append(entryList, entry)
	}

	// Get the last value if it exist because iterator not catch the end key
	bz := storeObj.Get(keyEnd)
	if bz != nil {
		entry, err := parseValue(bz)
		if err != nil {
			return nil, err
		}
		entryList = append(entryList, entry)
	}

	// return err if not have any value in range
	if len(entryList) == 0 {
		return entryList, errors.New("no values in range")
	}

	return entryList, nil
}

func RemoveValueInRange(storeObj store.KVStore, keyStart []byte, keyEnd []byte, reverseIterate bool) {
	iterator := makeIterator(storeObj, keyStart, keyEnd, reverseIterate)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		storeObj.Delete(iterator.Key())
	}
}

func makeIterator(storeObj store.KVStore, keyStart []byte, keyEnd []byte, reverse bool) store.Iterator {
	if reverse {
		return storeObj.ReverseIterator(keyStart, keyEnd)
	}
	return storeObj.Iterator(keyStart, keyEnd)
}
