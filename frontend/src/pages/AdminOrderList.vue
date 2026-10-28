<template>
  <div class="min-h-[calc(100vh-60px)] bg-[#fbfaf7] font-sans text-[#14181d]">
    <div class="mx-auto max-w-[1600px] px-6 py-9">
      <!-- Header -->
      <div class="mb-7 flex flex-col justify-between gap-5 sm:flex-row sm:items-end">
        <div>
          <p class="text-[11px] font-semibold uppercase tracking-[0.2em] text-[#ff6b35]">Order management</p>
          <h1 class="mt-2 text-[30px] font-semibold leading-none tracking-[-0.03em]">订单管理</h1>
          <p class="mt-3 text-[13.5px] leading-relaxed text-[#14181d]/50">
            展厅下单记录 · 共 {{ total }} 笔 · 待处理 {{ pendingCount }} 笔
          </p>
        </div>
        <button
          @click="load"
          class="inline-flex h-10 shrink-0 items-center justify-center gap-2 rounded-xl border border-[#14181d]/10 bg-white px-4 text-[13px] font-medium text-[#14181d]/65 transition-colors hover:border-[#14181d]/25 hover:text-[#14181d]"
        >
          <RotateCw class="h-3.5 w-3.5" :class="loading ? 'animate-spin' : ''" />
          刷新
        </button>
      </div>

      <!-- Filters -->
      <div class="mb-5 rounded-2xl border border-[#14181d]/[0.07] bg-white p-4 shadow-[0_1px_2px_rgba(20,24,29,0.03)]">
        <div class="flex flex-col items-end gap-3 md:flex-row">
          <div class="grid w-full flex-1 grid-cols-1 gap-3 md:grid-cols-3">
            <div class="space-y-1.5">
              <label class="text-[11.5px] font-medium tracking-wide text-[#14181d]/45">订单号 / 第三方订单号 / 收件人 / 电话</label>
              <div class="relative">
                <Search class="pointer-events-none absolute left-3 top-1/2 h-3.5 w-3.5 -translate-y-1/2 text-[#14181d]/30" />
                <input v-model="searchText" class="field" placeholder="搜索…" @keyup.enter="search" />
              </div>
            </div>
            <div class="space-y-1.5">
              <label class="text-[11.5px] font-medium tracking-wide text-[#14181d]/45">状态</label>
              <select v-model="statusFilter" class="field" @change="search">
                <option value="">全部状态</option>
                <option value="pending">待处理</option>
                <option value="converted">已转标签</option>
                <option value="cancelled">已取消</option>
              </select>
            </div>
            <div class="space-y-1.5">
              <label class="text-[11.5px] font-medium tracking-wide text-[#14181d]/45">每页</label>
              <select v-model.number="pageSize" class="field" @change="search">
                <option :value="10">10 条</option>
                <option :value="20">20 条</option>
                <option :value="50">50 条</option>
              </select>
            </div>
          </div>
          <div class="flex w-full gap-2 md:w-auto">
            <button
              @click="search"
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

      <!-- Table -->
      <div class="overflow-hidden rounded-2xl border border-[#14181d]/[0.07] bg-white shadow-[0_1px_2px_rgba(20,24,29,0.03)]">
        <div class="relative w-full overflow-x-auto">
          <table class="w-full caption-bottom text-[13px]">
            <thead>
              <tr class="border-b border-[#14181d]/[0.06] bg-[#fbfaf7]/80">
                <th class="th">订单号</th>
                <th class="th w-[150px]">第三方订单号</th>
                <th class="th">模板</th>
                <th class="th w-[110px]">数量 / 金额</th>
                <th class="th">收件人</th>
                <th class="th">电话</th>
                <th class="th max-w-[260px]">邮寄地址</th>
                <th class="th w-[110px]">状态</th>
                <th class="th w-[150px]">下单时间</th>
                <th class="th w-[230px] text-right">操作</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="(row, index) in items"
                :key="row.id"
                v-motion
                :initial="{ opacity: 0, x: -14 }"
                :enter="{ opacity: 1, x: 0, transition: { delay: index * 45 } }"
                class="row cursor-pointer"
                @click="openDetail(row)"
              >
                <td class="td whitespace-nowrap font-mono text-[12.5px] font-semibold tracking-tight text-[#ff6b35]">
                  {{ row.order_no }}
                </td>
                <td class="td whitespace-nowrap">
                  <span
                    v-if="row.external_order_no"
                    class="inline-flex items-center gap-1 rounded-md bg-[#7c3aed]/[0.08] px-2 py-0.5 font-mono text-[11.5px] font-medium text-[#6d28d9]"
                    :title="'第三方订单 ' + row.external_order_no"
                  >
                    <Link2 class="h-3 w-3" />
                    {{ row.external_order_no }}
                  </span>
                  <span v-else class="text-[12px] text-[#14181d]/25">—</span>
                </td>
                <td class="td">
                  <div class="flex items-center gap-1.5">
                    <span
                      class="flex h-5 w-5 flex-shrink-0 items-center justify-center rounded"
                      :style="{ background: typeMeta(row.template_type).iconBg, color: typeMeta(row.template_type).iconColor }"
                    >
                      <component :is="typeMeta(row.template_type).icon" class="h-3 w-3" />
                    </span>
                    <span class="whitespace-nowrap text-[12px]">{{ typeLabel(row.template_type) }}</span>
                    <span class="text-[#14181d]/25">·</span>
                    <span class="whitespace-nowrap text-[12px] text-[#14181d]/50">{{ styleLabel(row) }}</span>
                  </div>
                </td>
                <td class="td whitespace-nowrap">
                  <span class="text-[12.5px] font-medium">{{ row.quantity }} 枚</span>
                  <span class="ml-1.5 font-medium text-[#ff6b35]">{{ money(row.amount) }}</span>
                </td>
                <td class="td whitespace-nowrap text-[12.5px] font-medium">{{ row.contact_name }}</td>
                <td class="td whitespace-nowrap font-mono text-[12px] text-[#14181d]/60">{{ row.contact_phone }}</td>
                <td class="td max-w-[260px]">
                  <span class="block truncate text-[12px] text-[#14181d]/55" :title="row.contact_address">{{ row.contact_address }}</span>
                </td>
                <td class="td">
                  <span class="badge" :class="statusClass(row.status)">{{ statusLabel(row.status) }}</span>
                </td>
                <td class="td whitespace-nowrap font-mono text-[11.5px] text-[#14181d]/45">{{ fmtTime(row.created_at) }}</td>
                <td class="td text-right" @click.stop>
                  <div class="flex items-center justify-end gap-1.5">
                    <button
                      v-if="row.status === 'pending'"
                      :disabled="convertingId === row.id"
                      class="inline-flex h-8 items-center gap-1 rounded-lg bg-[#ff6b35] px-2.5 text-[11.5px] font-medium text-white transition-colors hover:bg-[#f25a22] disabled:opacity-50"
                      @click="convert(row)"
                    >
                      <Tag class="h-3.5 w-3.5" />
                      {{ convertingId === row.id ? '转换中…' : '转换成标签' }}
                    </button>
                    <span
                      v-else-if="row.status === 'converted'"
                      class="inline-flex h-8 items-center gap-1 rounded-lg bg-[#16a34a]/10 px-2.5 text-[11.5px] font-medium text-[#15803d]"
                      :title="row.nfc_id"
                    >
                      <Check class="h-3.5 w-3.5" />
                      {{ row.nfc_id || '已转换' }}
                    </span>
                    <button
                      class="icon-btn"
                      title="详情"
                      @click="openDetail(row)"
                    >
                      <Eye class="h-3.5 w-3.5" />
                    </button>
                    <button
                      class="icon-btn icon-danger"
                      title="删除"
                      @click="remove(row)"
                    >
                      <Trash2 class="h-3.5 w-3.5" />
                    </button>
                  </div>
                </td>
              </tr>

              <template v-if="loading">
                <tr v-for="i in 5" :key="'sk' + i" class="border-b border-[#14181d]/[0.04]">
                  <td class="px-4 py-4" v-for="j in 10" :key="j">
                    <div class="h-3.5 animate-pulse rounded-md bg-[#14181d]/[0.055]" :style="{ width: j === 5 ? '70%' : '100%' }"></div>
                  </td>
                </tr>
              </template>
              <tr v-else-if="items.length === 0">
                <td colspan="10" class="px-4 py-16 text-center">
                  <div class="mx-auto flex h-12 w-12 items-center justify-center rounded-2xl bg-[#14181d]/[0.04]">
                    <ShoppingBag class="h-5 w-5 text-[#14181d]/25" />
                  </div>
                  <p class="mt-3 text-[13px] font-medium text-[#14181d]/60">还没有订单</p>
                  <p class="mt-1 text-[12px] text-[#14181d]/35">客户在展厅点击「获取硬件」下单后会出现在这里</p>
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

    <!-- 详情抽屉 -->
    <transition name="drawer">
      <div v-if="detail" class="fixed inset-0 z-[60] flex justify-end">
        <div class="absolute inset-0 bg-[#14181d]/30 backdrop-blur-[2px]" @click="detail = null"></div>
        <aside class="relative flex h-full w-full max-w-[480px] flex-col bg-white shadow-[-20px_0_60px_-30px_rgba(20,24,29,0.45)]">
          <header class="flex items-start justify-between gap-3 border-b border-[#14181d]/[0.07] px-5 py-4">
            <div>
              <p class="text-[11px] font-semibold uppercase tracking-[0.18em] text-[#14181d]/35">Order detail</p>
              <h3 class="mt-1 font-mono text-[18px] font-semibold tracking-tight text-[#ff6b35]">{{ detail.order_no }}</h3>
            </div>
            <button class="icon-btn" @click="detail = null">
              <X class="h-4 w-4" />
            </button>
          </header>

          <div class="flex-1 space-y-5 overflow-y-auto px-5 py-5">
            <div class="flex items-center justify-between rounded-xl bg-[#fbfaf7] px-3.5 py-3">
              <span class="text-[12px] text-[#14181d]/45">状态</span>
              <span class="badge" :class="statusClass(detail.status)">{{ statusLabel(detail.status) }}</span>
            </div>

            <section>
              <h4 class="detail-title">模板</h4>
              <div class="mt-2 flex items-center gap-2 rounded-xl border border-[#14181d]/[0.07] px-3.5 py-3">
                <span
                  class="flex h-8 w-8 items-center justify-center rounded-lg"
                  :style="{ background: typeMeta(detail.template_type).iconBg, color: typeMeta(detail.template_type).iconColor }"
                >
                  <component :is="typeMeta(detail.template_type).icon" class="h-4 w-4" />
                </span>
                <div class="min-w-0">
                  <p class="text-[13px] font-medium">{{ typeLabel(detail.template_type) }}</p>
                  <p class="text-[11.5px] text-[#14181d]/45">{{ styleLabel(detail) }}</p>
                  <p v-if="detail.tag_name" class="mt-0.5 truncate text-[11px] text-[#14181d]/35">标签名：{{ detail.tag_name }}</p>
                </div>
                <a
                  :href="previewHref(detail)"
                  target="_blank"
                  rel="noopener"
                  class="ml-auto inline-flex h-7 flex-shrink-0 items-center gap-1 rounded-lg border border-[#14181d]/10 px-2 text-[11px] font-medium text-[#14181d]/55 transition-colors hover:border-[#ff6b35] hover:text-[#ff6b35]"
                >
                  <ExternalLink class="h-3 w-3" />
                  预览
                </a>
              </div>
            </section>

            <section v-if="detail.external_order_no">
              <h4 class="detail-title">第三方订单</h4>
              <div class="mt-2 flex items-center justify-between gap-3 rounded-xl bg-[#7c3aed]/[0.06] px-3.5 py-3">
                <div class="min-w-0">
                  <p class="truncate font-mono text-[13px] font-medium text-[#6d28d9]">{{ detail.external_order_no }}</p>
                  <p class="mt-0.5 text-[11px] text-[#6d28d9]/70">数量与金额来自该订单</p>
                </div>
                <Link2 class="h-4 w-4 flex-shrink-0 text-[#6d28d9]/60" />
              </div>
            </section>

            <section>
              <h4 class="detail-title">数量与费用</h4>
              <div class="mt-2 flex items-center justify-between rounded-xl border border-[#14181d]/[0.07] px-3.5 py-3">
                <div>
                  <p class="text-[12px] text-[#14181d]/45">共 {{ detail.quantity }} 枚 · 单价约 {{ unitMoney(detail.amount, detail.quantity) }}</p>
                  <p class="mt-0.5 text-[20px] font-semibold tracking-[-0.02em] text-[#ff6b35]">{{ money(detail.amount) }}</p>
                </div>
                <span class="rounded-full bg-[#14181d]/[0.05] px-2.5 py-1 text-[11px] text-[#14181d]/45">支付未接入</span>
              </div>
            </section>

            <section>
              <h4 class="detail-title">邮寄信息</h4>
              <dl class="mt-2 divide-y divide-[#14181d]/[0.06] rounded-xl border border-[#14181d]/[0.07] text-[12.5px]">
                <div class="flex gap-3 px-3.5 py-2.5">
                  <dt class="w-16 flex-shrink-0 text-[#14181d]/40">收件人</dt>
                  <dd class="font-medium">{{ detail.contact_name }}</dd>
                </div>
                <div class="flex gap-3 px-3.5 py-2.5">
                  <dt class="w-16 flex-shrink-0 text-[#14181d]/40">电话</dt>
                  <dd class="font-mono">{{ detail.contact_phone }}</dd>
                </div>
                <div class="flex gap-3 px-3.5 py-2.5">
                  <dt class="w-16 flex-shrink-0 text-[#14181d]/40">地址</dt>
                  <dd class="leading-relaxed">{{ detail.contact_address }}</dd>
                </div>
                <div v-if="detail.contact_zipcode" class="flex gap-3 px-3.5 py-2.5">
                  <dt class="w-16 flex-shrink-0 text-[#14181d]/40">邮编</dt>
                  <dd class="font-mono">{{ detail.contact_zipcode }}</dd>
                </div>
                <div v-if="detail.contact_note" class="flex gap-3 px-3.5 py-2.5">
                  <dt class="w-16 flex-shrink-0 text-[#14181d]/40">备注</dt>
                  <dd class="leading-relaxed">{{ detail.contact_note }}</dd>
                </div>
              </dl>
            </section>

            <section>
              <h4 class="detail-title">展示内容</h4>
              <pre class="mt-2 max-h-[280px] overflow-auto rounded-xl border border-[#14181d]/[0.07] bg-[#fbfaf7] p-3 font-mono text-[11px] leading-relaxed text-[#14181d]/65">{{ prettyPayload(detail) }}</pre>
            </section>

            <section v-if="detail.status === 'converted'">
              <h4 class="detail-title">已生成的标签</h4>
              <div class="mt-2 flex items-center justify-between rounded-xl bg-[#16a34a]/[0.07] px-3.5 py-3">
                <div>
                  <p class="font-mono text-[12.5px] font-semibold text-[#15803d]">{{ detail.nfc_id }}</p>
                  <p class="mt-0.5 text-[11px] text-[#15803d]/70">标签 #{{ detail.nfc_tag_id }}</p>
                </div>
                <router-link
                  to="/admin/nfc"
                  class="inline-flex h-7 items-center rounded-lg bg-white px-2.5 text-[11px] font-medium text-[#15803d]"
                >
                  去标签列表
                </router-link>
              </div>
            </section>
          </div>

          <footer class="border-t border-[#14181d]/[0.07] px-5 py-4">
            <div v-if="detail.status === 'pending'" class="space-y-3">
              <div class="grid grid-cols-2 gap-3">
                <div class="space-y-1.5">
                  <label class="text-[11.5px] font-medium text-[#14181d]/45">NFC ID（留空自动生成）</label>
                  <input v-model="convertForm.nfc_id" class="field" :placeholder="detail.order_no" />
                </div>
                <div class="space-y-1.5">
                  <label class="text-[11.5px] font-medium text-[#14181d]/45">标签名（留空自动生成）</label>
                  <input v-model="convertForm.name" class="field" :placeholder="detail.tag_name || '按订单内容生成'" />
                </div>
              </div>
              <button
                :disabled="convertingId === detail.id"
                class="inline-flex h-10 w-full items-center justify-center gap-2 rounded-xl bg-[#ff6b35] text-[13px] font-semibold text-white transition-colors hover:bg-[#f25a22] disabled:opacity-50"
                @click="convert(detail)"
              >
                <Tag class="h-4 w-4" />
                {{ convertingId === detail.id ? '转换中…' : '转换成标签' }}
              </button>
            </div>
            <p v-else class="text-center text-[12px] text-[#14181d]/35">该订单已处理完成</p>
          </footer>
        </aside>
      </div>
    </transition>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import {
  Search, RotateCw, Tag, Check, Eye, Trash2, X, ExternalLink, ShoppingBag, Link2,
} from 'lucide-vue-next'
import { ElMessage, ElMessageBox } from 'element-plus'
import { fetchOrderList, convertOrder, deleteOrder } from '../api/orders'
import { templateTypes } from '../templates/registry'
import { money, unitMoney } from '../config/pricing'

