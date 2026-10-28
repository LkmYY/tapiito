<template>
  <div class="af-page">
    <div class="af-shell">
      <header class="af-head">
        <div>
          <p class="af-eyebrow">Feedback management</p>
          <h1>用户意见</h1>
          <p class="af-sub">
            共 {{ stats.total }} 条意见，待查看 {{ stats.pending }} 条，已采纳 {{ stats.adopted }} 条。
          </p>
        </div>
        <button type="button" class="af-refresh" :disabled="loading" @click="load">
          <RotateCw class="h-4 w-4" :class="{ 'af-spin': loading }" />
          刷新
        </button>
      </header>

      <div class="af-stats">
        <div class="af-stat">
          <span>全部意见</span>
          <strong>{{ stats.total }}</strong>
        </div>
        <div class="af-stat af-stat-wait">
          <span>待查看</span>
          <strong>{{ stats.pending }}</strong>
        </div>
        <div class="af-stat af-stat-ok">
          <span>已采纳</span>
          <strong>{{ stats.adopted }}</strong>
        </div>
      </div>

      <section class="af-card">
        <div class="af-filters">
          <label class="af-field af-field-search">
            <span>搜索</span>
            <div class="af-input-wrap">
              <Search class="af-input-icon h-4 w-4" />
              <input
                v-model="filters.search"
                type="text"
                placeholder="搜索意见内容或订单号"
                @keyup.enter="search"
              />
            </div>
          </label>

          <label class="af-field">
            <span>状态</span>
            <select v-model="filters.status" @change="search">
              <option value="">全部状态</option>
              <option value="pending">待查看</option>
              <option value="adopted">已采纳</option>
            </select>
          </label>

          <label class="af-field af-field-size">
            <span>每页</span>
            <select v-model.number="filters.page_size" @change="search">
              <option :value="10">10 条</option>
              <option :value="20">20 条</option>
              <option :value="50">50 条</option>
            </select>
          </label>

          <div class="af-filter-actions">
            <button type="button" class="af-btn af-btn-dark" @click="search">查询</button>
            <button type="button" class="af-btn af-btn-light" @click="reset">重置</button>
          </div>
        </div>

        <div class="af-table-wrap">
          <table class="af-table">
            <thead>
              <tr>
                <th class="af-col-check">采纳</th>
                <th>意见内容</th>
                <th class="af-col-order">订单号</th>
                <th class="af-col-status">状态</th>
                <th class="af-col-time">提交时间</th>
              </tr>
            </thead>
            <tbody>
              <template v-if="loading">
                <tr v-for="n in 5" :key="`sk-${n}`" class="af-row af-row-skeleton">
                  <td><span class="af-sk af-sk-check"></span></td>
                  <td><span class="af-sk af-sk-long"></span><span class="af-sk af-sk-short"></span></td>
                  <td><span class="af-sk af-sk-mid"></span></td>
                  <td><span class="af-sk af-sk-pill"></span></td>
                  <td><span class="af-sk af-sk-mid"></span></td>
                </tr>
              </template>

              <tr v-else v-for="item in rows" :key="item.id" class="af-row">
                <td data-label="采纳" class="af-col-check">
                  <label class="af-check" :title="item.status === 'adopted' ? '取消采纳' : '标记为已采纳'">
                    <input
                      type="checkbox"
                      :checked="item.status === 'adopted'"
                      :disabled="updatingId !== null"
                      @change="toggleAdopt(item, $event)"
                    />
                    <span class="af-check-box">
                      <Loader2 v-if="updatingId === item.id" class="h-3 w-3 af-spin" />
                      <Check v-else class="h-3 w-3" />
                    </span>
                  </label>
                </td>

                <td data-label="意见内容" class="af-content-cell">
                  <p class="af-content">{{ item.content }}</p>
                  <p v-if="item.status === 'adopted' && item.adopted_at" class="af-adopted-at">
                    <Check class="h-3 w-3" />
                    采纳于 {{ fmtTime(item.adopted_at) }}
                  </p>
                </td>

                <td data-label="订单号" class="af-col-order">
                  <span v-if="item.order_no" class="af-order">{{ item.order_no }}</span>
                  <span v-else class="af-muted">未填写</span>
                </td>

                <td data-label="状态" class="af-col-status">
                  <span class="af-badge" :class="item.status === 'adopted' ? 'af-badge-ok' : 'af-badge-wait'">
                    {{ item.status === 'adopted' ? '已采纳' : '待查看' }}
                  </span>
                </td>

                <td data-label="提交时间" class="af-col-time">
                  <span class="af-time">{{ fmtTime(item.created_at) }}</span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <div v-if="!loading && !rows.length" class="af-empty">
          <span class="af-empty-icon"><Inbox class="h-6 w-6" /></span>
          <h3>暂时没有符合条件的意见</h3>
          <p>{{ filters.search || filters.status ? '换个筛选条件试试。' : '用户提交意见后会显示在这里。' }}</p>
        </div>

        <footer v-if="!loading && total" class="af-pager">
          <span>第 {{ filters.page }} / {{ totalPages }} 页 · 共 {{ total }} 条</span>
          <div class="af-pager-btns">
            <button
              type="button"
              class="af-pg"
              :disabled="filters.page <= 1"
              @click="go(filters.page - 1)"
            >
              上一页
            </button>
            <button
              type="button"
              class="af-pg"
              :disabled="filters.page >= totalPages"
              @click="go(filters.page + 1)"
            >
              下一页
            </button>
          </div>
        </footer>
      </section>
    </div>
  </div>
