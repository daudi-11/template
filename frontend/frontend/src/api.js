const API_ORIGIN = (import.meta.env.VITE_API_URL || 'http://localhost:3030').replace(/\/$/, '')

export class ApiError extends Error {
  constructor(message, status) {
    super(message)
    this.name = 'ApiError'
    this.status = status
  }
}

export async function apiFetch(path, { body, headers, token, ...options } = {}) {
  const requestHeaders = new Headers(headers)
  if (body !== undefined) requestHeaders.set('Content-Type', 'application/json')
  if (token) requestHeaders.set('Authorization', `Bearer ${token}`)

  // fetch includes the HTTP-only login cookie; the bearer token supports API clients that use it instead.
  const response = await fetch(`${API_ORIGIN}${path}`, {
    ...options,
    headers: requestHeaders,
    credentials: 'include',
    body: body === undefined ? undefined : JSON.stringify(body),
  })

  const result = await response.json().catch(() => ({}))
  if (!response.ok) {
    throw new ApiError(result.message || 'The request could not be completed.', response.status)
  }
  return result
}