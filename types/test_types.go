package types

type BarInterface interface {
	DeepCopyBarInterface() BarInterface
	String() string
}

type FuncType func()

// +k8s:deepcopy-gen:interfaces=github.com/lt90s/deepcopy-gen-demo/types.BarInterface
type FooType struct {
	a int
	b string
	c []byte
	d []string
	e [][]byte
	f map[string][]byte
	// not supported types
	//g chan struct{}
	//h FuncType
	//i interface{}
	j BarInterface
}

func (foo *FooType) String() string {
	return "A FooType"
}

// +k8s:deepcopy-gen=false
type BazType struct {
	a string
}
