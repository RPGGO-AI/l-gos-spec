# Goal & Status V1: Schema

## Introduction
For each chapter in a game, there is a pair of goals: a goal that allows the player to progress to the next chapter, or a goal that results in the game ending if failed. If the description of a goal is too complex or general, it can be broken down into smaller goals that are easier for a language model (LLM) to judge. A list of goals also helps players have a clear understanding of what they need to do in each chapter, and it can enhance the sense of game progression. 

A goal can be associated with 0 or more statuses. If the goal design is related to the numerical system, a status list maintained within the game can assist in determining whether the goal has been achieved, regardless of whether it is judged by LLM or hard-coded.
## Fields
The current format can be represented as this TypeScript definition:
```ts
enum StatusType {
    Player = "player",
    Game = "game",
    NPC = "npc"
}

// Basic unit of game numerical system
interface Status {
    id: string;
    name: string;
    affiliation?: CharacterCard; // Optional depending on the affiliationType
    affiliationType: StatusType;
    value: number;
}

interface Goal {
    id: string;
    description: string;
    statuses: Status[]; // Optional depending on the goal
    status_thresholds: number[]; // Optional depending on the goal
    subgoals: Goal[]; // Optional depending on the goal
}

interface GoalInfo {
    goals: Goal[];
    goal_setting: Record<string, any>; // Meta information for goals
}
```
or as this python definition:
```python
class StatusType(Enum):
    player = "player"
    game = "game"
    npc = "npc"


# Basic unit of game numerical system
class Status(BaseModel):
    id: str
    name: str
    affiliation: Optional[CharacterCard]  # Optional depending on the affiliationType
    affiliation_type: StatusType
    value: int
    

class Goal(BaseModel):
    id: str
    description: str
    statuses: Optional[list[Status]]  # Optional depending on the goal
    status_thresholds: Optional[list[int]]  # Optional depending on the goal
    subgoals: Optional[list[Goal]]  # Optional depending on the goal

    
class GoalInfo(BaseModel):
    sucess_goal: Goal
    fail_goal: Goal
    goal_setting: dict # Meta information for goals
```