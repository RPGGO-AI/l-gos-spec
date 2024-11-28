# ID Generation Rules and Descriptions

| **Field Name**    | **Generation Rule**                      | **Length** | **Example**       | **Description**                                                |
|--------------------|------------------------------------------|------------|-------------------|----------------------------------------------------------------|
| `game_id`         | `G` + Random 8 characters from `ID_ALPHABET` | 9          | `GX7Q3PX2M`       | Unique identifier for a game. Starts with `G`.                |
| `character_id`    | `C` + Random 8 characters from `ID_ALPHABET` | 9          | `C4E25YTLT`       | Unique identifier for a character. Starts with `C`.           |
| `module_id`       | `M` + Random 8 characters from `ID_ALPHABET` | 9          | `M3YG78607`       | Unique identifier for a module. Starts with `M`.              |
| `chapter_id`      | First 9 characters of a `uuidv4()`         | 9          | `_R-UvupEO`       | Unique identifier for a chapter, generated using `uuidv4()`.   |
| `entrie_id`       | Random 9 characters using `nanoid(9)`      | 9          | `DmGpL8eo2`       | Unique identifier for an entry, generated using `nanoid(9)`.   |

### Notes:
- **ID_ALPHABET**: Defined as `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ`, ensuring randomness and uniqueness.
- Each ID type has a distinct prefix or generation rule to avoid collisions and ensure easy identification of their purpose.
- All IDs are fixed at **9 characters** for uniformity across the system.
