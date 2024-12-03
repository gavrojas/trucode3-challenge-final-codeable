import type { OptionsApiCall } from '@/types/index'
import { useAuthStore } from '@/stores/auth'

const BASE_URL = import.meta.env.VITE_API_URL || 'https://trucode-gavrojas.serveftp.com'

type ApiCallOptions = OptionsApiCall | undefined

export async function apiHandleData(
  path: string,
  { method = 'GET', data, headers, notifyLogout }: ApiCallOptions = {}
) {
  const sessionToken = sessionStorage.getItem('token')
  const requestHeaders = {
    'Content-Type': 'application/json',
    Authorization: `Bearer ${sessionToken}`,
    ...headers,
  }
  const body = method !== 'GET' && data ? JSON.stringify(data) : undefined

  try {
    const response = await fetch(BASE_URL + path, {
      method,
      headers: requestHeaders,
      body,
    })

    if(notifyLogout){
      await fetch(BASE_URL + '/logout',{
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${sessionToken}`,
          'Content-Type': 'application/json'
        },
      })}

    if (!response.ok) {
      const errorData = await response.json()

      // Si el error es de sesión expirada (401), redirigir al login
      if (response.status === 401) {
        const authStore = useAuthStore()
        authStore.clearSession() 
        throw new Error('Session expired. Redirecting to login...')
      }

      throw new Error(errorData.error || 'Request failed')
    }

    if (method === 'DELETE') {
      return response.ok
    }

    return response.json()
  } catch (error) {
    throw new Error(`API call failed: ${error instanceof Error ? error.message : 'Unknown error'}`)
  }
}

export async function apiDownload(
  path: string,
  headers: Record<string, string> = {}
) {
  const sessionToken = sessionStorage.getItem('token')
  const requestHeaders = {
    Authorization: `Bearer ${sessionToken}`,
    ...headers,
  }

  try {
    const response = await fetch(BASE_URL + path, {
      method: 'GET',
      headers: requestHeaders,
    })

    if (!response.ok) {
      const errorData = await response.json()
      if (response.status === 401) {
        const authStore = useAuthStore()
        authStore.clearSession()
        throw new Error('Session expired. Redirecting to login...')
      }
      throw new Error(errorData.error || 'Request failed')
    }

    return response.blob()
  } catch (error) {
    throw new Error(`API call failed: ${error instanceof Error ? error.message : 'Unknown error'}`)
  }
}

export async function apiUploadFile(
  path: string,
  file: File, 
  headers: Record<string, string> = {}
) {
  const sessionToken = sessionStorage.getItem('token')
  const formData = new FormData()
  formData.append('file', file)

  const requestHeaders = {
    'Authorization': `Bearer ${sessionToken}`,
    ...headers
  }

  try {
    const response = await fetch(BASE_URL + path, {
      method: 'PATCH',
      headers: requestHeaders,
      body: formData
    })

    if (!response.ok) {
      const errorData = await response.json()
      if (response.status === 401) {
        const authStore = useAuthStore()
        authStore.clearSession()
        throw new Error('Session expired. Redirecting to login...')
      }
      throw new Error(errorData.error || 'Request failed')
    }
    
    return response.json()
  } catch (error) {
    throw new Error(`File upload failed: ${error instanceof Error ? error.message : 'Unknown error'}`)
  }
}