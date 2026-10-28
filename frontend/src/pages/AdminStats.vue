<template>
  <div class="min-h-[calc(100vh-60px)] bg-[#fbfaf7] font-sans text-[#14181d]">
    <div class="mx-auto max-w-[1600px] px-6 py-9">

      <!-- Page header -->
      <div class="mb-7 flex flex-wrap items-end justify-between gap-4">
        <div>
          <p class="mb-2 text-[11px] font-semibold uppercase tracking-[0.2em] text-[#ff6b35]">Insights</p>
          <h1 class="text-[30px] font-semibold leading-none tracking-[-0.03em]">访问统计</h1>
          <p class="mt-2.5 text-[13px] text-[#14181d]/50">
            每一枚标签的每一次碰触，都记在这里 · 累计 {{ statsOverview.total_visits }} 次
          </p>
        </div>
        <div class="flex items-center gap-2 text-[12px] text-[#14181d]/40">
          <span class="inline-flex h-2 w-2 rounded-full bg-[#22c55e]"></span>
          数据实时更新
        </div>
      </div>

      <!-- Overview cards -->
      <div class="mb-5 grid grid-cols-1 gap-4 md:grid-cols-3">
        <div
          v-for="(card, i) in overviewCards"
          :key="card.key"
          v-motion
          :initial="{ opacity: 0, y: 14 }"
          :enter="{ opacity: 1, y: 0, transition: { duration: 420, delay: i * 70 } }"
          class="card relative overflow-hidden p-5"
        >
          <div class="flex items-start justify-between">
            <span class="text-[12px] font-medium uppercase tracking-[0.12em] text-[#14181d]/40">{{ card.label }}</span>
            <span class="grid h-8 w-8 place-items-center rounded-[10px]" :style="{ background: card.tint, color: card.color }">
              <component :is="card.icon" class="h-4 w-4" :stroke-width="1.8" />
            </span>
          </div>
          <div v-if="statsLoading" class="mt-5 h-9 w-20 animate-pulse rounded-lg bg-[#14181d]/[0.06]"></div>
          <div v-else class="mt-5 flex items-baseline gap-2">
            <span class="text-[34px] font-semibold leading-none tracking-[-0.03em] tabular-nums">{{ card.value }}</span>
            <span class="text-[12px] text-[#14181d]/35">{{ card.unit }}</span>
          </div>
          <div class="mt-4 h-px w-full" :style="{ background: `linear-gradient(90deg, ${card.color}33, transparent)` }"></div>
          <p class="mt-3 text-[12px] text-[#14181d]/40">{{ card.hint }}</p>
        </div>
      </div>

      <!-- Date filter -->
      <div class="card mb-5 p-4">
        <div class="flex flex-col items-stretch gap-3 md:flex-row md:items-end">
          <div class="flex-1">
            <label class="field-label">开始日期</label>
            <input v-model="dateRange.start" type="date" class="field" />
          </div>
          <div class="flex-1">
            <label class="field-label">结束日期</label>
            <input v-model="dateRange.end" type="date" class="field" />
          </div>
          <div class="flex gap-2">
            <button class="btn-dark" @click="loadStats">
              <Search class="h-3.5 w-3.5" :stroke-width="2" />
              查看统计
            </button>
            <button class="btn-ghost" @click="resetDateRange">重置</button>
          </div>
        </div>
      </div>

      <!-- Daily chart -->
      <div class="card mb-5 p-6">
        <div class="mb-5 flex items-end justify-between">
          <div>
            <h2 class="text-[16px] font-semibold tracking-[-0.01em]">每日访问量</h2>
            <p class="mt-1 text-[12px] text-[#14181d]/40">按天统计的碰触次数</p>
          </div>
          <span class="text-[11px] uppercase tracking-[0.14em] text-[#14181d]/30">Daily taps</span>
        </div>

        <div v-if="statsOverview.daily_stats && statsOverview.daily_stats.length > 0" class="h-64">
          <div class="flex h-full items-end gap-3">
            <div v-for="day in statsOverview.daily_stats" :key="day.date" class="flex h-full flex-1 flex-col items-center gap-2">
              <div class="flex w-full flex-1 items-end">
                <div class="group relative h-full w-full">
                  <span
                    class="absolute inset-x-0 text-center text-[11px] font-medium tabular-nums text-[#14181d]/45 transition group-hover:text-[#ff6b35]"
                    :style="{ bottom: `calc(${Math.max(8, (day.visit_count / maxDailyVisits) * 100)}% + 6px)` }"
                  >
                    {{ day.visit_count }}
                  </span>
                  <div
                    class="bar absolute inset-x-0 bottom-0 w-full rounded-t-[7px] transition-all duration-500"
                    :style="{ height: `${Math.max(8, (day.visit_count / maxDailyVisits) * 100)}%` }"
                  ></div>
                </div>
              </div>
              <div class="text-[11px] tabular-nums text-[#14181d]/35">{{ formatDate(day.date) }}</div>
            </div>
          </div>
        </div>
        <div v-else class="flex h-64 flex-col items-center justify-center gap-3 text-[#14181d]/30">
          <BarChart2 class="h-8 w-8" :stroke-width="1.5" />
          <span class="text-[13px]">还没有访问数据</span>
        </div>
      </div>

      <!-- Ranking -->
      <div class="card mb-5 overflow-hidden">
        <div class="flex items-center justify-between border-b border-[#14181d]/[0.06] px-6 py-4">
          <div>
            <h2 class="text-[16px] font-semibold tracking-[-0.01em]">标签访问排名</h2>
            <p class="mt-1 text-[12px] text-[#14181d]/40">哪几枚标签被人碰得最多</p>
          </div>
          <span class="rounded-full bg-[#fbfaf7] px-3 py-1 text-[11px] tabular-nums text-[#14181d]/45">
            {{ statsOverview.top_tags.length }} 枚标签
          </span>
        </div>

        <div class="relative w-full overflow-auto">
          <table class="w-full text-[13px]">
            <thead>
              <tr class="border-b border-[#14181d]/[0.06]">
                <th class="th w-[80px] text-left">排名</th>
                <th class="th text-left">NFC ID</th>
                <th class="th text-left">访问量</th>
                <th class="th w-[140px] text-right">操作</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(tag, index) in statsOverview.top_tags" :key="tag.nfc_id" class="row">
                <td class="px-4 py-3.5">
                  <span
                    class="grid h-7 w-7 place-items-center rounded-[8px] text-[11px] font-semibold tabular-nums"
                    :class="index < 3 ? 'bg-[#ff6b35]/[0.09] text-[#ff6b35]' : 'bg-[#14181d]/[0.04] text-[#14181d]/40'"
                  >
                    {{ String(index + 1).padStart(2, '0') }}
                  </span>
                </td>
                <td class="px-4 py-3.5 font-mono text-[12.5px] tracking-tight">{{ tag.nfc_id }}</td>
                <td class="px-4 py-3.5">
                  <div class="flex items-center gap-3">
                    <span class="w-10 tabular-nums">{{ tag.visit_count }}</span>
                    <div class="h-1 w-28 overflow-hidden rounded-full bg-[#14181d]/[0.05]">
                      <div class="h-full rounded-full bg-[#ff6b35]" :style="{ width: `${(tag.visit_count / maxTagVisits) * 100}%` }"></div>
                    </div>
                  </div>
                </td>
                <td class="px-4 py-3.5 text-right">
                  <button class="btn-row" @click="viewTagStats(tag.nfc_id)">
                    <BarChart2 class="h-3.5 w-3.5" :stroke-width="1.9" />
                    查看详情
                  </button>
                </td>
              </tr>
              <tr v-if="statsOverview.top_tags.length === 0">
                <td colspan="4" class="px-4 py-14 text-center text-[13px] text-[#14181d]/30">暂无数据</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Tag detail -->
      <div v-if="selectedTag" class="card p-6">
        <div class="mb-5 flex items-center justify-between">
          <div>
            <p class="mb-1.5 text-[10.5px] font-semibold uppercase tracking-[0.18em] text-[#ff6b35]">Tag detail</p>
            <h2 class="text-[17px] font-semibold tracking-[-0.01em]">标签 {{ selectedTag }} 访问详情</h2>
          </div>
          <button class="grid h-8 w-8 place-items-center rounded-[10px] text-[#14181d]/40 transition hover:bg-[#14181d]/[0.05] hover:text-[#14181d]" @click="selectedTag = null">
            <X class="h-4 w-4" />
            <span class="sr-only">关闭</span>
          </button>
        </div>

        <div v-if="tagStatsData" class="space-y-6">
          <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
            <div class="rounded-xl bg-[#fbfaf7] p-4">
              <div class="text-[12px] text-[#14181d]/45">总访问量</div>
              <div class="mt-2 text-[26px] font-semibold leading-none tabular-nums">{{ tagStatsData.total_visits }}</div>
            </div>
            <div class="rounded-xl bg-[#fbfaf7] p-4">
              <div class="text-[12px] text-[#14181d]/45">独立 IP 数</div>
              <div class="mt-2 text-[26px] font-semibold leading-none tabular-nums">{{ tagStatsData.unique_ips }}</div>
            </div>
          </div>

          <!-- Device split -->
          <div>
            <h3 class="mb-3 text-[13px] font-medium text-[#14181d]/70">设备类型分布</h3>
            <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
              <div v-for="(count, type) in tagStatsData.device_types" :key="type" class="rounded-xl bg-[#fbfaf7] p-4">
                <div class="flex items-baseline justify-between">
                  <span class="text-[12px] text-[#14181d]/45">{{ type === 'mobile' ? '移动端' : '桌面端' }}</span>
                  <span class="text-[18px] font-semibold tabular-nums">{{ count }}</span>
                </div>
                <div class="mt-3 h-1.5 w-full overflow-hidden rounded-full bg-[#14181d]/[0.06]">
                  <div class="h-full rounded-full bg-[#ff6b35] transition-all duration-500" :style="{ width: `${(count / tagStatsData.total_visits) * 100}%` }"></div>
                </div>
              </div>
            </div>
          </div>

          <!-- Recent visits -->
          <div>
            <h3 class="mb-3 text-[13px] font-medium text-[#14181d]/70">最近访问</h3>
            <div class="-mx-6 overflow-x-auto px-6">
              <table class="w-full min-w-[520px] text-[13px]">
                <thead>
                  <tr class="border-b border-[#14181d]/[0.06]">
                    <th class="th text-left">IP 地址</th>
                    <th class="th text-left">用户代理</th>
                    <th class="th text-left">访问时间</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="visit in tagStatsData.visits.slice(0, 10)" :key="visit.id" class="row">
                    <td class="px-4 py-3 font-mono text-[12px] whitespace-nowrap">{{ visit.ip_address }}</td>
                    <td class="max-w-[200px] truncate px-4 py-3 text-[#14181d]/55">{{ visit.user_agent }}</td>
                    <td class="px-4 py-3 whitespace-nowrap tabular-nums text-[#14181d]/55">{{ formatDateTime(visit.created_at) }}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
        <div v-else class="flex items-center gap-2 py-8 text-[13px] text-[#14181d]/40">
          <span class="h-1.5 w-1.5 animate-pulse rounded-full bg-[#ff6b35]"></span>
          正在加载标签数据…
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { Search, BarChart2, X, Activity, Tag, Eye } from 'lucide-vue-next'
import { ElMessage } from 'element-plus'
import { fetchStatsOverview, fetchTagStats } from '../api/stats'

