# Amazeing
Maze-Solver in Go.

The maze is provided as a PNG file. The program searches for the shortest path from a blue pixel (start) to a green pixel (end). The legal paths are depicted as black.

## Example
<figure>
  <img src="https://user-images.githubusercontent.com/51135157/233867966-8583743e-0a23-43ee-b509-81a5f8c0fa57.png" alt="maze6.png" style="width: 200px; height: 200px; object-fit: none; image-rendering: pixelated;">
  <figcaption>maze6.png</figcaption>
</figure>

Executing `go run amazeing.go maze6.png` creates:

<figure>
  <img src="https://user-images.githubusercontent.com/51135157/233867989-157e9647-4d52-4102-890d-75ee1aa3912a.png" alt="maze6_solved.png" style="width: 200px; height: 200px; object-fit: none; image-rendering: pixelated;">
  <figcaption>maze6_solved.png</figcaption>
</figure>
