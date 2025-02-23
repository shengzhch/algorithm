package basic

/*
哈希 哈希表包含一个数组，支持通过key通过哈希运算获得下标来访问数组中的元素。
与各类的key相比，哈希表的条目数相应较少，因此，绝大多数哈希函数会将不同的key映射到表中相同的槽位上，称为哈希冲突
*/

/*
链式哈希表
根本是一组链表，每组链表称为一个桶，通过key的哈希运算确定属于哪个桶。目标是尽可能的是元素均匀的分布在每个桶上。
哈希表的负载因子 a = n/m 。n为表中元素的个数，m为桶的个数。一般情况下 我们往往会检索元素个数大于负载因子的桶。
*/

/*
选择哈希函数
h(k) = x . x为整型，表示地址

取余法
h(k) = k mod m ; k为键值，m为槽位(桶)的个数

乘法
k乘以常数A取结果的小数部分，然后乘以m取结果的整数部分。A常取 0.618。
h(k) = FLOOR(（k*A mod 1）* m )
*/

//链式哈希表 维护一个链表数组 取余法
type Chtable struct {
	buckets int
	size    int
<<<<<<< HEAD
	table   []List
=======
	table   []List //槽位
>>>>>>> origin/master

	//f(key)%buckets 确定槽位 辅助哈希函数
	f Hash

	//
	match Match
}

//初始化
func (l *Chtable) Init(b int, f Hash, match Match) {
	l.table = make([]List, b, b)
	l.buckets = b
	l.size = 0
	l.f = f
	l.match = match
}

//是否存在
func (l *Chtable) LookUp(data interface{}) bool {
	//确定所属桶的位置
	buc := l.f(data) % l.buckets

	//确定查找函数
	f := func(e *Node, args ...interface{}) bool {
		// 第一个参数要匹配的元素，第二个参数存储结果的地址
		if len(args) != 2 {
			return false
		}

		//匹配到，退出返回true用于结束遍历
		if l.match(e.data, args[0]) {
			*args[1].(*bool) = true
			return true
		}
		return false
	}

	if l.table[buc].len == 0 {
		return false
	}

	var rel = new(bool)
	//遍历查找
	l.table[buc].Traverse(f, data, rel)

	return *rel
}

//插入
func (l *Chtable) Insert(data interface{}) {
	if l.LookUp(data) {
		return
	}

	buc := l.f(data) % l.buckets

	l.table[buc].InsertAsTail(&Node{data: data})
	l.size++
	return
}

//删除
func (l *Chtable) Remove(data interface{}) {
	buc := l.f(data) % l.buckets
	for m := l.table[buc].head; m != nil; m = m.next {
		if l.match(m.data, data) {
			l.table[buc].DelNode(m)
			l.size--
			return
		}
	}
	return
}
