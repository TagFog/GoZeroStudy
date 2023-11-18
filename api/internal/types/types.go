// Code generated by goctl. DO NOT EDIT.
package types

type Register struct {
	Name     string `json:"name"`
	Password int32  `json:"password"`
}

type Token struct {
	Atoken  string `json:"atoken"`
	Rtoken  string `json:"rtoken"`
	Version int32  `json:"version"`
}

type Update struct {
	Name     string `json:"name"`
	Password int32  `json:"password"`
	Atoken   string `json:"atoken"`
}

type State struct {
	Onestring string `json:"onestring"`
}
