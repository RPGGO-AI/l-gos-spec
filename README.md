# General LLM-based Game Open Schema (L-GOS)

The vision of RPGGO is to democratize game creation together with UGC creators. This means that UGC creators have full copyright ownership of their games and the freedom to transfer them to any third-party platform.

This project aims to define an open data schema that any third-party LLM gaming platform can easily support and be compatible with. This way, UGC game creators can import and run their games on any third-party LLM gaming platform.

Only with the support of more and more third-party platforms can a truly decentralized gaming ecosystem emerge. Let's work together toward this goal!

## 🗞️ News

**UPDATE Sept 12th 2024:**
- init the schema.


## 📂 Repo Structure

L-GOS-SPEC/ <br>
├── [keyword_definitions.md](./keyword_definitions.md)    # Some definitions of the keywords used in the spec <br>
├── [schema_v1](./schema_v1/)/                            # V1 Schema folder <br>
│   ├── [module_v1.md](./schema_v1/module_v1.md)          # Spec for Module.  <br>
│   └── [character_v1.md](./schema_v1/character_v1.md)    # Spec for Character. <br>
├── [README.md](./README.md)                              # This document. An explanation of what L-GOS proposes, and the thinking behind the schema. <br>
└── [LICENSE](./LICENSE)                                  # The license of the project <br>

## 📑 Introduction

Briefly, the schema contains 3 levels:
- Game, the data structure for running a complete game with LLM. One game can contain many characters.
- Character, the data structure for runnig a single character with LLM. One character can contain many modules
- Module, like Lego blocks, designed to enhance the expression and performance of characters and the game itself.

![whiteboard_exported_image (11)](https://github.com/user-attachments/assets/0f9de6a9-1671-4eb6-91c4-2588aba9aa75)



## ⚖️ License

This project is under CC0, which means you can do anything you want.

<br>

## 🤝 Acknowledgments


<br>

## 📬 Contact

If you have any questions, feedback, or would like to get in touch, please feel free to reach out to us via email at [dev@rpggo.ai](mailto:dev@rpggo.ai)
