# Account Flow Checklist

- [ ] Open `/` while signed out. Confirm the public home offers registration and sign-in.
- [ ] Register a new account. Confirm it signs in and the home/header show the account email.
- [ ] Open another tab on the same site. Confirm it uses the same active account.
- [ ] Sign out in one tab. Confirm the other tab updates to the signed-out state.
- [ ] Sign in again, refresh, and confirm the session is restored.
- [ ] Confirm logout clears `holygrail.token` and `holygrail.user` from local storage and revokes the server-side session.

## Repeatable Checks

From the repository root, run `./scripts/verify.ps1` to execute Go tests, frontend lint, and the frontend production build.
