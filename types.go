package cmplxNN

// Layer is a matrix representation of a layer of a nueral net.
type Layer struct {
	// Weights is a matrix of complex128 weights, w[row][column] indexing.
	// the number of columns is the expected dimension of the input array
	// the number of rows is the dimension of the output array
	Weights [][]complex128
	// Bias is added to the output after adjusting based on the weight and before normalizing the values.
	// the length of the bias should match the number of rows in Weights
	Bias []complex128
	// Normalize is the function used to normalize the output of the Layer
	Normalizer Normalizer
}

type NueralNet struct {
	Layers []Layer
}

type Modal interface {
	Generate([]complex128) []complex128
}
