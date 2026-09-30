import { Route, Routes } from 'react-router-dom'
import MainLayout from '../layouts/MainLayout.jsx'
import ProtectedRoute from '../components/ProtectedRoute.jsx'
import Home from '../pages/HomePage.jsx'
import Page2 from '../pages/Page2.jsx'
import Page3 from '../pages/Page3.jsx'
import User from '../pages/User.jsx'
import Parent from '../pages/parent.jsx'
import Admin from '../pages/admin.jsx'

function AppRoutes() {
  return (
    <Routes>
      <Route element={<MainLayout />}>
        {/* These routes are public and can be opened without signing in. */}
        <Route index element={<Home />} />
        <Route path="page2" element={<Page2 />} />
        <Route path="page3" element={<Page3 />} />

        {/* Put pages that require a saved login token inside this group. */}
        <Route element={<ProtectedRoute />}>
          <Route path="user" element={<User />} />
          <Route path="parent" element={<Parent />} />
          <Route path="admin" element={<Admin />} />
        </Route>
        {/*
          To add a new page:
          1. Create the page inside /pages.
          2. Import it at the top of this file.
          3. Add a new Route inside MainLayout's Route below.
        */}
      </Route>
    </Routes>
  )
}

export default AppRoutes