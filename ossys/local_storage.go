package ossys

/*
* LocalStorage 定义本地键值存储能力。
* 该接口支持单项和批量读写操作，以及过期时间设置。
 */
type LocalStorage interface {
	/*
	* Keys 获取当前存储中的所有键。
	* 返回键名列表。
	 */
	Keys() []string

	/*
	* Get 获取指定 key 对应的值。
	* 如果该 key 不存在，返回值由实现决定。
	 */
	Get(key string) string

	/*
	* Set 设置键值对。
	* key 是键名，value 是对应的值。
	 */
	Set(key, value string)

	/*
	* SetEx 设置带过期时间的键值对。
	* expiredAt 是过期时间戳。
	 */
	SetEx(key, value string, expiredAt int64)

	/*
	* Del 删除一个或多个键。
	* 返回实际删除的键数量。
	 */
	Del(keys ...string) int

	/*
	* MGet 批量获取多个 key 的值。
	* 返回键值映射结果。
	 */
	MGet(keys ...string) map[string]string

	/*
	* MSet 批量设置多个键值对。
	* data 是要写入的键值映射。
	 */
	MSet(data map[string]string)

	/*
	* MSetEx 批量设置带过期时间的多个键值对。
	* expiredAt 是统一过期时间戳。
	 */
	MSetEx(data map[string]string, expiredAt int64)
}
