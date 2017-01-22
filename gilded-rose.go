package main

import (
	"bytes"
	"fmt"
	"io"
)

type Item struct {
	Name    string
	SellIn  int
	Quality int
}

func (i Item) String() string {
	return fmt.Sprintf("%s, %d, %d", i.Name, i.SellIn, i.Quality)
}

type GildedRose struct {
	Items []Item
}

func (g GildedRose) UpdateItems() {
	items := g.Items
	for i := 0; i < len(items); i++ {
		if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].Quality > 0 {
				if items[i].Name != "Sulfuras, Hand of Ragnaros" {
					items[i].Quality = items[i].Quality - 1
				}
			}
		} else {
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
				if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].SellIn < 11 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
					if items[i].SellIn < 6 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
				}
			}
		}

		if items[i].Name != "Sulfuras, Hand of Ragnaros" {
			items[i].SellIn = items[i].SellIn - 1
		}

		if items[i].SellIn < 0 {
			if items[i].Name != "Aged Brie" {
				if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].Quality > 0 {
						if items[i].Name != "Sulfuras, Hand of Ragnaros" {
							items[i].Quality = items[i].Quality - 1
						}
					}
				} else {
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				if items[i].Quality < 50 {
					items[i].Quality = items[i].Quality + 1
				}
			}
		}
	}
}

var items = []Item{
	{"+5 Dexterity Vest", 10, 20},
	{"Aged Brie", 2, 0},
	{"Elixir of the Mongoose", 5, 7},
	{"Sulfuras, Hand of Ragnaros", 0, 80},
	{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
}

func main() {
	buff := new(bytes.Buffer)
	fprintInventoryOverTime(buff, GildedRose{items})
	fmt.Print(buff.String())
}

func fprintInventoryOverTime(w io.Writer, g GildedRose) {
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "--- day %d ---\n", i + 1)
		g.UpdateItems()
		for _, item := range items {
			fmt.Fprintln(w, item)
		}
	}
}

