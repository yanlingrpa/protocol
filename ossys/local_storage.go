package ossys

/*
* LocalStorage defines local key-value storage capabilities.
* This interface supports single and batch read/write operations, as well as expiration settings.
 */
type LocalStorage interface {
	/*
	* Keys gets all keys in the current storage.
	* Returns a list of key names.
	 */
	Keys() []string

	/*
	* Get gets the value for the specified key.
	* If the key does not exist, the return value depends on the implementation.
	 */
	Get(key string) string

	/*
	* Set sets a key-value pair.
	* key is the key name, and value is the value.
	 */
	Set(key, value string)

	/*
	* SetEx sets a key-value pair with an expiration time.
	* expiredAt is the expiration timestamp.
	 */
	SetEx(key, value string, expiredAt int64)

	/*
	* Del deletes one or more keys.
	* Returns the number of keys actually deleted.
	 */
	Del(keys ...string) int

	/*
	* MGet gets values for multiple keys in batch.
	* Returns a key-value mapping result.
	 */
	MGet(keys ...string) map[string]string

	/*
	* MSet sets multiple key-value pairs in batch.
	* data is the key-value mapping to write.
	 */
	MSet(data map[string]string)

	/*
	* MSetEx sets multiple key-value pairs with expiration in batch.
	* expiredAt is the unified expiration timestamp.
	 */
	MSetEx(data map[string]string, expiredAt int64)
}
