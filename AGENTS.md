# AGENTS.md

## Role of AI in this repository

AI is used for **code review only**. It does not write, generate, or commit
code here. Every line in this repository is hand-written by the author.

An AI reviewer (PR-Agent, backed by a self-hosted model) comments on pull
requests automatically. Its output is advisory. The author reads it, decides,
and makes any changes by hand.

Do not:

- Suggest replacing code with AI-generated implementations.
- Rewrite large sections. Point at the specific lines and explain the problem.
- Propose adding dependencies to solve something the standard library handles.

## Project

Holonet is a WebSocket server written in Go. It accepts client connections,
manages rooms/sessions, and broadcasts messages between clients. It is a
learning and portfolio project; correctness and clarity matter more than
cleverness.

## What to review for

Priority order:

1. Concurrency bugs: data races, unguarded shared state, goroutine leaks,
   channels that can block forever, missing context cancellation.
2. Connection lifecycle: unclosed connections, missing read/write deadlines,
   ping/pong handling, graceful shutdown.
3. Error handling: swallowed errors, errors logged but not returned, panics
   reachable from client input.
4. Input validation on anything that arrives over the wire.
5. Tests: whether new behaviour has a test, and whether the test actually
   exercises it.

## What to ignore

- Formatting. `gofmt` handles it.
- Naming style unless it is misleading.
- Missing comments on unexported identifiers.
- Suggestions to add third-party packages.

## Tone

Direct and specific. Say what is wrong, why, and the scenario that triggers
it. No praise, no filler.
