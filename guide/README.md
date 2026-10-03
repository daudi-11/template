# UI Guides

These guides show how to extend the existing React frontend without replacing its structure. Examples use React function components, React Router, and the global styles in `frontend/src/index.css`.

## Choose a guide

- [Add a button and handle its click event](buttons-and-events.md)
- [Create a reusable card](cards.md)
- [Build a form and handle submission](forms.md)
- [Add an accessible icon](icons.md)
- [Add an image](images.md)
- [Add a page and route](pages-and-routes.md)
- [Create a modal and dialog flow](modals-and-dialogs.md)
- [Build a data table or list](tables-and-lists.md)
- [Set typography and shared theme tokens](typography-and-theme.md)
- [Show feedback and toast messages](feedback-and-toasts.md)
- [Learn CSS layout patterns](layouts.md)
- [Add a page and route](pages-and-routing.md)

### A simple workflow

1. Find the closest existing component or page and follow its conventions.
2. Add the smallest UI change that achieves the task.
3. Put reusable styles in `frontend/src/index.css`; use a clear class name rather than styling unrelated elements.
4. Check keyboard use, labels, and small-screen layout as you build.

The snippets are patterns to adapt. They are not new application features until you add them to a page and connect any required state or API behavior.
