module github.com/jajir/hellogo/example/call

go 1.18

require (
	github.com/ev3go/ev3dev v0.0.0-20210313113244-a5fda5c6a492
	github.com/jajir/hellogo/example/ev3control v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.0
)

require (
	github.com/ev3go/ev3 v0.0.0-20230218221813-265c69c34aaa // indirect
	golang.org/x/image v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
)

replace github.com/jajir/hellogo/example/ev3control => ../ev3control
