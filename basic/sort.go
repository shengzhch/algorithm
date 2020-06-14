package basic

import (
	"math"
)

//插入排序 稳定
//将数据直接插入到排好的序列当中
func InsertSort(data []interface{}, cf CF) {
	var i, j int
	var size = len(data)
	var key interface{}
	for j = 1; j < size; j++ {
		//取出结点
		key = data[j]
		i = j - 1

		//将当前结点插入到第一个不大于它的后面
		for i >= 0 && cf(data[i], key) > 0 {
			data[i+1], data[i] = data[i], data[i+1]
			i--
		}
	}
}

//冒泡排序 稳定
//最大的数据依次移动到序列的最后
func BuddleSort(data []interface{}, cf CF) {
	size := len(data)
	for i := 0; i < size-1; i++ {
		finish := true
		//最大的放到了最后
		for j := 0; j < size-1-i; j++ {
			if cf(data[j], data[j+1]) > 0 {
				data[j], data[j+1] = data[j+1], data[j]
				finish = false
			}
		}
		if finish {
			return
		}
	}
}

//计数排序 稳定 扩展了计数排序
//在对一定范围内的整数排序，以空间换时间
func CountSort(data []int, min, max int) {
	size := len(data)
	var count = make([]int, max-min+1) //取值范围区间
	var tmp = make([]int, size)

	var i, j int
	//min-max 映射到 0 max-min
	for i = 0; i <= max-min; i++ {
		count[i] = 0
	}

	//偏移量是元素对比最小值计算

	//记录了每个整数相对偏移量出现的次数，index代表偏移量，value代表出现的次数
	for j = 0; j < size; j++ {
		offset := data[j] - min
		count[offset] = count[offset] + 1
	}

	//每个位置加上前一位置的偏移量出现的次数,index代表偏移量，value代表出现的该偏移量在新数组中的位置
	for i = 1; i <= max-min; i++ {
		count[i] = count[i] + count[i-1]
	}

	for j = 0; j < size; j++ {
		//-1 是因为下标从0开始
		index := count[data[j]-min] - 1
		tmp[index] = data[j]
		count[data[j]-min] = count[data[j]-min] - 1
	}

	copy(data[:], tmp[:])
}

//基数排序 稳定
//先按个位排，在按十位排，一直排到最高位
func RadixSort(data []int, size, p, k int) {
	//size 数组大小
	//p 位数
	//k 基数 10 ， 10^0 =1，10^1 = 10 ，10^2 = 100

	var (
		count                = make([]int, k)
		tmp                  = make([]int, size)
		index, pval, i, j, n int
	)

	//对每一位上运行计数排序
	for n = 0; n < p; n++ {

		//每一位的计数数组清空
		for i = 0; i < k; i++ {
			count[i] = 0
		}

		//求基础
		pval = int(math.Pow(float64(k), float64(n))) // k^0 k^1 ... k^n

		for j = 0; j < size; j++ {
			//
			index = (data[j] / pval) % k
			count[index] = count[index] + 1
		}

		//偏移量
		for i = 1; i < k; i++ {
			count[i] = count[i] + count[i-1]
		}

		//先出现的index应该最小，所以从最大处开始赋值
		for j = size - 1; j >= 0; j-- {
			index = count[(data[j]/pval)%k] - 1
			tmp[index] = data[j]
			count[(data[j]/pval)%k] = count[(data[j]/pval)%k] - 1
		}

		copy(data[:], tmp[:])
	}
}

//归并排序 稳定

//合并到一起，j是分割的位置 i是开始位置，k是结束位置
func merge(data []interface{}, i, j, k int, cf CF) int {
	var m = make([]interface{}, k-i+1)
	ipos := i
	jpos := j + 1
	mpos := 0

	for (ipos <= j || jpos <= k) {
		//处理两组数据剩余的部分,肯定只有一组剩余
		if (ipos > j) {
			//把所有剩余的元素填充到m中
			for (jpos <= k) {
				m[mpos] = data[jpos]
				jpos++
				mpos++
			}

			continue

		} else if (jpos > k) {
			for (ipos <= j) {
				m[mpos] = data[ipos]
				ipos++
				mpos++
			}
			continue
		}

		//分别从两组的数据头开始，按顺序放入到m中
		if (cf(data[ipos], data[jpos]) < 0) {
			m[mpos] = data[ipos]
			ipos++
			mpos++
		} else {
			m[mpos] = data[jpos]
			jpos++
			mpos++
		}
	}

	copy(data[i:k+1], m[:])
	return 0
}

//i为0 k为size-1
func MgSort(data []interface{}, i, k int, cf CF) int {
	var j int
	if (i < k) {
		j = (i + k - 1) / 2 //j处于中间位置

		//处理左
		if MgSort(data, i, j, cf) < 0 {
			return -1
		}

		//处理右
		if MgSort(data, j+1, k, cf) < 0 {
			return -1
		}

		//合并到一起
		if merge(data, i, j, k, cf) < 0 {
			return -1
		}
	}
	return 0
}

//快速排序
func QuickSort(data []interface{}, i, k int, cf CF) int {
	return 0
}
