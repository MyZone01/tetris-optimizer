#   TETRIS OPTIMIZER

##  Description
This is a program that receives only one argument, a path to a text file which will contain a list of tetrominoes and assemble them in order to create the smallest square possible.

##  Steps
+   [x] Read in the text file specified in the argument
+   [x] Parse the file contents into a list of tetrominoes
+   [ ] Calculate the total area of all the tetrominoes combined
+   [ ] Calculate the minimum square size needed to fit all the tetrominoes by taking the square root of the total area and rounding up to the nearest integer
+   [ ] Create a 2D grid with dimensions equal to the minimum square size calculated in step 4
+   [ ] Attempt to place each tetromino on the grid in order, starting from the top-left corner and moving across and down
+   [ ] If a tetromino cannot be placed in its current orientation, try rotating it and/or moving it to a different position until it can be placed
+   [ ] If all orientations and positions have been tried and the tetromino still cannot be placed, backtrack and try a different placement for the previous tetromino
+   [ ] If all possible placements for all tetrominoes have been tried and none of them can be placed, return an error message indicating that it is impossible to assemble the tetrominoes into a square of the required size
+   [ ] If all tetrominoes have been successfully placed on the grid, return the completed grid as output

##  Source 
[Tetromino Wikipedia](https://en.wikipedia.org/wiki/Tetromino)
[Fillit: Solving for the Smallest Square of Tetrominoes](https://www.bing.com/ck/a?!&&p=2703119851278429JmltdHM9MTY3ODQwNjQwMCZpZ3VpZD0zOGQyYjVkZC01YTc0LTZiODUtMDExZS1hNDdjNWIzOTZhZTUmaW5zaWQ9NTE4NQ&ptn=3&hsh=3&fclid=38d2b5dd-5a74-6b85-011e-a47c5b396ae5&psq=fill+tetromino+to+a+min+square&u=a1aHR0cHM6Ly9tZWRpdW0uY29tL0BiZXRobmVubmlnZXIvZmlsbGl0LXNvbHZpbmctZm9yLXRoZS1zbWFsbGVzdC1zcXVhcmUtb2YtdGV0cm9taW5vcy1jNjMxNjAwNGY5MDk&ntb=1)