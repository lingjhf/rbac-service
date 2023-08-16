import { Axios } from 'axios'
import cookies from 'js-cookie'
import { useTenantStore } from '@/store'

const BASE_URL = '/api'

export function createAxios() {
  const axios = new Axios({
    baseURL: BASE_URL,
    headers: {
      'Content-Type': 'application/json',
    },
  })
  axios.interceptors.request.use((config) => {
    const tenantStore = useTenantStore()

    const token = cookies.get('token')
    if (token) {
      config.headers['Authorization'] = token
    }
    if (tenantStore.currentTenant) {
      config.headers['Tenant'] = tenantStore.currentTenant.id
    }
    if (
      config.method === 'post' &&
      config.headers['Content-Type'] === 'application/json'
    ) {
      config.data = JSON.stringify(config.data)
    }

    return config
  })
  axios.interceptors.response.use((response) => {
    const jsonData = JSON.parse(response.data)
    if (response.status !== 200) {
      return Promise.reject(jsonData.message)
    }
    return Promise.resolve(jsonData.data)
  })
  return axios
}

export const axios = createAxios()