const items = ref([])
const loading = ref(false)
const searchText = ref('')
const statusFilter = ref('')
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const totalPages = ref(1)
const pendingCount = ref(0)

const detail = ref(null)
const convertingId = ref(null)
const convertForm = reactive({ nfc_id: '', name: '' })

const typeMeta = (key) => templateTypes[key] || { label: key, icon: ShoppingBag, iconBg: '#f1f5f9', iconColor: '#64748b' }
const typeLabel = (key) => typeMeta(key).label
const styleLabel = (row) => {
  const styles = templateTypes[row.template_type]?.styles || []
  const hit = styles.find((s) => s.value === (row.custom_template || ''))
  return hit?.label || styles[0]?.label || '默认'
}

const STATUS = {
  pending: { text: '待处理', cls: 'bg-[#ff6b35]/10 text-[#c2410c]' },
  converted: { text: '已转标签', cls: 'bg-[#16a34a]/10 text-[#15803d]' },
  cancelled: { text: '已取消', cls: 'bg-[#14181d]/[0.055] text-[#14181d]/40' },
}
const statusLabel = (s) => STATUS[s]?.text || s
const statusClass = (s) => STATUS[s]?.cls || ''

/** 后端时间为 UTC naive 字符串，转成本地时间显示 */
const fmtTime = (v) => {
  if (!v) return '—'
  const d = new Date(String(v).endsWith('Z') ? v : v + 'Z')
  if (Number.isNaN(d.getTime())) return v
  const p = (n) => String(n).padStart(2, '0')
  return `${d.getFullYear()}-${p(d.getMonth() + 1)}-${p(d.getDate())} ${p(d.getHours())}:${p(d.getMinutes())}`
}

