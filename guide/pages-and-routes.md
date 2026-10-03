# Pages and routes

Use pages for full-screen views and use the router to connect them to URLs. Keep the page component simple, attach the route in `frontend/src/App.jsx`, and add navigation only when the page should be discoverable from the shared UI.

## [STEP 1.1.0] Add a new page

### Goal
Create a standalone screen that matches the rest of the app's structure.

### What this changes
This adds the new view and the page-level content needed for a feature or section.

### Files involved
- `frontend/src/pages/ExamplePage.jsx`
- `frontend/src/index.css`

### Instructions
1. Create a new file inside `frontend/src/pages/`.
2. Export a default component.
3. Return a `<main>` element with the page wrapper and heading.
4. Keep page-specific logic inside this page and not scattered across unrelated files.
5. Save the file.

```jsx
export default function ExamplePage() {
  return (
    <main className="page-shell">
      <section className="page-card">
        <p className="overline">EXAMPLE</p>
        <h1>Example page</h1>
        <p>Use this page for a product feature, landing section, or support flow.</p>
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
- Put large unrelated utilities into a single page file.
- Duplicate the same layout in several pages without extracting a pattern.

### Expected result
The new page exists as a clean, self-contained screen.

---

## [STEP 1.1.1] Register the route

### Goal
Attach the page to a URL so it is accessible in the app.

### What this changes
This transforms the new file into a real route that the browser can open.

### Files involved
- `frontend/src/App.jsx`

### Instructions
1. Open `frontend/src/App.jsx`.
2. Import the new page component.
3. Add a route entry with a unique path.
4. Keep the route in the same public or protected section as the rest of the app.
5. Save and test the URL.

```jsx
import ExamplePage from './pages/ExamplePage.jsx'

<Route path="/example" element={<ExamplePage />} />
```

### Do not
- Add duplicate paths for the same page.
- Remove the route fallback for unknown pages.

### Expected result
The page can be visited in the browser at its route.

---

## [STEP 1.1.2] Add navigation

### Goal
Expose the page from the app shell so users can discover it.

### What this changes
This makes the route easier to reach without typing the URL manually.

### Files involved
- `frontend/src/components/Navbar.jsx`
- `frontend/src/App.jsx`

### Instructions
1. Open the navigation component.
2. Add a `Link` with the correct `to` value.
3. Use a short, clear label.
4. Save and verify the link works from the browser.

```jsx
import { Link } from 'react-router-dom'

<Link className="nav-link" to="/example">Example</Link>
```

### Do not
- Link to a page that is not registered.
- Put page links in the wrong section if the page is protected.

### Expected result
The page is discoverable and reachable from the app's primary navigation.

---

## [STEP 1.1.3] Protect a route when needed

### Goal
Restrict access to signed-in users for private pages.

### What this changes
This ensures the page is only visible to the correct audience.

### Files involved
- `frontend/src/App.jsx`
- `frontend/src/components/ProtectedRoute.jsx`

### Instructions
1. Open `frontend/src/App.jsx`.
2. Wrap the protected route in the same `ProtectedRoute` pattern used elsewhere.
3. Keep the redirect consistent with the login flow.
4. Test as both a guest and an authenticated user.

```jsx
<Route element={<ProtectedRoute />}>
  <Route path="/dashboard" element={<DashboardPage />} />
</Route>
```

### Do not
- Protect a page that should be public.
- Break the guest redirect logic when changing route structure.

### Expected result
The page is only visible to the correct user state.

---

## Checkpoints & Troubleshooting

### Checkpoint
- [ ] The page file loads without errors.
- [ ] The route path matches the navigation target.
- [ ] The page renders in the browser.
- [ ] Protected pages redirect visitors correctly.
- [ ] The page style matches the rest of the app.

### If the route does not open
1. Confirm the imported component name matches the file export.
2. Check the route `path` and the `Link` `to` value for exact matches.
3. Open the browser console to check for missing imports or runtime errors.

### If the page is blank
1. Review the JSX returned by the component.
2. Confirm the main wrapper and child elements are correctly closed.
3. Check for runtime errors in the console.

### If a protected page is still visible to guests
1. Confirm the route is nested inside `ProtectedRoute`.
2. Check the auth hook state in the app context.
3. Test the redirect target by signing out and returning to the route.
