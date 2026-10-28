<template>
  <div class="min-h-[calc(100vh-60px)] bg-[#fbfaf7] font-sans text-[#14181d]">
    <div class="mx-auto max-w-[1600px] px-6 py-9">
      <!-- Header -->
      <div class="mb-7 flex flex-col justify-between gap-5 sm:flex-row sm:items-end">
        <div>
          <p class="text-[11px] font-semibold uppercase tracking-[0.2em] text-[#ff6b35]">Label management</p>
          <h1 class="mt-2 text-[30px] font-semibold leading-none tracking-[-0.03em]">标签管理</h1>
          <p class="mt-3 text-[13.5px] leading-relaxed text-[#14181d]/50">
            管理所有标签的链接、模板与状态 · 共 {{ total }} 枚
          </p>
        </div>
        <button
          @click="$router.push('/admin/nfc/new')"
          class="group inline-flex h-10 shrink-0 items-center justify-center gap-2 rounded-xl bg-[#14181d] px-4 text-[13px] font-medium text-white transition-all duration-200 hover:-translate-y-px hover:bg-[#ff6b35] active:translate-y-0 active:scale-[0.98]"
        >
          <Plus class="h-4 w-4 transition-transform duration-200 group-hover:rotate-90" />
          新建标签
        </button>
      </div>

      <!-- Filters -->
      <div class="mb-5 rounded-2xl border border-[#14181d]/[0.07] bg-white p-4 shadow-[0_1px_2px_rgba(20,24,29,0.03)]">
        <div class="flex flex-col items-end gap-3 md:flex-row">
          <div class="grid w-full flex-1 grid-cols-1 gap-3 md:grid-cols-4">
            <div class="space-y-1.5">
              <label class="text-[11.5px] font-medium tracking-wide text-[#14181d]/45">NFC ID / 手机号 / 标签名 / 订单号</label>
              <div class="relative">
                <Search class="pointer-events-none absolute left-3 top-1/2 h-3.5 w-3.5 -translate-y-1/2 text-[#14181d]/30" />
                <input
                  v-model="searchText"
                  class="field"
                  placeholder="搜索…"
                  @keyup.enter="load"
                />
              </div>
            </div>
            <div class="space-y-1.5">
              <label class="text-[11.5px] font-medium tracking-wide text-[#14181d]/45">状态</label>
              <select v-model="statusFilter" class="field">
                <option value="">全部状态</option>
                <option value="enabled">开启</option>
                <option value="disabled">关闭</option>
              </select>
            </div>
            <div class="space-y-1.5">
              <label class="text-[11.5px] font-medium tracking-wide text-[#14181d]/45">创建人</label>
              <select v-model="creatorFilter" class="field">
                <option value="">全部创建人</option>
                <option v-for="c in creators" :key="c.id" :value="c.id">{{ c.name }}</option>
              </select>
            </div>
          </div>
          <div class="flex w-full gap-2 md:w-auto">
            <button
              @click="load"
              class="inline-flex h-9 flex-1 items-center justify-center rounded-xl bg-[#14181d] px-4 text-[12.5px] font-medium text-white transition-colors hover:bg-[#ff6b35] md:flex-none"
            >
              搜索
            </button>
            <button
              @click="reset"
              class="inline-flex h-9 flex-1 items-center justify-center rounded-xl border border-[#14181d]/10 bg-white px-4 text-[12.5px] font-medium text-[#14181d]/65 transition-colors hover:border-[#14181d]/25 hover:text-[#14181d] md:flex-none"
            >
              重置
            </button>
          </div>
        </div>
      </div>

      <!-- Batch actions -->
      <div v-if="selectedIds.length > 0" class="mb-4 flex flex-wrap items-center gap-2">
        <span class="text-[12.5px] text-[#14181d]/45">已选 {{ selectedIds.length }} 项</span>
        <button @click="batchAction('enable')" class="rounded-full bg-[#16a34a]/10 px-3 py-1.5 text-[11.5px] font-medium text-[#15803d] transition-colors hover:bg-[#16a34a]/20">批量启用</button>
        <button @click="batchAction('disable')" class="rounded-full bg-[#14181d]/[0.06] px-3 py-1.5 text-[11.5px] font-medium text-[#14181d]/60 transition-colors hover:bg-[#14181d]/10">批量禁用</button>
        <button @click="batchAction('delete')" class="rounded-full bg-[#dc2626]/10 px-3 py-1.5 text-[11.5px] font-medium text-[#b91c1c] transition-colors hover:bg-[#dc2626]/20">批量删除</button>
      </div>

      <!-- Table -->
      <div class="overflow-hidden rounded-2xl border border-[#14181d]/[0.07] bg-white shadow-[0_1px_2px_rgba(20,24,29,0.03)]">
        <div class="relative w-full overflow-x-auto">
          <table class="w-full caption-bottom text-[13px]">
            <thead>
              <tr class="border-b border-[#14181d]/[0.06] bg-[#fbfaf7]/80">
                <th class="th w-[46px]">
                  <input type="checkbox" :checked="isAllSelected" @change="toggleSelectAll" class="checkbox" />
                </th>
                <th class="th w-[64px]">ID</th>
                <th class="th">NFC ID</th>
                <th class="th w-[120px]">订单号</th>
                <th class="th">标签名</th>
                <th class="th">链接</th>
                <th class="th">手机号</th>
                <th class="th">创建人</th>
                <th class="th">模板</th>
                <th class="th w-[92px]">状态</th>
                <th class="th w-[250px] text-right">操作</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="(row, index) in items"
                :key="row.id"
                v-motion
                :initial="{ opacity: 0, x: -14 }"
                :enter="{ opacity: 1, x: 0, transition: { delay: index * 45 } }"
                class="row"
              >
                <td class="td">
                  <input type="checkbox" :checked="selectedIds.includes(row.id)" @change="toggleSelect(row.id)" class="checkbox" />
                </td>
                <td class="td font-mono text-[11.5px] text-[#14181d]/35">{{ row.id }}</td>
                <td class="td whitespace-nowrap font-mono text-[13px] font-semibold tracking-tight">{{ row.nfc_id }}</td>
                <td class="td whitespace-nowrap">
                  <span v-if="row.order_no" class="font-mono text-[11.5px] font-medium text-[#ff6b35]" :title="'来源订单 ' + row.order_no">
                    {{ row.order_no }}
                  </span>
                  <span v-else class="text-[12px] text-[#14181d]/25">—</span>
                </td>
                <td class="td max-w-[200px]">
                  <span v-if="row.name" class="block truncate text-[13px] font-medium" :title="row.name">{{ row.name }}</span>
                  <span v-else class="text-[12.5px] text-[#14181d]/25">未命名</span>
                </td>
                <td class="td">
                  <div class="flex items-center gap-1">
                    <a
                      :href="publicLink(row)"
                      target="_blank"
                      rel="noopener"
                      :title="publicLink(row)"
                      class="inline-block max-w-[230px] truncate align-middle font-mono text-[11.5px] text-[#ff6b35] transition-colors hover:text-[#d94f1a] hover:underline"
                    >
                      {{ publicLink(row) }}
                    </a>
                    <button
                      @click="copyLink(row)"
                      title="复制链接"
                      class="inline-flex h-7 w-7 shrink-0 items-center justify-center rounded-lg text-[#14181d]/30 transition-colors hover:bg-[#14181d]/[0.05] hover:text-[#14181d]"
                    >
                      <Copy class="h-3.5 w-3.5" />
                    </button>
                  </div>
                </td>
                <td class="td whitespace-nowrap text-[12.5px] text-[#14181d]/60">{{ row.phone || '—' }}</td>
                <td class="td">
                  <span v-if="row.creator_name" class="inline-flex items-center whitespace-nowrap rounded-md bg-[#14181d]/[0.055] px-2 py-0.5 text-[11.5px] font-medium text-[#14181d]/65">{{ row.creator_name }}</span>
                  <span v-else class="text-[11.5px] text-[#14181d]/30">管理员</span>
                </td>
                <td class="td">
                  <span class="inline-flex items-center whitespace-nowrap rounded-md border border-[#14181d]/[0.07] bg-[#fbfaf7] px-2 py-0.5 text-[11.5px] font-medium text-[#14181d]/60">
                    {{ typeLabel(row.template_type) }}
                  </span>
                </td>
                <td class="td">
                  <span
                    class="inline-flex items-center gap-1.5 whitespace-nowrap rounded-full px-2.5 py-1 text-[11.5px] font-medium"
                    :class="row.status === 'enabled' ? 'bg-[#16a34a]/10 text-[#15803d]' : 'bg-[#14181d]/[0.055] text-[#14181d]/40'"
                  >
                    <span class="h-1.5 w-1.5 rounded-full" :class="row.status === 'enabled' ? 'bg-[#16a34a]' : 'bg-[#14181d]/25'" />
                    {{ row.status === 'enabled' ? '开启' : '关闭' }}
                  </span>
                </td>
                <td class="td text-right">
                  <div class="flex justify-end gap-1.5">
                    <button
                      @click="toggleStatus(row)"
                      :disabled="togglingId === row.id"
                      class="inline-flex h-8 items-center justify-center rounded-lg px-2.5 text-[11.5px] font-medium transition-colors disabled:opacity-40"
                      :class="row.status === 'enabled' ? 'bg-[#14181d]/[0.05] text-[#14181d]/60 hover:bg-[#14181d]/10' : 'bg-[#16a34a]/10 text-[#15803d] hover:bg-[#16a34a]/20'"
                    >
                      {{ row.status === 'enabled' ? '关闭' : '开启' }}
                    </button>
                    <button
                      @click="$router.push('/admin/nfc/' + row.id + '/edit')"
                      class="inline-flex h-8 w-8 items-center justify-center rounded-lg text-[#14181d]/50 transition-colors hover:bg-[#14181d]/[0.05] hover:text-[#14181d]"
                    >
                      <Edit class="h-3.5 w-3.5" />
                      <span class="sr-only">编辑</span>
                    </button>
                    <button
                      @click="remove(row)"
                      class="inline-flex h-8 w-8 items-center justify-center rounded-lg text-[#14181d]/30 transition-colors hover:bg-[#dc2626]/10 hover:text-[#b91c1c]"
                    >
                      <Trash2 class="h-3.5 w-3.5" />
                      <span class="sr-only">删除</span>
                    </button>
                  </div>
                </td>
              </tr>

              <template v-if="loading">
                <tr v-for="i in 5" :key="'sk' + i" class="border-b border-[#14181d]/[0.04]">
                  <td class="px-4 py-4" v-for="j in 11" :key="j">
                    <div class="h-3.5 animate-pulse rounded-md bg-[#14181d]/[0.055]" :style="{ width: j === 5 ? '70%' : (j === 10 ? '50%' : '100%') }"></div>
                  </td>
                </tr>
              </template>
              <tr v-else-if="items.length === 0">
                <td colspan="11" class="px-4 py-16 text-center">
                  <div class="mx-auto flex h-12 w-12 items-center justify-center rounded-2xl bg-[#14181d]/[0.04]">
                    <Search class="h-5 w-5 text-[#14181d]/25" />
                  </div>
                  <p class="mt-3 text-[13px] font-medium text-[#14181d]/60">没有找到标签</p>
                  <p class="mt-1 text-[12px] text-[#14181d]/35">换个搜索条件，或者新建第一枚标签</p>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Pagination -->
        <div class="flex flex-wrap items-center justify-between gap-3 border-t border-[#14181d]/[0.06] bg-[#fbfaf7]/60 px-4 py-3">
          <span class="text-[12px] text-[#14181d]/40">
            第 {{ page }} / {{ totalPages }} 页 · 共 {{ total }} 条
          </span>
          <div class="flex items-center gap-2">
            <button class="pg-btn" :disabled="page <= 1" @click="page--; load()">上一页</button>
            <button class="pg-btn" :disabled="page >= totalPages" @click="page++; load()">下一页</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Copy, Search } from 'lucide-vue-next'
