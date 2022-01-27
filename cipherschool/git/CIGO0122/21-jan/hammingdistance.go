package main

import (
	"errors"
	"fmt"
)

/*Calculate the Hamming Distance between two DNA strands.

Your body is made up of cells that contain DNA. Those cells regularly wear out and need replacing, which they achieve by dividing into daughter cells. In fact, the average human body experiences about 10 quadrillion cell divisions in a lifetime!

When cells divide, their DNA replicates too. Sometimes during this process mistakes happen and single pieces of DNA get encoded with the incorrect information. If we compare two strands of DNA and count the differences between them we can see how many mistakes occurred. This is known as the "Hamming Distance".

We read DNA using the letters C,A,G and T. Two strands might look like this:

    GAGCCTACTAACGGGAT
    CATCGTAATGACGGCCT
    ^ ^ ^  ^ ^    ^^

They have 7 differences, and therefore the Hamming Distance is 7.

The Hamming Distance is useful for lots of things in science, not just biology, so it's a nice phrase to be familiar with :)

## Implementation notes

The Hamming distance is only defined for sequences of equal length, so
an attempt to calculate it between sequences of different lengths should
not work.
*/
func main() {
	a := "GGAAACTAAT"
	b := "CAAAAGTTTT"
	dist, err := distance(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("hamming distance is", dist)
	}

}
func distance(a, b string) (int, error) {
	hammingDist := 0
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				hammingDist++
			}

		}
		return hammingDist, nil
	}
	return hammingDist, errors.New("both strings are not equal")
	panic("implement the distance function")
}
