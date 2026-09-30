const storagePrefix = 'pocket-deposit-progress:'

export function readDepositProgress(email) {
  if (!email) return {}

  try {
    return JSON.parse(localStorage.getItem(`${storagePrefix}${email.toLowerCase()}`) || '{}')
  } catch {
    return {}
  }
}

export function saveDepositProgress(email, progress) {
  if (!email) return
  localStorage.setItem(`${storagePrefix}${email.toLowerCase()}`, JSON.stringify(progress))
}
