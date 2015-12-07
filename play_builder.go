// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// Autogenerated by buildergenerator

package goexoml

//SetLoop sets Loop
func (__play__ *Play) SetLoop(loop int) *Play {
	__play__.Loop = loop
	return __play__
}

//SetDigits sets Digits
func (__play__ *Play) SetDigits(digits int) *Play {
	__play__.Digits = digits
	return __play__
}

//SetURL sets URL
func (__play__ *Play) SetURL(url string) *Play {
	__play__.URL = url
	return __play__
}

//NewPlay return a new Play pointer
func NewPlay() *Play {
	return new(Play)
}

//IPlay The interface that satisfies all the methods for this struct
type IPlay interface {
	SetLoop(loop int) *Play
	SetDigits(digits int) *Play
	SetURL(url string) *Play
}

//AddPlay appends the verb to response
func (r *Response) AddPlay(play IPlay) *Response {
	r.Response = append(r.Response, play)
	return r
}