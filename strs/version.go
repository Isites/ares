package strs

import (
	"strconv"
	"strings"

	"github.com/Isites/ares/lru"
)

// 从i开始返回下一个到.的位置
func diffPart(i int, rv []rune) (j int) {
	j = i
	for j = i; j < len(rv); j++ {
		// if rv[j] == '0' {
		// 	offset++
		// }
		if rv[j] == '.' {
			return j
		}
	}
	return j
}

// CompareVersion 比较两个appversion的大小
// return 0 means ver1 == ver2
// return 1 means ver1 > ver2
// return -1 means ver1 < ver2
func CompareVersion(ver1, ver2 string) (ret int) {
	defer func() {
		if ret > 0 {
			ret = 1
		} else if ret < 0 {
			ret = -1
		}
	}()
	if ver1 == ver2 {
		return 0
	}
	rv1, rv2 := []rune(ver1), []rune(ver2)
	var (
		l1, l2 = len(ver1), len(ver2)
		i, j   = 0, 0
	)
	for i < l1 && j < l2 {
		if rv1[i] == rv2[j] {
			i++
			j++
			continue
		}
		k := diffPart(i, rv1)
		rv1Str := string(rv1[i:k])
		curV1, e1 := strconv.Atoi(rv1Str)
		i = k
		k = diffPart(j, rv2)
		rv2Str := string(rv2[j:k])
		curV2, e2 := strconv.Atoi(rv2Str)
		j = k
		if e1 != nil || e2 != nil {
			ret = strings.Compare(rv1Str, rv2Str)
		} else {
			ret = curV1 - curV2
		}
		if ret != 0 {
			return ret
		}
	}
	if i < l1 {
		ret = 1
	} else if j < l2 {
		ret = -1
	}
	return ret
}

type cmVal struct {
	iv int
	sv string
	// 能否转换为整形
	canInt bool
}

func strs2cmVs(strs []string) []*cmVal {
	cmvs := make([]*cmVal, 0, len(strs))
	for _, v := range strs {
		it, e := strconv.Atoi(v)
		// 全部数据都保存
		cmvs = append(cmvs, &cmVal{it, v, e == nil})
	}
	return cmvs
}

var (
	// 暂时固定1000个版本大小的缓存
	cache = lru.New(1000)
)

// CompareVersion 比较两个appversion的大小
// return 0 means ver1 == ver2
// return 1 means ver1 > ver2
// return -1 means ver1 < ver2
func CompareVersionWithCache(ver1, ver2 string) int {
	// ver1 = strings.TrimSpace(ver1)
	// ver2 = strings.TrimSpace(ver2)
	// fast path
	if ver1 == ver2 {
		return 0
	}
	// slow path
	var (
		cmv1, cmv2             []*cmVal
		cmv1Exists, cmv2Exists bool
	)
	// read cache 1
	cmv, cmvExists := cache.Get(ver1)
	if cmvExists {
		cmv1, cmv1Exists = cmv.([]*cmVal)
	}
	if !cmv1Exists {
		// set val and cache
		cmv1 = strs2cmVs(strings.Split(ver1, "."))
		cache.Set(ver1, cmv1)
	}
	// read cache 2
	cmv, cmvExists = cache.Get(ver2)
	if cmvExists {
		cmv2, cmv2Exists = cmv.([]*cmVal)
	}
	if !cmv2Exists {
		// set val and cache
		cmv2 = strs2cmVs(strings.Split(ver2, "."))
		cache.Set(ver2, cmv2)
	}
	// compare ver str
	v1l, v2l := len(cmv1), len(cmv2)
	for i := 0; i < len(cmv1) && i < len(cmv2); i++ {
		res := 0
		// can use int compare
		if cmv1[i].canInt && cmv2[i].canInt {
			res = cmv1[i].iv - cmv2[i].iv
		} else {
			res = strings.Compare(cmv1[i].sv, cmv2[i].sv)
		}
		if res > 0 {
			return 1
		} else if res < 0 {
			return -1
		}
	}
	if v1l > v2l {
		return 1
	} else if v1l < v2l {
		return -1
	}
	return 0
}
