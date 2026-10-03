# Buttons and click events

Use a button for actions, and use a React Router link for navigation. Keep the event handler simple, avoid calling it during render, and give the button a clear accessible label.

## [STEP 1.1.0] Add a primary action button

### Goal
Create a visible action that triggers a real UI change.

### What this changes
This adds a working button to a page and connects it to a function.

### Files involved
- `frontend/src/pages/ExamplePage.jsx`
- `frontend/src/index.css`

### Instructions
1. Open the page where the action should appear.
2. Add a `button` element with `type="button"` when it is not submitting a form.
3. Add a `onClick` handler that changes state or triggers logic.
4. Use a label that matches the user action.
5. Save the file and refresh the page.

```jsx
import { useState } from 'react'

export default function ExamplePage() {
  const [message, setMessage] = useState('')

  function handleClick() {
    setMessage('The action was triggered.')
  }

  return (
    <main className="page-shell">
      <button type="button" className="button button-primary" onClick={handleClick}>
        Trigger action
      </button>
      {message && <p role="status">{message}</p>}
    </main>
  )
}
```

```css
.button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  border: none;
  border-radius: 10px;
  padding: 0.8rem 1.2rem;
  font-weight: 600;
  cursor: pointer;
}

.button-primary {
  background: var(--primary);
  color: white;
}
```

### Do not
- Call the handler during render with `onClick={handleClick()}`.
- Use a button for navigation if a React Router link is the correct pattern.

### Expected result
The button is clickable, visible, and triggers the intended state change.

---

## [STEP 1.1.1] Use a click handler correctly

### Goal
Keep event logic readable and predictable.

### What this changes
This controls when a handler runs and whether it receives data or changes state.

### Files involved
- `frontend/src/pages/ExamplePage.jsx`

### Instructions
1. Open the button component.
2. Define a separate function such as `handleSave()`.
3. Pass the function reference to `onClick`.
4. Use inline callbacks only for short one-off logic.
5. Test the result in the browser.

```jsx
function handleDelete() {
  console.log('Delete clicked')
}

<button type="button" onClick={handleDelete}>Delete</button>
```

```jsx
<button type="button" onClick={() => setOpen(true)}>Open</button>
```

### Do not
- Put long logic directly inside JSX.
- Run side effects from render time.

### Expected result
The event handler runs only when the user interacts with the button.

---

## [STEP 1.1.2] Use a button inside a form

### Goal
Handle form submission without accidentally triggering the wrong action.

### What this changes
This distinguishes a submit action from a non-submit action inside the same form.

### Files involved
- `frontend/src/pages/ExampleFormPage.jsx`
- `frontend/src/index.css`

### Instructions
1. Create a form with a submit button.
2. Use `type="submit"` on the main action button.
3. Use `type="button"` for secondary actions like Cancel or Reset.
4. Attach `onSubmit` to the form.
5. Save and test pressing Enter and the button.

```jsx
export default function ExampleFormPage() {
  function handleSubmit(event) {
    event.preventDefault()
    console.log('Submitted')
  }

  return (
    <form onSubmit={handleSubmit}>
      <label>
        <span>Email</span>
        <input type="email" />
      </label>
      <button type="submit" className="button button-primary">Submit</button>
      <button type="button" className="button button-secondary">Cancel</button>
    </form>
  )
}
```

### Do not
- Use `type="button"` on the main submit action.
- Attach `onClick` to the entire form when the form can submit natively.

### Expected result
The form works with both the button click and the Enter key.

---

## [STEP 1.1.3] Use navigation links correctly

### Goal
Navigate between routes without reloading the app.

### What this changes
This keeps in-app navigation smooth and consistent with the route setup.

### Files involved
- `frontend/src/components/Navbar.jsx`
- `frontend/src/App.jsx`

### Instructions
1. Open the component where navigation should appear.
2. Import `Link` from `react-router-dom`.
3. Use `to="/path"` with a URL matching the route.
4. Keep the destination consistent with the route path in `App.jsx`.
5. Save and test the link.

```jsx
import { Link } from 'react-router-dom'

<Link className="nav-link" to="/login">Sign in</Link>
```

### Do not
- Use a plain `<a href>` for internal routes when React Router is already in use.
- Link to a path that is not registered in `App.jsx`.

### Expected result
The user moves between screens without a full-page reload.

---

## Checkpoints & Troubleshooting

### Checkpoint
- [ ] Buttons are visible and labeled clearly.
- [ ] Click handlers run only when the button is activated.
- [ ] Form submit buttons do not accidentally behave as plain buttons.
- [ ] Navigation links point to valid routes.
- [ ] Buttons are testable with keyboard and mouse input.

### If the button does nothing
1. Confirm the handler is passed as a function reference, not as a function call.
2. Check that the component is re-rendering after state changes.
3. Look for a missing `type` attribute on a button inside a form.

### If the link does not change the page
1. Confirm the route exists in `frontend/src/App.jsx`.
2. Confirm the `Link` uses the correct path string.
3. Ensure the route is inside the router tree and the app is mounted correctly.

### If the button is hard to use on keyboard
1. Keep the button as a native `<button>` element.
2. Avoid replacing it with a `div` or clickable span.
3. Ensure the label remains visible and readable.
