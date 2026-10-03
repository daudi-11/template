# Modals and dialogs

Use a modal only for a focused task that must interrupt the current flow, such as confirmation, form entry, or a short decision. Keep the modal close to the trigger, ensure it is keyboard accessible, and return focus to the trigger when it closes.

## [STEP 1.1.0] Add a simple confirm modal

### Goal
Open a dialog for a single decision and keep the rest of the page dimmed or blocked until the user responds.

### What this changes
This changes the interaction pattern from an inline action to a focused confirmation dialog.

### Files involved
- `frontend/src/components/Modal.jsx`
- `frontend/src/pages/ExamplePage.jsx`
- `frontend/src/index.css`

### Instructions
1. Create a reusable modal component in `frontend/src/components/Modal.jsx`.
2. Accept `isOpen`, `title`, `onClose`, and `children` as props.
3. Render the modal only when `isOpen` is true.
4. Add an overlay and a close button.
5. Trap focus only if the modal is complex; for a simple confirmation modal, keep the implementation lightweight.
6. Add the modal trigger from the page that needs it.
7. Refresh the page and confirm the dialog opens and closes properly.

```jsx
import { useEffect } from 'react'

export default function Modal({ isOpen, title, onClose, children }) {
  useEffect(() => {
    if (!isOpen) return

    const handleKeyDown = (event) => {
      if (event.key === 'Escape') onClose()
    }

    window.addEventListener('keydown', handleKeyDown)
    return () => window.removeEventListener('keydown', handleKeyDown)
  }, [isOpen, onClose])

  if (!isOpen) return null

  return (
    <div className="modal-backdrop" onClick={onClose}>
      <div className="modal-card" role="dialog" aria-modal="true" aria-labelledby="modal-title" onClick={(event) => event.stopPropagation()}>
        <div className="modal-header">
          <h2 id="modal-title">{title}</h2>
          <button type="button" className="icon-button" onClick={onClose} aria-label="Close dialog">×</button>
        </div>
        <div className="modal-body">{children}</div>
      </div>
    </div>
  )
}
```

```css
.modal-backdrop {
  position: fixed;
  inset: 0;
  background: rgba(15, 23, 42, 0.55);
  display: grid;
  place-items: center;
  padding: 1rem;
  z-index: 1000;
}

.modal-card {
  width: min(100%, 480px);
  background: var(--paper);
  border: 1px solid var(--line);
  border-radius: 16px;
  box-shadow: var(--shadow-xl);
}

.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1rem 1.25rem;
  border-bottom: 1px solid var(--line);
}

.modal-body {
  padding: 1.25rem;
}
```

### Do not
- Open a modal for content that can live inline.
- Forget accessible labels on close buttons.
- Put a modal inside a form without a clear cancel/confirm flow.

### Expected result
The dialog opens with a clear title, has a close action, and can be dismissed without breaking page state.

---

## [STEP 1.1.1] Use the modal from a page

### Goal
Trigger the dialog from a page action and handle confirm/cancel behavior.

### What this changes
This connects the modal to real app data and lets a page open or close it intentionally.

### Files involved
- `frontend/src/pages/ExamplePage.jsx`
- `frontend/src/components/Modal.jsx`

### Instructions
1. Open the page where the action should occur.
2. Add `isModalOpen` state with `useState(false)`.
3. Add a button that sets the state to true.
4. Pass `onClose={() => setIsModalOpen(false)}` to the modal.
5. Put the actual action in the confirm button.
6. Save and test the flow.

```jsx
import { useState } from 'react'
import Modal from '../components/Modal.jsx'

export default function ExamplePage() {
  const [isOpen, setIsOpen] = useState(false)

  function handleConfirm() {
    setIsOpen(false)
    console.log('Confirmed')
  }

  return (
    <main className="page-shell">
      <button type="button" className="button button-primary" onClick={() => setIsOpen(true)}>
        Delete item
      </button>

      <Modal isOpen={isOpen} title="Confirm action" onClose={() => setIsOpen(false)}>
        <p>Are you sure you want to delete this item?</p>
        <div className="modal-actions">
          <button type="button" className="button button-secondary" onClick={() => setIsOpen(false)}>
            Cancel
          </button>
          <button type="button" className="button button-primary" onClick={handleConfirm}>
            Confirm
          </button>
        </div>
      </Modal>
    </main>
  )
}
```

### Do not
- Trigger the dialog from a hidden or disabled element without a clear user action.
- Use a modal for non-critical warnings that can remain visible in-place.

### Expected result
The page opens the confirmation dialog, confirms the action, and closes the modal when the user finishes.

---

## [STEP 1.1.2] Style modal actions

### Goal
Keep the modal visually aligned with the rest of the app and maintain a clear call-to-action hierarchy.

### What this changes
This keeps the modal consistent with the design system and the surrounding page.

### Files involved
- `frontend/src/index.css`

### Instructions
1. Open `frontend/src/index.css`.
2. Add styles for `.modal-actions` and any button alignment you want.
3. Keep the modal width and spacing consistent with the app's design tokens.
4. Use the same button classes as the rest of the project where possible.
5. Refresh and test on mobile screens.

```css
.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 0.75rem;
  margin-top: 1rem;
}

.icon-button {
  border: none;
  background: transparent;
  font-size: 1.5rem;
  line-height: 1;
  cursor: pointer;
  color: var(--ink-soft);
}
```

### Do not
- Add wildly different spacing from the rest of the UI.
- Place the action buttons in a single row on small screens without wrapping.

### Expected result
The modal looks intentional, readable, and consistent with the app's form and button system.

---

## Checkpoints & Troubleshooting

### Checkpoint
- [ ] The modal opens from the correct trigger.
- [ ] The modal closes on Escape or close button.
- [ ] The page behaves correctly after the modal closes.
- [ ] The modal is readable on mobile screens.
- [ ] The close button has an accessible label.

### If the modal does not open
1. Confirm the `isOpen` state is set to true before render.
2. Check that the modal component is imported correctly.
3. Confirm the trigger button is using `type="button"` to avoid accidental form submission.

### If the background page still scrolls
1. Check whether the modal is inside a page container that restricts overflow.
2. Add a fixed-position backdrop and set the body overflow if the page is meant to be locked.

### If the close button is not working
1. Check that `onClose` is passed and does not re-render incorrectly.
2. Confirm the click handler stops propagation when needed.
3. Test both button click and Escape key behavior.
