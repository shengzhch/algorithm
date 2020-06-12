package basic

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
func RadixSort(data []interface{}, cf CF) {

}

//归并排序 稳定
func MergeSort(data []interface{}, cf CF) {

}
