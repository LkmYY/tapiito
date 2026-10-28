<template>
  <div class="mx-auto max-w-[1440px] p-6">
    <!-- 页头 -->
    <div class="mb-6 flex flex-wrap items-end justify-between gap-4">
      <div>
        <h1 class="text-2xl font-semibold tracking-tight">{{ isEdit ? '编辑标签' : '新建标签' }}</h1>
        <p class="mt-1 text-sm text-muted-foreground">填写内容，右侧实时预览手机上的展示效果</p>
      </div>
      <div class="flex items-center gap-2">
        <button class="btn-ghost" @click="goBack">返回列表</button>
        <button class="btn-primary" :disabled="submitting" @click="submit">
          <Loader2 v-if="submitting" class="h-4 w-4 animate-spin" />
          {{ submitting ? '保存中…' : '保存' }}
        </button>
      </div>
    </div>

    <div v-if="pageLoading" class="grid gap-6 lg:grid-cols-[minmax(0,1fr)_400px]">
      <div class="space-y-4">
        <div v-for="i in 3" :key="i" class="h-40 animate-pulse rounded-xl border bg-muted/50"></div>
      </div>
      <div class="h-[600px] animate-pulse rounded-xl border bg-muted/50"></div>
    </div>

    <div v-else class="grid gap-6 lg:grid-cols-[minmax(0,1fr)_400px]">
      <!-- 左：表单 -->
      <div class="space-y-4">
        <!-- 基本信息 -->
        <section class="card">
          <h2 class="card-title">基本信息</h2>
          <div class="grid gap-4 sm:grid-cols-2">
            <div class="space-y-2 sm:col-span-2">
              <label class="field-label">标签名</label>
              <input v-model="form.name" class="field-input" placeholder="例如：林晓峰的名片、巷口面馆菜单" maxlength="128" />
              <p class="hint">显示在浏览器页签与列表页，方便区分不同标签</p>
            </div>
            <div class="space-y-2">
              <label class="field-label">NFC ID</label>
              <input v-model="form.nfc_id" :disabled="isEdit" class="field-input disabled:opacity-60" placeholder="芯片唯一标识，例如 1111" />
              <p class="hint">写入芯片的地址为 /p/{{ form.nfc_id || '…' }}</p>
            </div>
            <div class="space-y-2">
              <label class="field-label">绑定手机号</label>
              <input v-model="form.phone" class="field-input" placeholder="接收通知的手机号（选填）" />
              <p class="hint">用于异常提醒，不影响展示</p>
            </div>
          </div>
          <div class="mt-4 space-y-2">
            <label class="field-label">状态</label>
            <div class="flex gap-2">
              <button class="chip" :class="form.status === 'enabled' ? 'chip-on' : ''" @click="form.status = 'enabled'">开启</button>
              <button class="chip" :class="form.status === 'disabled' ? 'chip-off' : ''" @click="form.status = 'disabled'">关闭</button>
            </div>
          </div>
        </section>

        <!-- 模板类型 -->
        <section class="card">
          <h2 class="card-title">模板类型</h2>
          <div class="grid grid-cols-2 gap-2.5 sm:grid-cols-4">
            <button
              v-for="key in visibleTypeOrder"
              :key="key"
              class="group flex flex-col items-center gap-2 rounded-xl border-2 p-3 text-[13px] font-medium transition-all"
              :class="form.template_type === key ? 'border-primary bg-primary/[0.04]' : 'border-border hover:border-primary/30 hover:bg-accent/40'"
              @click="selectTemplateType(key)"
            >
              <span
                class="flex h-9 w-9 items-center justify-center rounded-lg transition-transform group-hover:scale-105"
                :style="{ background: templateTypes[key].iconBg, color: templateTypes[key].iconColor }"
              >
                <component :is="templateTypes[key].icon" class="h-[18px] w-[18px]" />
              </span>
              {{ templateTypes[key].label }}
            </button>
          </div>

          <!-- 风格 -->
          <div class="mt-5 space-y-2.5">
            <div class="flex items-baseline justify-between">
              <h3 class="text-[13px] font-semibold text-muted-foreground">选择风格</h3>
              <span class="text-[11px] text-muted-foreground/70">切换后各自保留已填内容</span>
            </div>
            <div class="-mx-1 flex gap-2.5 overflow-x-auto px-1 pb-2">
              <button
                v-for="style in currentStyles"
                :key="style.value"
                class="w-[136px] flex-shrink-0 overflow-hidden rounded-xl border text-left transition-all"
                :class="form.custom_template === style.value
                  ? 'border-primary/50 ring-2 ring-primary/10'
                  : 'border-border hover:-translate-y-0.5 hover:border-primary/30'"
                @click="selectStyle(style.value)"
              >
                <div class="relative h-14 w-full" :style="{ background: style.gradient }">
                  <div class="absolute inset-0 shadow-[inset_0_1px_0_rgba(255,255,255,0.3),inset_0_-14px_24px_rgba(0,0,0,0.16)]"></div>
                  <div v-if="form.custom_template === style.value" class="absolute right-1.5 top-1.5 flex h-4 w-4 items-center justify-center rounded-full bg-white/95">
                    <Check class="h-2.5 w-2.5 text-primary" :stroke-width="3" />
                  </div>
                </div>
                <div class="px-2.5 py-2">
                  <p class="truncate text-[13px] font-semibold tracking-tight">{{ style.label }}</p>
                  <p class="mt-0.5 truncate text-[11px] text-muted-foreground">{{ style.preview }}</p>
                </div>
              </button>
            </div>
          </div>
        </section>

        <!-- 访问控制（仅对当前标签生效） -->
        <section class="card">
          <h2 class="card-title">访问控制</h2>
          <p class="mt-1 text-[12.5px] leading-relaxed text-muted-foreground">
            仅对这一个标签生效，其它标签不受影响。不设置则任何人扫码都能直接查看。
          </p>

          <div class="mt-4 space-y-3">
            <label class="ac-row">
              <span class="ac-label">访问模式</span>
              <select v-model="form.access_mode" class="ac-input">
                <option value="none">不限制（直出）</option>
                <option value="passcode">仅暗号</option>
                <option value="ticket">仅一次性链接（扫码即开）</option>
                <option value="both">暗号 + 一次性链接</option>
              </select>
            </label>

            <template v-if="needPasscode">
              <label class="ac-row">
                <span class="ac-label">暗号</span>
                <input
                  v-model="form.passcode"
                  type="password"
                  class="ac-input"
                  autocomplete="new-password"
                  :placeholder="hasPasscode ? '已设置，留空表示不修改' : '设置一个暗号'"
                />
              </label>
              <label class="ac-row">
                <span class="ac-label">暗号提示</span>
                <input v-model="form.passcode_hint" class="ac-input" maxlength="64" placeholder="例如：今晚的暗号" />
              </label>
            </template>

            <template v-if="needTicket">
              <label class="ac-row">
                <span class="ac-label">链接有效期</span>
                <div class="ac-inline">
                  <input v-model.number="form.ticket_ttl_seconds" type="number" min="60" step="60" class="ac-input" />
                  <span class="ac-unit">秒</span>
                </div>
              </label>
            </template>

            <label class="ac-row">
              <span class="ac-label">自动跳回首页</span>
              <div class="ac-inline">
                <input v-model.number="form.auto_close_seconds" type="number" min="0" step="60" class="ac-input" />
                <span class="ac-unit">秒（0 = 不跳转）</span>
              </div>
            </label>

            <label class="ac-toggle">
              <input v-model="form.watermark_enabled" type="checkbox" />
              <span>显示防截图水印</span>
            </label>

            <p v-if="form.access_mode !== 'none'" class="ac-tip">
              已开启访问控制：转发出去的链接在对方设备上无法打开。
            </p>
          </div>
        </section>

        <section class="card">
          <div class="mb-4 flex items-start justify-between gap-3">
            <div>
              <h2 class="card-title !mb-1">展示内容</h2>
              <p class="text-xs text-muted-foreground">填写后右侧即时更新，无需了解 JSON</p>
            </div>
            <div class="flex flex-shrink-0 items-center gap-1.5">
              <button class="btn-mini" @click="fillSample">填充示例</button>
              <button class="btn-mini" @click="clearContent">清空</button>
            </div>
          </div>

          <div v-if="contentMissing" class="mb-4 flex items-start gap-2 rounded-lg border border-amber-300/60 bg-amber-50 px-3 py-2.5 text-[13px] text-amber-900">
            <Info class="mt-0.5 h-4 w-4 flex-shrink-0" />
            <span>当前风格缺少部分字段（例如医疗/宠物专属信息），已保留原有数据，切换回来即可继续编辑。</span>
          </div>

          <PayloadEditor
            ref="editorRef"
            :fields="currentSchema"
            :model-value="form.template_payload"
            :active-style="form.custom_template"
            @update:model-value="onPayloadChange"
          />
        </section>

        <!-- 高级：JSON -->
        <section class="card">
          <button class="flex w-full items-center justify-between text-left" @click="showJson = !showJson">
            <div>
              <h2 class="card-title !mb-0">高级：JSON 配置</h2>
              <p class="mt-0.5 text-xs text-muted-foreground">批量编辑或粘贴配置时使用，与上方表单双向同步</p>
            </div>
            <ChevronDown class="h-4 w-4 text-muted-foreground transition-transform" :class="showJson ? 'rotate-180' : ''" />
          </button>
          <div v-if="showJson" class="mt-3 space-y-2">
            <textarea v-model="jsonText" class="json-input" spellcheck="false" @blur="applyJson"></textarea>
            <div class="flex items-center justify-between">
              <span class="text-[11px]" :class="jsonError ? 'text-red-600' : 'text-muted-foreground'">
                {{ jsonError || '失焦时自动解析' }}
              </span>
              <button class="btn-mini" @click="applyJson">应用 JSON</button>
            </div>
          </div>
        </section>
      </div>


      <!-- 右：实时预览 -->
      <aside class="lg:sticky lg:top-6 lg:h-[calc(100vh-3rem)]">
        <div class="flex h-full flex-col overflow-hidden rounded-xl border bg-card">
          <div class="flex items-center justify-between border-b px-3 py-2">
            <div class="flex items-center gap-1.5">
              <span class="h-2 w-2 rounded-full bg-muted-foreground/30"></span>
              <span class="text-[11px] font-medium uppercase tracking-wider text-muted-foreground">实时预览</span>
            </div>
            <div class="flex items-center gap-1">
              <button class="btn-mini" @click="previewKey++" title="刷新预览">
                <RotateCw class="h-3 w-3" />
              </button>
              <a :href="previewHref" target="_blank" rel="noopener" class="btn-mini" title="新窗口打开">
                <ExternalLink class="h-3 w-3" />
              </a>
            </div>
          </div>
          <div class="flex-1 bg-muted/40 p-3">
            <div class="mx-auto h-full max-w-[375px] overflow-hidden rounded-[18px] bg-white shadow-sm">
              <iframe
                :key="previewKey"
                ref="previewFrame"
                :src="previewSrc"
                class="h-full w-full border-0"
                title="模板预览"
                @load="postPreview"
              ></iframe>
            </div>
          </div>
        </div>
      </aside>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, watch, onMounted, onBeforeUnmount, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Loader2, Check, ChevronDown, Info, RotateCw, ExternalLink } from 'lucide-vue-next'
