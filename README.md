# LLM-based Game Open Schema (L-GOS)

![Goose](https://github.com/user-attachments/assets/dc2c4123-1729-4065-ac26-bec60ad0a22b)

I pronunce 'L-GOS' as "Large Goose" :)


## 💡 Vision

The vision of RPGGO is to democratize game creation together with UGC creators. This means that UGC creators have full copyright ownership of their games and the freedom to transfer them to any third-party platform.

This project aims to define an open data schema that any third-party LLM gaming platform can easily support and be compatible with. This way, UGC game creators can import and run their games on any third-party LLM gaming platform.

Only with the support of more and more third-party platforms can a truly decentralized gaming ecosystem emerge. Let's work together toward this goal!

If your product or platform is compatible with this schema, please let us know.

## 🗞️ News

- **Dec 3th 2024:**, change the repo to public

- **Nov 27th 2024:**, finalize the game goal schema and attach samples.

- **Nov 12th 2024:**, init the game schema.

- **Nov 7th 2024:**, finalize the character schema.

- **Oct 10th 2024:**, finalize the module schema.

- **Sept 12th 2024:**, init the schema.


## 📂 Repo Structure

L-GOS-SPEC/ <br>
├── [schema_v1](./schema_v1/)/                                     # V1 Schema folder <br>
│   ├── [module_v1.md](./schema_v1/module_v1.md)                               # Spec for Module.  <br>
│   ├── [character_v1.md](./schema_v1/character_v1.md)                            # Spec for Character. <br>
│   └── [game_v1.md](./schema_v1/game_v1.md)              # Spec for Game. <br>
├── [compatible_with_sillytavern_card.md](./compatible_with_sillytavern_card.md)            # How it compatible with existing sillytavern cards <br>
├── [keyword_definitions.md](./keyword_definitions.md)                         # Some definitions of the keywords used in the spec <br>
├── [README.md](./README.md)                                      # An explanation of what L-GOS proposes, and the thinking behind the schema. <br>
└── [LICENSE](./LICENSE)                                  # The license of the project <br>

## 📑 Introduction

Briefly, the schema contains 3 levels:
- Game, the data structure for running a complete game with LLM. One game can contain many characters.
- Character, the data structure for runnig a single character with LLM. One character can contain many modules
- Module, like Lego blocks, designed to enhance the expression and performance of characters and the game itself.

![game_module](/image/game_module.png)


## ⚖️ License

This project is under CC0, which means you can do anything you want.

<br>

## 🤝 Acknowledgments


<br>

## 📬 Contact

If you have any questions, feedback, or would like to get in touch, please feel free to reach out to us via email at [dev@rpggo.ai](mailto:dev@rpggo.ai)

## ❤️ Thanks

We firmly believe that users should have complete ownership of the AI agents and AI games they create. In the era of generative AI, decentralized ownership is both essential and imperative. By further democratizing creative rights, we can nurture a broader space for innovation and sharing.

We have designed this schema to enable users to import and export their agent and game assets on our platform, with the flexibility to migrate to any AI platform of their choice. Our hope is that this schema will be adopted by more platforms, making cross-platform sharing a reality. 
 
