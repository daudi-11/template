import { Outlet } from 'react-router-dom'
import Footer from '../components/Footer.jsx'
import Navbar from '../components/Navbar.jsx'

function MainLayout() {
  return (
    <div className="site-layout">
      {/* This layout provides the shared Navbar; do not rebuild it on each page. */}
      <Navbar />
      <main className="site-main">
        <Outlet />
      </main>
      {/* This layout also provides the shared Footer; pages normally define only their own content. */}
      <Footer />
    </div>
  )
}

export default MainLayout