import { ElMessage } from 'element-plus'
import { fetchNfcDetail, createNfc, updateNfc } from '../api/nfc'
import PayloadEditor from '../templates/PayloadEditor.vue'
import {
  templateTypes, templateOrder, visibleTypeOrder, schemaFor, sampleFor,
} from '../templates/registry.js'

const route = useRoute()
const router = useRouter()

const form = reactive({
  nfc_id: '',
  name: '',
  phone: '',
  status: 'enabled',
  template_type: 'business_card',
  custom_template: '',
  template_payload: {},
  // 访问控制（标签级）
  access_mode: 'none',
  passcode: '',
  passcode_hint: '',
  ticket_ttl_seconds: 3600,
  auto_close_seconds: 0,
  watermark_enabled: false,
})

/** 编辑时后端只回传 has_passcode，不回传明文 */
const hasPasscode = ref(false)
const needPasscode = computed(() => form.access_mode === 'passcode' || form.access_mode === 'both')
const needTicket = computed(() => form.access_mode === 'ticket' || form.access_mode === 'both')

const pageLoading = ref(false)
const submitting = ref(false)
const showJson = ref(false)
const jsonText = ref('{}')
const jsonError = ref('')
const editorRef = ref(null)
const previewFrame = ref(null)
const previewKey = ref(0)