</template>

<script setup>
import { onMounted, onUnmounted, reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Check, Inbox, Loader2, RotateCw, Search } from 'lucide-vue-next'
import { adoptFeedback, fetchFeedbackList } from '../../api/feedback'

const rows = ref([])
const loading = ref(false)
const updatingId = ref(null)
const total = ref(0)
const totalPages = ref(1)
const stats = reactive({ total: 0, pending: 0, adopted: 0 })

const filters = reactive({
  page: 1,
  page_size: 10,
  search: '',
  status: '',
})

const fmtTime = (value) => {
  if (!value) return '—'
  const raw = String(value)
  const normalized = /Z$|[+-]\d{2}:?\d{2}$/.test(raw) ? raw : `${raw}Z`
  const date = new Date(normalized)
  if (Number.isNaN(date.getTime())) return raw.replace('T', ' ').slice(0, 16)
  const pad = (n) => String(n).padStart(2, '0')
  return `${date.getFullYear()}-${pad(date.getMonth() + 1)}-${pad(date.getDate())} ${pad(
    date.getHours()
  )}:${pad(date.getMinutes())}`
}

const load = async () => {
  loading.value = true
  try {
    const { data } = await fetchFeedbackList({
      page: filters.page,
      page_size: filters.page_size,
      search: filters.search.trim() || undefined,
      status: filters.status || undefined,
    })
    rows.value = data.items || []
    total.value = data.total || 0
    totalPages.value = data.total_pages || 1
    stats.total = data.total || 0
    stats.pending = data.pending_total || 0
    stats.adopted = data.adopted_total || 0
  } catch {
    rows.value = []
  } finally {
    loading.value = false
  }
}

const search = () => {
  filters.page = 1
  load()
}

const reset = () => {
  filters.page = 1
  filters.search = ''
  filters.status = ''
  filters.page_size = 10
  load()
}

const go = (page) => {
  if (page < 1 || page > totalPages.value) return
  filters.page = page
  load()
}

const toggleAdopt = async (item, event) => {
  if (updatingId.value !== null) return
  const next = event.target.checked
  updatingId.value = item.id
  try {
    const { data } = await adoptFeedback(item.id, next)
    item.status = data.status
    item.adopted_at = data.adopted_at
    stats.adopted += next ? 1 : -1
    stats.pending += next ? -1 : 1
    ElMessage.success(next ? '已采纳这条意见' : '已取消采纳')
    if (filters.status && (filters.status === 'adopted') !== next) {
      await load()
    }
  } catch {
    event.target.checked = item.status === 'adopted'
  } finally {
    updatingId.value = null
  }
}

onMounted(() => {
  document.title = '用户意见 · Tapiito'
  load()
})

onUnmounted(() => {
  document.title = 'Tapiito'
})
</script>

<style scoped>
.af-page {
  min-height: calc(100vh - 60px);
  background: #fbfaf7;
  color: #14181d;
  font-family: 'Geist', -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}

.af-shell {
  width: min(1440px, 100%);
  margin: 0 auto;
  padding: 36px 24px 48px;
}

