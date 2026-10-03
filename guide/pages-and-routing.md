# Pages and routing

Put full-page views in `frontend/src/pages/` and register them in `frontend/src/App.jsx`. Use React Router `Link` components for in-app navigation so the page does not reload the whole app.

## [STEP 1.1.0] Add a new page

### Goal
Create a new screen that fits the app's existing page structure.

### What this changes
This adds a new route and a dedicated view for a feature or simple section.

### Files involved
- `frontend/src/pages/ExamplePage.jsx`
- `frontend/src/App.jsx`
- `frontend/src/index.css`

### Instructions
1. Create a new file in `frontend/src/pages/`.
2. Export a default function component.
3. Return a `<main>` wrapper with a heading and content.
4. Keep the file focused on one page topic.
5. Save the component.

```jsx
export default function ExamplePage() {
  return (
    <main className="page-shell">
      <section className="page-card">
        <p className="overline">EXAMPLE</p>
        <h1>Example page</h1>
        <p>Use this page to stand up a new feature or section.</p>
      </section>
    </main>
  )
}
```

```css
.page-shell {
  max-width: 1100px;
  margin: 0 auto;
  padding: 2rem 1rem 4rem;
}

.page-card {
  background: var(--paper);
  border: 1px solid var(--line);
  border-radius: 16px;
  padding: 2rem;
  box-shadow: var(--shadow);
}
```

### Do not
- Put unrelated logic into a page component.
- Route directly to a page that is not imported into `App.jsx`.

### Expected result
A new page component exists and is ready to be attached to the router.

---

## [STEP 1.1.1] Register the route

### Goal
Attach the page to a URL path.

### What this changes
This makes the new page accessible by browser navigation.

### Files involved
- `frontend/src/App.jsx`

### Instructions
1. Open `frontend/src/App.jsx`.
2. Import the new page component.
3. Add a `<Route path="/example" element={<ExamplePage />} />` entry.
4. Keep the route next to other public or protected routes.
5. Save and test the URL in the browser.

```jsx
import ExamplePage from './pages/ExamplePage.jsx'

<Route path="/example" element={<ExamplePage />} />
```

### Do not
- Add duplicate route paths.
- Remove the fallback route that redirects unknown paths home.

### Expected result
The page loads at its route and can be accessed from the browser.

---

## [STEP 1.1.2] Add a navigation link

### Goal
Make the page reachable from the shared navigation.

### What this changes
This makes the page discoverable for users without manually typing the URL.

### Files involved
- `frontend/src/components/Navbar.jsx`
- `frontend/src/App.jsx`

### Instructions
1. Open `frontend/src/components/Navbar.jsx`.
2. Add a `Link` component with a `to` path matching your route.
3. Keep the label short and clear.
4. Save the file.
5. Click the link in the browser to confirm the route works.

```jsx
import { Link } from 'react-router-dom'

<Link className="nav-link" to="/example">Example</Link>
```

### Do not
- Point the link to a route that is not registered.
- Place the link in the wrong navigation section for public vs. signed-in pages.

### Expected result
Users can reach the new page from the existing app navigation.

---

## [STEP 1.1.3] Protect a page with auth

### Goal
Require a signed-in user before a page can render.

### What this changes
This restricts access to private content and redirects guests to login.

### Files involved
- `frontend/src/components/ProtectedRoute.jsx`
- `frontend/src/App.jsx`

### Instructions
1. Open `frontend/src/App.jsx`.
2. Wrap the private route in the same pattern used by the existing protected routes.
3. Keep the `Navigate` redirect consistent with the login flow.
4. Test as both a guest and an authenticated user.

```jsx
<Route element={<ProtectedRoute />}>
  <Route path="/dashboard" element={<DashboardPage />} />
</Route>
```

### Do not
- Protect a public page by mistake.
- Remove the sign-in redirect flow without testing guest access.

### Expected result
Only signed-in users can open the protected page.

---

## Checkpoints & Troubleshooting

### Checkpoint
- [ ] The page file exists and imports correctly.
- [ ] The route path matches the navigation link destination.
- [ ] The route works in the browser.
- [ ] Private pages redirect guests as expected.
- [ ] The page style is consistent with other screens.

### If the page does not load
1. Confirm the page is imported correctly in `App.jsx`.
2. Check the `path` value and `Link` `to` value match exactly.
3. Confirm the component is exported as default.

### If the page loads but styles look wrong
1. Confirm the CSS class names exist in `frontend/src/index.css`.
2. Check whether a shared rule is overriding your page-specific styles.
3. Inspect the browser DevTools to ensure the right selector is active.

### If guests can still access a protected page
1. Confirm the route is nested inside `ProtectedRoute`.
2. Review the `isAuthenticated` state in the auth context.
3. Ensure the redirect target is set correctly.
