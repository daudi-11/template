# HolyGrail Customization Manual

This guide is a practical customization roadmap for the HolyGrail template. It is organized as a step-by-step build flow, not as a file-reference summary.

Use it in order. Do not jump ahead unless a later step explicitly says it is safe.

## Master Roadmap

1.0.0 Project Setup
- 1.1.0 Prepare the development environment
- 1.1.1 Install prerequisites
- 1.1.2 Create the backend environment file
- 1.1.3 Start the backend
- 1.1.4 Start the frontend
- 1.1.5 Verify the app works

2.0.0 Understand the Template
- 2.1.0 Understand the frontend structure
- 2.1.1 Review the app shell and routes
- 2.1.2 Review the auth state
- 2.1.3 Review page-level UI files
- 2.2.0 Understand the backend structure
- 2.2.1 Review routes
- 2.2.2 Review controllers
- 2.2.3 Review config and environment handling
- 2.2.4 Review auth flow

3.0.0 Customize the Frontend
- 3.1.0 Customize the global layout
- 3.1.1 Edit the Navbar
- 3.1.2 Edit the app shell
- 3.1.3 Adjust the global page structure
- 3.2.0 Customize global styling
- 3.2.1 Change color palette
- 3.2.2 Change buttons and links
- 3.2.3 Change cards and spacing
- 3.2.4 Adjust mobile layout
- 3.3.0 Customize existing pages
- 3.3.1 Customize the Home page
- 3.3.2 Customize the Login page
- 3.3.3 Customize the Register page
- 3.4.0 Create a new page
- 3.5.0 Create reusable components

4.0.0 Customize Authentication
- 4.1.0 Customize registration
- 4.2.0 Customize login
- 4.3.0 Customize logout
- 4.4.0 Customize protected routes

5.0.0 Customize the Backend
- 5.1.0 Customize API routes
- 5.2.0 Customize controllers
- 5.3.0 Customize auth data and session behavior
- 5.4.0 Update environment and DB settings

6.0.0 Connect Frontend and Backend
- 6.1.0 Use the shared API helper
- 6.2.0 Create a complete feature end-to-end

7.0.0 Test and Prepare for Production
- 7.1.0 Test frontend behavior
- 7.2.0 Test auth flow
- 7.3.0 Test backend APIs
- 7.4.0 Build and verify production readiness

Start here: [STEP 1.1.0]

---

## [STEP 1.1.0] Prepare the Project

Current area: 1.0 Project Setup
Current step: 1.1.0
Next: 1.1.1 Install prerequisites

Goal:
Set up the local environment so both the Go API and the React frontend can run together.

Files involved:
- `backend/.env.example`
- `backend/config/config.go`
- `frontend/package.json`

What to do:
1. Install Go, Node.js, npm, and PostgreSQL.
2. Make sure the app can use PostgreSQL locally.
3. Create a copy of `backend/.env.example` and rename it to `backend/.env`.
4. Fill in the values required by the backend:
   - `AUTH_DB_URL` should point to your PostgreSQL instance.
   - `JWT_SECRET` must be at least 32 characters long.
5. Keep the file local and do not commit it to source control.

What this changes:
The app can connect to the database and sign JWT tokens correctly.

Do not:
- Use a short or shared JWT secret.
- Leave the database URL blank.
- Skip the `.env` file before trying to run the backend.

Expected result:
The project is ready to start in a local development environment.

---

## [STEP 1.1.1] Install Prerequisites

Goal:
Install the tools the app depends on.

Files:
- `backend/go.mod`
- `frontend/package.json`

Actions:
1. Install Go.
2. Install Node.js and npm.
3. Install PostgreSQL.
4. Open a terminal and confirm these commands work:

```powershell
go version
node -v
npm -v
```

Expected result:
Your machine is ready to run both sides of the template.

Requires:
Complete [STEP 1.1.0] before continuing.

---

## [STEP 1.1.2] Configure the Backend Environment

Goal:
Set the backend environment values that the app reads from `backend/.env`.

Files:
- `backend/.env`
- `backend/config/config.go`

Actions:
1. Open `backend/.env`.
2. Set `AUTH_DB_URL` to your PostgreSQL connection string.
3. Set `JWT_SECRET` to a private value with at least 32 characters.
4. Confirm the app is expecting values from the environment in `backend/config/config.go`.
5. Save the file.

Important:
The backend validates `AUTH_DB_URL` and `JWT_SECRET` before it starts.

