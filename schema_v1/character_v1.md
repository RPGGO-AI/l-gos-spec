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
    age?: string; // Optional
    appearance?: string; // Optional
    background?: string; // Optional
    opening_line?: string; // Optional
    dialogue_reference?: string; // Optional
    creator?: string; // Optional
    creator_notes?: string; // Optional
    tone?: string[]; // Optional
    tags?: string[]; // Optional
    modules?: Module[]; // Optional
};
```

also with Python definition:
```python
class CharacterCard(BaseModel):
        name: str
        traits: list[str]
        age: str = ""
        appearance: str = ""
        background: str = ""
        opening_line: str = ""
        dialogue_reference: str = ""
        creator: str = ""
        creator_notes: str = ""
        tone: list[str] = []
        tags: list[str] = []
        modules: list[Module] = []
```

All fields are mandatory and **MUST** default to the empty string, not null or absent/undefined.

In prompts sent to the AI, the fields `appearance`, `background`, `opening_line`, `dialogue_reference`, and `modules` **MUST** replace the following magic strings, with a **case-insensitive** search (e.g. `<BOT>` and `<bot>` both work):
- {{char}} or `<BOT>` to the value of the card's `name` field
- {{user}} or `<USER>` to the application's set display name.

A default value for the user's name **MUST** exist.

Whether {{user}} and `<USER>` should be replaced inside the `name` field is an **UNSPECIFIED** edge case.

Details for each field follows.


### `name`

Used to identify a character.

### `traits`



### `age`



### `appearance`


### `background`

### `opening_line`

### `dialogue_reference`

### `creator`

### `creator_notes`

### `tone`

### `tags`


### `modules`

