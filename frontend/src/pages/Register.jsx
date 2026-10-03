import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { useAuth } from '../context/useAuth.js'

export default function Register() {
  const navigate = useNavigate()
  const { register } = useAuth()
  const [form, setForm] = useState({ name: '', email: '', password: '' })
  const [error, setError] = useState('')
  const [busy, setBusy] = useState(false)

  async function handleSubmit(event) {
    event.preventDefault()
    setBusy(true)
    setError('')
    try {
      await register(form)
      navigate('/', { replace: true })
    } catch (requestError) {
      setError(requestError.message || 'Account creation failed. Please try again.')
    } finally {
      setBusy(false)
    }
  }

  return (
    <main className="auth-layout register-layout auth-layout-simple">

      <section className="auth-panel" aria-labelledby="register-title">
        <div className="auth-panel-heading">
          <p className="overline">GET STARTED</p>
          <h2 id="register-title">Create your account</h2>
          <p>Create an account to sign in and keep a session across tabs in this browser.</p>
        </div>
        {error && <div className="form-alert" role="alert">{error}</div>}
        <form className="auth-form" onSubmit={handleSubmit}>
          <label className="field"><span>Your name</span><input autoComplete="name" autoFocus required value={form.name} onChange={(event) => setForm({ ...form, name: event.target.value })} /></label>
          <label className="field"><span>Email address</span><input autoComplete="email" required type="email" value={form.email} onChange={(event) => setForm({ ...form, email: event.target.value })} /></label>
          <label className="field"><span>Password</span><input autoComplete="new-password" minLength="8" required type="password" value={form.password} onChange={(event) => setForm({ ...form, password: event.target.value })} /></label>
          <button className="button button-primary full-width" disabled={busy} type="submit">{busy ? 'Creating account…' : 'Create account'} <span aria-hidden="true">→</span></button>
        </form>
        <p className="auth-foot">Already registered? <Link to="/login">Sign in</Link></p>
      </section>
    </main>
  )
}
