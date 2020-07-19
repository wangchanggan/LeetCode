package Easy

func Sqrt(x int) int {
	//return int(math.Sqrt(float64(x)))
	for i:=1;i<=x/2;i++{
		if i*i==x{
			return i
		}else if i*i>x{
			return i-1
		}
	}
	return 0
}