/*
Copyright 2012 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package groupcache

// A Sink receives data from a Get call.
//
// Implementation of Getter must call exactly one of the Set methods
// on success.
type Sink interface {
	Set(v CacheEntry) error

	// view returns a frozen view of the bytes for caching.
	view() (CacheEntry, error)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func CacheEntrySink(dst *CacheEntry) Sink {
	return &cacheEntrySink{dst: dst}
}

type cacheEntrySink struct {
	dst *CacheEntry
}

func (s *cacheEntrySink) Set(v CacheEntry) error {
	*s.dst = CacheEntry{
		data: cloneBytes(v.data),
		meta: cloneBytes(v.meta),
	}
	return nil
}

func (s *cacheEntrySink) setView(v CacheEntry) error {
	*s.dst = v
	return nil
}

func (s *cacheEntrySink) view() (CacheEntry, error) {
	return *s.dst, nil
}

// AllocatingByteSliceSink returns a Sink that allocates
// a byte slice to hold the received value and assigns
// it to *dst. The memory is not retained by groupcache.
func AllocatingCacheEntrySink(data *[]byte, meta *[]byte) Sink {
	return &allocCacheEntrySink{data: data, meta: meta}
}

type allocCacheEntrySink struct {
	data *[]byte
	meta *[]byte
	v    CacheEntry
}

func (s *allocCacheEntrySink) view() (CacheEntry, error) {
	return s.v, nil
}

func (s *allocCacheEntrySink) setView(v CacheEntry) error {
	// TODO: is this supposed to clone or not? I don't get it.
	if v.data != nil {
		*s.data = cloneBytes(v.data)
	}
	if v.meta != nil {
		*s.meta = cloneBytes(v.meta)
	}
	s.v = v
	return nil
}

func (s *allocCacheEntrySink) Set(v CacheEntry) error {
	s.setView(CacheEntry{
		data: cloneBytes(v.data),
		meta: cloneBytes(v.meta),
	})
	return nil
}
