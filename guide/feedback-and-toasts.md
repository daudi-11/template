# Feedback and toasts

Use feedback patterns for status updates, success messages, validation errors, and short actions. A toast is best for a message that should appear briefly without interrupting the user's flow. A banner or alert is better for persistent information on the page.

## [STEP 1.1.0] Add a simple success/error message

### Goal
Tell the user what happened after a form or action completes.

### What this changes
This adds a visible response when the user completes or fails an action.

### Files involved
- `frontend/src/pages/ExampleFormPage.jsx`
- `frontend/src/index.css`

### Instructions
1. Open the page that triggers the action.
2. Add `status` and `error` state with `useState()`.
3. Set the message after a successful action or a failed one.
4. Render a message area with a class that matches the status type.
5. Save and verify the message is shown clearly.

```jsx
import { useState } from 'react'

export default function ExampleFormPage() {
  const [status, setStatus] = useState('')
  const [error, setError] = useState('')

  function handleSubmit() {
    setError('')
    setStatus('Saved successfully.')
  }

  return (
    <main className="page-shell">
      <button type="button" className="button button-primary" onClick={handleSubmit}>
        Save
      </button>

      {status && <p className="status-message success">{status}</p>}
      {error && <p className="status-message error">{error}</p>}
    </main>
  )
}
```

```css
.status-message {
  margin-top: 1rem;
  padding: 0.75rem 1rem;
  border-radius: 10px;
  font-weight: 600;
}

.status-message.success {
  background: rgba(22, 163, 74, 0.12);
  color: #166534;
}

.status-message.error {
  background: rgba(220, 38, 38, 0.1);
  color: #991b1b;
}
```

### Do not
- Hide the message immediately after the action without giving the user time to read it.
- Use a success state for errors or vice versa.

### Expected result
The user sees clear feedback after the action completes or fails.

---

## [STEP 1.1.1] Build a small toast component

### Goal
Show temporary feedback without blocking the page.

### What this changes
This adds a short-lived notification pattern for quick actions like save, copy, or remove.

### Files involved
- `frontend/src/components/Toast.jsx`
- `frontend/src/index.css`

### Instructions
1. Create a simple toast component in `frontend/src/components/Toast.jsx`.
2. Accept `message`, `type`, and an optional `onClose`.
3. Render the toast in a fixed position near the top right of the screen.
4. Use a short timeout to clear it automatically.
5. Trigger it from the action that needs feedback.

```jsx
import { useEffect } from 'react'

export default function Toast({ message, type = 'success', onClose }) {
  useEffect(() => {
    if (!message) return

    const timer = setTimeout(() => onClose?.(), 3000)
    return () => clearTimeout(timer)
  }, [message, onClose])

  if (!message) return null

  return (
    <div className={`toast toast-${type}`} role="status" aria-live="polite">
      {message}
    </div>
  )
}
```

```css
.toast {
  position: fixed;
  right: 1rem;
  top: 1rem;
  min-width: 220px;
  max-width: 320px;
  padding: 0.8rem 1rem;
  border-radius: 12px;
  box-shadow: var(--shadow);
  z-index: 1200;
}

.toast-success {
  background: #166534;
  color: white;
}

.toast-error {
  background: #991b1b;
  color: white;
}
```

### Do not
- Use a toast for long-form validation that needs to stay visible.
- Create multiple simultaneous toasts without a queue.

### Expected result
Users receive immediate but non-blocking feedback for small actions.

---

## [STEP 1.1.2] Use feedback patterns consistently

### Goal
Keep the app's feedback system coherent across success, error, and warning states.

### What this changes
This standardizes how the app communicates updates and failure states.

### Files involved
- `frontend/src/pages/`
- `frontend/src/components/`
- `frontend/src/index.css`

### Instructions
1. Pick one pattern for each case: inline message, banner, or toast.
2. Use inline messages for form validation.
3. Use toast messages for short successful actions.
4. Use alerts or modals for critical or destructive actions.
5. Keep the wording short and clear.

### Do not
- Mix message types without a clear rule.
- Use vague text such as "Something went wrong" when the user needs a specific action.

### Expected result
Users understand what happened and what to do next without confusion.

---

## Checkpoints & Troubleshooting

### Checkpoint
- [ ] Success and error states have distinct visual styles.
- [ ] Toast messages disappear after the expected time.
- [ ] Feedback is visible without blocking the main interaction.
- [ ] Messages are readable on both light and dark backgrounds.
- [ ] Error text tells the user what to do next.

### If the message never appears
1. Confirm the state is set before the component renders.
2. Check whether the correct conditional is used in JSX.
3. Confirm the element is not hidden behind a CSS rule.

### If the toast does not disappear
1. Check the timeout code and ensure it is cleaned up correctly.
2. Confirm the `onClose` callback is passed and updates the state.
3. Ensure the component is not recreated in a loop.

### If the message text is unreadable
1. Increase contrast between message background and text.
2. Confirm the color classes match the intended status.
3. Keep the message length short enough to remain scannable.
