import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { apiFetch } from '../api.js'
import WalletCards from '../components/WalletCards.jsx'
import { readDepositProgress, saveDepositProgress } from '../depositProgress.js'
import { useAuth } from '../useAuth.js'

const currency = new Intl.NumberFormat('en-US', {
  style: 'currency',
  currency: 'USD',
  maximumFractionDigits: 2,
})

function Parent() {
  const navigate = useNavigate()
  const { email, logout, token } = useAuth()
  const [children, setChildren] = useState([])
  const [selectedChildId, setSelectedChildId] = useState('')
  const [amount, setAmount] = useState('')
  const [transactionType, setTransactionType] = useState('credit')
  const [childForm, setChildForm] = useState({ name: '', email: '', password: '' })
  const [depositProgress, setDepositProgress] = useState(() => readDepositProgress(email))
  const [activity, setActivity] = useState([])
  const [loading, setLoading] = useState(true)
  const [sending, setSending] = useState(false)
  const [addingChild, setAddingChild] = useState(false)
  const [notice, setNotice] = useState(null)

  const totalBalance = children.reduce((total, child) => total + Number(child.balance || 0), 0)
  const selectedChild = children.find((child) => String(child.id) === selectedChildId)

  // Fetching the child list on entry restores saved wallets and balances from the backend after reload.
  async function fetchChildren() {
    const result = await apiFetch('/children', { token })
    const nextChildren = Array.isArray(result.children) ? result.children : []
    setChildren(nextChildren)
    setSelectedChildId((currentId) => currentId || String(nextChildren[0]?.id || ''))
    return nextChildren
  }

  useEffect(() => {
    let active = true

    async function loadChildren() {
      setLoading(true)
      try {
        const result = await apiFetch('/children', { token })
        if (!active) return
        const nextChildren = Array.isArray(result.children) ? result.children : []
        setChildren(nextChildren)
        setSelectedChildId((currentId) => currentId || String(nextChildren[0]?.id || ''))
      } catch (error) {
        if (active) setNotice({ type: 'error', text: error.message })
      } finally {
        if (active) setLoading(false)
      }
    }

    loadChildren()
    return () => { active = false }
  }, [token])

  async function handleSendAllowance(event) {
    event.preventDefault()
    if (!selectedChild || Number(amount) <= 0) return

    setSending(true)
    setNotice(null)
    try {
      const result = await apiFetch(`/${transactionType}`, {
        method: 'POST',
        token,
        body: { child_id: selectedChild.id, amount: Number(amount) },
      })
      const sentAmount = Number(result.amount)
      setChildren((current) => current.map((child) => (
        child.id === selectedChild.id ? { ...child, balance: Number(result.balance) } : child
      )))
      const nextProgress = { ...depositProgress }
      const previousDeposit = nextProgress[selectedChild.id]
      if (transactionType === 'credit') {
        nextProgress[selectedChild.id] = {
          amount: sentAmount,
          remaining: sentAmount,
          recordedAt: new Date().toISOString(),
        }
      } else if (previousDeposit) {
        nextProgress[selectedChild.id] = {
          ...previousDeposit,
          remaining: Math.max(0, Number(previousDeposit.remaining) - sentAmount),
          recordedAt: new Date().toISOString(),
        }
      }
      setDepositProgress(nextProgress)
      saveDepositProgress(email, nextProgress)
      setActivity((current) => [{
        id: result.transaction?.id || Date.now(),
        type: transactionType,
        childName: selectedChild.name,
        amount: sentAmount,
        date: new Date().toISOString(),
      }, ...current].slice(0, 8))
      setAmount('')
      setNotice({
        type: 'success',
        text: transactionType === 'credit'
          ? `${currency.format(sentAmount)} sent to ${selectedChild.name}.`
          : `${currency.format(sentAmount)} recorded from ${selectedChild.name}’s wallet.`,
      })
    } catch (error) {
      setNotice({ type: 'error', text: error.message })
    } finally {
      setSending(false)
    }
  }

  async function handleAddChild(event) {
    event.preventDefault()
    setAddingChild(true)
    setNotice(null)
    try {
      const result = await apiFetch('/children', {
        method: 'POST',
        token,
        body: childForm,
      })
      await fetchChildren()
      setSelectedChildId(String(result.child.id))
      setChildForm({ name: '', email: '', password: '' })
      setNotice({ type: 'success', text: `${result.child.name}’s wallet is ready.` })
    } catch (error) {
      setNotice({ type: 'error', text: error.message })
    } finally {
      setAddingChild(false)
    }
  }

  async function handleRefreshBalance(child) {
    try {
      const result = await apiFetch(`/balance?child_id=${child.id}`, { token })
      setChildren((current) => current.map((item) => (
        item.id === child.id ? { ...item, balance: Number(result.balance) } : item
      )))
      setNotice({ type: 'success', text: `${child.name}’s balance is up to date.` })
    } catch (error) {
      setNotice({ type: 'error', text: error.message })
    }
  }

  async function handleLogout() {
    try {
      await logout()
    } catch {
      // Auth state is cleared even if the logout request cannot reach the server.
    }
    navigate('/', { replace: true })
  }

  return (
    <div className="family-dashboard">
      <header className="dashboard-heading">
        <div>
          <p className="eyebrow">YOUR FAMILY, IN ONE PLACE</p>
          <h1>Good morning<span>.</span></h1>
          <p>Small steps today. More confident money choices tomorrow.</p>
        </div>
        <button className="button button-quiet sign-out-button" onClick={handleLogout}>Sign out <span aria-hidden="true">↗</span></button>
      </header>

      {notice && <div className={`dashboard-notice ${notice.type}`} role={notice.type === 'error' ? 'alert' : 'status'}>{notice.text}</div>}

      <section className="dashboard-summary" aria-label="Family wallet overview">
        <article className="summary-card summary-total">
          <span className="summary-label">TOTAL IN FAMILY WALLETS</span>
          <strong>{currency.format(totalBalance)}</strong>
          <span className="summary-foot">Across {children.length} {children.length === 1 ? 'wallet' : 'wallets'}</span>
          <span className="summary-decoration" aria-hidden="true">$</span>
        </article>
        <article className="summary-card">
          <span className="summary-label">CHILD WALLETS</span>
          <strong>{String(children.length).padStart(2, '0')}</strong>
          <span className="summary-foot">Managed from this account</span>
          <span className="summary-mark" aria-hidden="true">⌂</span>
        </article>
        <article className="summary-card">
          <span className="summary-label">ALLOWANCES SENT</span>
          <strong>{String(activity.length).padStart(2, '0')}</strong>
          <span className="summary-foot">This session</span>
          <span className="summary-mark mark-warm" aria-hidden="true">↗</span>
        </article>
      </section>

      <div className="dashboard-columns">
        <section className="dashboard-panel wallet-panel">
          <div className="panel-heading">
            <div><p className="eyebrow">SEPARATE POCKETS</p><h2>Children’s wallets</h2></div>
            <span className="panel-count">{children.length} {children.length === 1 ? 'wallet' : 'wallets'}</span>
          </div>

          {loading ? (
            <div className="loading-state" role="status">Loading your family wallets…</div>
          ) : children.length === 0 ? (
            <div className="empty-wallets">
              <span className="empty-mark" aria-hidden="true">＋</span>
              <strong>Your family’s first wallet starts here</strong>
              <p>Add a child below. Their wallet will appear here with a zero balance.</p>
            </div>
          ) : (
            <WalletCards
              deposits={depositProgress}
              onRefresh={handleRefreshBalance}
              onSelect={setSelectedChildId}
              selectedChildId={selectedChildId}
              totalBalance={totalBalance}
              wallets={children}
            />
          )}
        </section>

        <section className="dashboard-panel send-panel">
          <div className="panel-heading">
            <div><p className="eyebrow">A LITTLE SOMETHING</p><h2>Send allowance</h2></div>
            <span className="send-mark" aria-hidden="true">↗</span>
          </div>
          <p className="panel-description">Choose a wallet and add money in a few steps.</p>
          <form className="dashboard-form" onSubmit={handleSendAllowance}>
            <label className="form-field">
              <span>Send to</span>
              <select required value={selectedChildId} onChange={(event) => setSelectedChildId(event.target.value)} disabled={loading || children.length === 0}>
                <option value="">{children.length ? 'Choose a child' : 'Add a child wallet first'}</option>
                {children.map((child) => <option key={child.id} value={child.id}>{child.name}</option>)}
              </select>
            </label>
            <fieldset className="transaction-mode">
              <legend>Money movement</legend>
              <div className="transaction-mode-options">
                <button className={transactionType === 'credit' ? 'is-active' : ''} onClick={() => setTransactionType('credit')} type="button">Send allowance</button>
                <button className={transactionType === 'debit' ? 'is-active' : ''} onClick={() => setTransactionType('debit')} type="button">Record spending</button>
              </div>
            </fieldset>
            <label className="form-field">
              <span>Amount</span>
              <div className="money-input"><span aria-hidden="true">$</span><input required min="0.01" step="0.01" type="number" inputMode="decimal" placeholder="0.00" value={amount} onChange={(event) => setAmount(event.target.value)} /></div>
            </label>
            <button className="button button-primary send-button" disabled={sending || loading || children.length === 0} type="submit">
              {sending ? 'Processing…' : transactionType === 'credit' ? 'Send allowance' : 'Record spending'} <span aria-hidden="true">→</span>
            </button>
          </form>
          <div className="secure-note"><span className="secure-dot" />Progress tracks money moves made in this browser.</div>
        </section>
      </div>

      <div className="dashboard-lower-grid">
        <section className="dashboard-panel add-child-panel">
          <div className="panel-heading">
            <div><p className="eyebrow">GROW YOUR FAMILY WALLET</p><h2>Add a child</h2></div>
            <span className="panel-step">01</span>
          </div>
          <form className="dashboard-form" onSubmit={handleAddChild}>
            <label className="form-field"><span>Name</span><input autoComplete="name" required value={childForm.name} onChange={(event) => setChildForm({ ...childForm, name: event.target.value })} /></label>
            <label className="form-field"><span>Email</span><input autoComplete="email" required type="email" value={childForm.email} onChange={(event) => setChildForm({ ...childForm, email: event.target.value })} /></label>
            <label className="form-field"><span>Password</span><input autoComplete="new-password" required type="password" value={childForm.password} onChange={(event) => setChildForm({ ...childForm, password: event.target.value })} /></label>
            <button className="button button-secondary" disabled={addingChild} type="submit">{addingChild ? 'Creating wallet…' : 'Create child wallet'} <span aria-hidden="true">＋</span></button>
          </form>
        </section>

        <section className="dashboard-panel activity-panel">
          <div className="panel-heading"><div><p className="eyebrow">RECENT MOVEMENT</p><h2>Latest allowances</h2></div></div>
          {activity.length === 0 ? (
            <div className="activity-empty">Your recent allowance sends will appear here.</div>
          ) : (
            <div className="activity-list">
              {activity.map((item) => (
                <div className="activity-row" key={item.id}>
                  <span className={`activity-icon ${item.type}`} aria-hidden="true">{item.type === 'credit' ? '↗' : '↙'}</span>
                  <span className="activity-copy"><strong>{item.type === 'credit' ? 'Allowance for' : 'Spending from'} {item.childName}</strong><small>{new Date(item.date).toLocaleString([], { dateStyle: 'medium', timeStyle: 'short' })}</small></span>
                  <strong className={`activity-amount ${item.type}`}>{item.type === 'credit' ? '+' : '−'}{currency.format(item.amount)}</strong>
                </div>
              ))}
            </div>
          )}
          <p className="activity-footnote">Showing sends made during this session.</p>
        </section>
      </div>

      <p className="signed-in-note">Signed in as <strong>{email}</strong></p>
    </div>
  )
}

export default Parent
