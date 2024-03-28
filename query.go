package httpio

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Query struct {
	query map[string][]string
}

func NewQuery(r *http.Request) *Query {
	query := r.URL.Query()
	return &Query{query: query}
}

func (q *Query) GetStringFromQuery(key string) string {
	if q.query[key] == nil {
		return ""
	}
	return q.query[key][0]
}

func (q *Query) GetBoolFromQuery(key string) (bool, error) {
	if q.query[key] == nil {
		return false, nil
	}
	return strconv.ParseBool(q.query[key][0])
}

func (q *Query) GetTimeFromQuery(key string) (time.Time, error) {
	if q.query[key] == nil {
		return time.Time{}, nil
	}
	return time.Parse(time.RFC3339, q.query[key][0])
}

// GetIntFromQuery bitSize 0 for int, 32 for int32, 64 for int64
func (q *Query) GetIntFromQuery(key string, bitSize int) (int64, error) {
	if q.query[key] == nil {
		return 0, nil
	}
	return strconv.ParseInt(q.query[key][0], 10, bitSize)
}

// GetFloatFromQuery bitSize 32 for float32, 64 for float64
func (q *Query) GetFloatFromQuery(key string, bitSize int) (float64, error) {
	if q.query[key] == nil {
		return 0, nil
	}
	return strconv.ParseFloat(q.query[key][0], bitSize)
}

// GetUintFromQuery bitSize 0 for uint, 32 for uint32, 64 for uint64
func (q *Query) GetUintFromQuery(key string, bitSize int) (uint, error) {
	if q.query[key] == nil {
		return 0, nil
	}
	i, err := strconv.ParseUint(q.query[key][0], 10, bitSize)
	return uint(i), err
}

func (q *Query) GetStringSliceFromQuery(keyPrefix string) []string {
	var result []string
	for key, values := range q.query {
		if strings.HasPrefix(key, keyPrefix) {
			if len(values) > 0 {
				result = append(result, values[0])
			}
		}
	}
	return result
}

// GetIntSliceFromQuery bitSize 0 for int, 32 for int32, 64 for int64
func (q *Query) GetIntSliceFromQuery(keyPrefix string, bitSize int) []interface{} {
	var result []interface{}
	for key, values := range q.query {
		if strings.HasPrefix(key, keyPrefix) {
			if len(values) > 0 {
				i, err := strconv.ParseInt(values[0], 10, bitSize)
				if err != nil {
					continue
				}
				switch bitSize {
				case 8:
					result = append(result, int8(i))
				case 16:
					result = append(result, int16(i))
				case 32:
					result = append(result, int32(i))
				case 64:
					result = append(result, i)
				default:
					result = append(result, int(i))
				}
			}
		}
	}
	return result
}

// GetUintSliceFromQuery bitSize 0 for uint, 32 for uint32, 64 for uint64
func (q *Query) GetUintSliceFromQuery(keyPrefix string, bitSize int) []interface{} {
	var result []interface{}
	for key, values := range q.query {
		if strings.HasPrefix(key, keyPrefix) {
			if len(values) > 0 {
				i, err := strconv.ParseUint(values[0], 10, bitSize)
				if err != nil {
					continue
				}
				switch bitSize {
				case 8:
					result = append(result, uint8(i))
				case 16:
					result = append(result, uint16(i))
				case 32:
					result = append(result, uint32(i))
				case 64:
					result = append(result, i)
				default:
					result = append(result, uint(i))
				}
			}
		}
	}
	return result
}

// GetFloatSliceFromQuery bitSize 32 for float32, 64 for float64
func (q *Query) GetFloatSliceFromQuery(keyPrefix string, bitSize int) []interface{} {
	var result []interface{}
	for key, values := range q.query {
		if strings.HasPrefix(key, keyPrefix) {
			if len(values) > 0 {
				i, err := strconv.ParseFloat(values[0], bitSize)
				if err != nil {
					continue
				}
				switch bitSize {
				case 32:
					result = append(result, float32(i))
				case 64:
					result = append(result, i)
				default:
					result = append(result, float64(i))
				}
			}
		}
	}
	return result
}

// GetBoolSliceFromQuery bitSize 0 for bool
func (q *Query) GetBoolSliceFromQuery(keyPrefix string) []bool {
	var result []bool
	for key, values := range q.query {
		if strings.HasPrefix(key, keyPrefix) {
			if len(values) > 0 {
				i, err := strconv.ParseBool(values[0])
				if err != nil {
					continue
				}
				result = append(result, i)
			}
		}
	}
	return result
}

// GetTimeSliceFromQuery bitSize 0 for time.Time
func (q *Query) GetTimeSliceFromQuery(keyPrefix string) []time.Time {
	var result []time.Time
	for key, values := range q.query {
		if strings.HasPrefix(key, keyPrefix) {
			if len(values) > 0 {
				i, err := time.Parse(time.RFC3339, values[0])
				if err != nil {
					continue
				}
				result = append(result, i)
			}
		}
	}
	return result
}