Do not:
- Leave them blank.
- Use a weak secret.
- Commit `.env` changes to Git.

Expected result:
The backend can start without environment-related errors.

---

## [STEP 1.1.3] Start the Backend

Goal:
Run the Go server so the API becomes available.

Files:
- `backend/main.go`
- `backend/routes/routes.go`

Actions:
1. Open a terminal.
2. Change into the backend directory.
3. Run:

```powershell
cd backend
go run .
```

4. Wait for the server to start without crashing.

What this changes:
The API becomes available for the frontend to call.

Troubleshooting:
- If the server fails, check the `.env` values first.
- Confirm PostgreSQL is reachable.
- Confirm the port is valid.

Expected result:
The backend is running and ready for requests.

---

## [STEP 1.1.4] Start the Frontend

Goal:
Run the Vite React app.

Files:
- `frontend/package.json`
- `frontend/vite.config.js`

Actions:
1. Open a second terminal.
2. Change into the frontend folder.
3. Install dependencies if they are not present:

```powershell
cd frontend
npm install
```

4. Start the app:

```powershell
npm run dev
```

5. Open the local Vite URL in the browser.

What this changes:A
The React interface loads and can talk to the backend API.

Expected result:
The frontend loads and can render the login, registration, and home pages.

---

## [STEP 1.1.5] Verify the Application

Goal:
Confirm that the template works before changing anything.

Files:
- `frontend/src/App.jsx`
- `frontend/src/pages/HomePage.jsx`
- `frontend/src/pages/Login.jsx`
- `frontend/src/pages/Register.jsx`
- `backend/routes/routes.go`

Actions:
1. Open the app in the browser.
2. See whether the home page loads.
3. Create a new account.
4. Sign in with that account.
5. Confirm the Navbar changes to the signed-in state.
6. Sign out.
7. Confirm the app redirects correctly.

Expected result:
The base template works as intended and you are ready to customize it.

Before continuing:
Complete all steps in 1.1.0 through 1.1.5.

---

## [STEP 2.1.0] Understand the Frontend Structure

Current area: 2.0 Understand the Template
Current step: 2.1.0
Next: 2.1.1 Review the app shell and routes

Goal:
Learn how the React app is organized before changing layout or behavior.

Files:
- `frontend/src/App.jsx`
- `frontend/src/pages/`
- `frontend/src/components/`
- `frontend/src/context/`
- `frontend/src/services/api.js`
- `frontend/src/index.css`

What to review:
- `App.jsx` registers routes.
- `pages/` contains the page-level screens.
- `components/` contains shared UI pieces.
- `context/` stores authentication state.
- `services/api.js` is the shared API helper.
- `index.css` holds the global styling rules.

Expected result:
You can identify where to edit a page, route, shared component, or style without guessing.

---

## [STEP 2.1.1] Review the App Shell and Routes

Files:
- `frontend/src/App.jsx`

Actions:
1. Open `App.jsx`.
2. Review the route setup.
3. Notice the home route, login route, and register route.
4. Notice the `PublicOnly` wrapper that prevents signed-in users from visiting auth pages.
5. Add future pages in the same route structure.

What this changes:
This file decides which page loads at each URL.

Do not:
- Break the route nesting.
- Remove the `Navigate` fallback.

Expected result:
You can add or adjust pages without damaging navigation.

---

## [STEP 2.1.2] Review Authentication State

Files:
- `frontend/src/context/AuthProvider.jsx`
- `frontend/src/context/useAuth.js`

Actions:
1. Open `AuthProvider.jsx`.
2. Review how the app stores the token and user in localStorage.
3. Review how login, register, and logout functions call the backend.
4. Confirm that session state is kept in sync across tabs using the browser storage event.

What this changes:
This is the central logic that determines whether a user is signed in.

Expected result:
You know where to change account/session behavior.

---

## [STEP 2.1.3] Review the Existing Pages

Files:
- `frontend/src/pages/HomePage.jsx`
- `frontend/src/pages/Login.jsx`
- `frontend/src/pages/Register.jsx`

Actions:
1. Open each page file.
2. Check how the page decides between guest and signed-in states.
3. Notice how the forms submit and navigate after success.
4. Use the same patterns when creating a new page.

Expected result:
You can adapt the current UI copy, layout, and flow to your project.

---

## [STEP 2.2.0] Understand the Backend Structure

Files:
- `backend/routes/routes.go`
- `backend/controllers/authController.go`
- `backend/config/config.go`
- `backend/middleware/auth.go`
- `backend/models/`

