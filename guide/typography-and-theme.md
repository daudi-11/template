# Typography and theme

This guide covers the visual foundation of the UI: fonts, type scale, color, spacing, and design tokens. Small changes here affect every page, so make them deliberately and keep them consistent.

## [STEP 1.1.0] Define the theme variables

### Goal
Set a small design token system that all components can reuse.

### What this changes
This creates a consistent visual foundation instead of hard-coding colors and spacing on individual elements.

### Files involved
- `frontend/src/index.css`

### Instructions
1. Open `frontend/src/index.css`.
2. Find the root style block or theme variables.
3. Define color and spacing tokens for your app.
4. Reuse those variables across the rest of the stylesheet.
5. Refresh the app and confirm the design is consistent.

```css
:root {
  --bg: #f5f7fb;
  --paper: #ffffff;
  --ink: #0f172a;
  --ink-soft: #475569;
  --line: #dfe7f1;
  --primary: #2563eb;
  --primary-strong: #1d4ed8;
  --muted: #f8fafc;
  --success: #16a34a;
  --warning: #f59e0b;
  --danger: #dc2626;
  --radius: 12px;
  --shadow: 0 8px 24px rgba(15, 23, 42, 0.08);
}
```

### Do not
- Set raw colors in many component blocks without a shared token.
- Mix multiple unrelated color scales across the app.

### Expected result
The interface has a consistent starting palette and spacing system.

---

## [STEP 1.1.1] Set the base typography

### Goal
Define the fonts, text sizes, and text color used across the template.

### What this changes
This controls the default reading experience for the app.

### Files involved
- `frontend/src/index.css`

### Instructions
1. Open the CSS root styles.
2. Set the default font family and base text size.
3. Define headline and body styles.
4. Keep the text hierarchy readable and minimal.

```css
html {
  font-family: Inter, 'Segoe UI', sans-serif;
  color: var(--ink);
  background: var(--bg);
}

body {
  margin: 0;
  line-height: 1.6;
  font-size: 16px;
}

h1, h2, h3, h4 {
  margin: 0 0 0.6rem;
  color: var(--ink);
  line-height: 1.2;
}

p {
  margin: 0 0 1rem;
  color: var(--ink-soft);
}
```

### Do not
- Use too many font families.
- Make headings too large without a clear hierarchy.

### Expected result
Headings and paragraphs have a consistent text scale and a clean reading rhythm.

---

## [STEP 1.1.2] Update buttons, links, and labels to the theme

### Goal
Make common UI elements follow the app's new design system.

### What this changes
This updates the styling of buttons, links, and form labels without changing their logic.

### Files involved
- `frontend/src/index.css`

### Instructions
1. Open the shared button and link styles.
2. Replace color values with theme variables.
3. Keep states such as hover and focus consistent.
4. Update labels and text colors where needed.

```css
.button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  padding: 0.75rem 1.1rem;
  border-radius: 10px;
  border: 1px solid transparent;
  font-weight: 600;
  transition: 0.2s ease;
}

.button-primary {
  background: var(--primary);
  color: #fff;
}

.button-primary:hover {
  background: var(--primary-strong);
}

.link {
  color: var(--primary);
  text-decoration: none;
}
```

### Do not
- Add bright colors to every button without an intentional hierarchy.
- Remove focus styles, which affects keyboard accessibility.

### Expected result
All core controls use the same color and interaction model.

---

## [STEP 1.1.3] Tune spacing and layout rhythm

### Goal
Keep the interface consistent across cards, pages, lists, and forms.

### What this changes
This controls whitespace and layout density across the UI.

### Files involved
- `frontend/src/index.css`

### Instructions
1. Open the shared spacing styles.
2. Review the padding and margin values used across major sections.
3. Replace repeated pixel values with design tokens where possible.
4. Check the page layout for consistent vertical rhythm.

```css
.page-shell {
  max-width: 1100px;
  margin: 0 auto;
  padding: 2rem 1rem 3rem;
}

.section-gap {
  margin-top: 2rem;
}
```

### Do not
- Use random spacing values for each new section.
- Overcrowd the app with too little whitespace.

### Expected result
The layout feels intentional, balanced, and easier to scan.

---

## Checkpoints & Troubleshooting

### Checkpoint
- [ ] The color palette is consistent across page sections.
- [ ] Buttons and links follow the same theme.
- [ ] Headings have a clear hierarchy.
- [ ] Text remains readable on light and dark surfaces.
- [ ] Spacing feels even across the app.

### If the theme feels inconsistent
1. Check whether styles are using raw color values instead of theme variables.
2. Review whether some pages use different font or spacing defaults.
3. Confirm whether a component overrides the base rule unintentionally.

### If button styles feel off
1. Confirm the CSS class names match the component markup.
2. Check whether a more specific selector is overriding the shared button class.
3. Inspect hover and focus states to keep accessibility intact.

### If text is hard to read
1. Confirm `--ink` and `--ink-soft` are not too close in contrast.
2. Review text on colored cards or buttons for contrast issues.
3. Test on a few different components before publishing.
