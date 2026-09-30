import { useState } from 'react'
import { apiFetch } from './api.js'
import { AuthContext } from './AuthContext.jsx'

function AuthProvider({ children }) {
  const [token, setToken] = useState(() => localStorage.getItem('token') || '')
  const [email, setEmail] = useState(() => localStorage.getItem('parentEmail') || '')

  async function login(credentials) {
    const result = await apiFetch('/login', { method: 'POST', body: credentials })
    if (!result.token) throw new Error('The server did not return a login token.')

    localStorage.setItem('token', result.token)
    localStorage.setItem('parentEmail', credentials.email)
    setToken(result.token)
    setEmail(credentials.email)
    return result
  }

  async function logout() {
    try {
      await apiFetch('/logout', { method: 'POST', token })
    } finally {
      localStorage.removeItem('token')
      localStorage.removeItem('access_token')
      localStorage.removeItem('parentEmail')
      setToken('')
      setEmail('')
    }
  }

  return (
    <AuthContext.Provider value={{ token, email, login, logout }}>
      {children}
    </AuthContext.Provider>
  )
}

export default AuthProvider
