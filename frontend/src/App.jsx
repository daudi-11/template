import { Navigate, Outlet, Route, Routes } from 'react-router-dom'
import Navbar from './components/Navbar.jsx'
import { useAuth } from './context/useAuth.js'
import HomePage from './pages/HomePage.jsx'
import Login from './pages/Login.jsx'
import Register from './pages/Register.jsx'

function PublicOnly() {
  const { isAuthenticated } = useAuth()
  return isAuthenticated ? <Navigate to="/" replace /> : <Outlet />
}

function App() {
	// [STEP 1.2] Keep frontend route registration here as pages are added to the template.
  return (
    <div className="app-shell">
      <Navbar />
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route element={<PublicOnly />}>
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
        </Route>
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </div>
  )
}

export default App
