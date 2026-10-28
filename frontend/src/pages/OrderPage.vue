<template>
  <div class="min-h-screen bg-[#fbfaf7] font-sans text-[#14181d]">
    <!-- 顶部 -->
    <SiteHeader back-to="/showroom" back-label="返回展厅" cta-label="逛展厅" cta-to="/showroom" />

    <!-- 下单成功 -->
    <section v-if="done" class="mx-auto flex min-h-[calc(100vh-60px)] max-w-[720px] items-center px-4 py-16 sm:px-6">
      <div class="w-full overflow-hidden rounded-3xl border border-[#14181d]/[0.07] bg-white shadow-[0_24px_70px_-40px_rgba(20,24,29,0.4)]">
        <div class="border-b border-[#14181d]/[0.06] bg-[#fbfaf7] px-7 py-8 text-center">
          <span class="mx-auto flex h-14 w-14 items-center justify-center rounded-2xl bg-[#16a34a]/10">
            <CheckCircle2 class="h-7 w-7 text-[#16a34a]" :stroke-width="1.8" />
          </span>
          <h2 class="mt-4 text-[24px] font-semibold tracking-[-0.02em]">订单已提交</h2>
          <p class="mt-2 text-[13.5px] leading-relaxed text-[#14181d]/55">
            我们会尽快与你确认制作与邮寄细节，请保持手机畅通。
          </p>
        </div>

        <div class="px-7 py-7">
          <p class="text-[11px] font-semibold uppercase tracking-[0.18em] text-[#14181d]/35">订单号</p>
          <div class="mt-2 flex items-center gap-3">
            <span class="font-mono text-[26px] font-semibold tracking-tight text-[#ff6b35]">{{ done.order_no }}</span>
            <button class="copy-btn" @click="copyOrderNo">
              <Copy class="h-3.5 w-3.5" />
              复制
            </button>
          </div>

          <dl class="mt-6 grid gap-3 rounded-2xl bg-[#fbfaf7] p-4 text-[13px] sm:grid-cols-2">
            <div>
              <dt class="text-[#14181d]/40">模板</dt>
              <dd class="mt-0.5 font-medium">{{ typeLabel(form.template_type) }} · {{ styleLabel(form.template_type, form.custom_template) }}</dd>
            </div>
            <div>
              <dt class="text-[#14181d]/40">数量 / 金额</dt>
              <dd class="mt-0.5 font-medium">{{ done.quantity }} 枚 · {{ money(done.amount) }}</dd>
            </div>
            <div v-if="done.external_order_no">
              <dt class="text-[#14181d]/40">第三方订单号</dt>
              <dd class="mt-0.5 font-mono font-medium">{{ done.external_order_no }}</dd>
            </div>
            <div>
              <dt class="text-[#14181d]/40">收件人</dt>
              <dd class="mt-0.5 font-medium">{{ done.contact_name }} · {{ done.contact_phone }}</dd>
            </div>
            <div class="sm:col-span-2">
              <dt class="text-[#14181d]/40">邮寄地址</dt>
              <dd class="mt-0.5 font-medium">{{ done.contact_address }}</dd>
            </div>
          </dl>

          <div class="mt-7 flex flex-col gap-2 sm:flex-row">
            <button class="btn-dark flex-1" @click="startOver">
              <Plus class="h-4 w-4" />
              再下一单
            </button>
            <router-link to="/showroom" class="btn-ghost flex-1 justify-center">
              回到展厅
            </router-link>
          </div>
        </div>
      </div>
    </section>

    <!-- 下单主体 -->
    <template v-else>
      <section class="mx-auto max-w-[1400px] px-4 pb-6 pt-12 sm:px-6">
        <p class="eyebrow">Get hardware</p>
        <h1 class="mt-3 text-[32px] font-semibold leading-tight tracking-[-0.03em] md:text-[42px]">
          选一个模板，<br class="md:hidden" />贴上就能用
        </h1>
        <p class="mt-4 max-w-[560px] text-[15px] leading-relaxed text-[#14181d]/55">
          选类型、挑风格、填内容，再留下收件信息就行。下单后我们制作并寄出，购买即完成首次写入；含 1 次免费内容调整。
        </p>
      </section>

      <section class="mx-auto max-w-[1400px] px-4 pb-24 sm:px-6">
        <!-- 必须显式 grid-cols-1：否则移动端是隐式 auto 列，
             会被内部横向滚动区（30 张风格卡）的 max-content 撑到 4000px+ -->
        <div class="grid grid-cols-1 gap-6 lg:grid-cols-[minmax(0,1fr)_368px]">
          <!-- 左：表单 -->
          <div class="min-w-0 space-y-4">
            <!-- 01 类型 -->
            <section class="panel">
              <div class="panel-head">
                <span class="step">01</span>
                <div>
                  <h2 class="panel-title">选择类型</h2>
                  <p class="panel-sub">这款标签用来做什么</p>
                </div>
              </div>
              <div class="grid grid-cols-2 gap-2.5 sm:grid-cols-3 lg:grid-cols-5">
                <button
                  v-for="key in visibleTypeOrder"
                  :key="key"
                  class="type-btn"
                  :class="form.template_type === key ? 'type-btn-on' : ''"
                  @click="selectType(key)"
                >
                  <span
                    class="flex h-9 w-9 items-center justify-center rounded-lg"
                    :style="{ background: templateTypes[key].iconBg, color: templateTypes[key].iconColor }"
                  >
                    <component :is="templateTypes[key].icon" class="h-[18px] w-[18px]" />
                  </span>
                  <span class="mt-1.5 text-[12.5px] font-medium leading-tight">{{ templateTypes[key].label }}</span>
                  <span class="mt-0.5 text-[10.5px] text-[#14181d]/35">{{ templateTypes[key].styles.length }} 种风格</span>
                </button>
              </div>
            </section>

            <!-- 02 风格 -->
            <section class="panel">
              <div class="panel-head">
                <span class="step">02</span>
                <div class="min-w-0 flex-1">
                  <h2 class="panel-title">选择风格</h2>
                  <p class="panel-sub">{{ templateTypes[form.template_type]?.label }} · 共 {{ currentStyles.length }} 款，切换后已填内容会保留</p>
                </div>
              </div>
              <div class="-mx-1 flex gap-2.5 overflow-x-auto px-1 pb-2">
                <button
                  v-for="style in currentStyles"
                  :key="style.value"
                  class="style-card"
                  :class="form.custom_template === style.value ? 'style-card-on' : ''"
                  @click="selectStyle(style.value)"
                >
                  <div class="relative h-16 w-full" :style="{ background: style.gradient }">
                    <div class="absolute inset-0 shadow-[inset_0_1px_0_rgba(255,255,255,0.28),inset_0_-16px_26px_rgba(0,0,0,0.2)]"></div>
                    <span v-if="form.custom_template === style.value" class="absolute right-1.5 top-1.5 flex h-4 w-4 items-center justify-center rounded-full bg-white/95">
                      <Check class="h-2.5 w-2.5 text-[#ff6b35]" :stroke-width="3.2" />
                    </span>
                  </div>
                  <div class="px-2.5 py-2">
                    <p class="truncate text-[12.5px] font-semibold tracking-tight">{{ style.label }}</p>
                    <p class="mt-0.5 truncate text-[10.5px] text-[#14181d]/40">{{ style.preview }}</p>
                  </div>
                </button>
              </div>
            </section>

            <!-- 03 内容 -->
            <section class="panel">
              <div class="panel-head">
                <span class="step">03</span>
                <div class="min-w-0 flex-1">
                  <h2 class="panel-title">填写展示内容</h2>
                  <p class="panel-sub">手机上看到的就是这些内容，右侧可实时预览</p>
                </div>
                <div class="flex flex-shrink-0 items-center gap-1.5">
                  <button class="mini-btn" @click="fillSample">填充示例</button>
                  <button class="mini-btn" @click="clearContent">清空</button>
                </div>
              </div>

              <div class="mb-4 space-y-2">
                <label class="field-label">标签名</label>
                <input v-model="form.tag_name" class="field-input" maxlength="128" placeholder="方便你在后台区分这枚标签" />
              </div>

              <PayloadEditor
                ref="editorRef"
                :fields="currentSchema"
                :model-value="form.template_payload"
                :active-style="form.custom_template"
                @update:model-value="onPayloadChange"
              />
            </section>

            <!-- 04 第三方订单 -->
            <section class="panel">
              <div class="panel-head">
                <span class="step">04</span>
                <div class="min-w-0 flex-1">
                  <h2 class="panel-title">
                    第三方订单
                    <span class="opt-tag">选填</span>
                  </h2>
                  <p class="panel-sub">已经在淘宝 / 微信小店等平台下过单？填这里的单号，数量与金额以该订单为准</p>
                </div>
              </div>

              <div class="grid gap-4 sm:grid-cols-3">
                <div class="space-y-2">
                  <label class="field-label">第三方订单编号</label>
                  <input
                    v-model="form.external_order_no"
                    class="field-input"
                    maxlength="64"
                    placeholder="例如：TB2026102000123"
                  />
                </div>
                <div class="space-y-2">
                  <label class="field-label">数量（个）</label>
                  <input
                    v-model.number="externalQty"
                    type="number"
                    min="1"
                    max="9999"
                    class="field-input"
                    :disabled="!hasExternalOrder"
                    placeholder="该订单的数量"
                  />
                </div>
                <div class="space-y-2">
                  <label class="field-label">金额（元）</label>
                  <input
                    v-model.number="externalAmount"
                    type="number"
                    min="0"
                    step="0.1"
                    class="field-input"
                    :disabled="!hasExternalOrder"
                    placeholder="该订单的实付金额"
                  />
                </div>
              </div>

              <p
                v-if="hasExternalOrder"
                class="mt-3 rounded-xl bg-[#16a34a]/[0.07] px-3.5 py-2.5 text-[12px] leading-relaxed text-[#15803d]"
              >
                已关联第三方订单 <span class="font-mono font-medium">{{ form.external_order_no }}</span>，下面不用再选数量与费用。
              </p>
            </section>

            <!-- 05 数量与费用（填了第三方订单则不展示） -->
            <section v-if="!hasExternalOrder" class="panel">
              <div class="panel-head">
                <span class="step">05</span>
                <div class="min-w-0 flex-1">
                  <h2 class="panel-title">数量与费用</h2>
                  <p class="panel-sub">选一档数量，价格自动算；也可以自定义数量与金额</p>
                </div>
              </div>

              <div class="grid grid-cols-3 gap-2.5 sm:grid-cols-6">
                <button
                  v-for="preset in QUANTITY_PRESETS"
                  :key="preset.qty"
                  class="qty-btn"
                  :class="qtyMode === preset.qty ? 'qty-btn-on' : ''"
                  @click="selectQty(preset.qty)"
                >
                  <span class="qty-main">{{ preset.label }}</span>
                  <span class="qty-price">{{ money(preset.amount) }}</span>
                  <span class="qty-hint">{{ preset.hint }}</span>
                </button>
                <button
                  class="qty-btn"
                  :class="qtyMode === CUSTOM_MODE ? 'qty-btn-on' : ''"
                  @click="selectQty(CUSTOM_MODE)"
                >
                  <span class="qty-main">自定义</span>
                  <span class="qty-price">按阶梯</span>
                  <span class="qty-hint">自定数量/金额</span>
                </button>
              </div>

              <div v-if="qtyMode === CUSTOM_MODE" class="mt-4 grid gap-3 sm:grid-cols-2">
                <div class="space-y-2">
                  <label class="field-label">数量（个）</label>
                  <input
                    v-model.number="customQty"
                    type="number"
                    min="1"
                    max="9999"
                    class="field-input"
                    placeholder="例如：20"
                    @input="onCustomQtyInput"
                  />
                  <p class="text-[11px] text-[#14181d]/35">50 个以上 ¥6/个，100 个以上 ¥5.5/个</p>
                </div>
                <div class="space-y-2">
                  <label class="field-label">金额（元）</label>
                  <input
                    v-model.number="customAmount"
                    type="number"
                    min="0"
                    step="0.1"
                    class="field-input"
                    placeholder="按数量自动算，可修改"
                    @input="amountTouched = true"
                  />
                  <p class="text-[11px] text-[#14181d]/35">可手动调整，商务确认以实际报价为准</p>
                </div>
              </div>

              <div class="mt-4 flex flex-wrap items-end justify-between gap-3 rounded-xl bg-[#fbfaf7] px-4 py-3.5">
                <div>
                  <p class="text-[11.5px] text-[#14181d]/40">共 {{ quantity }} 枚 · 单价约 {{ unitMoney(amount, quantity) }}</p>
                  <p class="mt-1 text-[24px] font-semibold tracking-[-0.02em]">{{ money(amount) }}</p>
                </div>
                <p class="text-[11px] text-[#14181d]/35">支付暂未接入，先记录订单</p>
              </div>
            </section>

            <!-- 06 邮寄信息 -->
            <section class="panel">
              <div class="panel-head">
                <span class="step">06</span>
                <div>
                  <h2 class="panel-title">邮寄信息</h2>
                  <p class="panel-sub">标签做好后按这里寄出，请填真实信息</p>
                </div>
              </div>
              <div class="grid gap-4 sm:grid-cols-2">
                <div class="space-y-2">
                  <label class="field-label">收件人 <span class="req">*</span></label>
                  <input v-model="form.contact_name" class="field-input" maxlength="64" placeholder="姓名" />
                </div>
                <div class="space-y-2">
                  <label class="field-label">手机号 <span class="req">*</span></label>
                  <input v-model="form.contact_phone" class="field-input" maxlength="32" inputmode="tel" placeholder="用于快递联系" />
                </div>
                <div class="space-y-2 sm:col-span-2">
                  <label class="field-label">详细地址 <span class="req">*</span></label>
                  <textarea
                    v-model="form.contact_address"
                    rows="2"
                    maxlength="256"
                    class="field-input resize-y leading-relaxed"
                    placeholder="省 / 市 / 区 / 街道门牌，例如：江苏省苏州市姑苏区平江路 112 号 3 单元 501"
                  ></textarea>
                </div>
                <div class="space-y-2">
                  <label class="field-label">邮政编码</label>
                  <input v-model="form.contact_zipcode" class="field-input" maxlength="16" placeholder="选填" />
                </div>
                <div class="space-y-2">
                  <label class="field-label">备注</label>
                  <input v-model="form.contact_note" class="field-input" maxlength="256" placeholder="选填，例如：工作日送货" />
                </div>
              </div>
            </section>

            <!-- 下单 -->
            <section class="rounded-2xl border border-[#14181d]/[0.07] bg-[#14181d] p-5 text-white sm:p-6">
              <div class="flex flex-wrap items-center justify-between gap-4">
                <div class="min-w-0">
                  <p class="text-[11px] font-semibold uppercase tracking-[0.18em] text-white/35">Order summary</p>
                  <p class="mt-1.5 text-[15px] font-medium">
                    {{ typeLabel(form.template_type) }} · {{ styleLabel(form.template_type, form.custom_template) }}
                  </p>
                  <p class="mt-1 text-[12.5px] text-white/45">
                    <span v-if="hasExternalOrder" class="font-mono text-white/70">第三方订单 {{ form.external_order_no }} · </span>
                    {{ quantity }} 枚 · {{ money(amount) }} · 下单后客服联系确认配送
                  </p>
                </div>
                <button class="btn-orange" :disabled="submitting" @click="submitOrder">
                  <Loader2 v-if="submitting" class="h-4 w-4 animate-spin" />
                  {{ submitting ? '提交中…' : '下订单' }}
                  <ArrowRight v-if="!submitting" class="h-4 w-4" />
                </button>
              </div>
            </section>
          </div>

          <!-- 右：实时预览 -->
          <aside class="lg:sticky lg:top-[76px] lg:self-start">
            <div class="overflow-hidden rounded-2xl border border-[#14181d]/[0.07] bg-white">
              <div class="flex items-center justify-between border-b border-[#14181d]/[0.06] px-3 py-2">
                <span class="flex items-center gap-1.5">
                  <span class="h-1.5 w-1.5 rounded-full bg-[#16a34a]"></span>
                  <span class="text-[10.5px] font-medium uppercase tracking-[0.16em] text-[#14181d]/40">实时预览</span>
                </span>
                <div class="flex items-center gap-1">
                  <button class="mini-btn" title="刷新预览" @click="previewKey++">
                    <RotateCw class="h-3 w-3" />
                  </button>
                  <a :href="previewHref" target="_blank" rel="noopener" class="mini-btn" title="新窗口打开">
                    <ExternalLink class="h-3 w-3" />
                  </a>
                </div>
              </div>
              <div class="bg-[#f4f4f3] p-3">
                <div class="mx-auto h-[560px] max-w-[335px] overflow-hidden rounded-[18px] bg-white shadow-[0_10px_30px_-16px_rgba(20,24,29,0.4)]">
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
      </section>
    </template>
  </div>
