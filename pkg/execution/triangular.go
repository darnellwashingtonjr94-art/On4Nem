package execution

type OddsNode struct {
	BookName    string
	ImpliedProb float64
}

func DetectTriangularArbitrage(nodeA, nodeB, nodeC OddsNode) (bool, float64) {
	arbMargin := (1.0 / nodeA.ImpliedProb) + (1.0 / nodeB.ImpliedProb) + (1.0 / nodeC.ImpliedProb)
	if arbMargin < 1.0 {
		return true, 1.0 - arbMargin // Guaranteed percentage yield
	}
	return false, 0.0
}
