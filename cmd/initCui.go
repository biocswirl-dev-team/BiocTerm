/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jroimartin/gocui"
)

const delta = 1

var (
	views   = []string{}
	curView = -1
	idxView = 0
)

func terminal() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	//Set GUI Managers
	g.SetManagerFunc(layout)

	//Set Keybindings
	if err := initKeybindings(g); err != nil {
		log.Panicln(err)
	}

	////Mouse Bindings

	g.Mouse = true
	//if err := g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, fcn); err != nil {
	//	// handle error
	//}

	//Selection settings
	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorBlack
	g.SelBgColor = gocui.ColorGreen

	//Crash
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

//GUI Manager
func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("help", maxX-25, 0, maxX-1, 9)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "KEYBINDINGS")
		fmt.Fprintln(v, "Space: New View")
		fmt.Fprintln(v, "Tab: Next View")
		fmt.Fprintln(v, "← ↑ → ↓: Move View")
		fmt.Fprintln(v, "Backspace: Delete View")
		fmt.Fprintln(v, "t: Set view on top")
		fmt.Fprintln(v, "b: Set view on bottom")
		fmt.Fprintln(v, "^C: Exit")
	}

	//Directory Panel
	if v, err := g.SetView("Directory", -1, 0, int(0.2*float32(maxX)), maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.Editable = true
		v.Wrap = true
		v.Autoscroll = true

		// Placeholder text for demo
		v.Title = "Table of Contents"

		fmt.Fprintln(v, "\rGetting Started")
		fmt.Fprintln(v, "\rCourse Directory")
		fmt.Fprint(v, "\rSettings")
		fmt.Fprint(v, "\rR Configuration")
	}

	//BiocTerm Panel
	if v, err := g.SetView("biocTerm", int(0.2*float32(maxX)), 0, maxX, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "biocTerm"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.Editable = true
		v.Wrap = true
		v.Autoscroll = true

		// Placeholder for course directory and contents options
		fmt.Fprintln(v, "Example Text 1")
		fmt.Fprintln(v, "Example Question?")
		fmt.Fprintln(v, "Example Answer")

	}

	//Command Line Interface
	if v, err := g.SetView("cmdLine", -1, maxY-10, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Input Panel"
		v.Editable = true
		v.Wrap = true

		fmt.Fprint(v, "Welcome to BiocTerm! Click on the Table of Contents to get started.")

	}
	return nil
}

//Custom Editor

//Set Keybindings

func initKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return newView(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyBackspace2, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return delView(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return nextView(g, true)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, -delta, 0)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, delta, 0)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, 0, delta)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, 0, -delta)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 't', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			_, err := g.SetViewOnTop(views[curView])
			return err
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'b', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			_, err := g.SetViewOnBottom(views[curView])
			return err
		}); err != nil {
		return err
	}
	return nil
}

// Supplementary View Functions
func newView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	name := fmt.Sprintf("v%v", idxView)
	v, err := g.SetView(name, maxX/2-5, maxY/2-5, maxX/2+5, maxY/2+5)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		fmt.Fprintln(v, strings.Repeat(name+" ", 30))
	}
	if _, err := g.SetCurrentView(name); err != nil {
		return err
	}

	views = append(views, name)
	curView = len(views) - 1
	idxView += 1
	return nil
}

func delView(g *gocui.Gui) error {
	if len(views) <= 1 {
		return nil
	}

	if err := g.DeleteView(views[curView]); err != nil {
		return err
	}
	views = append(views[:curView], views[curView+1:]...)

	return nextView(g, false)
}

func nextView(g *gocui.Gui, disableCurrent bool) error {
	next := curView + 1
	if next > len(views)-1 {
		next = 0
	}

	if _, err := g.SetCurrentView(views[next]); err != nil {
		return err
	}

	curView = next
	return nil
}

func moveView(g *gocui.Gui, v *gocui.View, dx, dy int) error {
	name := v.Name()
	x0, y0, x1, y1, err := g.ViewPosition(name)
	if err != nil {
		return err
	}
	if _, err := g.SetView(name, x0+dx, y0+dy, x1+dx, y1+dy); err != nil {
		return err
	}
	return nil
}

// Quit
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
