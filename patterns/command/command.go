package main

import "fmt"

// Command - интерфейс команды
type Command interface {
	Execute()
}

// Light - класс света
type Light struct {
	isOn bool
}

func (l *Light) On() {
	l.isOn = true
	fmt.Println("Свет включен")
}

func (l *Light) Off() {
	l.isOn = false
	fmt.Println("Свет выключен")
}

// LightOnCommand - команда включения света
type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

// LightOffCommand - команда выключения света
type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

// RemoteControl - пульт управления
type RemoteControl struct {
	command Command
}

func (rc *RemoteControl) SetCommand(command Command) {
	rc.command = command
}

func (rc *RemoteControl) PressButton() {
	rc.command.Execute()
}

func main() {
	light := &Light{}

	lightOnCommand := NewLightOnCommand(light)
	lightOffCommand := NewLightOffCommand(light)

	remoteControl := &RemoteControl{}

	remoteControl.SetCommand(lightOnCommand)
	remoteControl.PressButton() // Включение света

	remoteControl.SetCommand(lightOffCommand)
	remoteControl.PressButton() // Выключение света
}