import { ElMessageBox, ElMessage } from 'element-plus'
import { fetchNfcList, fetchCreators, deleteNfc, batchNfc, updateNfc } from '../api/nfc'
import { templateTypes } from '../templates/registry'

const items = ref([])
const loading = ref(false)
const searchText = ref('')
const statusFilter = ref('')
const creatorFilter = ref('')
const creators = ref([])
const page = ref(1)
const pageSize = 10
const total = ref(0)
const totalPages = ref(1)
const selectedIds = ref([])
const togglingId = ref(null)

/** 模板类型 key → 中文名 */
const typeLabel = (key) => templateTypes[key]?.label || key

const isAllSelected = computed(() =>
  items.value.length > 0 && items.value.every(r => selectedIds.value.includes(r.id))
)

const toggleSelectAll = () => {
  if (isAllSelected.value) {
    selectedIds.value = []
  } else {
    selectedIds.value = items.value.map(r => r.id)
  }
}

const toggleSelect = (id) => {
  const idx = selectedIds.value.indexOf(id)
  if (idx >= 0) {
    selectedIds.value.splice(idx, 1)
  } else {
    selectedIds.value.push(id)
  }
}

const load = async () => {
  loading.value = true
  try {
    const params = { page: page.value, page_size: pageSize }
    if (searchText.value) params.search = searchText.value
    if (statusFilter.value) params.status = statusFilter.value
    if (creatorFilter.value !== '') params.creator_id = creatorFilter.value
    const { data } = await fetchNfcList(params)
    items.value = data.items
    total.value = data.total
    totalPages.value = data.total_pages
    selectedIds.value = []
  } catch {
    // 错误已由 request 拦截器处理
  } finally {
    loading.value = false
  }
}

