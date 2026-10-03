import { Navigate, Outlet, useLocation } from 'react-router-dom'
import { useAuth } from '../context/useAuth.js'

// [STEP 4.4] Guard nested routes by authentication and redirect guests to sign in.
export default function ProtectedRoute() {
  const location = useLocation()
  const { isAuthenticated } = useAuth()

  if (!isAuthenticated) {
    return <Navigate to="/login" replace state={{ from: location }} />
  }
  return <Outlet />
}
