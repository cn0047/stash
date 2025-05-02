Prompt Engineering
-

Prompt - straightforward question or instruction given to AI system.
Meta prompt - instruction that guides how AI should respond (context, output format).
Mega prompt - large, detailed prompt, with extensive information and comprehensive details to guide AI system.

Flow:
* Start with decomposition.
* Provide simple, clear and complete instructions. Assign role and audience.
* Place question at the end of prompt.
* Use separator characters for API calls.
* Specify output format.
* Use output indicators. Clarify questions.
* Control model response with inference parameters.

System architecture methodology:
* RAG (Retrieval-Augmented Generation) - technique to enhance performance of LLMs
  by combining them with public or external data at runtime.

## Prompting technique:

**COT** (Chain of Thought) - technique that encourages the model to break down complex problem into series
of logical steps before providing solution.
Structure: Thought → Thought → ... → Answer.

````
Question: [complex or multi-step]
Thought: [step-by-step reasoning]
Answer: [final answer]
````

**ReAct** (Reasoning and Acting) - technique that enables LLMs to generate reasoning traces
and task-specific actions in an interleaved manner.
Structure: Thought → Action → Observation → ... → Answer.
Actions: calculator, search, use public information, etc.

````
Question: [here]

Thought 1: [think through]
Action 1: [take action]
Observation 1: [result of action]

Thought 2: [...]
Action 2: [...]
Observation 2: [...]

...

Answer: [final answer]
````

**DSP** (Directional Stimulus Prompting) - technique to guide LLMs behavior or reasoning style by showing it how to think.
Structure: Direction → Input → Output → … → Answer.

````
Instruction: [task or question]
Direction: [style or perspective]
Example: # like hint
- Input: [sample input]
  Output: [sample response]

Then main input:
Input: [new input]
Output: [response in same direction]
````
