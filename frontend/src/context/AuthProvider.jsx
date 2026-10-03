import { useEffect, useState } from 'react'
import { apiFetch } from '../services/api.js'
import { AuthContext } from './AuthContext.jsx'

const TOKEN_KEY = 'holygrail.token'
const USER_KEY = 'holygrail.user'

function loadUser() {
  try {
    return JSON.parse(localStorage.getItem(USER_KEY) || 'null')
  } catch {
    return null
  }
}

export default function AuthProvider({ children }) {
  const [token, setToken] = useState(() => localStorage.getItem(TOKEN_KEY) || '')
  const [user, setUser] = useState(loadUser)

  // [STEP 4.1] Keep open tabs on the same browser session when another tab signs in or out.
  useEffect(() => {
    function syncSession(event) {
      if (event.key !== null && event.key !== TOKEN_KEY && event.key !== USER_KEY) return
      setToken(localStorage.getItem(TOKEN_KEY) || '')
      setUser(loadUser())
    }

    window.addEventListener('storage', syncSession)
    return () => window.removeEventListener('storage', syncSession)
  }, [])

  function saveSession(nextToken, nextUser) {
    localStorage.setItem(TOKEN_KEY, nextToken)
    localStorage.setItem(USER_KEY, JSON.stringify(nextUser))
    setToken(nextToken)
    setUser(nextUser)
  }

  async function login(credentials) {
    const result = await apiFetch('/api/auth/login', { method: 'POST', body: credentials })
    if (!result.token || !result.user?.email) {
      throw new Error('The server did not return a complete user session.')
    }
    saveSession(result.token, result.user)
    return result
  }

  async function register(credentials) {
    await apiFetch('/api/auth/register', { method: 'POST', body: credentials })
    return login({ email: credentials.email, password: credentials.password })
  }

  async function logout() {
    try {
      if (token) await apiFetch('/api/auth/logout', { method: 'POST', token })
    } catch {
      // Remove the local session even if the server cannot be reached.
    } finally {
      localStorage.removeItem(TOKEN_KEY)
      localStorage.removeItem(USER_KEY)
      setToken('')
      setUser(null)
    }
  }

  return (
    <AuthContext.Provider value={{ token, user, isAuthenticated: Boolean(token && user), login, register, logout }}>
      {children}
    </AuthContext.Provider>
  )
}
