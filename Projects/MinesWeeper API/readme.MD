# functional requirements

- Get a game: given an id, it should return the game associated.
- Create a new game: given a name, a size (N) and the number of bombs (B),
it should create and return a new game where its board has NxN cells and B bombs
in random positions.
- Reveal a cell: given a game id and cell position, it should reveal the cell and return the resulted game. If a bomb is found or the last empty cell is revealed then the game is over.