Actions:
1. Open `routes.go` to see the API route map.
2. Open `authController.go` to review registration, login, logout, and validation.
3. Open `config.go` to see required environment values.
4. Review the middleware layer for JWT verification and session authorization.
5. Review the models if you plan to add new fields or session data.

Expected result:
You know which part of the backend to edit for each type of change.

---

## [STEP 3.1.0] Customize the Global Layout

Current area: 3.0 Frontend Customization
Current step: 3.1.0
Next: 3.1.1 Edit the Navbar

Goal:
Create a consistent structure for every page in the app.

Files:
- `frontend/src/App.jsx`
- `frontend/src/components/Navbar.jsx`
- `frontend/src/index.css`

What to do:
1. Open `App.jsx`.
2. Review the app shell.
3. Open `Navbar.jsx` and decide whether the navigation should stay simple or be expanded.
4. Adjust the layout in `index.css` if you want broader changes to spacing, max width, or container behavior.

Do not:
- Remove the `Routes` wrapper.
- Break the order of global elements.

Expected result:
The app keeps a clean, reusable base layout.

---

## [STEP 3.1.1] Edit the Navbar

Files:
- `frontend/src/components/Navbar.jsx`
- `frontend/src/index.css`

Goal:
Change the top navigation without breaking routing and auth state.

Actions:
1. Open the Navbar file.
2. Change visible link text, such as Sign in or Create account.
3. Change the `to` values if a link should point to a different route.
4. Update the signed-in account section if you want a different user display.
5. Update the associated styles in `index.css`.
6. Refresh the browser and test every link.

What this changes:
The visible navigation and account controls for guests and signed-in users.

Do not:
- Remove the `Link` components without adding a valid route.
- Break the logout function.

Expected result:
The header still works and matches your branding.

---

## [STEP 3.1.2] Edit the App Shell

Files:
- `frontend/src/App.jsx`

Actions:
1. Open `App.jsx`.
2. Add new routes if you want more pages.
3. Keep the main wrapper and page structure stable.
4. If a page should be protected, wrap it with the existing auth pattern.

Expected result:
The app shell remains consistent while new pages are added.

---

## [STEP 3.2.0] Customize Global Styling

Files:
- `frontend/src/index.css`

Goal:
Change the template theme without rewriting the whole app.

Common edits:
- Change colors and accent values.
- Change button styling.
- Adjust form spacing.
- Adjust card design and typography.
- Update mobile-specific rules.

Actions:
1. Open `index.css`.
2. Find the existing classes for buttons, forms, pages, and navigation.
3. Update the colors, spacing, and typography to match your brand.
4. Refresh the app to test the results.

Do not:
- Remove base classes used by the current components.
- Rename a class without updating every matching component.

Expected result:
The application keeps the same structure while reflecting your design system.

---

## [STEP 3.3.0] Customize Existing Pages

Files:
- `frontend/src/pages/HomePage.jsx`
- `frontend/src/pages/Login.jsx`
- `frontend/src/pages/Register.jsx`
- `frontend/src/index.css`

Goal:
Adjust the current screens to match your product.

---

## [STEP 3.3.1] Customize the Home Page

Files:
- `frontend/src/pages/HomePage.jsx`
- `frontend/src/index.css`

Actions:
1. Open `HomePage.jsx`.
2. Change the welcome copy.
3. Change the call-to-action buttons.
4. Change the signed-in and guest states if needed.
5. Add your own branding or product text.
6. Refresh the browser.

What this changes:
The primary landing experience for visitors.

Expected result:
The homepage reflects your product or brand without breaking session logic.

---

## [STEP 3.3.2] Customize the Login Page

Files:
- `frontend/src/pages/Login.jsx`
- `frontend/src/index.css`

Actions:
1. Change the text shown on the sign-in form.
2. Add or remove form field labels as needed.
3. Adjust the button labels or layout.
4. Confirm the form still calls `login()` correctly.

Expected result:
The login screen matches the rest of the app while keeping authentication working.

---

## [STEP 3.3.3] Customize the Register Page

Files:
- `frontend/src/pages/Register.jsx`
- `frontend/src/index.css`

Actions:
1. Change the title and copy.
2. Edit the fields if you want to collect different registration information.
3. Make sure the form still calls `register()` from the auth context.
4. Test creating an account again.

Important:
If you change registration fields, you must also update the backend request validation and database model to match.

Expected result:
Your sign-up flow remains consistent with your design and backend requirements.

---

