import { axios } from './api'

export interface CreateTenantForm {
  name: string
  parentId: string | null
}

export interface TenantListQuery {
  name: string
  page: number
  limit: number
}

export interface TenantItem {
  id: string
  name: string
  parentId: string | null
}

export interface TenantListData {
  list: TenantItem[]
  count: number
}

export async function createTenant(data: CreateTenantForm) {
  return await axios.post('/tenant/create', data)
}

export async function getTenant(id: string): Promise<TenantItem> {
  const data = (await axios.get(`/tenant/${id}`)) as TenantItem
  return data
}

export async function getRootTenantList(
  query?: TenantListQuery
): Promise<TenantListData> {
  const data = (await axios.get('/tenant/root/list')) as TenantListData
  return data
}

export async function getTenantChildren(): Promise<TenantListData> {
  const data = (await axios.get('/tenant/children')) as TenantListData
  return data
}
