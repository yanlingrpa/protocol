package ossys

type LocalStorage interface {
	Keys() []string
	Get(key string) string
	Set(key, value string)
	SetEx(key, value string, expiredAt int64)
	Del(keys ...string) int
	MGet(keys ...string) map[string]string
	MSet(data map[string]string)
	MSetEx(data map[string]string, expiredAt int64)
}
