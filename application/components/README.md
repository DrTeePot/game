# Components

## Render

- Mesh
- Texture
- Colour
- Light
- Particles?

## Collision

- Rigidbody, which is also a mesh, however this one doesn't get loaded
    to opengl
- Terrain?

## Physics

- Gravity
- Flow
- Movement

## Functionality

- Identitier (static vs active vs dynamic, uuid, etc)
- StaticTransform (coordinates and orientation)
- DynamicTransform (position and rotation adn scale)
- Growth 
- Health
- Mana type thing
- Inventory
- Equipment
- AI
- Script (programmable in-game)

# Static vs Dynamic vs Active

Some entities move and some don't. Unmoving entities are `Static` entities
and they don't load when not in range, and they don't have anything that
actively runs on them. These are things like terrain, trees, structures, etc.

Moving entities are either `Dynamic` or `Active`. These entities have some 
functionality that is executing.

`Dynamic` entities will run and move  when the player is in range, but freeze 
their state and stop running when the player is out of range. These are 
things like animals. Animals and plants will have algorithmic growth cycles,
and when loaded they will check the time in order to update their growth state.

`Active` entities will be kept in a seperate list from the chunked entities,
and these will always be running even if they aren't being rendered. These
are things like machines. 