</template>

<script setup>
import { ref, reactive, computed, watch, nextTick } from 'vue'
import {
  Nfc, ArrowLeft, ArrowRight, Check, CheckCircle2, Copy, Loader2,
  Plus, RotateCw, ExternalLink,
} from 'lucide-vue-next'
import { ElMessage } from 'element-plus'
import { templateTypes, templateOrder, visibleTypeOrder, schemaFor, sampleFor } from '../templates/registry'
import { emptyFromFields } from '../templates/schema.js'
import {
  QUANTITY_PRESETS, CUSTOM_MODE, amountForQuantity, money, unitMoney,
} from '../config/pricing'
import PayloadEditor from '../templates/PayloadEditor.vue'
import { createOrder } from '../api/orders'
import SiteHeader from '../components/SiteHeader.vue'

const editorRef = ref(null)
const previewFrame = ref(null)
const previewKey = ref(0)

const submitting = ref(false)
const done = ref(null)

const form = reactive({
  template_type: templateOrder[0],
  custom_template: (templateTypes[templateOrder[0]].styles[0]?.value ?? ''),
  template_payload: {},
  tag_name: '',
  external_order_no: '',
  contact_name: '',
  contact_phone: '',
  contact_address: '',
  contact_zipcode: '',
  contact_note: '',
})

