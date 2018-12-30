package main

type TypeCommon struct {
	categ int
	qual  int
	align int
	size  int
	bty   Type
}

type Type struct {
	//TypeCommon
}

type symbolCommon struct {
	kind  int
	name  string
	aname string
	//ty Type
	level     int
	sclass    int
	ref       int
	defined   bool
	addressed bool
	needwb    bool
}
type symbol struct {
}
