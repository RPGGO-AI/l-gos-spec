# Game Card V1: Schema

This exists for reference in case of ambiguity, or for future new implementers.
[Read the definitions of MUST, SHOULD, and MAY](./keyword_definitions.md)

## Table of contents

- [Embedding methods](#embedding-methods)
- [Fields](#fields)
  - [Chapter Fields](#Chapter)
    * [name](#name)
    * [image](#image)
    * [background](#background)
    * [goals](#goals)
    * [intro](#intro)
    * [endings](#endings)
    * [characters](#characters)
  - [Game Fields](#Game)
    * [name](#name)
    * [category](#category)
    * [game_tags](#game_tags)
    * [mechanics](#mechanics)
    * [image](#image)
    * [lang](#lang)
    * [background](#background)
    * [intro](#intro)
    * [chapters](#chapters)
    * [worlds](#worlds)
    * [moderation_level](#moderation_level)
    * [background_musics](#background_musics)

    
## Embedding methods

- .json file with no image. Discouraged for user-friendliness.
- PNG/APNG: JSON string encoded in base64 inside the "MOD" EXIF metadata field. Highly recommended for better user-friendliness.

## Fields

The current format can be represented as this TypeScript definition:

```ts
type Chapter = {
    name: string;
    image: string;
    background: string;
    goals: Goal[];            
    intro: string;
    endings: string;
    characters: Character[];   
};

type Game = {
    name: string;
    category: string;
    game_tags: string[];
    mechanics: string;
    image: string;
    lang: string;
    background: string;
    intro: string;
    chapters: Chapter[]; 
    worlds: World[];
    moderation_level: string;
    background_musics: string;
};
```

or as this python definition:

```python
class Chapter(BaseModel):
    name: str
    image: str                # necessary?
    background: str
    goals: list[Goal]         # what should we do about the goals?
    intro: str
    endings: str
    characters: list[Character]
    
class Game(BaseModel):
    name: str
    category: str             # abandoned here?
    game_tags: list[str]
    mechanics: str            # limited choices?
    image: str                # poster? should each field name of image be aligned? 
    lang: str
    background: str
    intro: str                # this is part of prompt, different from field at character/module card
    chapters: list[Chapter]
    worlds: list[World]       # There are worlds? too complex?
    moderation_level: str     # abandoned?
    background_musics: str
```

Details for each field follows.

### Chapter

#### name
The title or name of the chapter.

#### image
A URL or path to an image that visually represents the chapter (e.g., cover art or thematic image).

#### background
A description of the setting and context for the chapter, providing insight into the world or situation in which the chapter takes place.

#### goals
An array of `Goal` objects that define the objectives or missions that the player or characters must accomplish during this chapter.

#### intro
A brief introductory text that sets the stage for the chapter, potentially outlining the beginning of the events or the mood.

#### endings
A description of the potential outcomes or resolutions for the chapter, based on player decisions or events that occur during the chapter.

#### characters
An array of `Character` objects, representing the characters that are involved or featured in the chapter. These may be player-controlled or NPCs.

---

### Game

#### name
The title or name of the game.

#### category
A category or genre that describes the game (e.g., RPG, adventure, puzzle, etc.).

#### game_tags
An array of strings that represent tags related to the game, typically used for sorting, categorizing, or filtering. These tags are case-insensitive.

#### mechanics
A description of the gameplay mechanics, explaining how the game is played (e.g., turn-based combat, puzzle-solving, resource management).

#### image
A URL or path to an image that represents the game, often used for promotional purposes or as a cover image.

#### lang
The language(s) in which the game is available, typically represented by a language code (e.g., "en" for English, "fr" for French).

#### background
A detailed description of the game's world, setting, and overarching story or theme. This could include the history, politics, and key events in the game's universe.

#### intro
A short text or narrative that introduces the game to the player, setting the stage for the adventure, and possibly giving an overview of the game's premise.

#### chapters
An array of `Chapter` objects, each representing a specific chapter in the game's storyline.

#### worlds
An array of `World` objects, representing the different worlds or environments the game spans, each potentially with its own unique setting, theme, and challenges.

#### moderation_level
A string that defines the level of content moderation or suitability for different audiences (e.g., "E" for Everyone, "M" for Mature).

#### background_musics
A string or array of strings representing background music tracks or themes used in the game. These may be used to set the mood or enhance gameplay.
