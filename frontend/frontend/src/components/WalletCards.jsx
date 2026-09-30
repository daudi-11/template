const currency = new Intl.NumberFormat('en-US', {
  style: 'currency',
  currency: 'USD',
  maximumFractionDigits: 2,
})

function WalletCards({ wallets, deposits, totalBalance, selectedChildId, onSelect, onRefresh }) {
  if (wallets.length === 0) {
    return (
      <div className="empty-wallets">
        <span className="empty-mark" aria-hidden="true">＋</span>
        <strong>No child wallets yet</strong>
        <p>Create a child wallet from the family page to see balances and allowance progress here.</p>
      </div>
    )
  }

  return (
    <div className="wallet-grid">
      {wallets.map((child, index) => {
        const balance = Number(child.balance || 0)
        const deposit = deposits[child.id]
        const depositedAmount = Number(deposit?.amount || 0)
        const amountLeft = Math.min(depositedAmount, Math.max(0, Number(deposit?.remaining || 0)))
        // The fill represents the remaining amount from this browser's most recent tracked credit.
        const remainingPercent = depositedAmount > 0 ? Math.round((amountLeft / depositedAmount) * 100) : 0
        const familyShare = totalBalance > 0 ? Math.round((balance / totalBalance) * 100) : 0
        const selected = String(child.id) === selectedChildId

        return (
          <article className={`child-wallet-card ${selected ? 'is-selected' : ''}`} key={child.id}>
            <div className="child-card-top">
              <span className={`child-avatar avatar-${index % 3}`}>{child.name?.charAt(0)?.toUpperCase() || '?'}</span>
              <span className="wallet-status">{selected ? 'SELECTED' : `FAMILY SHARE ${familyShare}%`}</span>
            </div>
            <h3>{child.name}</h3>
            <p className="child-email">{child.email}</p>
            <div className="child-balance-label">CURRENT BALANCE</div>
            <strong className="child-balance">{currency.format(balance)}</strong>
            <div className="progress-caption">
              <span>Left from last tracked deposit</span>
              <span>{remainingPercent}%</span>
            </div>
            <div className="progress-track" role="progressbar" aria-label={`${child.name}: amount remaining from last tracked deposit`} aria-valuemin="0" aria-valuemax="100" aria-valuenow={remainingPercent}>
              <div className="progress-fill" style={{ width: `${remainingPercent}%` }} />
            </div>
            <div className="deposit-progress-detail">
              {depositedAmount > 0
                ? `${currency.format(amountLeft)} of ${currency.format(depositedAmount)} remains`
                : 'No deposit tracked in this browser yet'}
            </div>
            <div className="wallet-card-actions">
              {onSelect && <button className="card-action-primary" onClick={() => onSelect(String(child.id))}>{selected ? 'Selected' : 'Choose wallet'}</button>}
              {onRefresh && <button className="card-action-secondary" onClick={() => onRefresh(child)} aria-label={`Refresh ${child.name}'s balance`}>Refresh balance</button>}
            </div>
          </article>
        )
      })}
    </div>
  )
}

export default WalletCards
