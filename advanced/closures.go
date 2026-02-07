package advanced

func IntSequence() func() int{
	i := 0
	return func() int{
		i++
		return i
	}
}