const statsOverview = ref({ total_visits: 0, top_tags: [], daily_stats: [] })
const statsLoading = ref(false)
const tagStatsData = ref(null)
const selectedTag = ref(null)
const dateRange = reactive({ start: '', end: '' })

const todayVisits = computed(() => {
  const today = new Date().toISOString().split('T')[0]
  return statsOverview.value.daily_stats
    .filter(day => day.date === today)
    .reduce((sum, day) => sum + day.visit_count, 0)
})

const overviewCards = computed(() => [
  { key: 'total', label: '总访问量', value: statsOverview.value.total_visits, unit: '次碰触', hint: '从第一枚标签上线至今', icon: Activity, color: '#ff6b35', tint: 'rgba(255,107,53,0.09)' },
  { key: 'tags', label: '活跃标签', value: statsOverview.value.top_tags.length, unit: '枚', hint: '产生过访问的标签数', icon: Tag, color: '#7bdff2', tint: 'rgba(123,223,242,0.16)' },
  { key: 'today', label: '今日访问', value: todayVisits.value, unit: '次', hint: '今天到目前为止', icon: Eye, color: '#14181d', tint: 'rgba(20,24,29,0.06)' }
])

const maxDailyVisits = computed(() => {
  if (statsOverview.value.daily_stats.length === 0) return 1
  return Math.max(...statsOverview.value.daily_stats.map(day => day.visit_count))
})