/** 用户是否手动改过标签名（改过就不再自动跟随类型/风格） */
const tagNameTouched = ref(false)

/* ---------- 数量与费用 ---------- */
const qtyMode = ref(QUANTITY_PRESETS[0].qty)
const customQty = ref(20)
const customAmount = ref(amountForQuantity(20))
const amountTouched = ref(false)

/* ---------- 第三方订单 ---------- */
/** 填了第三方订单编号时，数量与费用以第三方订单为准，不再展示数量费用栏 */
const externalQty = ref(1)
const externalAmount = ref(0)
const hasExternalOrder = computed(() => form.external_order_no.trim().length > 0)

const quantity = computed(() => {
  if (hasExternalOrder.value) return Math.max(1, Math.floor(Number(externalQty.value) || 1))
  return qtyMode.value === CUSTOM_MODE
    ? Math.max(1, Math.floor(Number(customQty.value) || 1))
    : qtyMode.value
})

const amount = computed(() => {
  if (hasExternalOrder.value) return Math.max(0, Number(externalAmount.value) || 0)
  if (qtyMode.value === CUSTOM_MODE) {
    return amountTouched.value
      ? Math.max(0, Number(customAmount.value) || 0)
      : amountForQuantity(customQty.value)
  }
  return QUANTITY_PRESETS.find((p) => p.qty === qtyMode.value)?.amount ?? 0
})

