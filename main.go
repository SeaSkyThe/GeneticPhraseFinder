package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
    RADIUS = 5.0
    SPEED  = 5.0
)

type Ball struct {
	ballPosition rl.Vector2
	ballSpeed    rl.Vector2
	ballRadius   int
}

func spawnRandomBall(balls []*Ball) []*Ball {
	var ballRadius int = 20
	var ballPosition rl.Vector2 = rl.Vector2{X: float32(rl.GetRandomValue(int32(ballRadius), int32(rl.GetScreenWidth()-ballRadius))), Y: float32(rl.GetRandomValue(int32(ballRadius), int32(rl.GetScreenHeight()-ballRadius)))}
	var ballSpeed rl.Vector2 = rl.Vector2{X: float32(rl.GetRandomValue(3, 8)), Y: float32(rl.GetRandomValue(3, 8))}

	balls = append(balls, &Ball{ballPosition: ballPosition, ballSpeed: ballSpeed, ballRadius: ballRadius})

	return balls
}

func checkBorderCollision(ballPosition rl.Vector2, ballRadius float32, ballSpeed *rl.Vector2) {
	if (ballPosition.X >= (float32(rl.GetScreenWidth()) - float32(ballRadius))) || (ballPosition.X <= float32(ballRadius)) {
		ballSpeed.X *= -1.0
	}

	if (ballPosition.Y >= (float32(rl.GetScreenHeight()) - float32(ballRadius))) || (ballPosition.Y <= float32(ballRadius)) {
		ballSpeed.Y *= -1.0
	}
}

func main() {

	rl.InitWindow(800, 450, "This is my future Genetic Algo Viz")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var pause bool = false
	var framesCounter int = 0

	var balls []*Ball
	balls = spawnRandomBall(balls)

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeySpace) {
			pause = !pause
		}

		if rl.IsKeyPressed(rl.KeyA) {
			balls = spawnRandomBall(balls)
		}

		if !pause {
			for _, ball := range balls {
				ball.ballPosition.X += ball.ballSpeed.X
				ball.ballPosition.Y += ball.ballSpeed.Y
				checkBorderCollision(ball.ballPosition, float32(ball.ballRadius), &ball.ballSpeed)
			}
		} else {
			framesCounter += 1
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		// rl.DrawText("This is my future Genetic Algo Viz!", 190, 200, 20, rl.DarkGray)
		//
		for _, ball := range balls {
			rl.DrawCircleV(ball.ballPosition, float32(ball.ballRadius), rl.Maroon)
		}

		if pause && ((framesCounter/30)%2) == 0 {
			rl.DrawText("PAUSED", 350, 200, 30, rl.Gray)
		}
		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
