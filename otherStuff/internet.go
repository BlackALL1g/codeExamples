package otherStuff

const N = 7

func Internet1() {
	switch {
	case N > 0:
		{
			if N > 5 {
				goto A
			}
		A:
		}
	case true:
		{
			if N > 5 {
				goto B
			}
		B:
		}
	case false:
		{
			if N > 5 {
				goto C
			}
		C:
		}
	}
}
