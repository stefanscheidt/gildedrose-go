package main

import (
	"bytes"
	"testing"
)

func Test_GildedRose(t *testing.T) {
	dest := new(bytes.Buffer)
	fprintInventoryOverTime(dest)
	if !(inventory_over_time == dest.String()) {
		t.Fatalf("inventory does not evolve as expected:\n" +
			"expected: %q\nactual:   %q", inventory_over_time, dest.String())
	}
}

const inventory_over_time = `--- day 1 ---
+5 Dexterity Vest, 10, 20
Aged Brie, 2, 0
Elixir of the Mongoose, 5, 7
Sulfuras, Hand of Ragnaros, 0, 80
Backstage passes to a TAFKAL80ETC concert, 15, 20
Conjured Mana Cake, 3, 6
--- day 2 ---
+5 Dexterity Vest, 9, 19
Aged Brie, 1, 1
Elixir of the Mongoose, 4, 6
Sulfuras, Hand of Ragnaros, 0, 80
Backstage passes to a TAFKAL80ETC concert, 14, 21
Conjured Mana Cake, 2, 5
--- day 3 ---
+5 Dexterity Vest, 8, 18
Aged Brie, 0, 2
Elixir of the Mongoose, 3, 5
Sulfuras, Hand of Ragnaros, 0, 80
Backstage passes to a TAFKAL80ETC concert, 13, 22
Conjured Mana Cake, 1, 4
--- day 4 ---
+5 Dexterity Vest, 7, 17
Aged Brie, -1, 4
Elixir of the Mongoose, 2, 4
Sulfuras, Hand of Ragnaros, 0, 80
Backstage passes to a TAFKAL80ETC concert, 12, 23
Conjured Mana Cake, 0, 3
--- day 5 ---
+5 Dexterity Vest, 6, 16
Aged Brie, -2, 6
Elixir of the Mongoose, 1, 3
Sulfuras, Hand of Ragnaros, 0, 80
Backstage passes to a TAFKAL80ETC concert, 11, 24
Conjured Mana Cake, -1, 1
--- day 6 ---
+5 Dexterity Vest, 5, 15
Aged Brie, -3, 8
Elixir of the Mongoose, 0, 2
Sulfuras, Hand of Ragnaros, 0, 80
Backstage passes to a TAFKAL80ETC concert, 10, 25
Conjured Mana Cake, -2, 0
--- day 7 ---
+5 Dexterity Vest, 4, 14
Aged Brie, -4, 10
Elixir of the Mongoose, -1, 0
Sulfuras, Hand of Ragnaros, 0, 80
Backstage passes to a TAFKAL80ETC concert, 9, 27
Conjured Mana Cake, -3, 0
--- day 8 ---
+5 Dexterity Vest, 3, 13
Aged Brie, -5, 12
Elixir of the Mongoose, -2, 0
Sulfuras, Hand of Ragnaros, 0, 80
Backstage passes to a TAFKAL80ETC concert, 8, 29
Conjured Mana Cake, -4, 0
--- day 9 ---
+5 Dexterity Vest, 2, 12
Aged Brie, -6, 14
Elixir of the Mongoose, -3, 0
Sulfuras, Hand of Ragnaros, 0, 80
Backstage passes to a TAFKAL80ETC concert, 7, 31
Conjured Mana Cake, -5, 0
--- day 10 ---
+5 Dexterity Vest, 1, 11
Aged Brie, -7, 16
Elixir of the Mongoose, -4, 0
Sulfuras, Hand of Ragnaros, 0, 80
Backstage passes to a TAFKAL80ETC concert, 6, 33
Conjured Mana Cake, -6, 0
`
