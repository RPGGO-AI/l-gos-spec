# Character Card V1: Specification

This exists for reference in case of ambiguity, or for future new implementers.
[Read the definitions of MUST, SHOULD, and MAY](./keyword_definitions.md)

## Table of contents

- [Embedding methods](#embedding-methods)
- [Fields](#fields)
  * [name](#name)
  * [avatar](#avatar)
  * [appearance](#appearance)
  * [background](#background)
  * [opening_line](#opening_line)
  * [creator](#creator)
  * [creator_notes](#creator_notes)
  * [intro](#intro)
  * [traits](#traits)
  * [tone](#tone)
  * [tags](#tags)
  * [modules](#modules)


## Embedding methods

- .json file with no image. Discouraged for user-friendliness.
- PNG/APNG: JSON string encoded in base64 inside the "MOD" EXIF metadata field. Highly recommended for better user-friendliness.


## Fields

The current format can be represented as this TypeScript type:

```ts
type CharacterCard = {
    name: string;
    avatar: string;
    appearance: string;
    background: string;
    opening_line: string;
    creator: string;
    creator_notes: string;
    intro: string;
    traits: string[];
    tone: string[];
    tags: string[];
    modules: Module[];
};
```

also with Python definition:
```python
class CharacterCard(BaseModel):
    name: str = ""
    avatar: str = ""
    appearance: str = ""
    background: str = ""
    opening_line: str = ""
    creator: str = ""
    creator_notes: str = ""
    intro: str = ""
    traits: list[str] = []
    tone: list[str] = []
    tags: list[str] = []                           
    modules: list[Module] = []
```

All fields are mandatory and **MUST** default to the empty string/list, not null or absent/undefined.

In prompts sent to the AI, the fields `appearance`, `background`, `opening_line`, `dialogue_reference`, and `modules` **MUST** replace the following magic strings, with a **case-insensitive** search (e.g. `<BOT>` and `<bot>` both work):
- {{char}} or `<BOT>` to the value of the card's `name` field
- {{user}} or `<USER>` to the application's set display name.

A default value for the user's name **MUST** exist.

Whether {{user}} and `<USER>` should be replaced inside the `name` field is an **UNSPECIFIED** edge case.

Details for each field follows.

### name
Used to identify a character.

### avatar
The URL or path to the character's avatar image.

### appearance
A description of the character's physical appearance.

### background
General background information and backstory about the character.

### opening_line
The character's first line of dialogue, which may be used to establish their persona or guide interaction with the player. **Should** be displayed in the chat interface.

### creator
Identifies the character's creator. This field **may** default to "Anonymous" if not specified.

### creator_notes
A field for the creator to include additional notes to bot developers.

### intro
A brief introduction to the character, which could be used to provide a summary of who they are. This field **Should** be displayed to players.

### traits
An array of adjectives that represent the character's personality traits.

### tone
An array of strings describing the character's tone of speech or manner of expression.

### tags
An array of strings that represent the character's roleplay features. These may be used for sorting or filtering on the frontend (case-insensitive), but **should not** be used for prompt engineering.

### modules
An array of modules, like building blocks, designed to enhance the expression and performance of characters in the system.


[Further Check on module](./module_v1.md)

![373492583-0f9de6a9-1671-4eb6-91c4-2588aba9aa75](https://github.com/user-attachments/assets/0a67748d-1982-401c-86e4-bf9f013ed2d8)