const selectQty = (value) => {
  qtyMode.value = value
  if (value === CUSTOM_MODE) {
    amountTouched.value = false
    customAmount.value = amountForQuantity(customQty.value)
  }
}

const onCustomQtyInput = () => {
  if (!amountTouched.value) customAmount.value = amountForQuantity(customQty.value)
}

const typeLabel = (key) => templateTypes[key]?.label || key

const styleLabel = (typeKey, styleValue) => {
  const styles = templateTypes[typeKey]?.styles || []
  const hit = styles.find((s) => s.value === (styleValue || ''))
  return hit?.label || styles[0]?.label || '默认'
}

const currentStyles = computed(() => templateTypes[form.template_type]?.styles || [])
const currentSchema = computed(() => schemaFor(form.template_type, form.custom_template))

// src 固定：切风格不再重载 iframe（之前 src 带 style，每次切换都白屏闪一下），
// 风格与内容通过 postMessage 实时更新
const previewSrc = ref(`/preview?type=${form.template_type}&nfcId=PREVIEW`)
const previewHref = computed(() => previewSrc.value)

/** 初始化 / 切换类型时重建内容对象：默认填入示例，避免右侧预览一片空白 */
const resetPayload = () => {
  const sample = sampleFor(form.template_type, form.custom_template)
  form.template_payload = sample
    ? JSON.parse(JSON.stringify(sample))
    : emptyFromFields(currentSchema.value)
}

