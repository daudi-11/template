import { useState } from 'react'
import { useLocation, useNavigate } from 'react-router-dom'
import { apiFetch } from '../api.js'
import { useAuth } from '../useAuth.js'

function HomePage() {
  const navigate = useNavigate()
  const location = useLocation()
  const { login, token } = useAuth()
  const [mode, setMode] = useState('login')
  const [form, setForm] = useState({ name: '', email: '', password: '' })
  const [busy, setBusy] = useState(false)
  const [error, setError] = useState('')

  async function handleSubmit(event) {
    event.preventDefault()
    setBusy(true)
    setError('')

    try {
      if (mode === 'register') {
        await apiFetch('/register', { method: 'POST', body: form })
      }
      await login({ email: form.email, password: form.password })
      const returnTo = location.state?.from?.pathname || '/parent'
      navigate(returnTo, { replace: true })
    } catch (requestError) {
      setError(requestError.message || 'Sign-in failed. Check your details and try again.')
    } finally {
      setBusy(false)
    }
  }

  if (token) {
    return (
      <section className="welcome-card">
        <div className="welcome-copy">
          <p className="eyebrow">YOUR FAMILY WALLET</p>
          <h1>Good money habits start small.</h1>
          <p>Send allowances, check balances, and keep every child’s wallet in one place.</p>
        </div>
        <button className="button button-primary" onClick={() => navigate('/parent')}>
          Open family wallet <span aria-hidden="true">→</span>
        </button>
      </section>
    )
  }

  return (
    <section className="home-page">
      <div className="home-intro">
        <p className="eyebrow">A LITTLE MORE INDEPENDENCE</p>
        <h1>Money lessons,<br />made everyday.</h1>
        <p className="home-description">Give your children a place to practice saving and spending, while you stay close to every balance.</p>

        <div className="home-note">
          <span className="home-note-mark" aria-hidden="true">↗</span>
          <div><strong>Thoughtful by design</strong><span>One family. Separate wallets. A clearer view.</span></div>
        </div>
      </div>

      <section className="auth-card" id="sign-in" aria-labelledby="auth-heading">
        <div className="auth-card-topline"><span className="card-chip" aria-hidden="true" /><span>POCKET FAMILY</span></div>
        <p className="eyebrow">{mode === 'login' ? 'WELCOME BACK' : 'GET STARTED'}</p>
        <h2 id="auth-heading">{mode === 'login' ? 'Sign in' : 'Create your account'}</h2>
        <p className="auth-intro">{mode === 'login' ? 'Your family wallet is just ahead.' : 'Start with your parent account.'}</p>

        {error && <p className="form-message error-message" role="alert">{error}</p>}

        <form className="auth-form" onSubmit={handleSubmit}>
          {mode === 'register' && (
            <label className="form-field">
              <span>Your name</span>
              <input autoComplete="name" required value={form.name} onChange={(event) => setForm({ ...form, name: event.target.value })} />
            </label>
          )}
          <label className="form-field">
            <span>Email</span>
            <input autoComplete="email" required type="email" value={form.email} onChange={(event) => setForm({ ...form, email: event.target.value })} />
          </label>
          <label className="form-field">
            <span>Password</span>
            <input autoComplete={mode === 'login' ? 'current-password' : 'new-password'} required type="password" value={form.password} onChange={(event) => setForm({ ...form, password: event.target.value })} />
          </label>
          <button className="button button-primary auth-submit" disabled={busy} type="submit">
            {busy ? 'One moment…' : mode === 'login' ? 'Sign in to Pocket' : 'Create account'}
            <span aria-hidden="true">→</span>
          </button>
        </form>

        <p className="auth-switch">
          {mode === 'login' ? 'New to Pocket?' : 'Already have an account?'}{' '}
          <button type="button" onClick={() => { setMode(mode === 'login' ? 'register' : 'login'); setError('') }}>
            {mode === 'login' ? 'Create an account' : 'Sign in'}
          </button>
        </p>
      </section>
    </section>
  )
}

export default HomePage