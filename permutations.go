package projecteuler

// Permutations calculates and returns permutations smaller than limit, or until f returns true
func Permutations(limit byte, f func(...interface{}) bool, args ...interface{}) (permutations [][]byte) {
	starting := make([]byte, 0, limit)
	changing := make([]byte, 0, limit)
	for i := byte(0); i < limit; i++ {
		changing = append(changing, i)
	}

	permutations, _ = doPerms(starting, changing, permutations, f, args...)
	return
}

func doPerms(starting, changing []byte, perms [][]byte, f func(...interface{}) bool, args ...interface{}) (
	[][]byte, bool) {

	var halt bool

	if len(changing) == 1 {
		currPerm := make([]byte, cap(starting))
		copy(currPerm, starting)
		currPerm[cap(starting)-1] = changing[0]
		return processPerm(currPerm, perms, f, args...)
	}

	for i := 0; i < len(changing); i++ {
		starting = append(starting, changing[i])
		newChanging := cutOutIndex(changing, i)
		perms, halt = doPerms(starting, newChanging, perms, f, args...)
		if halt {
			return perms, true
		}
		starting = starting[:len(starting)-1]
	}

	return perms, false
}

func processPerm(currPerm []byte, perms [][]byte, f func(...interface{}) bool, args ...interface{}) (
	[][]byte, bool) {

	perms = append(perms, currPerm)
	args = append(args, currPerm)

	retValue := false
	if f(args...) {
		retValue = true
	}
	args = args[:len(args)-1]

	return perms, retValue
}

func cutOutIndex(changing []byte, index int) (newChanging []byte) {
	start := make([]byte, index)
	copy(start, changing[:index])

	end := make([]byte, len(changing)-index-1)
	copy(end, changing[index+1:])

	newChanging = append(newChanging, start...)
	newChanging = append(newChanging, end...)
	return
}