/** 标签名默认值：类型 · 风格 */
const autoTagName = () =>
  `${typeLabel(form.template_type)} · ${styleLabel(form.template_type, form.custom_template)}`

const selectType = (key) => {
  if (form.template_type === key) return
  form.template_type = key
  form.custom_template = templateTypes[key].styles[0]?.value ?? ''
  resetPayload()
  if (!tagNameTouched.value) form.tag_name = autoTagName()
}

const selectStyle = (value) => {
  if (form.custom_template === value) return
  form.custom_template = value
  // 各风格字段可能不同（如自定义企业），补齐新字段但保留已填内容
  const base = emptyFromFields(schemaFor(form.template_type, value))
  form.template_payload = { ...base, ...(form.template_payload || {}) }
  if (!tagNameTouched.value) form.tag_name = autoTagName()
}

const onPayloadChange = (val) => {
  form.template_payload = val
}

const fillSample = () => {
  const sample = sampleFor(form.template_type, form.custom_template)
  form.template_payload = JSON.parse(JSON.stringify(sample || {}))
  ElMessage.success('已填入示例内容，可直接改')
}

const clearContent = () => {
  resetPayload()
  ElMessage.success('已清空')
}

/* ---------- 预览通信 ---------- */
const postPreview = () => {
  const win = previewFrame.value?.contentWindow
  if (!win) return
  win.postMessage(
    {
      __nfcPreview: true,
      type: form.template_type,
      style: form.custom_template || '',
      payload: JSON.parse(JSON.stringify(form.template_payload || {})),
      nfcId: 'PREVIEW',
    },
    '*'
  )
}