## [STEP 3.4.0] Create a New Page

Files:
- `frontend/src/pages/`
- `frontend/src/App.jsx`
- `frontend/src/components/Navbar.jsx`
- `frontend/src/index.css`

Goal:
Add a new page to the app and make it reachable from the UI.

Actions:
1. Create a new page component in `frontend/src/pages/`.
2. Import it into `App.jsx`.
3. Add a new `Route` in the main router.
4. Add a navigation link in `Navbar.jsx` if the page should be visible.
5. Apply component-specific styling in `index.css`.
6. Open the page in the browser and test the route.

Expected result:
The new page loads correctly and is reachable from the app.

Requires:
Complete [STEP 3.1.0] before adding new page routes.

---

## [STEP 3.5.0] Create a Reusable Component

Files:
- `frontend/src/components/`
- `frontend/src/pages/`
- `frontend/src/index.css`

Goal:
Factor repeated UI into a reusable part of the app.

Actions:
1. Decide what pattern repeats across pages.
2. Create a component in `frontend/src/components/`.
3. Pass dynamic content in with props.
4. Keep the component focused on one job.
5. Import and render it in one or more pages.
6. Add styling in `index.css`.

Expected result:
Your app becomes easier to maintain and extend.

---

## [STEP 4.1.0] Customize Registration

Current area: 4.0 Authentication Customization
Current step: 4.1.0
Next: 4.2.0 Customize login

Files:
- `frontend/src/pages/Register.jsx`
- `frontend/src/context/AuthProvider.jsx`
- `backend/controllers/authController.go`
- `backend/routes/routes.go`
- `backend/models/`

Actions:
1. Modify the registration form fields in `Register.jsx`.
2. Update `AuthProvider.register()` if the request shape changes.
3. Update the backend to accept the new fields.
4. Validate the data in the controller.
5. Save the new data in the database model.
6. Test account creation with a real browser flow.

Do not:
- Change the frontend and backend request shape separately.
- Remove required validation without updating the API.

Expected result:
Users can still create accounts successfully.

---

## [STEP 4.2.0] Customize Login

Files:
- `frontend/src/pages/Login.jsx`
- `frontend/src/context/AuthProvider.jsx`
- `backend/controllers/authController.go`

Actions:
1. Edit the login form fields in `Login.jsx`.
2. Review the backend login route and response shape.
3. Keep the token/session behavior consistent with the frontend auth provider.
4. Test signing in with a valid account.
5. Test signing in with an invalid account.

Expected result:
The app continues to authenticate and save the user session correctly.

---

## [STEP 4.3.0] Customize Logout

Files:
- `frontend/src/components/Navbar.jsx`
- `frontend/src/context/AuthProvider.jsx`
- `backend/controllers/authController.go`
- `backend/middleware/auth.go`

Actions:
1. Review the sign-out button in the Navbar.
2. Review the logout function in `AuthProvider.jsx`.
3. Check the API call to `/api/auth/logout`.
4. Confirm the local session is cleared from browser storage.
5. Test that a user is redirected to login after logout.

Expected result:
Users can end a session cleanly and the app does not keep stale auth data.

---

## [STEP 4.4.0] Customize Protected Routes

Files:
- `frontend/src/components/ProtectedRoute.jsx`
- `frontend/src/App.jsx`

Goal:
Control which pages require an active user session.

Actions:
1. Open `ProtectedRoute.jsx`.
2. Decide whether a page should be public only or protected.
3. Wrap routes in `App.jsx` using the same authentication pattern.
4. Test the behavior as a guest and as a signed-in user.

Expected result:
Guests are redirected to sign in, and signed-in users can access protected pages.

---

## [STEP 5.1.0] Customize API Routes

Files:
- `backend/routes/routes.go`

Actions:
1. Open the route registration file.
2. Add a new group or route if needed.
3. Connect the route to the correct controller method.
4. Decide whether the route should be public or protected.
5. Test the endpoint with the browser or Postman.

Expected result:
Your backend endpoints reflect the feature you are building.

---

## [STEP 5.2.0] Customize Controllers

Files:
- `backend/controllers/`

Actions:
1. Find the controller that matches the feature.
2. Add request validation.
3. Handle success and error responses consistently.
4. Return data in the shape expected by the frontend.
5. Test the request end-to-end.

Expected result:
The backend responds in a stable, frontend-friendly format.

---

## [STEP 5.3.0] Customize Authentication Data and Session Behavior