const isEdit = computed(() => !!route.params.id)

const currentStyles = computed(() => templateTypes[form.template_type]?.styles || [])
const currentSchema = computed(() => schemaFor(form.template_type, form.custom_template))

// 预览地址只决定"加载哪个页面"，不随风格变化——
// 之前 src 里带 style，每次切风格都会重载整个 iframe，表现为白屏闪烁。
// 现在固定 src，风格与内容全部通过 postMessage 更新（PreviewPage 支持动态切换）。
const previewSrc = ref('/preview?nfcId=PREVIEW')
const previewHref = computed(() => `/preview?${new URLSearchParams({
  type: form.template_type, style: form.custom_template || '', nfcId: 'PREVIEW',
}).toString()}`)

const contentMissing = computed(() => {
  // 风格专属字段存在但当前风格不展示 → 给出提示
  const styles = currentStyles.value.map((s) => s.value)
  const allKeys = new Set()
  for (const s of currentStyles.value) {
    for (const f of s.schema || []) {
      if (f.key) allKeys.add(f.key)
    }
  }
  const shown = new Set()
  for (const f of currentSchema.value) {
    if (f.key) shown.add(f.key)
  }
  const hidden = [...allKeys].filter((k) => !shown.has(k))
  return hidden.some((k) => form.template_payload?.[k] != null && form.template_payload[k] !== '')
})

