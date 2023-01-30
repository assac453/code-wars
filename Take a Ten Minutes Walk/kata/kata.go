package kata

// n   s     e     w
// up down right left
func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}
	var x, y int32
	for i := 0; i < len(walk); i++ {
		switch walk[i] {
		case 'n':
			y += 1
			break
		case 's':
			y -= 1
			break
		case 'e':
			x += 1
			break
		case 'w':
			x -= 1
			break
		}
	}
	if x == 0 && y == 0 {
		return true
	}
	return false
}
