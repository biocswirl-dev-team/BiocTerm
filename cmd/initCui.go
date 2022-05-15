/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
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
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
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

	if v, err := g.SetView("cmdline", -1, maxY-10, maxX, maxY); err != nil {
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

// Quit
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
