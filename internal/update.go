package internal

import (
	"math"
	"math/rand/v2"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (state State) Update() State {
	if !state.Player.Alive {
		return state
	}

	state = handleIndependent(state)
	handleInteractions(&state)
	return state
}

func handleIndependent(state State) State {
	var (
		newPlayer      Player
		newEnemies     EnemyList
		newProjectiles ProjectileList
	)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		newPlayer = updatePlayer(state.Player)
	}()
	go func() {
		defer wg.Done()
		newEnemies = updateEnemies(state)
	}()
	go func() {
		defer wg.Done()
		newProjectiles = updateProjectiles(state)
	}()

	wg.Wait()
	return State{newPlayer, newEnemies, newProjectiles}
}

func updatePlayer(player Player) Player {
	inputMap := map[int32]rl.Vector2{
		rl.KeyW: {X: 0, Y: -1},
		rl.KeyA: {X: -1, Y: 0},
		rl.KeyS: {X: 0, Y: 1},
		rl.KeyD: {X: 1, Y: 0},
	}
	newVelocity := player.Velocity
	acceleration := float32(8000)
	if player.ZPos != 0 {
		acceleration /= 4
	} else {
		newVelocity = rl.Vector2ClampValue(newVelocity, 0, max(rl.Vector2Length(newVelocity)-4000*rl.GetFrameTime(), 0))
	}
	for key, direction := range inputMap {
		if rl.IsKeyDown(key) {
			newVelocity = rl.Vector2Add(newVelocity, rl.Vector2Scale(direction, acceleration*rl.GetFrameTime()))
			if rl.IsKeyPressed(rl.KeyLeftShift) {
				newVelocity = rl.Vector2Add(newVelocity, rl.Vector2Scale(direction, 3000))
			}
		}
	}
	if rl.Vector2Length(newVelocity) > 800 && player.ZPos == 0 {
		newVelocity = rl.Vector2Scale(newVelocity, max(0.99, 0.992-rl.GetFrameTime()))
	}

	var newZVel float32
	if player.ZPos == 0 {
		if rl.IsKeyPressed(rl.KeySpace) {
			newZVel = 2
		} else {
			newZVel = 0
		}
	} else {
		newZVel = player.ZVel - 10*rl.GetFrameTime()
	}

	newZPos := max(0, player.ZPos+newZVel*rl.GetFrameTime())

	return Player{
		true,
		rl.Vector2Add(player.Position, rl.Vector2Scale(player.Velocity, rl.GetFrameTime())),
		newVelocity,
		player.Radius,
		newZPos,
		newZVel,
	}
}

func canSpawnEnemy(state State) bool {
	chance := rl.GetFrameTime()
	if rl.Vector2Length(state.Player.Position) > 10000 {
		chance *= 10
	}
	return rand.Float32() < chance
}
func updateEnemies(state State) EnemyList {
	newEnemies := EnemyList{}
	canSpawnEnemy := canSpawnEnemy(state)
	for i, e := range state.Enemies {
		if e.Alive {
			target := rl.Vector2Add(state.Player.Position, rl.Vector2Scale(state.Player.Velocity, e.Personality*3))
			newPos := rl.Vector2MoveTowards(e.Position, target, 600*rl.GetFrameTime())
			newEnemies[i] = Enemy{true, newPos, e.Radius, e.Personality}
		} else if canSpawnEnemy {
			canSpawnEnemy = false

			pos := rl.Vector2Add(state.Player.Position, rl.Vector2Rotate(rl.Vector2{X: 1000, Y: 0}, rand.Float32()*math.Pi))
			newEnemies[i] = Enemy{true, pos, 64, rand.Float32()}
		}
	}
	return newEnemies
}

func updateProjectiles(state State) ProjectileList {
	newProjectiles := ProjectileList{}
	playerShooting := rl.IsMouseButtonPressed(rl.MouseButtonLeft)
	for i, p := range state.Projectiles {
		if p.Alive {
			newProjectiles[i] = Projectile{
				Alive:    true,
				Position: rl.Vector2Add(p.Position, rl.Vector2Scale(p.Velocity, rl.GetFrameTime())),
				Velocity: p.Velocity,
				Hostile:  p.Hostile,
				Radius:   p.Radius,
			}
		} else if playerShooting {
			playerShooting = false

			delta := rl.Vector2Subtract(rl.GetMousePosition(), rl.Vector2{X: 800, Y: 450})
			velocity := rl.Vector2Scale(rl.Vector2Normalize(delta), 1200)
			newProjectiles[i] = Projectile{
				Alive:    true,
				Position: state.Player.Position,
				Velocity: velocity,
				Hostile:  false,
				Radius:   16,
			}
		}
	}
	return newProjectiles
}

func handleInteractions(state *State) {
	updateCollisions(state)
}

func updateCollisions(state *State) {
	for i, e := range state.Enemies {
		if !e.Alive {
			continue
		}

		if rl.Vector2Distance(state.Player.Position, e.Position) < e.Radius+state.Player.Radius {
			state.Player.Alive = false
			println("Player died from enemy at", e.Position.X, e.Position.Y)
		}

		for j, p := range state.Projectiles {
			if p.Alive && !p.Hostile && rl.Vector2Distance(e.Position, p.Position) < e.Radius+p.Radius {
				state.Enemies[i].Alive = false
				state.Projectiles[j].Alive = false
			}
		}
	}

	for j, p := range state.Projectiles {
		if p.Alive && p.Hostile && rl.Vector2Distance(state.Player.Position, p.Position) < state.Player.Radius+p.Radius {
			state.Player.Alive = false
			state.Projectiles[j].Alive = false
			println("Player died from projectile at", p.Position.X, p.Position.Y)
		}
	}
}
