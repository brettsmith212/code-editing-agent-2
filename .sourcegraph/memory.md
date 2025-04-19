# Code-Editing Agent Memory

## Overview

This is a Go CLI application that connects to Anthropic Claude 3.5 Sonnet and allows you to edit code using tools.

## Code Style Guidelines
- Error handling: Use descriptive error messages with fmt.Errorf and error wrapping with %w
- Function inputs validate parameters before processing
- Naming: Use PascalCase for exported funcs/vars, camelCase for unexported
- Imports are grouped by stdlib, external, then internal packages
- Tool definitions follow a consistent pattern with Name, Description, InputSchema, Function
- Errors are properly propagated and handled
- Use jsonschema_description tags for input struct fields documentation

## Project Structure
- agent/ - Agent implementation connecting to Claude
- tools/ - Tool implementations for file operations
- logger/ - Logging utilities
- logs/ - Log output directory

## Testing and Debugging

1. Tail logs in the `logs/` directory to view the chat history and debug issues from my session I am running.

## Instructinos when building agent tooling

### Rules for Building Effective AI Agents

These rules provide guidelines for creating effective, reliable, and maintainable AI agents. They are based on best practices from the "Building Effective Agents" article and are tailored for your code editor AI Agent to follow while helping you develop an AI Agent code helper.

**Note**: Tool development is critical for agent success. Dedicate significant effort to crafting tools that are intuitive for the LLM, with clear documentation and well-tested interfaces.

### General Principles

1. **Start Simple**: Begin with the simplest solution possible. Optimize single LLM calls using retrieval and in-context examples before moving to multi-step agentic systems.
2. **Evaluate Trade-offs**: Before implementing an agent, determine if the performance gain justifies the increased cost and latency.
3. **Understand Frameworks**: If using a framework, ensure you understand its underlying code to avoid problems caused by abstraction layers.
4. **Maintain Simplicity**: Design the agent to be as simple as possible, adding complexity only when it clearly improves results.
5. **Prioritize Transparency**: Ensure the agent’s planning steps are explicit and transparent to aid debugging and comprehension.
6. **Customize Capabilities**: Tailor the augmented LLM’s capabilities (retrieval, tools, memory) to match the specific needs of your coding agent.

### Workflow Patterns

7. **Prompt Chaining**: Use for tasks that can be broken into sequential steps, such as generating code and then refining it, to boost accuracy.
8. **Routing**: Apply when inputs (e.g., coding queries) can be categorized and directed to specialized processes, like bug fixes vs. feature additions.
9. **Parallelization**: Use to run independent coding subtasks simultaneously (e.g., reviewing different files) or to aggregate multiple outputs for better reliability.
10. **Orchestrator-Workers**: Implement for complex coding tasks where subtasks (e.g., editing multiple files) are determined dynamically by the agent.
11. **Evaluator-Optimizer**: Use for iterative coding improvements, such as refining code based on test feedback, when clear evaluation criteria exist.

### Agent Design

12. **Open-Ended Tasks**: Deploy agents for coding problems with unpredictable steps, such as resolving complex GitHub issues, where fixed workflows won’t suffice.
13. **Trustworthy Decision-Making**: Ensure the agent’s decisions are reliable by testing extensively in sandboxed environments and adding guardrails to limit errors.
14. **Tool Documentation**: Provide clear, detailed documentation for all tools to ensure the LLM can use them correctly during coding tasks.
15. **Adapt Patterns**: Combine and adjust workflow patterns based on performance metrics and the specific demands of your coding project.

### Tool Development

16. **Effective ACI**: Build a clear agent-computer interface (ACI) with well-documented, thoroughly tested tools to support coding tasks.
17. **Intuitive Interfaces**: Design tools with simple parameters and usage patterns that the LLM can easily understand and apply.
18. **Detailed Documentation**: Include examples, edge cases, and distinctions between similar tools in documentation to guide the LLM’s usage.
19. **Extensive Testing**: Test tools with the LLM using varied coding scenarios and refine them based on observed performance.
20. **Simplify Formats**: Use straightforward, natural formats for tool inputs and outputs (e.g., markdown over JSON for code) to minimize errors.
21. **Mistake-Proofing**: Design tools to prevent errors, such as requiring absolute file paths instead of relative ones, following poka-yoke principles.