watch(
  () => [form.template_type, form.custom_template, form.template_payload],
  () => nextTick(postPreview),
  { deep: true }
)

watch(
  () => form.tag_name,
  (v, old) => {
    if (v !== old && v !== autoTagName()) tagNameTouched.value = true
  }
)

/* ---------- 提交 ---------- */
const validateContact = () => {
  if (!form.contact_name.trim()) return '请填写收件人'
  if (!form.contact_phone.trim()) return '请填写手机号'
  if (!/^[\d+\-\s()]{6,}$/.test(form.contact_phone.trim())) return '手机号格式看起来不对'
  if (!form.contact_address.trim()) return '请填写详细地址'
  return ''
}

/** 第三方订单模式下，数量和金额取自该订单，必须填全 */
const validateExternal = () => {
  if (!hasExternalOrder.value) return ''
  if (!(Number(externalQty.value) >= 1)) return '请填写第三方订单的数量'
  if (!(Number(externalAmount.value) > 0)) return '请填写第三方订单的金额'
  return ''
}

const submitOrder = async () => {
  const externalError = validateExternal()
  if (externalError) {
    ElMessage.warning(externalError)
    return
  }
  const contentErrors = editorRef.value?.validate?.() || []
  if (contentErrors.length > 0) {
    ElMessage.warning(contentErrors[0])
    return
  }
  const contactError = validateContact()
  if (contactError) {
    ElMessage.warning(contactError)
    return
  }

  submitting.value = true
  try {
    const { data } = await createOrder({
      template_type: form.template_type,
      custom_template: form.custom_template || '',
      template_payload: JSON.parse(JSON.stringify(form.template_payload || {})),
      tag_name: form.tag_name.trim() || autoTagName(),
      external_order_no: form.external_order_no.trim(),
      quantity: quantity.value,
      amount: amount.value,
      contact_name: form.contact_name.trim(),
      contact_phone: form.contact_phone.trim(),
      contact_address: form.contact_address.trim(),
      contact_zipcode: form.contact_zipcode.trim(),
      contact_note: form.contact_note.trim(),
    })
    done.value = data
    window.scrollTo({ top: 0, behavior: 'smooth' })
  } catch {
    // 错误提示由 request 拦截器统一处理
  } finally {
    submitting.value = false
  }
}

