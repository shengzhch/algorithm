package basic

import (
	"fmt"
	"math"
	"math/rand"
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
		//最大的放到了最后,最后面的i个是之前排好序的
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

	//记录了每个整数相对偏移量出现的次数，index代表偏移量，value代表出现的次数
	for j := 0; j < size; j++ {
		offset := data[j] - min
		count[offset] ++
	}

	//每个位置加上前一位置的偏移量出现的次数,index代表偏移量，value代表出现的该偏移量在新数组中的位置
	for i := 0; i <= max-min; i++ {
		if i == 0 {
			//index从0开始
			count[i] = count[i] - 1
		} else {
			count[i] = count[i] + count[i-1]
		}
	}

	for j := size - 1; j >= 0; j-- {
		offset := data[j] - min
		tmp[count[offset]] = data[j]
		count[offset]  --
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
			count[index] ++
		}

		//偏移量
		for i = 0; i < k; i++ {
			if i == 0 {
				count[i] = count[i] - 1
			} else {
				count[i] = count[i] + count[i-1]
			}
		}

		//先出现的index应该最小，所以从最大处开始赋值，保证稳定
		for j = size - 1; j >= 0; j-- {
			index = (data[j] / pval) % k
			tmp[count[index]] = data[j]
			count[index] --
		}

		copy(data[:], tmp[:])
	}
}

//归并排序 稳定

//合并到一起，j是分割的位置 i是开始位置，k是结束位置 (数组下标)
func merge(data []interface{}, i, j, k int, cf CF) int {
	var m = make([]interface{}, k-i+1)
	ipos := i
	jpos := j + 1
	mpos := 0

	for ipos <= j || jpos <= k {
		//处理两组数据剩余的部分,肯定只有一组剩余
		if ipos > j {
			//把所有剩余的元素填充到m中
			for jpos <= k {
				m[mpos] = data[jpos]
				jpos++
				mpos++
			}

			continue

		} else if jpos > k {
			for ipos <= j {
				m[mpos] = data[ipos]
				ipos++
				mpos++
			}
			continue
		}

		//分别从两组的数据头开始，按顺序放入到m中
		if cf(data[ipos], data[jpos]) < 0 {
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
	if i < k {
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

/*
快速排序 解决一般问题的最佳排序算法 分治算法
分 ： 设定一个分割值将数据分成两部分
治 ： 分别在两部分用递归的方式继续使用快速排序算法
合 ： 对分割部分排序直至完成
*/

func randInt() int {
	return rand.Intn(2147483647)
}

//i,k 初试值设置为 0 size-1 分区 原书给的算法有误，有可能陷入死循环,已调整
func Partition(data []interface{}, i, k int, cf CF) int {
	if i == k {
		return i
	}
	var pval interface{}

	var r = make([]interface{}, 3, 3)

	r[0] = (randInt() % (k - i + 1)) + i
	r[1] = (randInt() % (k - i + 1)) + i
	r[2] = (randInt() % (k - i + 1)) + i

	InsertSort(r, CfInt)

	//应该也可以随机取一个值，这样取值的目的是保证绝对随机。
	pval = data[r[1].(int)] //取中位数的方法取出一个值作为哨兵

	//不稳定，交换时不能保证交换到前面的值时第一个出现的值
	for {
		//从右端开始向左搜索找到第一个小于或者等于哨兵的值
		for cf(data[k], pval) > 0 {
			k--
		}

		//从左端开始向右搜索找到第一个大于或者等于哨兵的值
		for cf(data[i], pval) < 0 {
			i++
		}

		//
		if i >= k {
			//则说明哨兵已经把数组分成大值和小值两组，
			break
		} else {
			//肯定是data[i]大于等于data[k])
			if cf(data[i], data[k]) != 0 {
				data[i], data[k] = data[k], data[i]
			} else {
				//防止陷入死循环
				del := k - i
				if del == 1 {
					//紧挨着，说明可以跳出循环
					return i
				} else {
					//更换值
					pval = data[i+1]
				}
			}
		}
	}
	return i
}

//成功返回0，失败返回-1；分区方法调用，确定左边递归调用，右边一直确定新的哨兵，直到i的值移到最后
func QuickSort1(data []interface{}, i, k int, cf CF) int {
	var j int
	for i < k {
		if j = Partition(data, i, k, cf); j < 0 {
			return -1
		}

		if QuickSort1(data, i, j, cf) < 0 {
			return -1
		}
		i = j + 1
	}
	return 0
}

//成功返回0，失败返回-1；分区方法调用
func QuickSort2(data []interface{}, i, k int, cf CF) int {
	if i >= k {
		return 0
	}
	j := Partition(data, i, k, cf)

	if QuickSort2(data, i, j, cf) < 0 {
		return -1
	}

	if QuickSort2(data, j+1, k, cf) < 0 {
		return -1
	}

	return 0
}

/*
归并排序和快速排序的区别和联系
1.联系
原理都是基于分而治之，首先把待排序的数组分为两组，然后分别对两组排序，最后把两组结果合并起来

2.区别
进行分组的策略不同，合并的策略也不同。
归并的分组策略：是假设待排序的元素存放在数组中，那么把数组前面的一半元素作为一组，后面一半作为另一组。
快排的分组厕率：是根据元素的值来分的，大于某个值的元素一组，小于某个值的元素一组。
快速排序在分组的时候已经根据元素的大小来分组了，而合并时，只需要把两个分组合并起来就可以了，归并排序则需要对两个有序的数组根据元素大小合并
*/

//希尔排序 多次插入插入排序 充分利用插入排序对处理整体有序的数据的优越性
//g 最后一次为1 保底

//间隔为g的数组有序
func insortwithg(data []interface{}, size, g int, cf CF) int {
	var i, j int
	var key interface{}
	for j = g; j < size; j++ {
		key = data[j]

		i = j - g

		//一步一步把key移动到合适的位置，前面的数据是已经排好序的
		for i >= 0 && cf(data[i], key) > 0 {
			data[i+g], data[i] = data[i], data[i+g]
			i -= g
		}
	}
	return 0
}

func ShellSort(data []interface{}, size int, cf CF) {
	var g = make([]int, 0)
	//构造合适的g
	for h := 1; ; {
		if h > size {
			break
		}
		g = append(g, h)
		h = 3*h + 1
	}


	for i := len(g) - 1; i >= 0; i-- {
		insortwithg(data, size, g[i], cf)
		fmt.Println(data)
	}
}
