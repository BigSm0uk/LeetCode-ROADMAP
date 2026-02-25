package task88

func merge(nums1 []int, m int, nums2 []int, n int) {
	out := make([]int, 0, m+n)
	l, r := 0, 0

    for l < m && r < n {
        if nums1[l] < nums2[r] {
            out = append(out, nums1[l])
            l++
        } else {
            out = append(out, nums2[r])
            r++
        }
    }
    for l < m {
        out = append(out, nums1[l])
        l++
    }
    for r < n {
        out = append(out, nums2[r])
        r++
    }
    copy(nums1, out)
    
}

