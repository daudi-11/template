import { Link, useNavigate } from 'react-router-dom'
import { useAuth } from '../context/useAuth.js'

export default function Navbar() {
  const navigate = useNavigate()
  const { isAuthenticated, user, logout } = useAuth()

  async function handleLogout() {
    await logout()
    navigate('/login', { replace: true })
  }

  return (
    <header className="topbar">
      <div className="topbar-inner">
        <Link className="brand" to="/" aria-label="Home">
          <span className="brand-mark">H</span>
          <span>holygrail<span className="brand-period">.</span></span>
        </Link>
        {/* [STEP 4.2] Keep navigation focused on home and account actions for guests and signed-in users. */}
        {isAuthenticated ? (
          <>
            <div className="account-tools">
              <span className="role-pill">Signed in</span>
              <span className="account-email">{user.email}</span>
              <button className="sign-out" onClick={handleLogout} type="button">
                <span aria-hidden="true">↗</span><span>Sign out</span>
              </button>
            </div>
          </>
        ) : (
          <nav className="main-nav public-nav" aria-label="Account navigation">
            <Link className="nav-link" to="/login">Sign in</Link>
            <Link className="nav-join" to="/register">Create account <span aria-hidden="true">↗</span></Link>
          </nav>
        )}
      </div>
    </header>
  )
}
