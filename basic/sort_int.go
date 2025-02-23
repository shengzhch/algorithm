package basic

import (
	"math"
)

//基数排序
func RadixSortInt(data []int, p, k int) {
	size := len(data)

	var (
		tmp = make([]int, size)
	)

	//对每一位上运行计数排序
	for n := 0; n < p; n++ {
		count := make([]int, k)

		//10 100 1000 10000
		p := int(math.Pow(float64(k), float64(n)))

		//12 11 123 => 1 1 3
		for j := 0; j < size; j++ {
			weishu := (data[j] / p) % k
			count[weishu]++
		}

		for j := 0; j < k; j++ {
			if j == 0 {
				count[j] = count[j] - 1
			} else {
				count[j] = count[j] + count[j-1]
			}
		}

		for j := size - 1; j >= 0; j++ {
			weishu := (data[j] / p) % k
			tmp[count[weishu]] = data[j]
			count[weishu]--
		}

		copy(data[:], tmp[:])
	}
}

//挖坑法
func quickInt(data []int, start, end int) {
	if start < end {
		pval := data[start]
		i := start
		j := end
		for i < j {
			//从右端开始向左搜索找到第一个小于等于哨兵的值
			for i < j && data[j] > pval {
				j--
			}
			if i < j {
				data[i] = data[j]
				i++
			}
			//从左端开始向右搜索找到第一个大于等于哨兵的值
			for i < j && data[i] < pval {
				i++
			}

			if i < j {
				data[j] = data[i]
				j--
			}
		}
		data[i] = pval
		quickInt(data, start, i-1)
		quickInt(data, i+1, end)
	}
}

func QuickInt(arr []int) {
	quickInt(arr, 0, len(arr)-1)
}
