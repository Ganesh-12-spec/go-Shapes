# go-shapes

A shape calculator built in Go, using interfaces for polymorphism.

## Features
- `Shape` interface (Area, Perimeter) — Circle, Rectangle, Triangle
- `Solid` interface (Volume) — Sphere, Cylinder, Cone
- `TotalArea()` and `LargestArea()` helper functions
- JSON HTTP API: `POST /calculate` — send shape data, get area/perimeter back

## Usage
