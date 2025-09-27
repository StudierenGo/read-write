package cloud

type Cloud struct {
	url string
}

func NewCloud(url string) *Cloud {
	return &Cloud{
		url: url,
	}
}

func (cloud *Cloud) Read() ([]byte, error) {
	return []byte{}, nil
}

func (cloud *Cloud) Write(content []byte) {
}
