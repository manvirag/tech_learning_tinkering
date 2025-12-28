package utils

const BASE62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GetBase62(val int64) string {
	if val == 0 {
		return string(BASE62[0])
	}
	var result string
	for val > 0 {
		remainder := val % 62
		result = string(BASE62[remainder]) + result
		val /= 62
	}
	return result

}
