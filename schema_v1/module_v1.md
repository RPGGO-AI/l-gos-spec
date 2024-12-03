# Module Card V1: Schema

This exists for reference in case of ambiguity, or for future new implementers.
[Read the definitions of MUST, SHOULD, and MAY](./keyword_definitions.md)

## Table of contents

- [Embedding methods](#embedding-methods)
- [Fields](#fields)
  * [module_id](#module_id)
  * [name](#name)
  * [creator](#creator)
  * [intro](#intro)
  * [category](#category)
  * [entries](#entries)
  * [cover_url](#cover_url)


## Embedding methods

- .json file with no image. Discouraged for user-friendliness.
- PNG/APNG: JSON string encoded in base64 inside the "MOD" EXIF metadata field. Highly recommended for better user-friendliness.


## Fields

The current format can be represented as this TypeScript definition:

```ts
type Entry = {
    entry_id: string;
    keys: string[];
    content: string;
};

type Module = {
    module_id: string;
    name: string;
    creator: string;
    intro: string;
    category: string
    cover_url: string;
    entries: Entry[];
};
```

or as this python definition:

```python
from pydantic import BaseModel, Field

class Entry(BaseModel):
    entry_id: str = ""
    keys: list[str] = Field(default_factory=list)
    content: str = ""
    
class Module(BaseModel):
    module_id: str = ""
    name: str = ""
    creator: str = ""
    intro: str = ""
    category: str = ""
    cover_url: str = ""
    entries: list[Entry] = Field(default_factory=list)
```

**Note**: All fields are mandatory and **MUST** default to the empty string, not null or absent/undefined.

A default value for the module's `name` **MUST** exist.

Details for each field follows.

### `module_id`
A unique identifier for each module. The `module_id` is generated using the following logic:

- **ID_ALPHABET**: A predefined set of characters consisting of `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ`.
- The ID **MUST** start with the prefix `M` followed by 8 randomly selected characters from `ID_ALPHABET`.

This ensures uniqueness and immutability across all modules.
Example: "M1A2B3C4D"

[Check all id rules here](/id_spec.md)
### `name`
Used to identify a module.

### `creator`
Used to identify the module creator in case he/she wants to be remembered.

### `intro`
Introduction of the module, where creators can include information to player, and **SHOULD NOT** be used in prompt. 

### `category`
Used to identify the usage purpose of a module. It **MUST** be one of the following values:

### `cover_url`
The URL or path to the Module's cover image.

```ts
enum Module_Category {
  World_Knowledge = 'World Knowledge',     
  Event_Lore = 'Event Lore',                
  Character_Lore = 'Character Lore',        
  Motion_Enhance = 'Motion Enhance',        
  Language_Enhance = 'Language Enhance',    
  Output_Format = 'Output Format',          
  User_Persona = 'User Persona',           
  Others = "Others"                        
};
```
where,
- World_Knowledge - Identify the module as a world book prompt set
- Event_Lore - Identify the module as a event lore set
- Character_Lore - Identify the module as additional details and backstory for AI characters
- Motion_Enhance - Identify the module to serve as the enhancement of the character's motion
- Language_Enhance - Identify the module to enhancement the AI robot's language expression abilities, such as word choice, sentence structure, stylization, and conversational tone
- Output_Format - Identify the module to improve the AI robot's output format of the conversation
- User_Persona - Identify the module as additional details of AI character's personality
- Others - Used if the module cannot be identified by previous options.

### `entries`
The prompt set to be sent to the AI.

Each entry contains :
- entry_id, which is a unique identifier for each entry within a module. The `entry_id` is generated using the **nanoid** library with a length of 9 characters. Nanoid generates a compact, random, and URL-friendly ID that minimizes collisions while being concise.
- keys, which is used as tags to identify the following content, which **MAY** be used for searching
- content, which is the real prompt. 
**Note** In prompts sent to the AI, the fields `content` **MUST** replace the following magic strings, with a **case-insensitive** search (e.g. `<BOT>` and `<bot>` both work):
- {{char}} or `<BOT>` to the name of the character which is applied with this module
- {{user}} or `<USER>` to the game player's set display name.

some examples of the entry:
- "resume, profile, childhood": "{{char}} is borned in United States and currently 19 years old, blablabla"
- "secret, tone": "{{char}} should use a gentle, alluring tone, which makes other feel there are some secrets hided"

![module_entry](/image/module_entry.png)