const previewHref = (row) => {
  const q = new URLSearchParams({ type: row.template_type, style: row.custom_template || '', nfcId: row.nfc_id || 'PREVIEW' })
  return `/preview?${q.toString()}`
}

const prettyPayload = (row) => {
  try {
    return JSON.stringify(row.template_payload || {}, null, 2)
  } catch {
    return '—'
  }
}

const load = async () => {
  loading.value = true
  try {
    const params = { page: page.value, page_size: pageSize.value }
    if (searchText.value) params.search = searchText.value
    if (statusFilter.value) params.status = statusFilter.value
    const { data } = await fetchOrderList(params)
    items.value = data.items
    total.value = data.total
    totalPages.value = data.total_pages || 1
    // 单独取一次待处理总量，用于页头提示
    try {
      const { data: pending } = await fetchOrderList({ page: 1, page_size: 1, status: 'pending' })
      pendingCount.value = pending.total
    } catch {
      pendingCount.value = 0
    }
  } catch {
    // 错误已由 request 拦截器处理
  } finally {
    loading.value = false
  }
}

const search = () => {
  page.value = 1
  load()
}

const reset = () => {
  searchText.value = ''
  statusFilter.value = ''
  page.value = 1
  load()
}

const openDetail = (row) => {
  detail.value = row
  convertForm.nfc_id = ''
  convertForm.name = ''
}

