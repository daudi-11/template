const currentYear = new Date().getFullYear()

function Footer() {
  return (
    <footer className="site-footer">
      <div className="site-footer-inner">
        <div>
          {/* Replace this with a short site name or shared footer message. */}
          <strong>Pocket Family</strong>
          {/* Add support email, phone, or other shared contact details here. */}
        </div>

        <nav className="footer-links" aria-label="Footer links">
          {/* Add shared social links or other footer menu items here. */}
        </nav>

        {/* Update the site name and add copyright details that should appear on every page. */}
        <small>© {currentYear} Pocket Family</small>
      </div>
    </footer>
  )
}

export default Footer