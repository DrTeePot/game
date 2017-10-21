# Engine Architecture

See plan.md as well, these documents need to be reconciled. 

Potential name for the React-like rendering engine might be `Oxygen`,
`Carbon`, or `Fluorine` (flourine is highly reactive, oxygen causes a very
common reaction, and carbon because of hexagons). 

Potential name for the game engine as a whole might be `Benzene`, `Graphite`,
or some similar carbon-based name (due to the hexagons). 

## Philosophy

The purpose of this project is to try and create a game engine functional
ideas of data streams rather than object manipulation. 

Using Go's channels, we will create a somewhat complex version of Redux, 
although our state will be held in multiple objects (reducers) that will work
on their own domain area. 

The state of the engine will be maintained and operated on by a series
of pure functions, called `reducers`, that will take the current state and
an atomic action as input and return the new state.

## Areas of Concern

In a game engine, there are some main components: 

- Game logic
- Rendering
- Audio engine
- Physics engine
- Artificial intelligence
- Scripting Engine

These areas of concern will each have their own state object, along with a 
series of reducers that will act on the state. Actions will be created 
seperately, and each module will need to import the actions that
they will be listening to. There will also be global state. Reducers may
act on one or the other, but will always recieve both state objects as input.

In the future, the state architecture may change to only include a global 
state for the engine, and a seperate state object for the application logic.

Actions must have no dependencies, and be stateless. 

The rendering engine, audio engine, physics engine, and artificial 
intelligence engine will all be abstracted into their own libraries. Likely
rendering, physics, and audio will occupy one `reactive` library that will
be used to build game components. This `reactive` library will observe the
prop changes of the components in the game and update accordingly.

### Game (Application) Logic

Similar to how React and Redux work together, games will be made of components
that are connected to the over-arching `redux` system. Components will have 
properties that define which parts of the application state they are 
interested in, and properties that will dispatch actions.

Components will not have state; however, the application as a whole will
have state. Components will be "re-rendered" when changes to this application
state happen that effect state objects they are subscribed to. 

A likely action to take is to have an object change it's 3-dimensional 
coordinates, which is an action that will have an object-id and the new
coordinates. A reducer will then modify the objects coordinates in state, 
and the render engine (which is listening on components) will re-render the 
component. 

### Rendering Engine

This will provide an interface to build 3d and 2d models with. These models
can then be connected and composed to create complex game objects. 

### Audio Engine

This will provide an interface to add sounds to the world and perform 
a variety of actions on these sounds.

### Physics Engine

This will provide an interface to add physics properties to an object. This 
will involve sending large amounts of actions to the application state object,
so the sending and modification of state will need to be very efficient. 

### Aritificial Intelligence

This will be implemented far in the future, when I begin to have a need for
NPC's. 

### Scripting Engine

This is a cool feature that will allow players to add their own logic to the 
game. Also will be implemented far in the future when I figure out what 
this actually will mean in practise. 

## Structure

With that architecture, there will be a library for utilities around implmenting
the state and engine structure, and then the game code itself.

The proposed tree is:

- engine (fluorine?):
  - render
  - world generator (honeycomb?, benzene?)
  - audio
  - physics
  - scripting
  - AI (when created)
  - reducers
  - actions
- game logic
  - main.go (entrypoint)
  - components
    - game.go - main component
    - worlds
    - terrains
      - probably just different generator options and some custom logic
    - liquids
    - foliage
    - structures
    - animals
      - probably a generator that just creates a bunch from a map? 
    - etc
  - reducers
    - application logic
  - actions
    - atomic custom actions that can happen in the game, note that engine 
        related actions will come from the engine
  - middleware
    - side effects from actions that can't be handled with state, like logging
      or syncronization