const convert = async (row) => {
  if (convertingId.value) return
  convertingId.value = row.id
  try {
    const payload = {}
    if (convertForm.nfc_id.trim()) payload.nfc_id = convertForm.nfc_id.trim()
    if (convertForm.name.trim()) payload.name = convertForm.name.trim()
    const { data } = await convertOrder(row.id, payload)
    ElMessage.success(`已生成标签：${data.nfc_id}`)
    if (detail.value && detail.value.id === row.id) {
      detail.value.status = 'converted'
      detail.value.nfc_id = data.nfc_id
      detail.value.nfc_tag_id = data.nfc_tag_id
    }
    await load()
  } catch {
    // 错误已由 request 拦截器处理
  } finally {
    convertingId.value = null
  }
}

const remove = async (row) => {
  try {
    await ElMessageBox.confirm(
      row.status === 'converted'
        ? '删除订单不会影响已生成的标签，确认删除？'
        : '确认删除这笔订单吗？',
      '提示',
      { type: 'warning', confirmButtonText: '删除', cancelButtonText: '取消' }
    )
    await deleteOrder(row.id)
    ElMessage.success('已删除')
    if (detail.value && detail.value.id === row.id) detail.value = null
    await load()
  } catch {
    // cancel or error
  }
}

onMounted(load)
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
input.field { padding-left: 0.75rem; }
div.relative input.field { padding-left: 2.1rem; }

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

