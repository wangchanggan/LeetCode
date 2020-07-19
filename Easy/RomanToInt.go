package Easy

func RomanToInt(s string) int{
	var sum,i int
	for i=0;i<len(s)-1;i++{
		switch s[i] {
		case 'I':
			if s[i+1]=='V'{
				sum += 4
				i++
			}else if s[i+1]=='X'{
				sum += 9
				i++
			}else{
				sum += 1
			}
		case 'V':
			sum += 5
		case 'X':
			if s[i+1]=='L'{
				sum += 40
				i++
			}else if s[i+1]=='C'{
				sum += 90
				i++
			}else{
				sum += 10
			}
		case 'L':
			sum += 50
		case 'C':
			if s[i+1]=='D'{
				sum += 400
				i++
			}else if s[i+1]=='M'{
				sum += 900
				i++
			}else {
				sum += 100
			}
		case 'D':
			sum += 500
		case 'M':
			sum += 1000
		}
	}
    if i+1==len(s){
		switch s[i] {
		case 'I':
			sum += 1
		case 'V':
			sum += 5
		case 'X':
			sum += 10
		case 'L':
			sum += 50
		case 'C':
			sum += 100
		case 'D':
			sum += 500
		case 'M':
			sum += 1000
		}
	}
	return sum
}