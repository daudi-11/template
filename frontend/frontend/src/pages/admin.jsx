import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { apiFetch } from '../api.js'
import WalletCards from '../components/WalletCards.jsx'
import { readDepositProgress } from '../depositProgress.js'
import { useAuth } from '../useAuth.js'

const currency = new Intl.NumberFormat('en-US', {
  style: 'currency',
  currency: 'USD',
  maximumFractionDigits: 2,
})

function Admin() {
  const { email, token } = useAuth()
  const [wallets, setWallets] = useState([])
  const [deposits] = useState(() => readDepositProgress(email))
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const totalBalance = wallets.reduce((sum, wallet) => sum + Number(wallet.balance || 0), 0)
  const trackedAmount = Object.values(deposits).reduce((sum, item) => sum + Number(item.amount || 0), 0)
  const trackedRemaining = Object.values(deposits).reduce((sum, item) => sum + Number(item.remaining || 0), 0)
  const trackedPercent = trackedAmount > 0 ? Math.round((trackedRemaining / trackedAmount) * 100) : 0

  useEffect(() => {
    let active = true

    apiFetch('/children', { token })
      .then((result) => {
        if (active) setWallets(Array.isArray(result.children) ? result.children : [])
      })
      .catch((requestError) => {
        if (active) setError(requestError.message)
      })
      .finally(() => {
        if (active) setLoading(false)
      })

    return () => { active = false }
  }, [token])

  async function refreshWallet(wallet) {
    try {
      const result = await apiFetch(`/balance?child_id=${wallet.id}`, { token })
      setWallets((current) => current.map((item) => (
        item.id === wallet.id ? { ...item, balance: Number(result.balance) } : item
      )))
    } catch (requestError) {
      setError(requestError.message)
    }
  }

  return (
    <section className="subpage-content">
      <header className="dashboard-heading">
        <div>
          <p className="eyebrow">FAMILY ACCOUNT · {email}</p>
          <h1>Family insights<span>.</span></h1>
          <p>A quick read on wallets and allowance progress in this account.</p>
        </div>
        <Link className="button button-quiet" to="/parent">Manage wallets <span aria-hidden="true">→</span></Link>
      </header>

      {error && <div className="dashboard-notice error" role="alert">{error}</div>}

      <div className="dashboard-summary">
        <article className="summary-card summary-total"><span className="summary-label">TOTAL WALLET BALANCE</span><strong>{currency.format(totalBalance)}</strong><span className="summary-foot">Current funds across this family</span></article>
        <article className="summary-card"><span className="summary-label">CHILDREN</span><strong>{String(wallets.length).padStart(2, '0')}</strong><span className="summary-foot">Wallets in this parent account</span></article>
        <article className="summary-card"><span className="summary-label">TRACKED FUNDS LEFT</span><strong>{currency.format(trackedRemaining)}</strong><span className="summary-foot">{trackedPercent}% of browser-tracked deposits</span></article>
      </div>

      <section className="dashboard-panel subpage-panel">
        <div className="panel-heading"><div><p className="eyebrow">WALLET HEALTH</p><h2>Allowance progress</h2></div></div>
        {loading ? <div className="loading-state">Loading family data…</div> : (
          <WalletCards deposits={deposits} onRefresh={refreshWallet} totalBalance={totalBalance} wallets={wallets} />
        )}
        <p className="progress-disclaimer">This view covers this parent account only. Progress is tracked in this browser, not server transaction history.</p>
      </section>
    </section>
  )
}

export default Admin