const maxTagVisits = computed(() => {
  if (statsOverview.value.top_tags.length === 0) return 1
  return Math.max(...statsOverview.value.top_tags.map(tag => tag.visit_count))
})

const formatDate = (dateStr) => {
  const date = new Date(dateStr)
  return `${date.getMonth() + 1}/${date.getDate()}`
}

const formatDateTime = (dateStr) => new Date(dateStr).toLocaleString()

const resetDateRange = () => {
  dateRange.start = ''
  dateRange.end = ''
  loadStats()
}

const buildParams = () => {
  const params = {}
  if (dateRange.start) params.start_date = dateRange.start
  if (dateRange.end) params.end_date = dateRange.end
  return params
}

const loadStats = async () => {
  statsLoading.value = true
  try {
    const { data } = await fetchStatsOverview(buildParams())
    statsOverview.value = data
  } catch {
    // 错误已由 request 拦截器处理
  } finally {
    statsLoading.value = false
  }
}

const viewTagStats = async (nfcId) => {
  selectedTag.value = nfcId
  try {
    const { data } = await fetchTagStats(nfcId, buildParams())
    tagStatsData.value = data
  } catch {
    ElMessage.error('加载标签统计数据失败')
    selectedTag.value = null
  }
}

onMounted(loadStats)
</script>

