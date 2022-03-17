package a

func Case1(a int) {}

type MyInt int

func Case2(b MyInt, c string) {}

func Case3(b []byte, ss string) {}

func Case4(e error) {}

func Case5(b []byte) {}

//not generated because this function cannot be exported by other package
func case6(b []byte) {}

func Hoge() {}
