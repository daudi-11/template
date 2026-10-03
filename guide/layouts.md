# CSS Layout Guide

This guide shows how to build a clean page layout in the project using CSS and the existing React structure. Use it when you want to design a top navbar, a page shell, a form section, or a dashboard-style layout without rebuilding the whole app.

## [STEP 1.1.0] Start with a page shell

### Goal
Create a clean container that keeps page content organized and centered.

### What this changes
This sets the outer space, width, and spacing for every page.

### Files involved
- `frontend/src/index.css`
- `frontend/src/pages/HomePage.jsx`

### Instructions
1. Open `frontend/src/index.css`.
2. Add a shared container class for pages.
3. Use `max-width`, `margin`, and `padding` to control the page shape.
4. Apply the class to the main wrapper on a page.
5. Refresh the browser and check spacing.

```css
:root {
  --bg: #f8fafc;
  --paper: #ffffff;
  --ink: #0f172a;
  --muted: #475569;
  --line: #dfe7f1;
  --primary: #2563eb;
  --shadow: 0 12px 28px rgba(15, 23, 42, 0.08);
}

.page-shell {
  max-width: 1180px;
  margin: 0 auto;
  padding: 2rem 1rem 4rem;
}
```

```jsx
<main className="page-shell">
  <h1>Welcome</h1>
</main>
```

### Do not
- Put all layout styles directly on every page without a reusable class.
- Use very large margins that make the page feel broken on small screens.

### Expected result
Your pages have a predictable content area that stays centered and readable.

---

## [STEP 1.1.1] Design a navbar

### Goal
Create a top navigation bar that looks clean and works on desktop and mobile.

### What this changes
This controls branding, nav links, and the account area in the header.

### Files involved
- `frontend/src/components/Navbar.jsx`
- `frontend/src/index.css`

### Instructions
1. Open `frontend/src/components/Navbar.jsx`.
2. Check the existing header structure and keep the semantic order.
3. Add a `header` wrapper with a logo block and nav actions.
4. Use flexbox to align items left and right.
5. Add styles in `frontend/src/index.css` for `.topbar`, `.topbar-inner`, `.brand`, and `.main-nav`.
6. Test with guest and signed-in states.

```jsx
<header className="topbar">
  <div className="topbar-inner">
    <Link className="brand" to="/">HolyGrail</Link>

    <nav className="main-nav" aria-label="Main navigation">
      <Link className="nav-link" to="/">Home</Link>
      <Link className="nav-link" to="/login">Sign in</Link>
      <Link className="nav-join" to="/register">Create account</Link>
    </nav>
  </div>
</header>
```

```css
.topbar {
  background: var(--paper);
  border-bottom: 1px solid var(--line);
  position: sticky;
  top: 0;
  z-index: 20;
}

.topbar-inner {
  max-width: 1180px;
  margin: 0 auto;
  padding: 1rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
}

.main-nav {
  display: flex;
  align-items: center;
  gap: 0.8rem;
}

.brand {
  color: var(--ink);
  text-decoration: none;
  font-weight: 700;
}

.nav-link,
.nav-join {
  text-decoration: none;
  color: var(--ink);
  font-weight: 600;
}

.nav-join {
  background: var(--primary);
  color: white;
  padding: 0.7rem 1rem;
  border-radius: 10px;
}
```

### Do not
- Put the navigation links in random order without matching the app flow.
- Use a fixed width that breaks the layout on small screens.

### Expected result
The navbar stays clear, balanced, and consistent across the app.

---

## [STEP 1.1.2] Build a two-column content layout

### Goal
Create a section with text on one side and media or form content on the other.

### What this changes
This is the base pattern for landing pages, dashboards, and content-rich sections.

### Files involved
- `frontend/src/index.css`
- `frontend/src/pages/HomePage.jsx`

### Instructions
1. Open the page component.
2. Create a wrapper with two child columns.
3. Use CSS Grid or Flexbox to assign widths.
4. Keep the text column and media column aligned vertically.
5. Test on desktop and mobile.

```jsx
<section className="hero-grid">
  <div className="hero-copy">
    <p className="overline">A SIMPLE START</p>
    <h1>Build your product faster.</h1>
    <p>Use this layout to explain your app or idea.</p>
  </div>

  <div className="hero-panel">
    <div className="mini-card">Stats</div>
    <div className="mini-card">Forms</div>
  </div>
</section>
```

```css
.hero-grid {
  display: grid;
  grid-template-columns: 1.2fr 0.8fr;
  gap: 1.5rem;
  align-items: center;
}

.hero-panel {
  display: grid;
  gap: 1rem;
}

.mini-card {
  background: var(--paper);
  border: 1px solid var(--line);
  border-radius: 14px;
  padding: 1.25rem;
  box-shadow: var(--shadow);
}
```

