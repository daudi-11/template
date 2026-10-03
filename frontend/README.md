# Frontend

This React app contains the guest home page, registration and login forms, and the signed-in account view.

## Run locally

```powershell
npm install
npm run dev
```

The app calls `http://localhost:3030` by default. Set `VITE_API_URL` to override the backend URL.

## Extend the app

- Add views in `src/pages/` and reusable UI in `src/components/`.
- Register routes in `src/App.jsx`.
- Use `src/context/` for shared auth state and `src/services/api.js` for backend requests.
- Follow `[STEP 4.X]` comments for the auth provider, home states, and protected-route extension point.

See the [root developer guide](../README.md) for backend setup, auth endpoints, and the full step index.
