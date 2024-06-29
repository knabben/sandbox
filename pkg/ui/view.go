package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/knabben/showcase/api/v1alpha1"
	"github.com/knabben/showcase/pkg/cmd"
	"github.com/rivo/tview"
	"strings"
)

func Run(demo *v1alpha1.Demo) error {
	var (
		currentStep = 0
		running     bool
		outChannel  = make(chan string)
		errChannel  = make(chan string)
	)
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetChangedFunc(func() { app.Draw() })
	frame := tview.NewFrame(textView)
	textView.SetBorder(true)
	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			if currentStep == len(demo.Spec.Steps) {
				app.Stop()
				return event
			}
			go func() {
				if !running {
					step := demo.Spec.Steps[currentStep]
					textView.SetTitle(step.Description)
					frame.Clear()
					frame.AddText(strings.Join(step.Command, " "), true, tview.AlignCenter, tcell.ColorBlue)
					running = true
					c := cmd.NewCommand(step.Command)
					err := c.Execute(&outChannel, &errChannel)
					if err != nil {
						fmt.Fprintf(textView, "%s\n", err)
					} else {
						var output string
						go EnableOutput(&output, outChannel, textView)
						go EnableOutput(&output, errChannel, textView)
					}
					running = false
					currentStep += 1
				}
			}()
		}
		return event
	})
	app.SetRoot(frame, true).EnableMouse(true)
	return app.Run()
}

func EnableOutput(output *string, std chan string, textView *tview.TextView) {
	var outlist []string
	for {
		select {
		case n, ok := <-std:
			if !ok {
				if output != nil {
					*output = strings.Join(outlist, " ")
					textView.Clear()
				}
				break
			}
			if output != nil {
				outlist = append(outlist, n)
			}
			fmt.Fprintf(textView, n+"\n")
		}
	}
}
