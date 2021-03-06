package gocb

// RemoveMt performs a Remove operation and includes MutationToken in the results.
func (b *Bucket) RemoveMt(key string, cas Cas) (Cas, MutationToken, error) {
	if !b.mtEnabled {
		panic("You must use OpenBucketMt with Mt operation variants.")
	}
	return b.remove(key, cas)
}

// UpsertMt performs a Upsert operation and includes MutationToken in the results.
func (b *Bucket) UpsertMt(key string, value interface{}, expiry uint32) (Cas, MutationToken, error) {
	if !b.mtEnabled {
		panic("You must use OpenBucketMt with Mt operation variants.")
	}
	return b.upsert(key, value, expiry)
}

// InsertMt performs a Insert operation and includes MutationToken in the results.
func (b *Bucket) InsertMt(key string, value interface{}, expiry uint32) (Cas, MutationToken, error) {
	if !b.mtEnabled {
		panic("You must use OpenBucketMt with Mt operation variants.")
	}
	return b.insert(key, value, expiry)
}

// ReplaceMt performs a Replace operation and includes MutationToken in the results.
func (b *Bucket) ReplaceMt(key string, value interface{}, cas Cas, expiry uint32) (Cas, MutationToken, error) {
	if !b.mtEnabled {
		panic("You must use OpenBucketMt with Mt operation variants.")
	}
	return b.replace(key, value, cas, expiry)
}

// AppendMt performs a Append operation and includes MutationToken in the results.
func (b *Bucket) AppendMt(key, value string) (Cas, MutationToken, error) {
	if !b.mtEnabled {
		panic("You must use OpenBucketMt with Mt operation variants.")
	}
	return b.append(key, value)
}

// PrependMt performs a Prepend operation and includes MutationToken in the results.
func (b *Bucket) PrependMt(key, value string) (Cas, MutationToken, error) {
	if !b.mtEnabled {
		panic("You must use OpenBucketMt with Mt operation variants.")
	}
	return b.prepend(key, value)
}

// CounterMt performs a Counter operation and includes MutationToken in the results.
func (b *Bucket) CounterMt(key string, delta, initial int64, expiry uint32) (uint64, Cas, MutationToken, error) {
	if !b.mtEnabled {
		panic("You must use OpenBucketMt with Mt operation variants.")
	}
	return b.counter(key, delta, initial, expiry)
}