.af-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 20px;
}

.af-eyebrow {
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.18em;
  text-transform: uppercase;
  color: #ff6b35;
}

.af-head h1 {
  margin-top: 6px;
  font-size: 30px;
  font-weight: 650;
  letter-spacing: -0.035em;
}

.af-sub {
  margin-top: 7px;
  font-size: 12.5px;
  color: rgba(20, 24, 29, 0.45);
}

.af-refresh {
  display: inline-flex;
  flex: 0 0 auto;
  align-items: center;
  gap: 7px;
  height: 36px;
  border: 1px solid rgba(20, 24, 29, 0.12);
  border-radius: 10px;
  background: #fff;
  padding: 0 14px;
  font-family: inherit;
  font-size: 12.5px;
  font-weight: 550;
  color: rgba(20, 24, 29, 0.7);
  cursor: pointer;
  transition: border-color 0.2s ease, color 0.2s ease;
}

.af-refresh:hover:not(:disabled) {
  border-color: rgba(20, 24, 29, 0.3);
  color: #14181d;
}

.af-refresh:disabled {
  cursor: wait;
  opacity: 0.6;
}

.af-stats {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 12px;
  margin-top: 22px;
}

.af-stat {
  border: 1px solid rgba(20, 24, 29, 0.08);
  border-radius: 14px;
  background: #fff;
  padding: 14px 16px;
}

.af-stat span {
  font-size: 11px;
  color: rgba(20, 24, 29, 0.42);
}

.af-stat strong {
  display: block;
  margin-top: 4px;
  font-size: 24px;
  font-weight: 650;
  letter-spacing: -0.03em;
  font-variant-numeric: tabular-nums;
}

.af-stat-wait strong {
  color: #d97706;
}

.af-stat-ok strong {
  color: #15803d;
}

.af-card {
  margin-top: 16px;
  border: 1px solid rgba(20, 24, 29, 0.08);
  border-radius: 18px;
  background: #fff;
  box-shadow: 0 18px 48px -42px rgba(20, 24, 29, 0.35);
  overflow: hidden;
}

.af-filters {
  display: flex;
  flex-wrap: wrap;
  align-items: flex-end;
  gap: 12px;
  border-bottom: 1px solid rgba(20, 24, 29, 0.07);
  background: rgba(251, 250, 247, 0.6);
  padding: 16px;
}

.af-field {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.af-field > span {
  font-size: 10.5px;
  font-weight: 600;
  color: rgba(20, 24, 29, 0.48);
}

.af-field-search {
  flex: 1 1 260px;
  max-width: 380px;
}

.af-field-size {
  width: 92px;
}

.af-input-wrap {
  position: relative;
}

.af-input-icon {
  position: absolute;
  top: 50%;
  left: 11px;
  color: rgba(20, 24, 29, 0.3);
  transform: translateY(-50%);
  pointer-events: none;
}

.af-field input,
.af-field select {
  height: 36px;
  width: 100%;
  border: 1px solid rgba(20, 24, 29, 0.1);
  border-radius: 10px;
  background: #fbfaf7;
  padding: 0 11px;
  font-family: inherit;
  font-size: 12.5px;
  color: #14181d;
  outline: none;
}

.af-field input {
  padding-left: 34px;
}

.af-field input:focus,
.af-field select:focus {
  border-color: rgba(255, 107, 53, 0.55);
  background: #fff;
  box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.08);
}

.af-filter-actions {
  display: flex;
  gap: 8px;
}

.af-btn {
  height: 36px;
  border-radius: 10px;
  padding: 0 16px;
  font-family: inherit;
  font-size: 12.5px;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s ease, border-color 0.2s ease;
}

.af-btn-dark {
  border: 1px solid #14181d;
  background: #14181d;
  color: #fff;
}

.af-btn-dark:hover {
  background: #ff6b35;
  border-color: #ff6b35;
}

.af-btn-light {
  border: 1px solid rgba(20, 24, 29, 0.12);
  background: #fff;
  color: rgba(20, 24, 29, 0.6);
}

.af-btn-light:hover {
  border-color: rgba(20, 24, 29, 0.28);
  color: #14181d;
}

.af-table-wrap {
  width: 100%;
  overflow-x: auto;
}

