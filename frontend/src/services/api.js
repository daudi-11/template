const API_ORIGIN = (import.meta.env.VITE_API_URL || 'http://localhost:3030').replace(/\/$/, '')
const TOKEN_KEY = 'holygrail.token'

export class ApiError extends Error {
  constructor(message, status) {
    super(message)
    this.name = 'ApiError'
    this.status = status
  }
}

// [STEP 4.1] Centralize API requests, attach the saved bearer token, and parse response/error payloads.
export async function apiFetch(path, { body, headers, token, ...options } = {}) {
  const requestHeaders = new Headers(headers)
  const savedToken = token ?? localStorage.getItem(TOKEN_KEY)

  if (body !== undefined && !(body instanceof FormData)) {
    requestHeaders.set('Content-Type', 'application/json')
  }
  if (savedToken) requestHeaders.set('Authorization', `Bearer ${savedToken}`)

  const response = await fetch(`${API_ORIGIN}${path}`, {
    ...options,
    headers: requestHeaders,
    credentials: 'omit',
    body: body === undefined ? undefined : body instanceof FormData ? body : JSON.stringify(body),
  })
  const text = await response.text()
  let result = {}
  try {
    result = text ? JSON.parse(text) : {}
  } catch {
    result = { message: text }
  }

  if (!response.ok) {
    throw new ApiError(result.message || 'The request could not be completed.', response.status)
  }
  return result
}
