import { useState } from 'react'
import { Link, useLocation, useNavigate } from 'react-router-dom'
import { useAuth } from '../context/useAuth.js'

export default function Login() {
  const navigate = useNavigate()
  const location = useLocation()
  const { login } = useAuth()
  const [form, setForm] = useState({ email: '', password: '' })
  const [error, setError] = useState('')
  const [busy, setBusy] = useState(false)

  async function handleSubmit(event) {
    event.preventDefault()
    setBusy(true)
    setError('')
    try {
      await login(form)
      navigate(location.state?.from?.pathname || '/', { replace: true })
    } catch (requestError) {
      setError(requestError.message || 'Sign-in failed. Check your details and try again.')
    } finally {
      setBusy(false)
    }
  }

  return (
    <main className="auth-layout auth-layout-simple">

      <section className="auth-panel" aria-labelledby="login-title">
        <div className="auth-panel-heading">
          <p className="overline">WELCOME BACK</p>
          <h2 id="login-title">Sign in to your account</h2>
          <p>Sign in to continue. Your session is shared across tabs in this browser.</p>
        </div>
        {error && <div className="form-alert" role="alert">{error}</div>}
        <form className="auth-form" onSubmit={handleSubmit}>
          <label className="field"><span>Email address</span><input autoComplete="email" autoFocus required type="email" value={form.email} onChange={(event) => setForm({ ...form, email: event.target.value })} /></label>
          <label className="field"><span>Password</span><input autoComplete="current-password" required type="password" value={form.password} onChange={(event) => setForm({ ...form, password: event.target.value })} /></label>
          <button className="button button-primary full-width" disabled={busy} type="submit">{busy ? 'Signing in…' : 'Sign in'} <span aria-hidden="true">→</span></button>
        </form>
        <p className="auth-foot">New here? <Link to="/register">Create an account</Link></p>
      </section>
    </main>
  )
}
