# Tables and lists

Tables and lists help people scan structured information. Use a table when fields align in rows and columns. Use a list when the content is more narrative or grouped by category. Keep the layout readable on mobile by allowing content to wrap or stack when necessary.

## [STEP 1.1.0] Build a simple data table

### Goal
Display tabular data in a clean, readable format.

### What this changes
This creates a simple read-only data presentation for records, status items, or user lists.

### Files involved
- `frontend/src/pages/ExampleTablePage.jsx`
- `frontend/src/index.css`

### Instructions
1. Open or create the page that should show a table.
2. Create a small array of data objects in the component.
3. Map over the array to render rows.
4. Add a table header for each column.
5. Keep the table readable with short labels and consistent data formatting.
6. Save and review the layout.

```jsx
const rows = [
  { name: 'Project A', owner: 'Ava', status: 'Active' },
  { name: 'Project B', owner: 'Noah', status: 'Draft' },
  { name: 'Project C', owner: 'Mila', status: 'Archived' },
]

export default function ExampleTablePage() {
  return (
    <main className="page-shell">
      <h1>Projects</h1>
      <table className="data-table">
        <thead>
          <tr>
            <th>Name</th>
            <th>Owner</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          {rows.map((row) => (
            <tr key={row.name}>
              <td>{row.name}</td>
              <td>{row.owner}</td>
              <td>{row.status}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </main>
  )
}
```

```css
.data-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 1rem;
  background: var(--paper);
  border: 1px solid var(--line);
  border-radius: 12px;
  overflow: hidden;
}

.data-table th,
.data-table td {
  padding: 0.85rem 1rem;
  border-bottom: 1px solid var(--line);
  text-align: left;
}

.data-table th {
  background: var(--muted);
}
```

### Do not
- Use a table for content that is not truly tabular.
- Put long text in a narrow cell without wrapping.
- Forget to include a unique key when mapping arrays.

### Expected result
The table is readable, consistent, and easy to scan.

---

## [STEP 1.1.1] Build a vertical list layout

### Goal
Display grouped items without a table, especially when the content is a list of cards or actions.

### What this changes
This creates a more flexible layout for a stack of items that may include actions, tags, or metadata.

### Files involved
- `frontend/src/pages/ExampleListPage.jsx`
- `frontend/src/index.css`

### Instructions
1. Create a list of objects with a title and optional metadata.
2. Use `.map()` to render a list item for each object.
3. Keep the structure semantic with `<ul>` and `<li>`.
4. Add a small action button or link if needed.
5. Save and verify the layout.

```jsx
const items = [
  { id: 1, title: 'Account settings', meta: 'Security' },
  { id: 2, title: 'Notifications', meta: 'Email' },
  { id: 3, title: 'Billing', meta: 'Invoices' },
]

export default function ExampleListPage() {
  return (
    <main className="page-shell">
      <h1>Preferences</h1>
      <ul className="stack-list">
        {items.map((item) => (
          <li key={item.id} className="stack-item">
            <div>
              <strong>{item.title}</strong>
              <p>{item.meta}</p>
            </div>
            <button type="button" className="button button-secondary">Open</button>
          </li>
        ))}
      </ul>
    </main>
  )
}
```

```css
.stack-list {
  list-style: none;
  padding: 0;
  margin: 1rem 0 0;
  display: grid;
  gap: 0.75rem;
}

.stack-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  padding: 1rem;
  border: 1px solid var(--line);
  border-radius: 12px;
  background: var(--paper);
}
```

### Do not
- Use a list when a table would be clearer.
- Add non-semantic `div` wrappers when a `<ul>` or `<li>` is the natural structure.

### Expected result
The item list is easy to scan and visually consistent with the rest of the UI.

---

## [STEP 1.1.2] Make tables mobile-friendly

### Goal
Keep data readable on narrow screens without breaking the layout.

### What this changes
This improves the mobile experience so tables and list layouts do not overflow or become unreadable.

### Files involved
- `frontend/src/index.css`

### Instructions
1. Open `frontend/src/index.css`.
2. Add mobile-specific rules for tables and lists.
3. Reduce padding and allow cells to wrap when needed.
4. Test on a narrow viewport.

```css
@media (max-width: 640px) {
  .data-table th,
  .data-table td {
    padding: 0.7rem 0.5rem;
    font-size: 0.95rem;
  }

  .stack-item {
    flex-direction: column;
    align-items: flex-start;
  }
}
```

### Do not
- Force a horizontal scroll on phones unless the content truly requires it.
- Reduce font size so much that the text becomes hard to read.

### Expected result
The table or list remains usable and readable on both desktop and mobile screens.

---

## Checkpoints & Troubleshooting

### Checkpoint
- [ ] The table headers match the row content.
- [ ] Each row has a unique key when mapped.
- [ ] Lists remain readable on small screens.
- [ ] Buttons inside lists are easy to target.
- [ ] Long values wrap or truncate cleanly.

### If the table looks broken
1. Confirm the `<thead>` and `<tbody>` structure matches the markup.
2. Check the CSS for collapsed borders or missing `border-collapse`.
3. Review the `key` values in the array map.

### If the list runs together
1. Check spacing in `.stack-item` and `.stack-list`.
2. Confirm the item container has a `display` and gap value.
3. Ensure the list is not missing `list-style: none` if you want a clean card style.

### If text overflows on small screens
1. Reduce padding or switch to column layout.
2. Add wrapping rules to table cells or long text wrappers.
3. Test with realistic content lengths before shipping.
