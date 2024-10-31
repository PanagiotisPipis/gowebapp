package watcher

type Counter struct {
	Iteration int `json:"iteration"`
	HexString string `json:"value"`
}

type CounterReset struct{}