const copyOrderNo = async () => {
  const text = done.value?.order_no || ''
  try {
    if (navigator.clipboard && window.isSecureContext) {
      await navigator.clipboard.writeText(text)
    } else {
      const ta = document.createElement('textarea')
      ta.value = text
      ta.style.position = 'fixed'
      ta.style.opacity = '0'
      document.body.appendChild(ta)
      ta.select()
      document.execCommand('copy')
      document.body.removeChild(ta)
    }
    ElMessage.success('订单号已复制')
  } catch {
    ElMessage.error('复制失败，请手动记录订单号')
  }
}

const startOver = () => {
  done.value = null
  form.contact_address = ''
  form.contact_note = ''
  form.contact_zipcode = ''
  tagNameTouched.value = false
  form.tag_name = autoTagName()
  form.external_order_no = ''
  externalQty.value = 1
  externalAmount.value = 0
  qtyMode.value = QUANTITY_PRESETS[0].qty
  customQty.value = 20
  customAmount.value = amountForQuantity(20)
  amountTouched.value = false
  resetPayload()
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

/* ---------- 初始化 ---------- */
form.tag_name = autoTagName()
resetPayload()
</script>

<style scoped>
.brand-chip {
  display: grid;
  place-items: center;
  width: 32px;
  height: 32px;
  border-radius: 10px;
  background: #14181d;
  transition: transform 0.3s cubic-bezier(0.22, 1, 0.36, 1);
}
.brand-chip:hover { transform: rotate(-6deg); }

.nav-back {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  height: 34px;
  padding: 0 14px;
  border-radius: 999px;
  font-size: 13px;
  color: rgba(20, 24, 29, 0.65);
  border: 1px solid rgba(20, 24, 29, 0.1);
  transition: all 0.25s ease;
}
.nav-back:hover {
  color: #ff6b35;
  border-color: rgba(255, 107, 53, 0.4);
  transform: translateX(-2px);
}

.eyebrow {
  font-size: 11px;
  font-weight: 600;
  letter-spacing: 0.2em;
  text-transform: uppercase;
  color: #ff6b35;
}

.panel {
  border-radius: 18px;
  border: 1px solid rgba(20, 24, 29, 0.07);
  background: #fff;
  padding: 1.15rem;
  box-shadow: 0 1px 2px rgba(20, 24, 29, 0.03);
}
@media (min-width: 640px) {
  .panel { padding: 1.35rem; }
}

.panel-head {
  display: flex;
  align-items: center;
  gap: 0.7rem;
  margin-bottom: 1rem;
}
.panel-title {
  font-size: 15px;
  font-weight: 600;
  letter-spacing: -0.01em;
}
.panel-sub {
  margin-top: 1px;
  font-size: 11.5px;
  color: rgba(20, 24, 29, 0.4);
}

.step {
  display: grid;
  place-items: center;
  width: 26px;
  height: 26px;
  flex-shrink: 0;
  border-radius: 8px;
  background: rgba(255, 107, 53, 0.1);
  color: #ff6b35;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 11px;
  font-weight: 700;
}

.opt-tag {
  display: inline-flex;
  align-items: center;
  height: 18px;
  margin-left: 6px;
  padding: 0 7px;
  border-radius: 999px;
  background: rgba(20, 24, 29, 0.055);
  font-size: 10.5px;
  font-weight: 500;
  letter-spacing: 0;
  color: rgba(20, 24, 29, 0.45);
  vertical-align: 2px;
}

.type-btn {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 0;
  border-radius: 14px;
  border: 1.5px solid rgba(20, 24, 29, 0.08);
  padding: 0.65rem 0.4rem;
  text-align: center;
  transition: all 0.2s ease;
  cursor: pointer;
}
.type-btn:hover {
  border-color: rgba(255, 107, 53, 0.35);
  background: rgba(255, 107, 53, 0.03);
  transform: translateY(-1px);
}
.type-btn-on {
  border-color: #ff6b35;
  background: rgba(255, 107, 53, 0.055);
  box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.1);
}

.style-card {
  width: 132px;
  flex-shrink: 0;
  overflow: hidden;
  border-radius: 13px;
  border: 1px solid rgba(20, 24, 29, 0.09);
  text-align: left;
  transition: all 0.22s ease;
  cursor: pointer;
}
.style-card:hover {
  transform: translateY(-2px);
  border-color: rgba(255, 107, 53, 0.35);
}
.style-card-on {
  border-color: rgba(255, 107, 53, 0.55);
  box-shadow: 0 0 0 2.5px rgba(255, 107, 53, 0.12);
}

