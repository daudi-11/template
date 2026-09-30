import { useEffect, useState } from 'react'
import { apiFetch } from '../api.js'
import WalletCards from '../components/WalletCards.jsx'
import { readDepositProgress } from '../depositProgress.js'
import { useAuth } from '../useAuth.js'

const currency = new Intl.NumberFormat('en-US', {
  style: 'currency',
  currency: 'USD',
  maximumFractionDigits: 2,
})

function User() {
  const { email, token } = useAuth()
  const [wallets, setWallets] = useState([])
  const [deposits] = useState(() => readDepositProgress(email))
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const totalBalance = wallets.reduce((sum, wallet) => sum + Number(wallet.balance || 0), 0)

  useEffect(() => {
    let active = true

    // This view reads wallets belonging only to the signed-in parent account.
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
          <p className="eyebrow">YOUR FAMILY WALLETS</p>
          <h1>Wallet overview<span>.</span></h1>
          <p>Balances and tracked allowance progress for your children.</p>
        </div>
      </header>

      {error && <div className="dashboard-notice error" role="alert">{error}</div>}

      <div className="dashboard-summary">
        <article className="summary-card summary-total"><span className="summary-label">TOTAL AVAILABLE</span><strong>{currency.format(totalBalance)}</strong><span className="summary-foot">Across your family wallets</span></article>
        <article className="summary-card"><span className="summary-label">CHILD WALLETS</span><strong>{String(wallets.length).padStart(2, '0')}</strong><span className="summary-foot">Connected to your account</span></article>
        <article className="summary-card"><span className="summary-label">TRACKED DEPOSITS</span><strong>{Object.keys(deposits).length}</strong><span className="summary-foot">Recorded in this browser</span></article>
      </div>

      <section className="dashboard-panel subpage-panel">
        <div className="panel-heading"><div><p className="eyebrow">BALANCE BY WALLET</p><h2>Children’s wallets</h2></div></div>
        {loading ? <div className="loading-state">Loading wallets…</div> : (
          <WalletCards deposits={deposits} onRefresh={refreshWallet} totalBalance={totalBalance} wallets={wallets} />
        )}
        <p className="progress-disclaimer">Progress reflects allowance sends and spending recorded in this browser.</p>
      </section>
    </section>
  )
}

export default User
