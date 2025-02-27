# Pokedex-REPL 🎮🦖

## Authors 🙋‍♂️

- [Joel Joseph](https://www.github.com/joeljosephwebdev)

## Description

This is a fun command line tool, written in go, inspired by the pokedex from the pokemon games.

### Prerequisites 🚀

This project was built using go version 1.23.5. But any version of go 1.2+ should work.

```sh
>> go version
go version go1.23.5 darwin/amd64
```

## Getting Started 💫

To use this tool, simply clone the repo, navigate to the scripts folder and execute the run.sh script. This starts the pokedex, from here you can use the help command to learn about each of the other built in commands.

```sh
  Pokedex > help
  Welcome to the Pokedex!
  Usage:

  map: Display next 20 pokemon locations
  mapb: Display the previous 20 pokemon locations
  pokedex: See all of your caught pokemon
  catch: Throw a pokeball at a specified pokemon. The stronger it is, the harder it is to catch!
  exit: Exits the pokedex
  explore: Displays all pokemon that can be found at provided location.
  help: Displays a help message
  inspect: Display details about a specified pokemon
```

### Example 
```sh
scripts. >> ./run.sh 
Pokedex > help
Welcome to the Pokedex!
Usage:
   - map: Display next 20 pokemon locations
   - mapb: Display the previous 20 pokemon locations
   - pokedex: See all of your caught pokemon
   - catch: Throw a pokeball at a specified pokemon. The stronger it is, the harder it is to catch!
   - exit: Exits the pokedex
   - explore: Displays all pokemon that can be found at the provided location.
   - help: Displays a help message
   - inspect: Display details about a specified pokemon

Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior

Pokedex > explore eterna-forest-area
Exploring eterna-forest
Found Pokemon:
 - caterpie
 - metapod
 - weedle
 - kakuna
 - gastly
 - hoothoot
 - murkrow
 - misdreavus
 - pineco
 - wurmple
 - silcoon
 - beautifly
 - cascoon
 - dustox
 - seedot
 - slakoth
 - nincada
 - bidoof
 - kricketot
 - budew
 - buneary

Pokedex > catch caterpie
Throwing a Pokeball at caterpie...
caterpie was caught!

Pokedex > inspect caterpie
Name: caterpie
Height: 3
Weight: 29
Stats:
   -hp: 45
   -attack: 30
   -defense: 35
   -special-attack: 20
   -special-defense: 20
   -speed: 45
Types:
   -bug

Pokedex > catch pidgey
Throwing a Pokeball at pidgey...
pidgey was caught!
Pokedex > pokedex

Your Pokedex: 
   - pidgey
   - caterpie
Pokedex > exit
Closing the Pokedex... Goodbye!
```
