package common


//CheckInNumberSlice 检查target是否在slice中
func CheckInNumberSlice(target uint64, arr []uint64) bool {
	for _, v := range arr {
		if target == v {
			return true
		}
	}
	
	return false
}


// DelEleInSlice 在slice中删除target, 返回新的slice, 只适用于slice中元素不重复的情况
func DelEleInSlice(target uint64, arr []uint64) []uint64{
	newSlice := make([]uint64, 0, len(arr))

	for i, v := range arr {
		if v == target {
			newSlice = append(newSlice, arr[:i]...)
			newSlice = append(newSlice, arr[i+1:]...)
		}else {
			newSlice = append(newSlice, v)
		}
	}


	return newSlice
}