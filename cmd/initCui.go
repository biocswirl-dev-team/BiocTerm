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

	//Selection colours
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
	if v, err := g.SetView("seqView", -1, -1, int(0.2*float32(maxX)), maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "bioSyntax"
	}

	if v, err := g.SetView("biocTerm", int(0.2*float32(maxX)), -1, maxX, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "biocTerm"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Item 1")
		fmt.Fprintln(v, "Item 2")
		fmt.Fprintln(v, "Item 3")
		fmt.Fprint(v, "\rWill be")
		fmt.Fprint(v, "deleted\rItem 4\nItem 5")
	}

	if v, err := g.SetView("cmdline", -1, maxY-5, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "BiocTerm Input"
	}
	return nil
}

//Custom Editor

// Quit
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
