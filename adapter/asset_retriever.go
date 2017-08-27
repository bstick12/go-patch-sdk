package adapter

//go:generate counterfeiter ./ AssetRetriever
type AssetRetriever interface {
	Asset(string) ([]byte, error)
}

type assertRetrieverImpl struct{}

func NewAssertRetriever() AssetRetriever {
	return assertRetrieverImpl{}
}

func (assertRetrieverImpl) Asset(assetName string) ([]byte, error) {
	return Asset(assetName)
}