const reset = () => {
  searchText.value = ''
  statusFilter.value = ''
  creatorFilter.value = ''
  page.value = 1
  load()
}

const loadCreators = async () => {
  try {
    const { data } = await fetchCreators()
    creators.value = data || []
  } catch {
    // 错误已由 request 拦截器处理
  }
}

const remove = async (row) => {
  try {
    await ElMessageBox.confirm('确认要删除该标签吗？', '提示', { type: 'warning', confirmButtonText: '删除', cancelButtonText: '取消' })
    await deleteNfc(row.id)
    ElMessage.success('已删除')
    await load()
  } catch {
    // cancel or error
  }
}

const batchAction = async (action) => {
  const actionLabels = { enable: '启用', disable: '禁用', delete: '删除' }
  try {
    await ElMessageBox.confirm(`确认要批量${actionLabels[action]} ${selectedIds.value.length} 个标签吗？`, '提示', {
      type: action === 'delete' ? 'warning' : 'info',
      confirmButtonText: '确认',
      cancelButtonText: '取消'
    })
    await batchNfc({ ids: selectedIds.value, action })
    ElMessage.success(`已批量${actionLabels[action]}`)
    await load()
  } catch {
    // cancel or error
  }
}

/** 生成写入 NFC 芯片用的公开链接：当前访问域名 + /p/{nfc_id} */
const publicLink = (row) => `${window.location.origin}/p/${row.nfc_id}`

