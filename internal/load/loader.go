package load

type Loader interface {
	Load([][]any) (int64, error)
}
