package _interface

type IFile interface {
	Read(buf []byte)(n int, err error)
	Write(buf []byte)(n int, err error)
	Seek(off int64, whence int)(pos int64, err error)
}

type IReader interface {

}

type IWriter interface {

}

type ICloser interface {
	
}