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
}

func (l *Layer) Generate(in []complex128) (out []complex128) {
	out = make([]complex128, len(l.Weights))
	for i := range out {
		for j := range in {
			out[i] += l.Weights[i][j] * in[j]
		}
		out[i] += l.Bias[i]
		out[i] = Normalize(out[i])
	}
	return out
}

type NueralNet struct {
	Layers []Layer
}

func (nn *NueralNet) Generate(in []complex128) (out []complex128) {
	out = in
	for _, l := range nn.Layers {
		out = l.Generate(out)
	}
	return out
}

type Model interface {
	Generate([]complex128) []complex128
	// Mutate(stepsize float64) Model
}
