# Images and local assets

Use local assets for images included in the source, and keep them organized in `frontend/src/assets/` or `frontend/public/` depending on how they are used. A component can import an image directly, while public assets are referenced by URL.

## [STEP 1.1.0] Import a local image into a component

### Goal
Add an image that is bundled with the frontend build.

### What this changes
This adds a branded or illustrative image to the UI without managing a custom asset pipeline.

### Files involved
- `frontend/src/assets/`
- `frontend/src/pages/ExamplePage.jsx`
- `frontend/src/index.css`

### Instructions
1. Place the image file in `frontend/src/assets/`.
2. Import it at the top of the component file.
3. Use the imported variable as the `src` value.
4. Add a useful `alt` attribute for accessibility.
5. Save the component and refresh the browser.

```jsx
import heroImage from '../assets/hero.png'

export default function ExamplePage() {
  return (
    <main className="page-shell">
      <section className="hero">
        <img className="hero-image" src={heroImage} alt="Illustration of a product dashboard" />
        <div>
          <p className="overline">WELCOME</p>
          <h1>Start building your product</h1>
        </div>
      </section>
    </main>
  )
}
```

```css
.hero {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1.5rem;
  align-items: center;
}

.hero-image {
  display: block;
  width: 100%;
  height: auto;
  border-radius: 16px;
}
```

### Do not
- Add a broken path or a file that is not actually inside the image folder.
- Use a blank `alt` value when the image is informative.

### Expected result
The image loads and is displayed correctly in the component.

---

## [STEP 1.1.1] Use a decorative image correctly

### Goal
Handle purely decorative imagery without adding unnecessary screen-reader noise.

### What this changes
This improves accessibility by letting assistive tech skip visual-only content.

### Files involved
- `frontend/src/pages/ExamplePage.jsx`

### Instructions
1. Open the component with the decorative image.
2. Set `alt=""` if the image is not informative.
3. Keep the image visible for design purposes but skip it for assistive technology.
4. Save and verify the page still reads clearly.

```jsx
<img src="/pattern.svg" alt="" className="decorative-pattern" />
```

### Do not
- Add a long alt description to purely decorative images.
- Use empty alt text when the image conveys real information.

### Expected result
The page remains accessible without noisy or repetitive screen-reader output.

---

## [STEP 1.1.2] Use a public asset when needed

### Goal
Serve an asset by URL from the public folder.

### What this changes
This makes a file available at a predictable public path, such as `/logo.svg`.

### Files involved
- `frontend/public/`
- `frontend/src/pages/ExamplePage.jsx`

### Instructions
1. Place the asset in `frontend/public/`.
2. Use a URL like `/logo.svg` in the component.
3. Use this pattern for app-wide brand assets or favicon-style files.
4. Save and verify the image loads from the browser.

```jsx
<img src="/logo.svg" alt="Project logo" />
```

### Do not
- Put a public asset in `src/assets` when it is meant to be served directly from the root.
- Hard-code a path without confirming the public asset exists.

### Expected result
The image loads reliably from the browser root and is easy to reference from multiple pages.

---

## Checkpoints & Troubleshooting

### Checkpoint
- [ ] The image file exists in the correct folder.
- [ ] The import path matches the actual file path.
- [ ] The `alt` text is useful and not missing when needed.
- [ ] The image does not overflow the page.
- [ ] The browser renders the expected asset.

### If the image does not load
1. Check the relative import path from the component file.
2. Confirm the file extension matches the import.
3. Verify the file exists in `frontend/src/assets/` or `frontend/public/`.

### If the image is too large
1. Add a width or `max-width` rule in `frontend/src/index.css`.
2. Set `height: auto` to preserve the image ratio.
3. Review the parent layout for overflow or fixed box sizes.

### If the image is decorative but still announced by screen readers
1. Confirm the `alt` attribute is the empty string `""`.
2. Remove any extraneous descriptive alt text from decorative images.
3. Keep informative images labeled clearly.
