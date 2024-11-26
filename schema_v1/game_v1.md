# Game Card V1: Schema

This exists for reference in case of ambiguity, or for future new implementers.
[Read the definitions of MUST, SHOULD, and MAY](./keyword_definitions.md)

## Table of contents

- [Embedding methods](#embedding-methods)
- [Fields](#fields)
  - [Chapter Fields](#Chapter)
    * [name](#name)
    * [image_url](#image_url)
    * [goal_info](#goals)
    * [statuses](#statuses)
    * [background](#background)
    * [intro](#intro)
    * [endings](#endings)
    * [characters](#characters)
  - [Game Fields](#Game)
    * [name](#name)
    * [game_tags](#game_tags)
    * [mechanics](#mechanics)
    * [image_url](#image_url)
    * [lang](#lang)
    * [background](#background)
    * [intro](#intro)
    * [chapters](#chapters)
    * [background_musics](#background_musics)

    
## Embedding methods

- .json file with no image. Discouraged for user-friendliness.
- PNG/APNG: JSON string encoded in base64 inside the "MOD" EXIF metadata field. Highly recommended for better user-friendliness.

## Fields

The current format can be represented as this TypeScript definition:

```ts
type Chapter = {
    name: string;
    image_url: string;
    goal_info: GoalInfo;
    statuses: Status[];
    background: string;
    intro: string;
    endings: string;
    characters: Character[];   
};

type Game = {
    name: string;
    game_tags: string[];
    mechanics: string;
    image_url: string;
    lang: string;
    background: string;
    intro: string;
    chapters: Chapter[];
    background_musics: string[];
};
```

or as this python definition:

```python
from pydantic import BaseModel, Field

class Chapter(BaseModel):
    name: str = ""
    image_url: str = ""                
    goal_info: GoalInfo = ""
    statuses: list[Status] = Field(default_factory=list)
    background: str = ""
    intro: str = ""
    endings: str = ""
    characters: list[Character] = Field(default_factory=list)
    
class Game(BaseModel):
    name: str = ""
    image_url: str = ""
    game_tags: list[str] = Field(default_factory=list)
    mechanics: str = ""
    lang: str = ""
    background: str = ""
    intro: str = ""
    chapters: list[Chapter] = Field(default_factory=list)
    background_musics: list[str] = Field(default_factory=list)
```

**Note**: All fields are mandatory and **MUST** default to the empty string, not null or absent/undefined.

A default value for the game and chapter's `name` **MUST** exist.

Details for each field follows.

### Chapter

#### `name`
The title or name of the chapter.

#### `image_url`
A URL or path to an image that visually represents the chapter (e.g., cover art or thematic image).

#### `goal_info`
Contains information about the goals for the chapter. This includes two goals which are `success_goal` and `fail_goal`, also `goal_setting` which is a dictionary for meta information for goals.
[Further Check on GoalInfo](./goal_v1.md)

#### `statuses`
An array of `Status` objects, representing the numerical values or statuses that are tracked or affected by the chapter's events.
[Further Check on Status](./goal_v1.md)

#### `background`
A description of the setting and context for the chapter, potentially outlining the beginning of the events.

#### `intro`
Introduction of the game, where creators can include information to player, This field **Should** be displayed to players, and **SHOULD NOT** be used in prompt. 

#### `endings`
A description of the potential outcomes or resolutions for the chapter, based on player decisions or events that occur at the end of the chapter.

This field **Should** be displayed to players., **SHOULD NOT** be used in prompt. 

#### `characters`
An array of `Character` objects, representing the characters that are involved or featured in the chapter.


[Further Check on Character](./character_v1.md)

---

### Game

#### `name`
The title or name of the game.

#### `game_tags`
An array of strings that represent tags related to the game, typically used for sorting, categorizing, or filtering. These tags are case-insensitive.

#### `mechanics`
Field used to support and implement gameplay mechanics in Algorithm.
```ts
enum Mechanics {
  Adventure = 'Adventure',
  RolePlay = 'Role-Play',
  Mystery = 'Mystery',
  Simulation = 'Simulation',
  Strategy = 'Strategy',
  CYOA = 'CYOA'
};
```

#### `image_url`
A URL or path to an image that represents the game, often used for promotional purposes or as a cover image.

#### `lang`
The language(s) in which the game is available, typically represented by a language code (e.g., "en" for English, "fr" for French).

#### `background`
A detailed description of the game's world, setting, and overarching story or theme. This could include the history, politics, and key events in the game's universe.

#### `intro`
A short text or narrative that introduces the game to the player, setting the stage for the adventure, and possibly giving an overview of the game's premise. This field **Should** be displayed to players. **SHOULD NOT** be used in prompt. 

#### `chapters`
An array of `Chapter` objects, each representing a specific chapter in the game's storyline.

#### `background_musics`
An array of strings representing background music tracks or themes used in the game. These may be used to set the mood or enhance gameplay.

![game_character](/image/game_character.png)