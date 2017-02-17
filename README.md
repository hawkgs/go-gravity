# go-gravity
Small Golang library for simulating 2D particles with the help of SDL2.

![GoGravity](https://raw.githubusercontent.com/hAWKdv/go-gravity/master/logo.png)

## Setup
In order to run/try the code in `rendering` package, you must install SDL2 along with the [Golang bindings](https://github.com/veandco/go-sdl2). Of course, you are free to choose whatever rendering method you want since the core of the library is decoupled from SDL2.

## Vector
The vector struct is one of the cores of the library. You can perform basic operations over it like addition, multiplication, normalization and many more.

## Force
The library has a set of forces (vectors) like wind, gravity, kinetic friction, etc.

## Mover
The movers represent a moving particle which is affected by the forces you apply on it. It has its own location, velocity and mass.


```golang
var location *Vector // Initial location
var container *Vector // The boundaries where the mover can move
var obj interface{} // Represents the object of rendering engine you are using (eg. SDL2_Rect)

// Our mover
mover := vectors.NewMover(&obj, location, container)

// Some forces
gravity := forces.CreateGravity()
push := forces.CreatePush(315, 0.2, 30)
friction := forces.CreateKineticFriction(mover)

// This represents the rendering loop. Each iteration = new frame.
// It is all rendering engine-specific.
running := true
for running {
    // ...

    mover.ApplyForce(gravity.GetForce())
    mover.ApplyForce(push.GetForce())
    mover.ApplyForce(friction.GetForce())

    mover.Update()
    // Update mover's obj

    mover.BounceOff()

    // ...
}
```

## Particle System
Is a set of particles (Particle, a Mover wrapper) with predefined forces applied to them.
