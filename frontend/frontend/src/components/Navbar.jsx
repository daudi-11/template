import { Link, NavLink, useNavigate } from 'react-router-dom'
import { useAuth } from '../useAuth.js'

function Navbar() {
  const navigate = useNavigate()
  const { email, logout, token } = useAuth()

  async function handleLogout() {
    try {
      await logout()
    } catch {
      // The auth provider clears local session state even if the server cannot be reached.
    }
    navigate('/', { replace: true })
  }

  return (
    <header className="site-header">
      <div className="site-header-inner">
        {/* Replace this text with your own logo or brand component. */}
        <Link className="site-logo" to="/" aria-label="Pocket home">
          <span className="site-logo-mark" aria-hidden="true">p</span>
          <span>Pocket<span className="logo-period">.</span></span>
        </Link>

        <nav className="site-nav" aria-label="Main navigation">
          <NavLink className={({ isActive }) => isActive ? 'site-nav-link is-active' : 'site-nav-link'} end to="/">Home</NavLink>
          {token && (
            <>
              <NavLink className={({ isActive }) => isActive ? 'site-nav-link is-active' : 'site-nav-link'} to="/parent">Family wallet</NavLink>
              <NavLink className={({ isActive }) => isActive ? 'site-nav-link is-active' : 'site-nav-link'} to="/user">Wallets</NavLink>
              <NavLink className={({ isActive }) => isActive ? 'site-nav-link is-active' : 'site-nav-link'} to="/admin">Insights</NavLink>
            </>
          )}
        </nav>

        {token ? (
          <div className="nav-account">
            <span className="nav-email">{email}</span>
            <button className="nav-signout" onClick={handleLogout}>Sign out</button>
          </div>
        ) : (
          <a className="nav-signin" href="#sign-in">Sign in <span aria-hidden="true">↗</span></a>
        )}
      </div>
    </header>
  )
}

export default Navbar