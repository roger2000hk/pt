package main

import "github.com/fogleman/pt/pt"

// http://graphics.cs.williams.edu/data/meshes/dabrovic-sponza.zip

func main() {
	scene := pt.Scene{}
	material := pt.GlossyMaterial(pt.HexColor(0xFCFAE1), 1.5, pt.Radians(20))
	mesh, _ := pt.LoadOBJ("examples/dabrovic-sponza/sponza.obj", material)
	mesh.MoveTo(pt.Vector{}, pt.Vector{0.5, 0, 0.5})
	scene.Add(mesh)
	scene.Add(pt.NewSphere(pt.Vector{0, 20, 0}, 3, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	camera := pt.LookAt(pt.Vector{-10, 2, 0}, pt.Vector{0, 4, 0}, pt.Vector{0, 1, 0}, 45)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 2560/2, 1440/2, -1, 16, 3)
}