.af-table {
  width: 100%;
  border-collapse: collapse;
  table-layout: fixed;
}

.af-th,
.af-table th {
  padding: 11px 14px;
  border-bottom: 1px solid rgba(20, 24, 29, 0.07);
  background: rgba(251, 250, 247, 0.5);
  text-align: left;
  font-size: 10.5px;
  font-weight: 650;
  letter-spacing: 0.04em;
  color: rgba(20, 24, 29, 0.42);
  white-space: nowrap;
}

.af-table td {
  padding: 15px 14px;
  border-bottom: 1px solid rgba(20, 24, 29, 0.055);
  vertical-align: top;
}

.af-table tbody tr:last-child td {
  border-bottom: 0;
}

.af-row {
  transition: background 0.18s ease;
}

.af-row:hover {
  background: rgba(255, 107, 53, 0.028);
}

.af-col-check {
  width: 66px;
}

.af-col-order {
  width: 150px;
}

.af-col-status {
  width: 94px;
}

.af-col-time {
  width: 140px;
}

.af-check {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  cursor: pointer;
  user-select: none;
}

.af-check input {
  position: absolute;
  width: 1px;
  height: 1px;
  opacity: 0;
  pointer-events: none;
}

.af-check-box {
  display: inline-flex;
  width: 22px;
  height: 22px;
  align-items: center;
  justify-content: center;
  border: 1.5px solid rgba(20, 24, 29, 0.18);
  border-radius: 7px;
  background: #fff;
  color: transparent;
  transition: background 0.18s ease, border-color 0.18s ease, color 0.18s ease;
}

.af-check:hover .af-check-box {
  border-color: rgba(255, 107, 53, 0.6);
}

.af-check input:checked + .af-check-box {
  border-color: #15803d;
  background: #15803d;
  color: #fff;
}

.af-check input:focus-visible + .af-check-box {
  box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.16);
}

.af-content {
  font-size: 13px;
  line-height: 1.75;
  color: rgba(20, 24, 29, 0.76);
  white-space: pre-wrap;
  overflow-wrap: anywhere;
}

.af-adopted-at {
  display: flex;
  align-items: center;
  gap: 4px;
  margin-top: 7px;
  font-size: 10.5px;
  color: #15803d;
}

.af-order {
  display: inline-block;
  max-width: 100%;
  border-radius: 6px;
  background: rgba(20, 24, 29, 0.05);
  padding: 3px 7px;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 10.5px;
  color: rgba(20, 24, 29, 0.62);
  overflow-wrap: anywhere;
}

.af-muted {
  font-size: 11px;
  color: rgba(20, 24, 29, 0.28);
}

.af-badge {
  display: inline-flex;
  align-items: center;
  border-radius: 999px;
  padding: 4px 9px;
  font-size: 10px;
  font-weight: 650;
  white-space: nowrap;
}

.af-badge-wait {
  background: rgba(217, 119, 6, 0.09);
  color: #b45309;
}

.af-badge-ok {
  background: rgba(22, 163, 74, 0.09);
  color: #15803d;
}

.af-time {
  font-size: 10.5px;
  color: rgba(20, 24, 29, 0.4);
  font-variant-numeric: tabular-nums;
  white-space: nowrap;
}

.af-row-skeleton .af-sk {
  display: block;
  height: 10px;
  margin-bottom: 7px;
  border-radius: 999px;
  background: linear-gradient(90deg, rgba(20, 24, 29, 0.05), rgba(20, 24, 29, 0.1), rgba(20, 24, 29, 0.05));
  background-size: 200% 100%;
  animation: af-shimmer 1.4s ease infinite;
}

.af-sk-check {
  width: 22px;
  height: 22px !important;
  border-radius: 7px !important;
}

.af-sk-long {
  width: 82%;
}

.af-sk-short {
  width: 54%;
}

.af-sk-mid {
  width: 70%;
}

.af-sk-pill {
  width: 54px;
  height: 20px !important;
  border-radius: 999px !important;
}

.af-empty {
  border-top: 1px solid rgba(20, 24, 29, 0.06);
  padding: 52px 24px;
  text-align: center;
}

.af-empty-icon {
  display: inline-flex;
  width: 48px;
  height: 48px;
  align-items: center;
  justify-content: center;
  border-radius: 14px;
  background: rgba(20, 24, 29, 0.05);
  color: rgba(20, 24, 29, 0.35);
}