### Do not
- Make both columns equal on wide screens when one side should dominate.
- Overuse grid if the layout only needs simple stacking.

### Expected result
The hero section or lead area reads nicely and keeps visual balance.

---

## [STEP 1.1.3] Layout cards and forms together

### Goal
Build a real page section that combines content cards and a form.

### What this changes
This pattern is common for landing pages, account pages, and app settings.

### Files involved
- `frontend/src/index.css`
- `frontend/src/pages/Login.jsx`
- `frontend/src/pages/Register.jsx`

### Instructions
1. Create a section that wraps cards and form content.
2. Use a grid with one large form column and one supporting details column.
3. Keep spacing consistent across both columns.
4. Match the shared classes used in the project.

```jsx
<section className="split-layout">
  <div className="info-column">
    <div className="info-card">Feature description</div>
    <div className="info-card">Support details</div>
  </div>

  <form className="auth-form">
    <label className="field">
      <span>Email</span>
      <input type="email" />
    </label>
    <button type="submit" className="button button-primary">Continue</button>
  </form>
</section>
```

```css
.split-layout {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1.5rem;
}

.info-column {
  display: grid;
  gap: 1rem;
}
```

### Do not
- Put the form and content blocks in random order without a visual hierarchy.
- Make the form wider than the available space on mobile.

### Expected result
The form and supporting content align into a structured, easy-to-scan section.

---

## [STEP 1.1.4] Make it responsive

### Goal
Ensure the layout keeps working on tablet and mobile screens.

### What this changes
This prevents overflow and keeps content readable on smaller devices.

### Files involved
- `frontend/src/index.css`

### Instructions
1. Open `frontend/src/index.css`.
2. Add media queries for small screens.
3. Change grid columns from two columns to one column when the screen is narrow.
4. Reduce padding and spacing where needed.
5. Refresh in browser responsive mode.

```css
@media (max-width: 760px) {
  .topbar-inner,
  .hero-grid,
  .split-layout {
    grid-template-columns: 1fr;
    flex-direction: column;
  }

  .main-nav {
    width: 100%;
    justify-content: space-between;
  }
}
```

### Do not
- Keep desktop-only widths on small screens.
- Hide important nav items without a good mobile alternative.

### Expected result
The layout flows naturally on mobile without broken spacing or hidden content.

---

## [STEP 1.1.5] Build a dashboard layout

### Goal
Create a workspace or dashboard with nav, content, and utility blocks.

### What this changes
This is a bigger layout pattern for app management screens or analytics pages.

### Files involved
- `frontend/src/index.css`
- `frontend/src/pages/ExampleDashboardPage.jsx`

### Instructions
1. Create a page wrapper with header and content layout.
2. Add a sidebar or summary panel for navigation or info.
3. Add a main content column for cards and tables.
4. Keep the layout using CSS Grid with a side panel and content panel.
5. Save and preview.

```jsx
<main className="dashboard-shell">
  <aside className="dashboard-sidebar">Menu</aside>
  <section className="dashboard-main">
    <div className="stat-card">Total users</div>
    <div className="stat-card">Weekly activity</div>
  </section>
</main>
```

```css
.dashboard-shell {
  display: grid;
  grid-template-columns: 220px 1fr;
  gap: 1.5rem;
  max-width: 1180px;
  margin: 0 auto;
  padding: 2rem 1rem;
}

.dashboard-sidebar,
.stat-card {
  background: var(--paper);
  border: 1px solid var(--line);
  border-radius: 14px;
  padding: 1rem;
}
```

### Do not
- Use dashboard sizing without thinking about mobile behavior.
- Make the sidebar too narrow for readable labels.

### Expected result
The dashboard layout feels organized, balanced, and easier to scale into more sections.

---

## Checkpoints & Troubleshooting

### Checkpoint
- [ ] The page shell is centered and readable.
- [ ] The navbar is aligned and easy to scan.
- [ ] Cards and forms sit in balanced sections.
- [ ] The layout collapses well on smaller screens.
- [ ] The page keeps consistent spacing and visual rhythm.

### If the page feels stretched or too wide
1. Check `max-width` and `padding` in `.page-shell`.
2. Confirm the grid columns are not too wide on smaller screens.
3. Add a media query if needed.

### If the navbar links are uneven
1. Review `display: flex`, `gap`, and alignment values.
2. Ensure the left and right groups share the same layout direction.
3. Check the spacing on the logo, nav, and actions.

### If the form or card section looks squashed
1. Confirm the grid columns are set correctly.
2. Reduce the number of columns on smaller screens.
3. Check whether nested cards or panels also need a `gap` value.

### If the mobile layout breaks
1. Add a media query at the correct breakpoint.
2. Change grid columns from 2 to 1 where needed.
3. Test the page with narrow widths before finalizing design changes.
