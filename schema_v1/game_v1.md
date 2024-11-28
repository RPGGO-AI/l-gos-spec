# Game Card V1: Schema

This exists for reference in case of ambiguity, or for future new implementers.
[Read the definitions of MUST, SHOULD, and MAY](./keyword_definitions.md)

## Table of Contents

- [Embedding Methods](#embedding-methods)
- [Fields](#fields)
  - [Chapter Fields](#chapter-fields)
    - [chapter_id](#chapter_id)
    - [name](#name-1)
    - [cover_url](#cover_url)
    - [background](#background)
    - [intro](#intro)
    - [endings](#endings)
    - [init_dialogues](#init_dialogues)
    - [lore_list](#lore_list)
    - [character_info](#character_info)
    - [characters](#characters)
  - [Game Fields](#game-fields)
    - [game_id](#game_id)
    - [name](#name)
    - [game_tags](#game_tags)
    - [mechanics](#mechanics)
    - [cover_url](#cover_url-1)
    - [lang](#lang)
    - [background](#background-1)
    - [intro](#intro-1)
    - [chapters](#chapters)
    - [background_musics](#background_musics)


    
## Embedding methods

- .json file with no image. Discouraged for user-friendliness.
- PNG/APNG: JSON string encoded in base64 inside the "GAME" EXIF metadata field. Highly recommended for better user-friendliness.

## Fields

The current format can be represented as this TypeScript definition:

```ts
type Chapter = {
    chapter_id: string;
    name: string;
    cover_url: string;
    background: string;
    intro: string;
    goal_info: GoalInfo;
    statuses: Status[];
    endings: string;
    init_dialogues: InitDialogue[];
    lore_list: Lore[];
    character_info: CharacterInfo[];
    characters: Character[];   
};

type Game = {
    game_id: string;
    name: string;
    game_tags: string[];
    mechanics: string;
    cover_url: string;
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
    chapter_id: str = ""
    name: str = ""
    cover_url: str = ""                
    background: str = ""
    intro: str = ""
    goal_info: GoalInfo = ""
    statuses: list[Status] = Field(default_factory=list)
    endings: str = ""
    init_dialogues: list[InitDialogue] = Field(default_factory=list)
    lore_list: list[Lore] = Field(default_factory=list)
    character_info: list[CharacterInfo] = Field(default_factory=list)
    characters: list[Character] = Field(default_factory=list)
    
class Game(BaseModel):
    game_id: str = ""
    name: str = ""
    cover_url: str = ""
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

#### `chapter_id`
A unique identifier for each chapter in the game. The `chapter_id` is generated using a customized **uuidv4()** method, but restricted to **9 characters** for brevity. 

The first 9 characters of a standard uuidv4 string are extracted to serve as the chapter ID.

Example: "550e8400e"
[Check all id rules here](/id_spec.md)

#### `name`
The title or name of the chapter.

#### `cover_url`
A URL or path to an image that visually represents the chapter (e.g., cover art or thematic image).

#### `background`
A description of the setting and context for the chapter, potentially outlining the beginning of the events.

#### `intro`
Introduction of the game, where creators can include information for players. This field **Should** be displayed to players, and **SHOULD NOT** be used in the prompt.

#### `goal_info`
Contains information about the goals for the chapter. This includes two goals which are `success_goal` and `fail_goal`, also `goal_setting` which is a dictionary for meta information for goals.
[Further Check on GoalInfo](./goal_v1.md)

#### `statuses`
An array of `Status` objects, representing the numerical values or statuses that are tracked or affected by the chapter's events.
[Further Check on Status](./goal_v1.md)

#### `endings`
A description of the potential outcomes or resolutions for the chapter, based on player decisions or events that occur at the end of the chapter.

This field **Should** be displayed to players and **SHOULD NOT** be used in the prompt.

#### `init_dialogues`
An array of InitDialogue objects that define the initial conversations that players encounter at the start of the chapter. The `InitDialogue` structure is defined as follows:
This will overwrite the `opening_line` of the character.
```ts
type InitDialogue = {
    name: string;              // The speaker or source of the dialogue
    index: number;             // The order of the dialogue in the sequence, which **Should** be the display order of frontend.
    message: string;           // The text content of the dialogue
    character_id: string;      // The unique identifier for the character speaking
};
```

#### `lore_list`
An array of `Lore` objects, providing additional narrative details and game-play elements related to the chapter. The `Lore` structure is defined as follows:

```ts
type Lore = {
    details: string;           // lore text content
    character_id: string[];    // Array of character IDs associated with the lore: "Who knows this lore".
};
```

#### `character_info`
An array of `CharacterInfo` objects, providing in-depth details about the characters featured in the chapter. The `CharacterInfo` structure is defined as follows:

```ts
type CharacterInfo = {
    emotion: string;           // Current emotional state of the character
    character_id: string;      // Unique identifier of the character
    recent_ongoing: string;    // Summary of the character's recent actions
    personal_setting: string;  // Additional context or personality details
};
```

#### `characters`
An array of `Character` objects, representing the characters that are involved or featured in the chapter.
[Further Check on Character](./character_v1.md)

---

### Game

#### `game_id`
A unique identifier for the game. The `game_id` is generated using the following logic:

- **ID_ALPHABET**: A predefined set of characters consisting of `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ`.
- The ID **MUST** start with the prefix `G` followed by 8 randomly selected characters from `ID_ALPHABET`.

This ID guarantees uniqueness across all games and **MUST** remain immutable once assigned.
Example: GX7Q3PX2M

#### `name`
The title or name of the game.

#### `game_tags`
An array of strings that represent tags related to the game, typically used for sorting, categorizing, or filtering. These tags are case-insensitive.

#### `mechanics`
Field used to support and implement gameplay mechanics in the algorithm.

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

#### `cover_url`
A URL or path to an image that represents the game, often used for promotional purposes or as a cover image.

#### `lang`
The language(s) in which the game is available, typically represented by a language code (e.g., "en" for English, "fr" for French).

#### `background`
A detailed description of the game's world, setting, and overarching story or theme. This could include the history, politics, and key events in the game's universe.

#### `intro`
A short text or narrative that introduces the game to the player, setting the stage for the adventure, and possibly giving an overview of the game's premise. This field **Should** be displayed to players and **SHOULD NOT** be used in the prompt.

#### `chapters`
An array of `Chapter` objects, each representing a specific chapter in the game's storyline. 

#### `background_musics`
An array of strings representing background music tracks or themes used in the game. These may be used to set the mood or enhance gameplay.


![game_character](/image/game_character.png)