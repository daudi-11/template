import { BrowserRouter } from 'react-router-dom'
import AuthProvider from './AuthProvider.jsx'
import AppRoutes from './routes/routes.jsx'

const App = () => {
  return (
    <AuthProvider>
      <BrowserRouter>
        <AppRoutes />
      </BrowserRouter>
    </AuthProvider>
  )
}

export default App