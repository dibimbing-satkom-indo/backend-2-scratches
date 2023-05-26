package main

import "fmt"

type Application struct {
	notifier Notifier
}

type Notifier interface {
	Send(msg string)
}

func (a *Application) SetNotifier(notifier Notifier) {
	a.notifier = notifier
}

func (a Application) DoSomething() {
	a.notifier.Send("I'm doing something")
}

type EmailNotifier struct {}

func (n EmailNotifier) Send(msg string) {
	fmt.Println("Send by email:", msg)
}

type FacebookDecorator struct {
	wrappee Notifier
}

func (fb FacebookDecorator) Send(msg string) {
	fb.wrappee.Send(msg)
	fmt.Println("Send by Facebook:", msg)
}

type SlackDecorator struct {
	wrappee Notifier
}

func (sl SlackDecorator) Send(msg string) {
	sl.wrappee.Send(msg)
	fmt.Println("Send by Slack:", msg)
}

func main() {
	var stack Notifier
	stack = EmailNotifier{}
	facebookEnabled := true
	slackEnabled := true
	if facebookEnabled {
		stack = FacebookDecorator{wrappee: stack}
	}
	if slackEnabled {
		stack = SlackDecorator{wrappee: stack}
	}
	app := Application{}
	app.SetNotifier(stack)
	app.DoSomething()
}