Files:
- `backend/controllers/authController.go`
- `backend/models/`
- `backend/middleware/auth.go`

Actions:
1. Review how users are saved.
2. Review how passwords are hashed.
3. Review how JWT tokens are issued and validated.
4. Review how session tokens are stored and revoked.
5. Change only the parts required for your app’s rules.

Do not:
- Break the JWT validation flow.
- Remove the hashing check.

Expected result:
User accounts and sessions remain secure and consistent.

---

## [STEP 5.4.0] Update Environment and Database Settings

Files:
- `backend/config/config.go`
- `backend/.env`

Actions:
1. Review the required environment values.
2. Add new variables only when needed.
3. Keep defaults reasonable for local development.
4. Update the config validation if you add a new required setting.

Expected result:
The backend remains predictable and easy to run across environments.

---

## [STEP 6.1.0] Use the Shared API Helper

Files:
- `frontend/src/services/api.js`

Goal:
Keep all frontend requests consistent.

Actions:
1. Open `api.js`.
2. Review how it adds the bearer token and JSON headers.
3. Use `apiFetch()` for requests instead of writing raw `fetch()` calls across the app.
4. Handle request errors in the same way everywhere.

Expected result:
Your app has one consistent API pattern for all frontend requests.

---

## [STEP 6.2.0] Create a Complete Feature End-to-End

Goal:
Build a full feature using the template’s current architecture.

Typical flow:
1. Create a page in `frontend/src/pages/`.
2. Add the route in `frontend/src/App.jsx`.
3. Add navigation in `frontend/src/components/Navbar.jsx` if needed.
4. Create or edit a frontend component.
5. Add API route in `backend/routes/routes.go`.
6. Add controller logic in `backend/controllers/`.
7. Update the model or DB structure if required.
8. Test the end-to-end feature.

Expected result:
The feature works across the UI, API, and database layers.

---

## [STEP 7.1.0] Test Frontend Behavior

Files:
- `frontend/src/`

Check:
- Navigation works.
- Links open the correct route.
- Buttons trigger the correct action.
- The layout still looks correct on smaller screens.
- No console errors appear.

Expected result:
The frontend remains stable and usable after your changes.

---

## [STEP 7.2.0] Test Authentication Flow

Files:
- `frontend/src/pages/Login.jsx`
- `frontend/src/pages/Register.jsx`
- `frontend/src/components/Navbar.jsx`
- `frontend/src/context/AuthProvider.jsx`

Check:
- Account creation works.
- Login works.
- Logout works.
- Guests are blocked from protected pages.
- Signed-in users can access the correct pages.

Expected result:
The app’s auth experience is consistent and secure.

---

## [STEP 7.3.0] Test Backend APIs

Files:
- `backend/routes/routes.go`
- `backend/controllers/`

Check:
- Routes respond as expected.
- Validation errors are returned correctly.
- Successful requests return the correct response shape.
- Unauthorized requests are blocked when required.

Expected result:
The backend remains reliable and frontend-compatible.

---

## [STEP 7.4.0] Build and Prepare for Production

Files:
- `backend/.env`
- `frontend/.env`
- `frontend/package.json`

Actions:
1. Review environment variables before release.
2. Ensure secrets are never committed.
3. Run a frontend build check.
4. Verify the backend still runs with the expected environment.
5. Confirm the app works with the production configuration.

Expected result:
The project is ready for a production-style deployment check.

---

## Quick Reference: Where the Real Template Logic Lives

- Frontend routes: `frontend/src/App.jsx`
- Home page: `frontend/src/pages/HomePage.jsx`
- Login page: `frontend/src/pages/Login.jsx`
- Register page: `frontend/src/pages/Register.jsx`
- Global navigation: `frontend/src/components/Navbar.jsx`
- Protected-page guard: `frontend/src/components/ProtectedRoute.jsx`
- Shared auth state: `frontend/src/context/AuthProvider.jsx`
- Shared API helper: `frontend/src/services/api.js`
- Global styles: `frontend/src/index.css`
- Backend routes: `backend/routes/routes.go`
- Auth controllers: `backend/controllers/authController.go`
- Environment settings: `backend/config/config.go`
- Middleware/auth checks: `backend/middleware/auth.go`

## Final Rule

When you customize HolyGrail, change one layer at a time:

1. UI in the frontend
2. auth logic in the client context
3. API routes and handlers in the backend
4. database and environment data
5. test the flow again

The guide is meant to be used as a working plan, not as a static summary. Start at [STEP 1.1.0] and continue in order.
