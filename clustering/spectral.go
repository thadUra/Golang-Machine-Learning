package cluster

/**
 * Spectral()
 * Performs spectral clustering on data and returns the label for each data point
 * Clusters data into 2 partitions
 * Radius value determines connectivity
 */
func Spectral(data [][]float64, radius float64) ([]string, error) {
	var label []string

	// Generate Laplacian matrix from a Radius Neighbors Graph
	laplacian := make([][]float64, len(data))
	for i := range data {
		laplacian[i] = Euclidean(data[i], data)
		sum := 0.0
		for j := range laplacian[i] {
			if laplacian[i][j] > radius {
				laplacian[i][j] = 0.0
			}
			sum += laplacian[i][j]
			laplacian[i][j] *= -1
		}
		laplacian[i][i] = sum
	}

	// Generate eigenvectors from laplacian matrix
	_, eigen_vectors, err := Eigen(laplacian)
	if err != nil {
		return label, err
	}

	// Determine the cluster for each data point
	// NOTE: No need to sort via eigenvalues since Eigen() returns them sorted in ascending order
	for i := range data {
		if eigen_vectors[i][1] >= 0 {
			label = append(label, "Cluster One")
		} else {
			label = append(label, "Cluster Two")
		}
	}
	return label, nil
}