.af-empty h3 {
  margin-top: 14px;
  font-size: 16px;
  font-weight: 600;
}

.af-empty p {
  margin-top: 5px;
  font-size: 12px;
  color: rgba(20, 24, 29, 0.4);
}

.af-pager {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  border-top: 1px solid rgba(20, 24, 29, 0.07);
  padding: 13px 16px;
  font-size: 11px;
  color: rgba(20, 24, 29, 0.42);
}

.af-pager-btns {
  display: flex;
  gap: 7px;
}

.af-pg {
  height: 30px;
  border: 1px solid rgba(20, 24, 29, 0.1);
  border-radius: 8px;
  background: #fff;
  padding: 0 11px;
  font-family: inherit;
  font-size: 11px;
  color: rgba(20, 24, 29, 0.58);
  cursor: pointer;
}

.af-pg:hover:not(:disabled) {
  border-color: rgba(20, 24, 29, 0.26);
  color: #14181d;
}

.af-pg:disabled {
  cursor: not-allowed;
  opacity: 0.4;
}

.af-spin {
  animation: af-spin 0.9s linear infinite;
}

@keyframes af-spin {
  to {
    transform: rotate(360deg);
  }
}

@keyframes af-shimmer {
  from {
    background-position: 200% 0;
  }
  to {
    background-position: -200% 0;
  }
}

@media (max-width: 820px) {
  .af-shell {
    padding: 24px 16px 40px;
  }

  .af-head h1 {
    font-size: 26px;
  }

  .af-stats {
    gap: 8px;
  }

  .af-stat {
    padding: 11px 12px;
  }

  .af-stat strong {
    font-size: 20px;
  }

  .af-filters {
    gap: 10px;
    padding: 13px;
  }

  .af-field-search {
    flex-basis: 100%;
    max-width: none;
  }

  .af-table thead {
    display: none;
  }

  .af-table,
  .af-table tbody,
  .af-table tr,
  .af-table td {
    display: block;
    width: 100%;
  }

  .af-table {
    table-layout: auto;
  }

  .af-row {
    position: relative;
    border-bottom: 1px solid rgba(20, 24, 29, 0.07);
    padding: 14px 15px 12px 52px;
  }

  .af-row:last-child {
    border-bottom: 0;
  }

  .af-table td {
    border: 0;
    padding: 0;
  }

  .af-table td[data-label]::before {
    display: block;
    margin-bottom: 3px;
    font-size: 9.5px;
    font-weight: 650;
    letter-spacing: 0.04em;
    color: rgba(20, 24, 29, 0.34);
    content: attr(data-label);
  }

  .af-table .af-col-check {
    position: absolute;
    top: 14px;
    left: 15px;
    width: 22px;
  }

  .af-table .af-col-check::before {
    display: none;
  }

  .af-table .af-content-cell {
    padding-bottom: 10px;
  }

  .af-table .af-col-order,
  .af-table .af-col-status,
  .af-table .af-col-time {
    display: inline-block;
    width: auto;
    margin-right: 14px;
    vertical-align: top;
  }

  .af-table .af-col-order::before,
  .af-table .af-col-status::before,
  .af-table .af-col-time::before {
    display: none;
  }

  .af-row-skeleton {
    padding-left: 52px;
  }

  .af-row-skeleton td {
    display: block !important;
    width: 100% !important;
    margin-bottom: 7px;
  }

  .af-row-skeleton td:first-child {
    position: absolute;
    top: 14px;
    left: 15px;
    width: 22px !important;
  }

  .af-pager {
    flex-direction: column;
    align-items: flex-start;
  }
}

@media (max-width: 480px) {
  .af-stats {
    grid-template-columns: 1fr 1fr 1fr;
  }

  .af-stat span {
    font-size: 10px;
  }

  .af-stat strong {
    font-size: 18px;
  }

  .af-field {
    flex: 1 1 calc(50% - 5px);
  }

  .af-field-size {
    width: auto;
  }

  .af-filter-actions {
    width: 100%;
  }

  .af-btn {
    flex: 1;
  }
}

@media (prefers-reduced-motion: reduce) {
  .af-spin,
  .af-row-skeleton .af-sk {
    animation: none;
  }
}
</style>
