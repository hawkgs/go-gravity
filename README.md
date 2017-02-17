# go-gravity
Small Golang library for simulating 2D particles with the help of SDL2.

![GoGravity](https://raw.githubusercontent.com/hAWKdv/go-gravity/master/logo.png)

## Setup
In order to run/try the code in `rendering` package, you must install SDL2 along with the [Golang bindings](https://github.com/veandco/go-sdl2). Of course, you are free to choose whatever rendering method you want since the core of the library is decoupled from SDL2.

## Vector
The `Vector` struct is one of the cores of the library. You can perform basic operations over it like addition, multiplication, normalization and many more.

## Force
The library has a set of forces (vectors) like wind, gravity, kinetic friction, etc.

## Mover
The movers represent a moving particle which is affected by the forces you apply on it. It has its own location, velocity and mass.


```golang
var location *Vector // Initial location
var container *Vector // The boundaries where the mover can move
var obj interface{} // Represents the object of the rendering engine you are using (eg. SDL2 Rect)

// ...

// Our mover
mover := vectors.NewMover(&obj, location, container)

// Some forces
gravity := forces.CreateGravity()
push := forces.CreatePush(315, 0.2, 30)
friction := forces.CreateKineticFriction(mover)

// Rendering loop - each iteration represents a new frame.
// It is all rendering engine-specific stuff...
running := true
for running {
    // ...

    mover.ApplyForce(gravity.GetForce())
    mover.ApplyForce(push.GetForce())
    mover.ApplyForce(friction.GetForce())

    mover.Update()
    // Update mover's obj

    // Bounce, if an end has been reached
    mover.BounceOff()

    // ...
}
```

## Particle System
Is a set of particles (`Particle`, a `Mover` wrapper) with predefined forces applied to them.


**SDL2 Demo**
```golang
// Create group of rects
var rectGroup []*sdl.Rect
// ...

// Create configuration and the particle system itself
psConf := particles.NewConf(false, 100, 100, location, container)
ps := particles.NewParticleSystem(rectGroup, psConf)

// Rendering loop
running := true
for running {
    // Break on quit
    // ...

    // Clear the window on every frame
    renderer.Clear()
    renderer.SetDrawColor(255, 255, 255, 255)
    renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: WindowWidth, H: WindowHeight})

    // Update the particle system and all of the SDL2 Rects
    ps.UpdateSystemSdl2(func(obj interface{}) {
        renderer.SetDrawColor(0, 0, 0, 255)
        switch obj.(type) {
        case *sdl.Rect:
            renderer.DrawRect(obj.(*sdl.Rect))
            break
        }
    })

    renderer.Present()
    sdl.Delay(1000 / FrameRate)
}

```