/* ---------- 草稿缓存：按「类型::风格」保留各自内容 ---------- */
const drafts = reactive({})
const comboKey = (t, s) => `${t}::${s || ''}`

const applySample = (t, s) => {
  const sample = sampleFor(t, s)
  form.template_payload = JSON.parse(JSON.stringify(sample || {}))
  syncJson()
}

const switchCombo = (newType, newStyle) => {
  const oldKey = comboKey(form.template_type, form.custom_template)
  const newKey = comboKey(newType, newStyle)
  if (oldKey === newKey) return

  drafts[oldKey] = JSON.parse(JSON.stringify(form.template_payload || {}))
  form.template_type = newType
  form.custom_template = newStyle

  const cached = drafts[newKey]
  if (cached !== undefined) {
    form.template_payload = JSON.parse(JSON.stringify(cached))
  } else {
    applySample(newType, newStyle)
  }
  syncJson()
}

const selectTemplateType = (key) => {
  const first = templateTypes[key]?.styles?.[0]?.value ?? ''
  switchCombo(key, first)
}
const selectStyle = (value) => switchCombo(form.template_type, value)

const onPayloadChange = (val) => {
  form.template_payload = val
  syncJson()
}

const fillSample = () => {
  applySample(form.template_type, form.custom_template)
  ElMessage.success('已填入示例数据')
}

const clearContent = () => {
  form.template_payload = {}
  syncJson()
  ElMessage.success('已清空，可重新填写')
}

/* ---------- JSON 双向同步 ---------- */
const syncJson = () => {
  jsonText.value = JSON.stringify(form.template_payload || {}, null, 2)
  jsonError.value = ''
}
const applyJson = () => {
  try {
    const parsed = JSON.parse(jsonText.value || '{}')
    if (!parsed || typeof parsed !== 'object' || Array.isArray(parsed)) {
      throw new Error('根节点必须是对象')
    }
    form.template_payload = parsed
    jsonError.value = ''
    ElMessage.success('JSON 已应用')
  } catch (e) {
    jsonError.value = `JSON 格式错误：${e.message}`
  }
}

