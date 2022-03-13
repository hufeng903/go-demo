package main

type Talk interface {
	Hello(userName string) string
	Say(heard string) (start string, end string)
}

type myTalk string

func (myTalk myTalk) Hello(userName string) string {
	return ""
}

func (myTalk myTalk) Say(heard string) (start string, end string) {
	return "", ""
}

func main() {
	var a myTalk

	a.Hello("abc")
	a.Say("ccc")
}
