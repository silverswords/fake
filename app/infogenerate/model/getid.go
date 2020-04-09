package model

//GetID is to get ids
func GetID(number int) []int {
	var id []int

	for i := 0; i < number; i++ {
		id = append(id, i)
	}

	return id
}