/* ---------- 预览通信 ---------- */
const postPreview = () => {
  const frame = previewFrame.value
  if (!frame?.contentWindow) return
  frame.contentWindow.postMessage({
    __nfcPreview: true,
    type: form.template_type,
    style: form.custom_template || '',
    // 响应式 Proxy 无法结构化克隆，先转成普通对象
    payload: JSON.parse(JSON.stringify(form.template_payload || {})),
    nfcId: form.nfc_id || 'PREVIEW',
  }, '*')
}

// iframe 重新加载（切换类型/风格导致 src 变化）后，它准备好的那一刻再推一次数据
const onFrameMessage = (e) => {
  if (e.data && e.data.__nfcPreviewReady === true) postPreview()
}

watch(
  () => [form.template_type, form.custom_template, form.template_payload, form.nfc_id],
  () => nextTick(postPreview),
  { deep: true }
)

/* ---------- 保存 ---------- */
const submit = async () => {
  if (!isEdit.value && !form.nfc_id.trim()) {
    ElMessage.warning('请输入 NFC ID')
    return
  }
  const errors = editorRef.value?.validate?.() || []
  if (errors.length > 0) {
    ElMessage.warning(errors[0])
    return
  }

  submitting.value = true
  try {
    const payload = {
      nfc_id: form.nfc_id,
      name: form.name || '',
      phone: form.phone,
      status: form.status,
      template_type: form.template_type,
      template_payload: form.template_payload,
      custom_template: form.custom_template || '',
      // 访问控制
      access_mode: form.access_mode,
      passcode_hint: form.passcode_hint || '',
      ticket_ttl_seconds: Number(form.ticket_ttl_seconds) || 3600,
      auto_close_seconds: Number(form.auto_close_seconds) || 0,
      watermark_enabled: !!form.watermark_enabled,
    }
    // 暗号：新建时必传（可为空）；编辑时留空表示不修改
    if (form.passcode) {
      payload.passcode = form.passcode
    } else if (!isEdit.value) {
      payload.passcode = ''
    }

    if (isEdit.value) {
      await updateNfc(route.params.id, payload)
    } else {
      await createNfc(payload)
    }

    ElMessage.success('保存成功')
    router.push('/admin/nfc')
  } catch {
    // 拦截器已提示
  } finally {
    submitting.value = false
  }
}

const goBack = () => {
  router.push('/admin/nfc')
}

/* ---------- 初始化 ---------- */
onMounted(async () => {
  window.addEventListener('message', onFrameMessage)
  if (isEdit.value) {
    pageLoading.value = true
    try {
      const { data } = await fetchNfcDetail(route.params.id)
      form.nfc_id = data.nfc_id || ''
      form.name = data.name || ''
      form.phone = data.phone || ''
      form.status = data.status || 'enabled'
      form.template_type = data.template_type || 'business_card'
      form.custom_template = data.custom_template || ''
      form.template_payload = data.template_payload || {}
      form.access_mode = data.access_mode || 'none'
      form.passcode_hint = data.passcode_hint || ''
      form.ticket_ttl_seconds = data.ticket_ttl_seconds || 3600
      form.auto_close_seconds = data.auto_close_seconds || 0
      form.watermark_enabled = !!data.watermark_enabled
      hasPasscode.value = !!data.has_passcode
      syncJson()
    } catch {
      // 拦截器已提示
    } finally {
      pageLoading.value = false
    }
  } else {
    applySample(form.template_type, form.custom_template)
  }
  nextTick(postPreview)
})

onBeforeUnmount(() => {
  window.removeEventListener('message', onFrameMessage)
})
</script>

