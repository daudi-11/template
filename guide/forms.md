# Forms and submit events

Use a form when the user needs to send information to the app. The form should use a real `onSubmit`, real labels, and accessible validation. Keep the submission flow predictable and test it with both mouse and keyboard input.

## [STEP 1.1.0] Build a basic form

### Goal
Create a simple input form with a submit action.

### What this changes
This adds the page structure, form fields, and submit logic needed for a real data entry flow.

### Files involved
- `frontend/src/pages/ExampleFormPage.jsx`
- `frontend/src/index.css`

### Instructions
1. Open the page you want to use for the form.
2. Add a `<form>` element and `onSubmit` callback.
3. Add a `<label>` and input for each field.
4. Use `value` and `onChange` for controlled inputs.
5. Add a submit button.
6. Save and test the form with the browser.

```jsx
import { useState } from 'react'

export default function ExampleFormPage() {
  const [email, setEmail] = useState('')
  const [status, setStatus] = useState('')

  function handleSubmit(event) {
    event.preventDefault()
    setStatus(`Thanks. We received ${email}.`)
  }

  return (
    <main className="page-shell">
      <form className="auth-form" onSubmit={handleSubmit}>
        <label className="field">
          <span>Email</span>
          <input
            type="email"
            name="email"
            value={email}
            onChange={(event) => setEmail(event.target.value)}
            autoComplete="email"
            required
          />
        </label>

        <button type="submit" className="button button-primary">Send</button>
      </form>

      {status && <p role="status">{status}</p>}
    </main>
  )
}
```

```css
.auth-form {
  display: grid;
  gap: 1rem;
  max-width: 480px;
}

.field {
  display: grid;
  gap: 0.5rem;
}

.field input {
  padding: 0.8rem 0.9rem;
  border: 1px solid var(--line);
  border-radius: 10px;
  font: inherit;
}
```

### Do not
- Forget `event.preventDefault()` in the submit handler.
- Use placeholder-only labels without a real `<label>`.

### Expected result
The form collects the input and submits it without a hard page reload.

---

## [STEP 1.1.1] Validate a form before submission

### Goal
Block bad input early and give the user immediate feedback.

### What this changes
This improves the form quality by using browser validation and clear error handling.

### Files involved
- `frontend/src/pages/ExampleFormPage.jsx`

### Instructions
1. Add `required` to inputs when necessary.
2. Use matching HTML input types such as `email`, `password`, or `number`.
3. Add a validation error state if the backend or the form logic fails.
4. Disable the submit button when a request is already in progress.
5. Test the flow with valid and invalid input.

```jsx
import { useState } from 'react'

export default function ExampleFormPage() {
  const [form, setForm] = useState({ email: '' })
  const [error, setError] = useState('')
  const [busy, setBusy] = useState(false)

  async function handleSubmit(event) {
    event.preventDefault()
    setError('')
    setBusy(true)

    try {
      if (!form.email) {
        throw new Error('Email is required.')
      }
      console.log('Submitted', form)
    } catch (requestError) {
      setError(requestError.message)
    } finally {
      setBusy(false)
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <label>
        <span>Email</span>
        <input
          type="email"
          value={form.email}
          onChange={(event) => setForm({ ...form, email: event.target.value })}
          required
        />
      </label>
      {error && <p role="alert">{error}</p>}
      <button type="submit" disabled={busy}> {busy ? 'Sending...' : 'Send'} </button>
    </form>
  )
}
```

### Do not
- Rely only on vague messages like "There was an error." without context.
- Remove `required` when the field must be filled.

### Expected result
The form blocks invalid input and communicates clear errors.

---

## [STEP 1.1.2] Match the project form styles

### Goal
Keep forms consistent with the app's design system.

### What this changes
This aligns your form styling with the rest of the template without changing the behavior.

### Files involved
- `frontend/src/index.css`
- `frontend/src/pages/Login.jsx`
- `frontend/src/pages/Register.jsx`

### Instructions
1. Open the CSS file that styles form elements.
2. Match the spacing, field border radius, and button styles used elsewhere.
3. Reuse `.field`, `.auth-form`, and `.button` where possible.
4. Keep the same look across the login and register screens.
5. Save and check the layout.

### Do not
- Create a one-off style for every form without reusing shared classes.
- Add wildly different spacing across similar forms.

### Expected result
The form feels like part of the same application rather than a separate pattern.

---

## Checkpoints & Troubleshooting

### Checkpoint
- [ ] Every field has a real label.
- [ ] The form submits with Enter and button click.
- [ ] Validation catches missing or invalid values.
- [ ] The submit button is disabled during a request.
- [ ] Error messages are clear and visible.

### If the form does not submit
1. Confirm `onSubmit` is attached to the `<form>` element.
2. Check that `event.preventDefault()` is called when needed.
3. Ensure the submit button uses `type="submit"`.

### If the field value does not update
1. Confirm the input is controlled with `value` and `onChange`.
2. Check the state update function keeps the other fields intact.
3. Inspect whether the form resets unexpectedly when the page re-renders.

### If the button is still clickable during loading
1. Add a `busy` state and disable the button while the request is pending.
2. Verify the state changes back after the request resolves.
3. Confirm the submit branch sets the state in a `finally` block.
