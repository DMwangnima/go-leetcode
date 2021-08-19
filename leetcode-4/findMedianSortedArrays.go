package leetcode_4

import "sort"

// 暴力法，将nums2加入nums1中，随后将num1排序，得到中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for _, n := range nums2 {
		nums1 = append(nums1, n)
	}
	sort.Sort(arr(nums1))
	if len(nums1) % 2 == 0 {
		return float64(nums1[len(nums1)/2-1]+nums1[len(nums1)/2]) / 2.0
	}
	return float64(nums1[len(nums1)/2])
}

type arr []int

func (a arr) Len() int {
	return len(a)
}

func (a arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a arr) Less(i, j int) bool {
	return a[i] < a[j]
}

// 使用类似归并排序的做法，ind1和ind2分别遍历nums1和nums2，ind记录遍历过的长度
// 由于总长度分奇偶，所以需要遍历到(len(nums1)+len(nums2))/2+1
// 并且偶数长度的中位数需要前一位，所以维护prev与cur两个值
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	sum := len(nums1) + len(nums2)
	length := sum / 2 + 1
    var ind, ind1, ind2, prev, cur int
    for ind < length {
    	prev = cur
    	if ind1 == len(nums1) {
    		cur = nums2[ind2]
    		ind2 += 1
    		ind += 1
    		continue
		}
		if ind2 == len(nums2) {
			cur = nums1[ind1]
			ind1 += 1
			ind += 1
			continue
		}
		if nums1[ind1] < nums2[ind2] {
			cur = nums1[ind1]
			ind1 += 1
		} else {
			cur = nums2[ind2]
			ind2 += 1
		}
		ind += 1
	}
	if sum % 2 != 0 {
		return float64(cur)
	}
	return float64(cur+prev) / 2
}