# Icons and accessible labels

Use icons to reinforce meaning, not to replace text. For a small interface, a compact inline SVG is usually the cleanest choice. For decorative icons, hide them from assistive technology; for purely symbolic buttons, add an accessible name.

## [STEP 1.1.0] Add a labeled icon button

### Goal
Add a simple icon that supports a real action and remains accessible.

### What this changes
This adds a visible action with a clear text label and icon treatment.

### Files involved
- `frontend/src/components/ExampleActionButton.jsx`
- `frontend/src/index.css`

### Instructions
1. Create a button element.
2. Add an inline SVG and mark it decorative with `aria-hidden="true"`.
3. Keep the visible text alongside the icon.
4. Use a `type="button"` when it is not a form submit.
5. Save and test it in the browser.

```jsx
export default function ExampleActionButton({ onClick }) {
  return (
    <button type="button" className="button button-secondary" onClick={onClick}>
      <svg aria-hidden="true" viewBox="0 0 20 20" width="16" height="16">
        <path d="M16 10a6 6 0 1 1-1.8-4.3M16 4v4h-4" fill="none" stroke="currentColor" strokeWidth="1.7" />
      </svg>
      Refresh
    </button>
  )
}
```

### Do not
- Use an icon alone without a label if it is interactive.
- Give the icon a visual style that makes it harder to see.

### Expected result
The button communicates both the action and the meaning clearly.

---

## [STEP 1.1.1] Add an icon-only button with an accessible name

### Goal
Create a compact action such as a close or delete button without losing accessibility.

### What this changes
This gives an icon-only control a readable label for assistive technology.

### Files involved
- `frontend/src/components/IconButton.jsx`

### Instructions
1. Create an icon-only button.
2. Add `aria-label` with a short, specific action name.
3. Keep the SVG decorative with `aria-hidden="true"`.
4. Use the same button semantics across the app.
5. Test keyboard focus and screen reader behavior.

```jsx
export default function IconButton({ onClose }) {
  return (
    <button type="button" className="icon-button" aria-label="Close dialog" onClick={onClose}>
      <svg aria-hidden="true" viewBox="0 0 20 20" width="18" height="18">
        <path d="m5 5 10 10M15 5 5 15" fill="none" stroke="currentColor" strokeWidth="2" />
      </svg>
    </button>
  )
}
```

### Do not
- Add a button with only a symbol and no label.
- Use non-button elements for interactive icon controls.

### Expected result
The button is usable by keyboard and screen-reader users alike.

---

## [STEP 1.1.2] Reuse icon styling across the app

### Goal
Keep icon appearance consistent with the rest of the interface.

### What this changes
This centralizes sizing and color rules so icons match text and buttons.

### Files involved
- `frontend/src/index.css`

### Instructions
1. Open `frontend/src/index.css`.
2. Add or adjust icon styling for inline button icons and icon-only controls.
3. Use `currentColor` so icons inherit text color.
4. Keep spacing consistent across icon buttons.
5. Refresh the app and verify the visual rhythm.

```css
.icon-button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 2.25rem;
  height: 2.25rem;
  border: none;
  background: transparent;
  color: var(--ink-soft);
  cursor: pointer;
}

.icon-button svg {
  width: 1rem;
  height: 1rem;
}
```

### Do not
- Hard-code colors on each SVG if they can inherit `currentColor`.
- Make icons too large or too small compared with nearby text.

### Expected result
Icons look intentional, aligned, and readable without changing the app's visual system.

---

## Checkpoints & Troubleshooting

### Checkpoint
- [ ] Icons are visible and aligned with surrounding text.
- [ ] Interactive icons are labeled properly.
- [ ] Decorative icons are hidden from assistive tech.
- [ ] Icon-only buttons remain keyboard reachable.
- [ ] SVG size matches the surrounding UI.

### If the icon does not display
1. Confirm the SVG path is valid and the element is not missing a closing tag.
2. Check the `viewBox` and dimensions.
3. Ensure the SVG is not being hidden by a parent CSS rule.

### If the button is not accessible
1. Add `aria-label` if the icon is the only visible content.
2. Keep the icon decorative with `aria-hidden="true"` when text already explains it.
3. Use a real `<button>` instead of a clickable `<div>`.

### If the icon color looks wrong
1. Set `stroke="currentColor"` or `fill="currentColor"`.
2. Confirm the button or parent element defines a color value.
3. Avoid hard-coded colors unless they are truly intentional.