.td { padding: 0.8rem 1rem; vertical-align: middle; }

.row {
  border-bottom: 1px solid rgba(20, 24, 29, 0.045);
  transition: background-color 0.16s ease;
}
.row:last-child { border-bottom: 0; }
.row:hover { background: rgba(255, 107, 53, 0.028); }

.badge {
  display: inline-flex;
  align-items: center;
  gap: 0.35rem;
  border-radius: 999px;
  padding: 0.25rem 0.6rem;
  font-size: 11.5px;
  font-weight: 500;
}

.icon-btn {
  display: inline-flex;
  height: 2rem;
  width: 2rem;
  align-items: center;
  justify-content: center;
  border-radius: 0.5rem;
  color: rgba(20, 24, 29, 0.5);
  transition: background-color 0.16s ease, color 0.16s ease;
}
.icon-btn:hover { background: rgba(20, 24, 29, 0.05); color: #14181d; }
.icon-danger:hover { background: rgba(220, 38, 38, 0.1); color: #b91c1c; }

.detail-title {
  font-size: 11px;
  font-weight: 600;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: rgba(20, 24, 29, 0.35);
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

.drawer-enter-active,
.drawer-leave-active { transition: opacity 0.22s ease; }
.drawer-enter-from,
.drawer-leave-to { opacity: 0; }
.drawer-enter-active aside,
.drawer-leave-active aside { transition: transform 0.28s cubic-bezier(0.22, 1, 0.36, 1); }
.drawer-enter-from aside,
.drawer-leave-to aside { transform: translateX(28px); }
</style>
