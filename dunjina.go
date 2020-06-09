package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	// select
	selectblock                = -1
	selectblockh, selectblockv int
	// player
	player                                                 = 4012
	playerh, playerv, playernext, playernexth, playernextv int
	// mouse
	mouseblock = -1
	// level
	levelw   = 1000
	levelh   = 1000
	levela   = levelh * levelw
	levelmap = make([]string, levela)
	// core
	drawblock, drawblocknext, drawa, drawblockv, drawblockh    int
	monh32, monw32                                             int32
	monitorh, monitorw, monitornum, blocksw, blocksh, blocknum int
	grid16on, grid4on, debugon, lrg, sml                       bool
	framecount                                                 int
	mousepos                                                   rl.Vector2
	camera                                                     rl.Camera2D
)

func timers() { // MARK: timers

}
func getpositions() { // MARK:getpositions()
	// horizontal vertical
	drawblockh = drawblocknext / levelw
	drawblockv = drawblocknext - (drawblockh * levelw)
	selectblockh = selectblock / levelw
	selectblockv = selectblock - (selectblockh * levelw)
	playerh = player / levelw
	playerv = player - (playerh * levelw)
	playernexth = playernext / levelw
	playernextv = playernext - (playernexth * levelw)
	// mouse block position
	xchange := float32(0)
	ychange := float32(0)
	ycount := 0
	for b := 0; b < blocksh; b++ {
		if mousepos.Y > 0+ychange && mousepos.Y < 16+ychange {
			for a := 0; a < blocksw; a++ {
				if mousepos.X > 0+xchange && mousepos.X < 16+xchange {
					mouseblock = a + ycount + drawblocknext
				}
				xchange += 16
			}
		}
		ychange += 16
		ycount += levelw
	}
}
func screenposition() { // MARK: screenposition()

	if playerh-drawblockh < 33 {
		if drawblockh > 0 {
			drawblocknext -= levelw
		}
	} else if playerh-drawblockh > 33 {
		if drawblockh < levelh-(blocksh+1) {
			drawblocknext += levelw
		}
	}

	if playerv-drawblockv < 60 {
		if drawblockv > 0 {
			drawblocknext--
		}
	} else if playerv-drawblockv > 60 {
		if drawblockv < levelw-(blocksw+1) {
			drawblocknext++
		}
	}

}
func updateall() { // MARK: updateall()

	getpositions()
	screenposition()
	moveplayer()

	if grid16on {
		grid16()
	}
	if grid4on {
		grid4()
	}
	timers()
}
func moveplayer() { // MARK: moveplayer()
	if playernext != player {

		if playernexth > playerh {
			player += levelw
		} else if playernexth < playerh {
			player -= levelw
		}

		if playernextv > playerv {
			player++
		} else if playernextv < playerv {
			player--
		}

	}
}
func createlevel() { // MARK: createlevel()

	for a := 0; a < levela; a++ {
		levelmap[a] = "."
	}

	levelmap[1] = "#"
	levelmap[3] = "#"
	levelmap[7] = "#"

	roomblock := 2010

	rooml := rInt(15, 25)
	roomw := rInt(15, 25)
	rooma := rooml * roomw
	count := 0

	for b := 0; b < 5; b++ {

		for a := 0; a < rooma; a++ {
			levelmap[roomblock] = "^"
			roomblock++
			count++
			if count == rooml {
				count = 0
				roomblock += levelw - rooml
			}
		}
		roomblock += rooml
		roomblock -= rInt(4, 7) * levelw

		count = 0
		passagel := rInt(5, 10)
		passagew := rInt(2, 5)
		passagea := passagel * passagew

		for a := 0; a < passagea; a++ {
			levelmap[roomblock] = "^"
			roomblock++
			count++
			if count == passagel {
				count = 0
				roomblock += levelw - passagel
			}
		}
		roomblock -= rInt(4, 7) * levelw
		rooml = rInt(15, 25)
		roomw = rInt(15, 25)
		rooma = rooml * roomw
		count = 0
	}

}
func startgame() { // MARK: startgame()
	createlevel()
}
func main() { // MARK: main()
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLog(rl.LogError)      // hides INFO window
	startsettings()
	raylib()
}
func raylib() { // MARK: raylib()
	rl.InitWindow(monw32, monh32, "dunjina")

	setscreen()
	startgame()
	rl.CloseWindow()
	rl.InitWindow(monw32, monh32, "dunjina")

	rl.SetExitKey(rl.KeyEnd) // key to end the game and close window
	//	imgs = rl.LoadTexture("imgs.png") // load images
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() { // MARK: WindowShouldClose

		mousepos = rl.GetMousePosition()
		framecount++
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		// rl.DrawTexture(backimg, 0, 0, rl.Red) // MARK: draw backimg
		rl.BeginMode2D(camera)

		// MARK: draw map layer 1

		drawblock = drawblocknext
		linecount := 0
		drawx := int32(0)
		drawy := int32(0)

		for a := 0; a < blocknum; a++ {

			checklevel := levelmap[drawblock]

			switch checklevel {

			case ".":
				rl.DrawRectangleLines(drawx, drawy, 15, 15, rl.Fade(rl.Brown, 0.1))
			case "#":
				rl.DrawRectangleLines(drawx, drawy, 15, 15, rl.Magenta)
			case "^":
				rl.DrawRectangleLines(drawx, drawy, 15, 15, rl.Fade(rl.Orange, 0.2))
			}

			if mouseblock == drawblock {
				rl.DrawRectangle(drawx, drawy, 15, 15, rl.Orange)
			}
			if player == drawblock {
				rl.DrawRectangle(drawx, drawy, 15, 15, rl.Red)
			}
			if playernext == drawblock {
				rl.DrawRectangle(drawx, drawy, 15, 15, rl.Blue)
			}

			linecount++
			drawblock++
			drawx += 16

			if linecount == blocksw {
				linecount = 0
				drawx = 0
				drawy += 16
				drawblock += levelw - blocksw
			}

		}

		// MARK: draw map layer 2
		rl.EndMode2D() // MARK: draw no camera

		if debugon {
			debug()
		}

		rl.EndDrawing()
		input()
		updateall()
	}
	rl.CloseWindow()
}
func setscreen() { // MARK: setscreen()
	monitornum = rl.GetMonitorCount()
	monitorh = rl.GetScreenHeight()
	monitorw = rl.GetScreenWidth()
	monh32 = int32(monitorh)
	monw32 = int32(monitorw)
	rl.SetWindowSize(monitorw, monitorh)
	setsizes()
}
func setsizes() { // MARK: setsizes()
	if monitorw >= 1600 {
		lrg = true
		sml = false
	} else if monitorw < 1600 && monitorw >= 1280 {
		lrg = false
		sml = true
	}
	blocksw = (monitorw / 16) + 1
	blocksh = (monitorh / 16) + 1
	blocknum = blocksh * blocksw
}
func startsettings() { // MARK: start
	camera.Zoom = 1.0
	camera.Target.X = 0.0
	camera.Target.Y = 0.0
	debugon = true
	//grid16on = true
	//selectedmenuon = true
}
func debug() { // MARK: debug
	rl.DrawRectangle(monw32-300, 0, 500, monw32, rl.Fade(rl.Black, 0.9))
	rl.DrawFPS(monw32-290, monh32-100)

	monitorwTEXT := strconv.Itoa(monitorw)
	monitorhTEXT := strconv.Itoa(monitorh)
	blockswTEXT := strconv.Itoa(blocksw)
	blockshTEXT := strconv.Itoa(blocksh)
	mouseposXTEXT := fmt.Sprintf("%.0f", mousepos.X)
	mouseposYTEXT := fmt.Sprintf("%.0f", mousepos.Y)
	drawblockvTEXT := strconv.Itoa(drawblockv)
	drawblockhTEXT := strconv.Itoa(drawblockh)
	blocknumTEXT := strconv.Itoa(blocknum)
	mouseblockTEXT := strconv.Itoa(mouseblock)
	selectblockhTEXT := strconv.Itoa(selectblockh)
	selectblockvTEXT := strconv.Itoa(selectblockv)

	rl.DrawText(monitorwTEXT, monw32-290, 10, 10, rl.White)
	rl.DrawText("monitorw", monw32-200, 10, 10, rl.White)
	rl.DrawText(monitorhTEXT, monw32-290, 20, 10, rl.White)
	rl.DrawText("monitorh", monw32-200, 20, 10, rl.White)
	rl.DrawText(blockswTEXT, monw32-290, 30, 10, rl.White)
	rl.DrawText("blocksw", monw32-200, 30, 10, rl.White)
	rl.DrawText(blockshTEXT, monw32-290, 40, 10, rl.White)
	rl.DrawText("blocksh", monw32-200, 40, 10, rl.White)
	rl.DrawText(mouseposXTEXT, monw32-290, 50, 10, rl.White)
	rl.DrawText("mouseposX", monw32-200, 50, 10, rl.White)
	rl.DrawText(mouseposYTEXT, monw32-290, 60, 10, rl.White)
	rl.DrawText("mouseposY", monw32-200, 60, 10, rl.White)
	rl.DrawText(drawblockvTEXT, monw32-290, 70, 10, rl.White)
	rl.DrawText("drawblockv", monw32-200, 70, 10, rl.White)
	rl.DrawText(drawblockhTEXT, monw32-290, 80, 10, rl.White)
	rl.DrawText("drawblockh", monw32-200, 80, 10, rl.White)
	rl.DrawText(blocknumTEXT, monw32-290, 90, 10, rl.White)
	rl.DrawText("blocknum", monw32-200, 90, 10, rl.White)
	rl.DrawText(mouseblockTEXT, monw32-290, 100, 10, rl.White)
	rl.DrawText("mouseblock", monw32-200, 100, 10, rl.White)
	rl.DrawText(selectblockhTEXT, monw32-290, 110, 10, rl.White)
	rl.DrawText("selectblockh", monw32-200, 110, 10, rl.White)
	rl.DrawText(selectblockvTEXT, monw32-290, 120, 10, rl.White)
	rl.DrawText("selectblockv", monw32-200, 120, 10, rl.White)

}
func input() { // MARK: keys input
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		playernext = mouseblock
	}
	if rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyDown(rl.KeyRight) {
		if drawblockv < (levelw - (blocksw + 1)) {
			drawblocknext++
		}
	}
	if rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyDown(rl.KeyLeft) {
		if drawblockv > 0 {
			drawblocknext--
		}
	}
	if rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyDown(rl.KeyUp) {
		if drawblockh > 0 {
			drawblocknext -= levelw
		}
	}
	if rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyDown(rl.KeyDown) {
		if drawblockh < levelh-(blocksh+1) {
			drawblocknext += levelw
		}
	}
	if rl.IsKeyPressed(rl.KeyKpAdd) {
		if camera.Zoom == 1.0 {
			camera.Zoom = 2.0
		} else if camera.Zoom == 2.0 {
			camera.Zoom = 4.0
		}
	}
	if rl.IsKeyPressed(rl.KeyKpSubtract) {
		if camera.Zoom == 4.0 {
			camera.Zoom = 2.0
		} else if camera.Zoom == 2.0 {
			camera.Zoom = 1.0
		}
	}
	if rl.IsKeyPressed(rl.KeyF1) {
		if grid16on {
			grid16on = false
		} else {
			grid16on = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF2) {
		if grid4on {
			grid4on = false
		} else {
			grid4on = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debugon {
			debugon = false
		} else {
			debugon = true
		}
	}

}
func grid16() { // MARK: grid16()
	for a := 0; a < monitorw; a += 16 {
		a32 := int32(a)
		rl.DrawLine(a32, 0, a32, monh32, rl.Fade(rl.Green, 0.1))
	}
	for a := 0; a < monitorh; a += 16 {
		a32 := int32(a)
		rl.DrawLine(0, a32, monw32, a32, rl.Fade(rl.Green, 0.1))
	}
}
func grid4() { // MARK: grid4()
	for a := 0; a < monitorw; a += 4 {
		a32 := int32(a)
		rl.DrawLine(a32, 0, a32, monh32, rl.Fade(rl.DarkGreen, 0.1))
	}
	for a := 0; a < monitorh; a += 4 {
		a32 := int32(a)
		rl.DrawLine(0, a32, monw32, a32, rl.Fade(rl.DarkGreen, 0.1))
	}
}

// random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
