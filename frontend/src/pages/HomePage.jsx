import { Link } from 'react-router-dom'
import { useAuth } from '../context/useAuth.js'

function displayName(user) {
  const name = user?.name || user?.email?.split('@')[0] || 'there'
  return name.replace(/[._-]+/g, ' ').replace(/\b\w/g, (letter) => letter.toUpperCase())
}

export default function HomePage() {
  const { isAuthenticated, user } = useAuth()

  // [STEP 4.2] Guests get a public welcome with account links; signed-in users get a compact account home.
  if (!isAuthenticated) {
    return (
      <main className="simple-home">
        <section className="simple-home-card">
          <p className="simple-home-kicker">A SIMPLE PROJECT TEMPLATE</p>
          <h1>A good place to start.</h1>
          <p className="simple-home-copy">Explore the site as a guest, or create an account to keep your own session across this browser.</p>
          <div className="simple-home-actions">
            <Link className="button button-primary" to="/register">Create account</Link>
            <Link className="simple-home-secondary" to="/login">Sign in</Link>
          </div>
        </section>
      </main>
    )
  }

  // [STEP 4.2] The shared header shows the signed-in state and account controls on every page.
  return (
    <main className="simple-home">
      <section className="simple-home-card simple-account-card">
        <p className="simple-home-kicker"><span className="status-dot" /> ACCOUNT</p>
        <h1>Welcome, {displayName(user)}.</h1>
        <p className="simple-home-copy">You are signed in as <strong>{user?.email}</strong>.</p>
        <p className="simple-session-note">Your account session is shared with other tabs in this browser.</p>
      </section>
    </main>
  )
}
