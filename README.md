# Cartographers: A Roll Player Tale

This is a clone of the boardgame: [Cartographers: A Roll Player Tale](https://www.thunderworksgames.com/cartographers.html).

The official ways to play this game are:
- [Happle Meeple](https://www.happymeeple.com/en/board-games/cartographers/overview/)
- [Android](https://play.google.com/store/apps/details?id=de.brettspielwelt.cartographers)
- [iOS](https://apps.apple.com/us/app/cartographers/id1515991485)

## Motivations

A software developer who wants to play Cartographers with his partner, in a remote way.

Both of us enjoy the drawing part, so we only want something to prepare the game for us, and to control the current game state.

## Project Goals

- Create a game setup generator, which prepares the game and sorts the 4 objectives
- Have a simple way to draw the next card
- Have a simple game state control, which:
  - knows what season it is
  - change season when needed
  - store the 4 objectives (A, B, C and D)
- Create a beatiful UI to present the current game state
