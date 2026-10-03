# Cards and reusable content

A card is a small, reusable content container. Use it to group text, links, actions, or metadata that belong together. Keep the structure simple, and only make a reusable component when the same pattern appears more than once.

## [STEP 1.1.0] Build a simple card component

### Goal
Create a card that can be reused across pages without repeating markup.

### What this changes
This adds a consistent content block for product summaries, feature descriptions, or informational sections.

### Files involved
- `frontend/src/components/InfoCard.jsx`
- `frontend/src/index.css`

### Instructions
1. Create a new component in `frontend/src/components/`.
2. Accept `title`, `children`, and optional `id` props.
3. Render a semantic `<section>` and use the title as a heading.
4. Keep the content inside a class-based wrapper for styling.
5. Import the component into the page where it is needed.

```jsx
export default function InfoCard({ title, children, id }) {
  return (
    <section className="info-card" aria-labelledby={id ? `${id}-title` : undefined}>
      <h2 id={id ? `${id}-title` : undefined}>{title}</h2>
      <div className="info-card-content">{children}</div>
    </section>
  )
}
```

```css
.info-card {
  background: var(--paper);
  border: 1px solid var(--line);
  border-radius: 16px;
  padding: 1.5rem;
  box-shadow: var(--shadow);
}

.info-card h2 {
  margin: 0 0 0.75rem;
}

.info-card-content {
  display: grid;
  gap: 0.75rem;
}
```

### Do not
- Build a reusable component for a one-off card that never appears again.
- Put unrelated page logic inside a generic card component.

### Expected result
The card component is reusable and visually consistent with the rest of the app.

---

## [STEP 1.1.1] Use the card in a page

### Goal
Place the reusable card on a page and fill it with the correct content.

### What this changes
This turns the generic card into a real page section with meaningful content.

### Files involved
- `frontend/src/pages/ExamplePage.jsx`
- `frontend/src/components/InfoCard.jsx`

### Instructions
1. Open the page that needs the card.
2. Import the component from `frontend/src/components/InfoCard.jsx`.
3. Use it with a title and content.
4. Keep the content concise and product-focused.
5. Save and view the page in the browser.

```jsx
import InfoCard from '../components/InfoCard.jsx'

export default function ExamplePage() {
  return (
    <main className="page-shell">
      <InfoCard id="getting-started" title="Getting started">
        <p>Use this template to build a clean, authenticated product shell.</p>
        <a href="/guide" className="link">Read the guide</a>
      </InfoCard>
    </main>
  )
}
```

### Do not
- Add too much content to a single card.
- Mix multiple unrelated ideas inside one card.

### Expected result
The card displays the intended information and stays readable.

---

## [STEP 1.1.2] Create a card grid layout

### Goal
Display multiple cards in an even, responsive layout.

### What this changes
This arranges cards into a grid that adapts to wider screens while staying readable on mobile.

### Files involved
- `frontend/src/index.css`

### Instructions
1. Open `frontend/src/index.css`.
2. Add a layout class for a card grid.
3. Use `repeat(auto-fit, minmax(240px, 1fr))` to keep a flexible layout.
4. Test the grid across screen sizes.

```css
.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 1rem;
}
```

### Do not
- Force every card to be the same height if the content is different.
- Create a grid that becomes cramped on small phones.

### Expected result
The cards arrange themselves cleanly without overloading the page.

---

## Checkpoints & Troubleshooting

### Checkpoint
- [ ] Each card has a clear heading.
- [ ] The content is grouped logically.
- [ ] The layout adapts on smaller screens.
- [ ] The component is reused where it makes sense.
- [ ] The card remains readable without too much vertical density.

### If the card looks stretched
1. Confirm the parent container is not using a fixed height.
2. Check grid or flex rules for `align-items` behavior.
3. Review whether the content includes a very large image or long block of text.

### If the content overflows
1. Check the spacing values and width constraints.
2. Make sure the card container has a proper `max-width` or responsive layout wrapper.
3. Reduce long text lines or break content into smaller cards.

### If the card appears too repetitive
1. Reuse the same component where the same pattern appears across pages.
2. Change only the content and heading, not the structure.
3. Keep the visual treatment consistent so the app feels intentional.
