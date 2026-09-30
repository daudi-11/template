import { Navigate, Outlet, useLocation } from 'react-router-dom'
import { useAuth } from '../useAuth.js'

function ProtectedRoute() {
  const location = useLocation()
  const { token } = useAuth()

  if (!token) {
    // Keep the requested URL so a login page can send the user back after authentication.
    return <Navigate to="/" replace state={{ from: location }} />
  }

  // A token only allows the page to render; the backend must still validate it for protected data.
  return <Outlet />
}

export default ProtectedRoute