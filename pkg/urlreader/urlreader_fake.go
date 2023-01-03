package filereader

type FakeUrlReader struct {
	Contents []byte
	Err      error
}

func NewFakeUrlReader(contents []byte, err error) *FakeUrlReader {
	return &FakeUrlReader{
		contents,
		err,
	}
}

func (f *FakeUrlReader) ReadFile(url string) ([]byte, error) {
	return f.Contents, f.Err
}