<style scoped>
.card {
  border-radius: 0.75rem;
  border: 1px solid hsl(var(--border));
  background-color: hsl(var(--card));
  padding: 1.25rem;
}
.card-title {
  margin-bottom: 1rem;
  font-size: 1rem;
  font-weight: 600;
  letter-spacing: -0.01em;
}
.field-label {
  font-size: 0.8125rem;
  font-weight: 500;
}
.field-input {
  display: flex;
  height: 2.375rem;
  width: 100%;
  border-radius: 0.375rem;
  border: 1px solid hsl(var(--input));
  background-color: hsl(var(--background));
  padding: 0 0.75rem;
  font-size: 0.875rem;
  outline: none;
  transition: border-color 0.15s, box-shadow 0.15s;
}
.field-input::placeholder { color: hsl(var(--muted-foreground)); }
.field-input:focus-visible {
  border-color: hsl(var(--ring));
  box-shadow: 0 0 0 2px hsl(var(--ring) / 0.12);
}
.hint { font-size: 0.6875rem; color: hsl(var(--muted-foreground)); }
.chip {
  border-radius: 0.375rem;
  border: 1px solid hsl(var(--input));
  padding: 0.375rem 0.875rem;
  font-size: 0.8125rem;
  font-weight: 500;
  color: hsl(var(--muted-foreground));
  transition: all 0.15s;
}
.chip:hover { background-color: hsl(var(--accent)); }
.chip-on {
  border-color: rgb(34 197 94 / 0.5);
  background-color: rgb(34 197 94 / 0.12);
  color: rgb(21 128 61);
}
.chip-off {
  border-color: rgb(107 114 128 / 0.5);
  background-color: rgb(107 114 128 / 0.12);
  color: rgb(55 65 81);
}
.btn-primary {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.375rem;
  height: 2.25rem;
  border-radius: 0.375rem;
  background-color: hsl(var(--primary));
  padding: 0 1rem;
  font-size: 0.875rem;
  font-weight: 500;
  color: hsl(var(--primary-foreground));
  box-shadow: 0 1px 2px rgb(0 0 0 / 0.06);
  transition: opacity 0.15s, transform 0.1s;
}
.btn-primary:hover { opacity: 0.9; }
.btn-primary:active { transform: scale(0.98); }
.btn-primary:disabled { opacity: 0.5; }
.btn-ghost {
  display: inline-flex;
  align-items: center;
  height: 2.25rem;
  border-radius: 0.375rem;
  border: 1px solid hsl(var(--input));
  padding: 0 1rem;
  font-size: 0.875rem;
  font-weight: 500;
  transition: background-color 0.15s;
}
.btn-ghost:hover { background-color: hsl(var(--accent)); }
.btn-mini {
  display: inline-flex;
  align-items: center;
  gap: 0.25rem;
  border-radius: 0.375rem;
  border: 1px solid hsl(var(--border));
  padding: 0.25rem 0.5rem;
  font-size: 0.6875rem;
  font-weight: 500;
  color: hsl(var(--muted-foreground));
  transition: all 0.15s;
}
.btn-mini:hover {
  background-color: hsl(var(--accent));
  color: hsl(var(--foreground));
}
.json-input {
  min-height: 220px;
  width: 100%;
  border-radius: 0.5rem;
  border: 1px solid hsl(var(--input));
  background-color: #0b0f16;
  padding: 0.75rem;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 0.75rem;
  line-height: 1.6;
  color: #e2e8f0;
  outline: none;
  resize: vertical;
}

/* 访问控制表单 */
.ac-row {
  display: block;
}
.ac-label {
  display: block;
  margin-bottom: 6px;
  font-size: 12.5px;
  color: rgba(20, 24, 29, 0.6);
}
.ac-input {
  width: 100%;
  height: 38px;
  padding: 0 12px;
  border: 1px solid rgba(20, 24, 29, 0.12);
  border-radius: 9px;
  background: #fff;
  font-size: 13.5px;
  outline: none;
  transition: border-color 0.2s ease, box-shadow 0.2s ease;
}
.ac-input:focus {
  border-color: rgba(255, 107, 53, 0.55);
  box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.09);
}
.ac-inline {
  display: flex;
  align-items: center;
  gap: 8px;
}
.ac-inline .ac-input {
  max-width: 140px;
}
.ac-unit {
  font-size: 12px;
  color: rgba(20, 24, 29, 0.45);
  white-space: nowrap;
}
.ac-toggle {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: rgba(20, 24, 29, 0.75);
  cursor: pointer;
}
.ac-toggle input {
  width: 15px;
  height: 15px;
  accent-color: #ff6b35;
}
.ac-tip {
  margin-top: 4px;
  font-size: 12px;
  color: rgba(20, 24, 29, 0.45);
}
</style>
