# Adaptive Planner

## Architecture

- Go is the primary application language.
- PostgreSQL is accessed through SQLC.
- Server-side HTML uses Go html/template.
- HTMX handles normal interactive HTTP behavior.
- Vanilla JavaScript is reserved for browser-specific behavior.
- Do not introduce React, Vue, Svelte, TypeScript, Node, npm, or frontend build tooling without explicit approval.

## Frontend

- Prefer semantic HTML.
- Prefer server-rendered HTML over client-side rendering.
- Keep HTMX behavior close to the HTML that triggers it.
- Keep custom JavaScript minimal.
- UI must work on desktop and mobile.
- Accessibility is required.
- Avoid unnecessary dependencies.

## Changes

- Make small focused changes.
- Do not refactor unrelated code.
- Explain architectural changes before making them.
- Run Go tests after Go changes.
