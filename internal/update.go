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
		newProjectiles = updateProjectiles(state.Projectiles)
	}()

	return State{newPlayer, newEnemies, newProjectiles}
}

func updatePlayer(player Player) Player {
	inputMap := map[int32]rl.Vector2{rl.KeyW: {X: 1, Y: 0}}
	newVelocity := player.Velocity
	for key, direction := range inputMap {
		if rl.IsKeyDown(key) {
			newVelocity = rl.Vector2Add(newVelocity, direction)
		}
	}
	return Player{
		true,
		rl.Vector2Add(player.Position, player.Velocity),
		newVelocity,
		player.Size,
	}
}

func updateEnemies(state State) EnemyList {
	newEnemies := EnemyList{}
	canSpawnEnemy := rand.Float32() < float32(1)/1000
	for i, e := range state.Enemies {
		if e.Alive {
			newPos := rl.Vector2MoveTowards(e.Position, state.Player.Position, 20)
			newEnemies[i] = Enemy{true, newPos, e.Size}
		} else if canSpawnEnemy {
			canSpawnEnemy = false

			pos := rl.Vector2Add(state.Player.Position, rl.Vector2Rotate(rl.Vector2{X: 1000, Y: 0}, rand.Float32()*math.Pi))
			newEnemies[i] = Enemy{true, pos, 64}
		}
	}
	return newEnemies
}

func updateProjectiles(projectiles ProjectileList) ProjectileList {
	newProjectiles := ProjectileList{}
	for i, p := range projectiles {
		if p.Alive {
			newProjectiles[i] = Projectile{
				Alive:    true,
				Position: rl.Vector2Add(p.Position, p.Velocity),
				Velocity: p.Velocity,
				Hostile:  p.Hostile,
				Size:     16,
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
		if rl.Vector2Distance(state.Player.Position, e.Position) < e.Size+state.Player.Size {
			state.Player.Alive = false
		}

		for j, p := range state.Projectiles {
			if !p.Hostile && rl.Vector2Distance(e.Position, p.Position) < e.Size+p.Size {
				state.Enemies[i].Alive = false
				state.Projectiles[j].Alive = false
			}
		}
	}

	for j, p := range state.Projectiles {
		if p.Hostile && rl.Vector2Distance(state.Player.Position, p.Position) < state.Player.Size+p.Size {
			state.Player.Alive = false
			state.Projectiles[j].Alive = false
		}
	}
}
