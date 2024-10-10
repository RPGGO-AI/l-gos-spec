# Character Card V1: Specification

This exists for reference in case of ambiguity, or for future new implementers.
[Read the definitions of MUST, SHOULD, and MAY](./keyword_definitions.md)

## Table of contents

- [Embedding methods](#embedding-methods)
- [Fields](#fields)
  * [name](#name)
  * [description](#description)
  * [personality](#personality)
  * [scenario](#scenario)
  * [first_mes](#first_mes)
  * [mes_example](#mes_example)


## Embedding methods

- .json file with no image. Discouraged for user-friendliness.
- PNG/APNG: JSON string encoded in base64 inside the "MOD" EXIF metadata field. Highly recommended for better user-friendliness.


## Fields

The current format can be represented as this TypeScript type:

```ts
type CharacterCard = {
    name: string;
    traits: string[];
    age: string;
    appearance: string;
    background: string;
    opening_line: string;
    dialogue_reference: string;
    creator: string;
    creator_notes: string;
    tone: string[];
    tags: string[];
    modules: Module[];
};
```

also with Python definition:
```python
class CharacterCard(BaseModel):
        name: str = ""
        age: str = ""
        appearance: str = ""
        background: str = ""
        opening_line: str = ""
        dialogue_reference: str = ""
        creator: str = ""
        creator_notes: str = ""
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


### `name`

Used to identify a character.

### `traits`

An Array of adjective words represent the character's personality and trait.

### `age`

Used to identify the character's age.

### `appearance`

Used to describe the character's physical looking.

### `background`

Used to generally describe the character's backstory and additional information.

### `opening_line`

First dialogue of the character, can be used to establish the character's persona or to guide the player's interaction. **SHOULD** be displayed in the chat interface.

### `dialogue_reference`

Sample dialogues to guide the AI's speaking style.

### `creator`

Used to identify the character's creator. This field MAY default to "Anonymous".

### `creator_notes`

A field where creators can include information.

### `tone`

An Array of Strings representing the character's tone of speech.

### `tags`

An array of strings, representing the character's roleplay feature, with no restrictions.

This field **MAY** be used for frontend sorting/filtering purposes (case-insensitive recommended).

This field **SHOULD NOT** be used for the prompt engineering.

### `modules`

An array of modules like Lego blocks, designed to enhance the expression and performance of characters.

[Further Check on module](./module_v1.md)

![373492583-0f9de6a9-1671-4eb6-91c4-2588aba9aa75](https://github.com/user-attachments/assets/0a67748d-1982-401c-86e4-bf9f013ed2d8)