<style scoped>
.card {
  @apply rounded-2xl border border-[#14181d]/[0.07] bg-white;
  box-shadow: 0 1px 2px rgba(20, 24, 29, 0.03);
}

.field-label {
  @apply mb-1.5 block text-[11px] font-medium uppercase tracking-[0.12em] text-[#14181d]/40;
}

.field {
  @apply h-9 w-full rounded-xl border border-[#14181d]/10 bg-[#fbfaf7] px-3 text-[13px] text-[#14181d] outline-none transition;
}
.field:focus {
  border-color: #ff6b35;
  box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.14);
}

.btn-dark {
  @apply inline-flex h-9 items-center justify-center gap-1.5 rounded-xl bg-[#14181d] px-4 text-[13px] font-medium text-white transition;
}
.btn-dark:hover {
  background: #ff6b35;
  transform: translateY(-1px);
}

.btn-ghost {
  @apply inline-flex h-9 items-center justify-center rounded-xl border border-[#14181d]/[0.12] bg-white px-4 text-[13px] font-medium text-[#14181d]/70 transition;
}
.btn-ghost:hover {
  border-color: rgba(20, 24, 29, 0.22);
  background: #fbfaf7;
}

.th {
  @apply px-4 py-3 text-[10.5px] font-semibold uppercase tracking-[0.13em] text-[#14181d]/35;
}

.row {
  @apply border-b border-[#14181d]/[0.05] transition-colors;
}
.row:hover {
  background: rgba(255, 107, 53, 0.028);
}

.btn-row {
  @apply inline-flex h-8 items-center gap-1.5 rounded-[10px] px-2.5 text-[12px] font-medium text-[#14181d]/55 transition;
}
.btn-row:hover {
  background: rgba(255, 107, 53, 0.09);
  color: #ff6b35;
}

.bar {
  background: linear-gradient(180deg, #ff8a5c 0%, #ff6b35 55%, #e85a26 100%);
  box-shadow: 0 6px 16px -6px rgba(255, 107, 53, 0.45);
  min-height: 4px;
}
.bar:hover {
  filter: brightness(1.06);
}
</style>
