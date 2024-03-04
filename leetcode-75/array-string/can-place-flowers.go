package array_string

func canPlaceFlowers(flowerbed []int, n int) bool {
	var count int
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			if i == 0 {
				if i+1 < len(flowerbed) && flowerbed[i+1] == 0 || i+1 == len(flowerbed) {
					flowerbed[0] = 1
					count++
				}
			} else if i == len(flowerbed)-1 {
				if flowerbed[len(flowerbed)-2] == 0 {
					flowerbed[len(flowerbed)-1] = 1
					count++
				}
			} else {
				if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					count++
				}
			}
		}
	}
	return count >= n
}
