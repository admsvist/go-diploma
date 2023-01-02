package filereader

type FakeFileReader struct {
	Contents []byte
	Err      error
}

func NewFakeFileReader(contents []byte, err error) *FakeFileReader {
	return &FakeFileReader{
		contents,
		err,
	}
}

func (f *FakeFileReader) ReadFile(path string) ([]byte, error) {
	return f.Contents, f.Err
}
