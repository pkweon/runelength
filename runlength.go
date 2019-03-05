package main

import (
    "github.com/mattn/go-runewidth"
)

func main() {
    println(runewidth.StringWidth("あ"))
}