.qty-btn {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 1px;
  border-radius: 13px;
  border: 1.5px solid rgba(20, 24, 29, 0.08);
  background: #fff;
  padding: 0.6rem 0.35rem;
  transition: all 0.2s ease;
  cursor: pointer;
}
.qty-btn:hover {
  border-color: rgba(255, 107, 53, 0.35);
  background: rgba(255, 107, 53, 0.03);
  transform: translateY(-1px);
}
.qty-btn-on {
  border-color: #ff6b35;
  background: rgba(255, 107, 53, 0.06);
  box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.1);
}
.qty-main {
  font-size: 13.5px;
  font-weight: 600;
  letter-spacing: -0.01em;
  white-space: nowrap;
}
.qty-price {
  font-size: 11.5px;
  font-weight: 500;
  color: #ff6b35;
  font-variant-numeric: tabular-nums;
  white-space: nowrap;
}
.qty-hint {
  font-size: 10px;
  color: rgba(20, 24, 29, 0.35);
  white-space: nowrap;
}

.field-label {
  font-size: 12px;
  font-weight: 500;
  color: rgba(20, 24, 29, 0.55);
}
.req { color: #dc2626; }

.field-input {
  width: 100%;
  border-radius: 10px;
  border: 1px solid rgba(20, 24, 29, 0.11);
  background: #fbfaf7;
  padding: 0.5rem 0.75rem;
  font-size: 13px;
  line-height: 1.4;
  color: #14181d;
  outline: none;
  transition: border-color 0.16s ease, box-shadow 0.16s ease, background 0.16s ease;
}
.field-input::placeholder { color: rgba(20, 24, 29, 0.28); }
.field-input:hover { border-color: rgba(20, 24, 29, 0.2); }
.field-input:focus {
  border-color: #ff6b35;
  background: #fff;
  box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.13);
}

.mini-btn {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  height: 26px;
  padding: 0 9px;
  border-radius: 7px;
  border: 1px solid rgba(20, 24, 29, 0.1);
  background: #fff;
  font-size: 11.5px;
  font-weight: 500;
  color: rgba(20, 24, 29, 0.55);
  transition: all 0.18s ease;
}
.mini-btn:hover {
  border-color: rgba(20, 24, 29, 0.24);
  color: #14181d;
}

.btn-orange {
  display: inline-flex;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  gap: 7px;
  height: 42px;
  padding: 0 22px;
  border-radius: 12px;
  background: #ff6b35;
  color: #fff;
  font-size: 13.5px;
  font-weight: 600;
  transition: all 0.2s ease;
}
.btn-orange:hover:not(:disabled) {
  background: #f25a22;
  transform: translateY(-1px);
  box-shadow: 0 10px 24px -12px rgba(255, 107, 53, 0.9);
}
.btn-orange:disabled { opacity: 0.6; cursor: not-allowed; }

.btn-dark {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  height: 42px;
  border-radius: 12px;
  background: #14181d;
  color: #fff;
  font-size: 13.5px;
  font-weight: 600;
  transition: all 0.2s ease;
}
.btn-dark:hover { background: #ff6b35; transform: translateY(-1px); }

.btn-ghost {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  height: 42px;
  padding: 0 18px;
  border-radius: 12px;
  border: 1px solid rgba(20, 24, 29, 0.12);
  background: #fff;
  color: rgba(20, 24, 29, 0.68);
  font-size: 13.5px;
  font-weight: 500;
  transition: all 0.2s ease;
}
.btn-ghost:hover { border-color: rgba(20, 24, 29, 0.3); color: #14181d; }

.copy-btn {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  height: 30px;
  padding: 0 11px;
  border-radius: 8px;
  border: 1px solid rgba(20, 24, 29, 0.1);
  background: #fff;
  font-size: 11.5px;
  font-weight: 500;
  color: rgba(20, 24, 29, 0.6);
  transition: all 0.18s ease;
}
.copy-btn:hover { border-color: #ff6b35; color: #ff6b35; }

@media (prefers-reduced-motion: reduce) {
  .type-btn, .style-card, .btn-orange, .btn-dark, .nav-back { transition: none !important; }
}
</style>