/** 复制文本到剪贴板，兼容非安全上下文（如局域网 IP 访问）与剪贴板权限受限场景 */
const copyText = async (text) => {
  if (navigator.clipboard && window.isSecureContext) {
    try {
      await navigator.clipboard.writeText(text)
      return true
    } catch {
      // 权限或焦点问题 → 落到 execCommand 兜底
    }
  }
  try {
    const ta = document.createElement('textarea')
    ta.value = text
    ta.setAttribute('readonly', '')
    ta.style.position = 'fixed'
    ta.style.top = '-9999px'
    ta.style.opacity = '0'
    document.body.appendChild(ta)
    ta.select()
    ta.setSelectionRange(0, ta.value.length)
    const ok = document.execCommand('copy')
    document.body.removeChild(ta)
    return ok
  } catch {
    return false
  }
}

const copyLink = async (row) => {
  const ok = await copyText(publicLink(row))
  if (ok) {
    ElMessage.success('链接已复制')
  } else {
    ElMessage.error('复制失败，请手动复制')
  }
}

/** 切换单个标签的开启/关闭状态 */
const toggleStatus = async (row) => {
  const target = row.status === 'enabled' ? 'disabled' : 'enabled'
  togglingId.value = row.id
  try {
    const { data } = await updateNfc(row.id, { status: target })
    row.status = data.status
    ElMessage.success(target === 'enabled' ? '已开启' : '已关闭')
  } catch {
    // 错误已由 request 拦截器处理
  } finally {
    togglingId.value = null
  }
}

onMounted(() => {
  load()
  loadCreators()
})
</script>

<style scoped>
.field {
  display: flex;
  height: 2.25rem;
  width: 100%;
  border-radius: 0.75rem;
  border: 1px solid rgba(20, 24, 29, 0.1);
  background: #fbfaf7;
  padding: 0 0.75rem;
  font-size: 12.5px;
  color: #14181d;
  outline: none;
  transition: border-color 0.18s ease, box-shadow 0.18s ease, background-color 0.18s ease;
}
.field::placeholder { color: rgba(20, 24, 29, 0.28); }
.field:hover { border-color: rgba(20, 24, 29, 0.18); }
.field:focus {
  border-color: #ff6b35;
  background: #fff;
  box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.13);
}
input.field { padding-left: 2.1rem; }

.th {
  height: 2.75rem;
  padding: 0 1rem;
  text-align: left;
  font-size: 10.5px;
  font-weight: 600;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: rgba(20, 24, 29, 0.35);
  white-space: nowrap;
}

.td {
  padding: 0.85rem 1rem;
  vertical-align: middle;
}

.row {
  border-bottom: 1px solid rgba(20, 24, 29, 0.045);
  transition: background-color 0.16s ease;
}
.row:last-child { border-bottom: 0; }
.row:hover { background: rgba(255, 107, 53, 0.028); }

.checkbox {
  height: 14px;
  width: 14px;
  border-radius: 4px;
  accent-color: #ff6b35;
  cursor: pointer;
}

.pg-btn {
  display: inline-flex;
  height: 2rem;
  align-items: center;
  justify-content: center;
  border-radius: 0.6rem;
  border: 1px solid rgba(20, 24, 29, 0.1);
  background: #fff;
  padding: 0 0.8rem;
  font-size: 12px;
  font-weight: 500;
  color: rgba(20, 24, 29, 0.65);
  transition: border-color 0.18s ease, color 0.18s ease;
}
.pg-btn:hover:not(:disabled) { border-color: rgba(20, 24, 29, 0.3); color: #14181d; }
.pg-btn:disabled { opacity: 0.35; cursor: not-allowed; }
</style>
