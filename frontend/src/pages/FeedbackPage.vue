<template>
  <div class="fb-page">
    <div class="fb-grain" aria-hidden="true"></div>
    <div class="fb-glow fb-glow-orange" aria-hidden="true"></div>
    <div class="fb-glow fb-glow-sun" aria-hidden="true"></div>

    <header class="fb-header">
      <div class="fb-shell fb-nav">
        <router-link to="/" class="fb-brand">
          <span class="fb-brand-mark">
            <img src="/tapiito-icon-white.png" alt="Tapiito" class="h-5 w-5" draggable="false" />
          </span>
          <span>Tapiito</span>
        </router-link>

        <nav class="fb-nav-links" aria-label="主导航">
          <router-link to="/showroom" class="fb-nav-link fb-hide-mobile">展厅</router-link>
          <router-link to="/about" class="fb-nav-link fb-hide-mobile">关于我们</router-link>
          <router-link to="/feedback" class="fb-nav-link fb-nav-active">我有意见</router-link>
          <router-link to="/consult" class="fb-nav-cta">
            预约咨询
            <ArrowUpRight class="h-3.5 w-3.5" />
          </router-link>
        </nav>
      </div>
    </header>

    <main>
      <!-- Hero -->
      <section class="fb-hero">
        <div class="fb-hero-grid" aria-hidden="true"></div>
        <div class="fb-shell">
          <p class="fb-kicker fb-rise" style="--d: 0s">FEEDBACK · 我有意见</p>
          <h1 class="fb-hero-title fb-rise" style="--d: 0.08s">
            你的每一句话，<br />
            都会<span class="fb-serif">被看见</span>。
          </h1>
          <p class="fb-hero-lead fb-rise" style="--d: 0.16s">
            用得不顺手、想要一个新功能、遇到订单问题，或者只是有个想法——
            都可以写在这里。被采纳的意见，会出现在本页最下方。
          </p>
          <div class="fb-hero-tags fb-rise" style="--d: 0.24s">
            <span>产品建议</span>
            <span>使用体验</span>
            <span>订单问题</span>
            <span>合作想法</span>
          </div>
        </div>
      </section>

      <!-- Form + process -->
      <section class="fb-shell fb-main">
        <div class="fb-grid">
          <div ref="formCard" class="fb-form-card fb-rise" style="--d: 0.28s">
            <div v-if="submitted" class="fb-success" aria-live="polite">
              <span class="fb-success-mark">
                <Check class="h-7 w-7" :stroke-width="2.4" />
              </span>
              <p class="fb-success-kicker">THANK YOU</p>
              <h2 class="fb-success-title">已经收到你的意见。</h2>
              <p class="fb-success-copy">
                我们会认真读完。如果这条意见被采纳，它会出现在下方「已被采纳」区域。
              </p>
              <div class="fb-success-actions">
                <button type="button" class="fb-submit" @click="writeAnother">
                  再写一条
                  <ArrowUpRight class="h-4 w-4" />
                </button>
                <router-link to="/showroom" class="fb-text-link">去展厅看看</router-link>
              </div>
            </div>

            <template v-else>
              <div class="fb-form-head">
                <div>
                  <p class="fb-card-kicker">TELL US</p>
                  <h2 class="fb-form-title">说说你的想法</h2>
                </div>
                <span class="fb-optional">订单号可选</span>
              </div>

              <form class="fb-form" novalidate @submit.prevent="submit">
                <div class="fb-field">
                  <label for="fb-order-no">
                    订单号
                    <span>选填</span>
                  </label>
                  <div class="fb-input-wrap">
                    <ShoppingBag class="fb-input-icon h-4 w-4" />
                    <input
                      id="fb-order-no"
                      v-model="form.order_no"
                      type="text"
                      maxlength="64"
                      autocomplete="off"
                      placeholder="例如 TP202610210001"
                    />
                  </div>
                  <p class="fb-hint">如果和某笔订单有关，可以填上，方便我们定位。</p>
                </div>

                <div class="fb-field">
                  <label for="fb-content">
                    你的意见
                    <span>必填</span>
                  </label>
                  <textarea
                    id="fb-content"
                    ref="contentEl"
                    v-model="form.content"
                    :maxlength="MAX_LEN"
                    rows="7"
                    placeholder="可以具体一点：你在哪个页面、做什么操作时遇到了什么，或者你希望我们增加什么。"
                    :aria-invalid="!!contentError"
                    aria-describedby="fb-content-help"
                  ></textarea>
                  <div class="fb-field-foot">
                    <p id="fb-content-help" class="fb-hint" :class="{ 'fb-hint-error': contentError }">
                      {{ contentError || '至少写 5 个字，越具体越有帮助。' }}
                    </p>
                    <span class="fb-count" :class="{ 'fb-count-warn': form.content.length > MAX_LEN - 100 }">
                      {{ form.content.length }} / {{ MAX_LEN }}
                    </span>
                  </div>
                </div>

                <div class="fb-form-actions">
                  <p class="fb-privacy-note">
                    提交即表示你同意我们按
                    <router-link to="/privacy">隐私政策</router-link>
                    处理这条内容。
                  </p>
                  <button type="submit" class="fb-submit" :disabled="submitting">
                    <Loader2 v-if="submitting" class="h-4 w-4 fb-spin" />
                    <Send v-else class="h-4 w-4" />
                    {{ submitting ? '正在提交' : '提交意见' }}
                  </button>
                </div>
              </form>
            </template>
          </div>

          <aside class="fb-aside">
            <div class="fb-process fb-rise" style="--d: 0.34s">
              <p class="fb-process-kicker">WHAT HAPPENS NEXT</p>
              <h2>我们会怎么处理</h2>
              <ol class="fb-steps">
                <li>
                  <span class="fb-step-num">01</span>
                  <div>
                    <strong>收下</strong>
                    <p>意见进入后台列表，订单号会帮我们找到相关记录。</p>
                  </div>
                </li>
                <li>
                  <span class="fb-step-num">02</span>
                  <div>
                    <strong>读完</strong>
                    <p>我们会判断它是个例、共性问题，还是一个新的方向。</p>
                  </div>
                </li>
                <li>
                  <span class="fb-step-num">03</span>
                  <div>
                    <strong>采纳</strong>
                    <p>管理员在后台勾选「已采纳」，这条意见就有了结果。</p>
                  </div>
                </li>
                <li>
                  <span class="fb-step-num">04</span>
                  <div>
                    <strong>展示</strong>
                    <p>被采纳的内容会出现在本页下方，所有人都能看到。</p>
                  </div>
                </li>
              </ol>
            </div>

            <div class="fb-note fb-rise" style="--d: 0.4s">
              <span class="fb-note-icon">
                <ShieldCheck class="h-5 w-5" />
              </span>
              <div>
                <strong>不会公开你的订单号</strong>
                <p>公开展示时只显示意见内容与采纳时间，订单号仅管理员可见。</p>
              </div>
            </div>
          </aside>
        </div>
      </section>

      <!-- Adopted feedback -->
      <section class="fb-adopted">
        <div class="fb-shell">
          <div class="fb-adopted-head">
            <div>
              <p class="fb-kicker">ADOPTED · 已被采纳</p>
              <h2>这些意见，<span class="fb-serif">真的变成了改变。</span></h2>
            </div>
            <span v-if="adopted.length" class="fb-adopted-count">{{ adopted.length }} 条</span>
          </div>

          <div v-if="loadingAdopted" class="fb-adopted-grid" aria-label="正在加载已采纳意见">
            <div v-for="n in 3" :key="n" class="fb-quote-card fb-skeleton">
              <span></span><span></span><span></span>
            </div>
          </div>

          <div v-else-if="adopted.length" class="fb-adopted-grid">
            <article v-for="(item, index) in adopted" :key="item.id" class="fb-quote-card">
              <div class="fb-quote-top">
                <span class="fb-quote-num">{{ String(index + 1).padStart(2, '0') }}</span>
                <span class="fb-adopted-badge">
                  <Check class="h-3 w-3" />
                  已采纳
                </span>
              </div>
              <p class="fb-quote-text">{{ item.content }}</p>
              <p class="fb-quote-date">
                <Clock class="h-3.5 w-3.5" />
                {{ formatDate(item.adopted_at || item.created_at) }}
              </p>
            </article>
          </div>

          <div v-else class="fb-empty">
            <span class="fb-empty-icon">
              <Inbox class="h-6 w-6" />
            </span>
            <h3>还没有已采纳的意见</h3>
            <p>第一条被采纳的意见，可能就来自你。</p>
            <button type="button" class="fb-text-link" @click="focusForm">写下你的想法</button>
          </div>
        </div>
      </section>
    </main>

    <footer class="fb-footer">
      <div class="fb-shell fb-footer-inner">
        <div class="fb-footer-brand">
          <span class="fb-brand-mark">
            <img src="/tapiito-icon-white.png" alt="Tapiito" class="h-5 w-5" draggable="false" />
          </span>
          <span>Tapiito</span>
          <span class="fb-footer-tagline">Tap it alive.</span>
        </div>
        <nav class="fb-footer-links" aria-label="页脚导航">
          <router-link to="/">首页</router-link>
          <router-link to="/about">关于我们</router-link>
          <router-link to="/privacy">隐私政策</router-link>
          <router-link to="/disclaimer">免责声明</router-link>
        </nav>
        <p class="fb-footer-copy">© {{ new Date().getFullYear() }} Tapiito</p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { computed, nextTick, onMounted, onUnmounted, reactive, ref } from 'vue'
import {
  ArrowUpRight,
  Check,
  Clock,
  Inbox,
  Loader2,
  Nfc,
  Send,
  ShieldCheck,
  ShoppingBag,
} from 'lucide-vue-next'
import { createFeedback, fetchAdoptedFeedback } from '../api/feedback'

const MAX_LEN = 2000

const form = reactive({
  order_no: '',
  content: '',
})
const touched = ref(false)
const submitting = ref(false)
const submitted = ref(false)
const adopted = ref([])
const loadingAdopted = ref(true)
const adoptedError = ref(false)
const formCard = ref(null)
const contentEl = ref(null)

const contentError = computed(() => {
  if (!touched.value) return ''
  if (form.content.trim().length < 5) return '意见内容至少需要 5 个字。'
  return ''
})

const formatDate = (value) => {
  if (!value) return ''
  const raw = String(value)
  const normalized = /Z$|[+-]\d{2}:?\d{2}$/.test(raw) ? raw : `${raw}Z`
  const date = new Date(normalized)
  if (Number.isNaN(date.getTime())) return raw.slice(0, 10)
  return `${date.getFullYear()}.${String(date.getMonth() + 1).padStart(2, '0')}.${String(
    date.getDate()
  ).padStart(2, '0')}`
}

const loadAdopted = async () => {
  loadingAdopted.value = true
  adoptedError.value = false
  try {
    const { data } = await fetchAdoptedFeedback({ limit: 12 })
    adopted.value = Array.isArray(data) ? data : []
  } catch {
    adoptedError.value = true
    adopted.value = []
  } finally {
    loadingAdopted.value = false
  }
}

const submit = async () => {
  touched.value = true
  if (contentError.value) {
    contentEl.value?.focus()
    return
  }

  submitting.value = true
  try {
    await createFeedback({
      order_no: form.order_no.trim() || null,
      content: form.content.trim(),
    })
    submitted.value = true
    form.order_no = ''
    form.content = ''
    touched.value = false
    await nextTick()
    formCard.value?.scrollIntoView({ behavior: 'smooth', block: 'center' })
  } catch {
    // 全局请求层已经给出错误提示
  } finally {
    submitting.value = false
  }
}

const writeAnother = () => {
  submitted.value = false
  nextTick(() => contentEl.value?.focus())
}

const focusForm = () => {
  formCard.value?.scrollIntoView({ behavior: 'smooth', block: 'center' })
  nextTick(() => contentEl.value?.focus())
}

onMounted(() => {
  document.title = '我有意见 · Tapiito'
  loadAdopted()
})

onUnmounted(() => {
  document.title = 'Tapiito'
})
</script>

<style scoped>
.fb-page {
  --ink: #14181d;
  --cream: #fbfaf7;
  --orange: #ff6b35;
  --sun: #ffd166;
  position: relative;
  min-height: 100vh;
  overflow-x: clip;
  background: var(--cream);
  color: var(--ink);
  font-family: 'Geist', -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}

.fb-grain {
  position: fixed;
  inset: 0;
  z-index: 0;
  pointer-events: none;
  opacity: 0.32;
  background-image: url("data:image/svg+xml,%3Csvg viewBox='0 0 180 180' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='.82' numOctaves='3' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)' opacity='.045'/%3E%3C/svg%3E");
}

.fb-glow {
  position: absolute;
  z-index: 0;
  border-radius: 999px;
  filter: blur(12px);
  pointer-events: none;
}

.fb-glow-orange {
  top: -170px;
  left: -140px;
  width: 520px;
  height: 520px;
  background: radial-gradient(circle, rgba(255, 107, 53, 0.16), transparent 68%);
}

.fb-glow-sun {
  top: 70px;
  right: -180px;
  width: 540px;
  height: 540px;
  background: radial-gradient(circle, rgba(255, 209, 102, 0.22), transparent 70%);
}

.fb-shell {
  position: relative;
  z-index: 1;
  width: min(1180px, calc(100% - 48px));
  margin: 0 auto;
}

.fb-header {
  position: sticky;
  top: 0;
  z-index: 30;
  border-bottom: 1px solid rgba(20, 24, 29, 0.07);
  background: rgba(251, 250, 247, 0.84);
  backdrop-filter: blur(16px);
}

.fb-nav {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 64px;
}

.fb-brand {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  font-size: 18px;
  font-weight: 650;
  letter-spacing: -0.03em;
  color: var(--ink);
}

.fb-brand-mark {
  display: inline-flex;
  width: 32px;
  height: 32px;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  background: var(--ink);
  color: var(--orange);
}

.fb-nav-links {
  display: flex;
  align-items: center;
  gap: 4px;
}

.fb-nav-link {
  display: inline-flex;
  align-items: center;
  border-radius: 9px;
  padding: 8px 12px;
  font-size: 12.5px;
  font-weight: 500;
  color: rgba(20, 24, 29, 0.55);
  transition: color 0.2s ease, background 0.2s ease;
}

.fb-nav-link:hover {
  color: var(--ink);
  background: rgba(20, 24, 29, 0.045);
}

.fb-nav-active {
  color: var(--ink);
  background: rgba(255, 107, 53, 0.1);
  font-weight: 650;
}

.fb-nav-cta {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  margin-left: 4px;
  border-radius: 9px;
  background: var(--ink);
  padding: 8px 13px;
  font-size: 12.5px;
  font-weight: 600;
  color: #fff;
  transition: transform 0.2s ease, background 0.2s ease;
}

.fb-nav-cta:hover {
  background: var(--orange);
  transform: translateY(-1px);
}

.fb-hero {
  position: relative;
  padding: 94px 0 62px;
}

.fb-hero-grid {
  position: absolute;
  inset: 0;
  opacity: 0.45;
  background-image: linear-gradient(rgba(20, 24, 29, 0.055) 1px, transparent 1px),
    linear-gradient(90deg, rgba(20, 24, 29, 0.055) 1px, transparent 1px);
  background-size: 46px 46px;
  mask-image: linear-gradient(to bottom, rgba(0, 0, 0, 0.8), transparent 88%);
}

.fb-kicker {
  font-size: 11px;
  font-weight: 650;
  letter-spacing: 0.2em;
  color: var(--orange);
}

.fb-hero-title {
  max-width: 820px;
  margin-top: 20px;
  font-size: clamp(44px, 6.7vw, 82px);
  font-weight: 680;
  line-height: 1.04;
  letter-spacing: -0.055em;
}

.fb-serif {
  font-family: 'Playfair Display', Georgia, 'Songti SC', serif;
  font-style: italic;
  font-weight: 500;
  color: var(--orange);
  letter-spacing: -0.04em;
}

.fb-hero-lead {
  max-width: 620px;
  margin-top: 24px;
  font-size: 16px;
  line-height: 1.9;
  color: rgba(20, 24, 29, 0.58);
}

.fb-hero-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-top: 26px;
}

.fb-hero-tags span {
  border: 1px solid rgba(20, 24, 29, 0.1);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.55);
  padding: 7px 13px;
  font-size: 12px;
  color: rgba(20, 24, 29, 0.55);
}

.fb-main {
  padding-bottom: 92px;
}

.fb-grid {
  display: grid;
  grid-template-columns: minmax(0, 1.32fr) minmax(300px, 0.68fr);
  gap: 24px;
  align-items: start;
}

.fb-form-card {
  border: 1px solid rgba(20, 24, 29, 0.09);
  border-radius: 26px;
  background: rgba(255, 255, 255, 0.84);
  padding: 32px;
  box-shadow: 0 24px 60px -42px rgba(20, 24, 29, 0.35);
  backdrop-filter: blur(8px);
}

.fb-form-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 16px;
}

.fb-card-kicker {
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.2em;
  color: var(--orange);
}

.fb-form-title {
  margin-top: 7px;
  font-size: 28px;
  font-weight: 650;
  letter-spacing: -0.04em;
}

.fb-optional {
  flex: 0 0 auto;
  border-radius: 999px;
  background: rgba(20, 24, 29, 0.05);
  padding: 6px 10px;
  font-size: 11px;
  color: rgba(20, 24, 29, 0.45);
}

.fb-form {
  margin-top: 28px;
}

.fb-field + .fb-field {
  margin-top: 22px;
}

.fb-field label {
  display: flex;
  align-items: center;
  gap: 6px;
  margin-bottom: 9px;
  font-size: 13px;
  font-weight: 600;
}

.fb-field label span {
  font-size: 10px;
  font-weight: 500;
  color: rgba(20, 24, 29, 0.38);
}

.fb-input-wrap {
  position: relative;
}

.fb-input-icon {
  position: absolute;
  top: 50%;
  left: 14px;
  color: rgba(20, 24, 29, 0.32);
  transform: translateY(-50%);
  pointer-events: none;
}

.fb-field input,
.fb-field textarea {
  width: 100%;
  border: 1px solid rgba(20, 24, 29, 0.12);
  border-radius: 13px;
  background: rgba(251, 250, 247, 0.7);
  padding: 13px 15px;
  font-family: inherit;
  font-size: 14px;
  color: var(--ink);
  outline: none;
  transition: border-color 0.2s ease, box-shadow 0.2s ease, background 0.2s ease;
}

.fb-field input {
  padding-left: 40px;
}

.fb-field textarea {
  min-height: 180px;
  line-height: 1.8;
  resize: vertical;
}

.fb-field input::placeholder,
.fb-field textarea::placeholder {
  color: rgba(20, 24, 29, 0.3);
}

.fb-field input:focus,
.fb-field textarea:focus {
  border-color: rgba(255, 107, 53, 0.55);
  background: #fff;
  box-shadow: 0 0 0 4px rgba(255, 107, 53, 0.08);
}

.fb-field-foot {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 12px;
  margin-top: 8px;
}

.fb-hint {
  font-size: 11px;
  line-height: 1.6;
  color: rgba(20, 24, 29, 0.38);
}

.fb-hint-error {
  color: #d94f1a;
}

.fb-count {
  flex: 0 0 auto;
  font-size: 11px;
  color: rgba(20, 24, 29, 0.35);
  font-variant-numeric: tabular-nums;
}

.fb-count-warn {
  color: var(--orange);
}

.fb-form-actions {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 18px;
  margin-top: 26px;
  padding-top: 22px;
  border-top: 1px solid rgba(20, 24, 29, 0.08);
}

.fb-privacy-note {
  max-width: 300px;
  font-size: 11px;
  line-height: 1.7;
  color: rgba(20, 24, 29, 0.4);
}

.fb-privacy-note a {
  color: rgba(20, 24, 29, 0.68);
  text-decoration: underline;
  text-underline-offset: 2px;
}

.fb-submit {
  display: inline-flex;
  flex: 0 0 auto;
  align-items: center;
  gap: 7px;
  border: 0;
  border-radius: 11px;
  background: var(--orange);
  padding: 12px 20px;
  font-family: inherit;
  font-size: 13.5px;
  font-weight: 650;
  color: #fff;
  cursor: pointer;
  transition: transform 0.2s ease, background 0.2s ease, opacity 0.2s ease;
}

.fb-submit:hover:not(:disabled) {
  background: #f25a22;
  transform: translateY(-1px);
}

.fb-submit:disabled {
  cursor: wait;
  opacity: 0.68;
}

.fb-spin {
  animation: fb-spin 0.9s linear infinite;
}

.fb-success {
  padding: 26px 6px;
  text-align: center;
}

.fb-success-mark {
  display: inline-flex;
  width: 62px;
  height: 62px;
  align-items: center;
  justify-content: center;
  border-radius: 999px;
  background: rgba(22, 163, 74, 0.1);
  color: #15803d;
}

.fb-success-kicker {
  margin-top: 20px;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.2em;
  color: #15803d;
}

.fb-success-title {
  margin-top: 8px;
  font-size: 30px;
  font-weight: 650;
  letter-spacing: -0.04em;
}

.fb-success-copy {
  max-width: 390px;
  margin: 12px auto 0;
  font-size: 14px;
  line-height: 1.8;
  color: rgba(20, 24, 29, 0.55);
}

.fb-success-actions {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: center;
  gap: 16px;
  margin-top: 24px;
}

.fb-text-link {
  border: 0;
  background: none;
  padding: 0;
  font-family: inherit;
  font-size: 12.5px;
  font-weight: 550;
  color: rgba(20, 24, 29, 0.55);
  text-decoration: underline;
  text-underline-offset: 3px;
  cursor: pointer;
}

.fb-aside {
  display: grid;
  gap: 16px;
}

.fb-process {
  border-radius: 24px;
  background: var(--ink);
  padding: 30px 26px;
  color: #fff;
}

.fb-process-kicker {
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.18em;
  color: var(--orange);
}

.fb-process h2 {
  margin-top: 9px;
  font-size: 22px;
  font-weight: 600;
  letter-spacing: -0.03em;
}

.fb-steps {
  display: grid;
  gap: 18px;
  margin-top: 26px;
  padding: 0;
  list-style: none;
}

.fb-steps li {
  display: flex;
  gap: 13px;
}

.fb-step-num {
  flex: 0 0 auto;
  padding-top: 2px;
  font-family: 'Playfair Display', Georgia, serif;
  font-size: 13px;
  font-style: italic;
  color: rgba(255, 255, 255, 0.35);
}

.fb-steps strong {
  display: block;
  font-size: 13px;
  font-weight: 600;
}

.fb-steps p {
  margin-top: 4px;
  font-size: 11.5px;
  line-height: 1.7;
  color: rgba(255, 255, 255, 0.48);
}

.fb-note {
  display: flex;
  gap: 12px;
  border: 1px solid rgba(20, 24, 29, 0.08);
  border-radius: 18px;
  background: rgba(255, 255, 255, 0.58);
  padding: 18px;
}

.fb-note-icon {
  display: inline-flex;
  width: 34px;
  height: 34px;
  flex: 0 0 auto;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  background: rgba(255, 107, 53, 0.1);
  color: var(--orange);
}

.fb-note strong {
  display: block;
  font-size: 12px;
  font-weight: 650;
}

.fb-note p {
  margin-top: 4px;
  font-size: 11px;
  line-height: 1.7;
  color: rgba(20, 24, 29, 0.45);
}

.fb-adopted {
  border-top: 1px solid rgba(20, 24, 29, 0.07);
  background: rgba(255, 255, 255, 0.38);
  padding: 72px 0 90px;
}

.fb-adopted-head {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  gap: 20px;
}

.fb-adopted-head h2 {
  max-width: 720px;
  margin-top: 10px;
  font-size: clamp(30px, 4vw, 44px);
  font-weight: 650;
  line-height: 1.15;
  letter-spacing: -0.045em;
}

.fb-adopted-count {
  flex: 0 0 auto;
  border-radius: 999px;
  background: var(--ink);
  padding: 7px 12px;
  font-size: 11px;
  font-weight: 600;
  color: #fff;
}

.fb-adopted-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 14px;
  margin-top: 32px;
}

.fb-quote-card {
  position: relative;
  border: 1px solid rgba(20, 24, 29, 0.08);
  border-radius: 18px;
  background: rgba(251, 250, 247, 0.88);
  padding: 22px;
  transition: transform 0.25s ease, border-color 0.25s ease, box-shadow 0.25s ease;
}

.fb-quote-card:hover {
  border-color: rgba(255, 107, 53, 0.25);
  box-shadow: 0 18px 40px -30px rgba(20, 24, 29, 0.35);
  transform: translateY(-3px);
}

.fb-quote-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.fb-quote-num {
  font-family: 'Playfair Display', Georgia, serif;
  font-size: 15px;
  font-style: italic;
  color: rgba(20, 24, 29, 0.3);
}

.fb-adopted-badge {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  border-radius: 999px;
  background: rgba(22, 163, 74, 0.09);
  padding: 5px 9px;
  font-size: 10px;
  font-weight: 650;
  color: #15803d;
}

.fb-quote-text {
  margin-top: 18px;
  font-size: 14px;
  line-height: 1.85;
  color: rgba(20, 24, 29, 0.75);
  white-space: pre-wrap;
  overflow-wrap: anywhere;
}

.fb-quote-date {
  display: flex;
  align-items: center;
  gap: 5px;
  margin-top: 18px;
  font-size: 10.5px;
  color: rgba(20, 24, 29, 0.34);
  font-variant-numeric: tabular-nums;
}

.fb-skeleton span {
  display: block;
  height: 10px;
  margin-top: 12px;
  border-radius: 999px;
  background: linear-gradient(90deg, rgba(20, 24, 29, 0.05), rgba(20, 24, 29, 0.1), rgba(20, 24, 29, 0.05));
  background-size: 200% 100%;
  animation: fb-shimmer 1.4s ease infinite;
}

.fb-skeleton span:first-child {
  width: 28%;
  margin-top: 0;
}

.fb-skeleton span:nth-child(2) {
  width: 92%;
}

.fb-skeleton span:nth-child(3) {
  width: 64%;
}

.fb-empty {
  margin-top: 30px;
  border: 1px dashed rgba(20, 24, 29, 0.14);
  border-radius: 20px;
  padding: 48px 24px;
  text-align: center;
}

.fb-empty-icon {
  display: inline-flex;
  width: 48px;
  height: 48px;
  align-items: center;
  justify-content: center;
  border-radius: 14px;
  background: rgba(20, 24, 29, 0.05);
  color: rgba(20, 24, 29, 0.35);
}

.fb-empty h3 {
  margin-top: 16px;
  font-size: 18px;
  font-weight: 600;
}

.fb-empty p {
  margin-top: 6px;
  font-size: 12.5px;
  color: rgba(20, 24, 29, 0.42);
}

.fb-empty .fb-text-link {
  display: inline-block;
  margin-top: 14px;
  color: var(--orange);
}

.fb-footer {
  border-top: 1px solid rgba(20, 24, 29, 0.07);
}

.fb-footer-inner {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: space-between;
  gap: 14px 20px;
  padding-top: 26px;
  padding-bottom: 26px;
}

.fb-footer-brand {
  display: flex;
  align-items: center;
  gap: 9px;
  font-size: 13px;
  font-weight: 600;
}

.fb-footer-brand .fb-brand-mark {
  width: 28px;
  height: 28px;
  border-radius: 8px;
}

.fb-footer-tagline {
  font-weight: 400;
  color: rgba(20, 24, 29, 0.38);
}

.fb-footer-links {
  display: flex;
  flex-wrap: wrap;
  gap: 8px 18px;
  font-size: 11.5px;
  color: rgba(20, 24, 29, 0.45);
}

.fb-footer-links a:hover {
  color: var(--ink);
}

.fb-footer-copy {
  font-size: 11px;
  color: rgba(20, 24, 29, 0.3);
}

.fb-rise {
  animation: fb-rise 0.8s cubic-bezier(0.22, 1, 0.36, 1) both;
  animation-delay: var(--d, 0s);
}

@keyframes fb-rise {
  from {
    opacity: 0;
    transform: translateY(18px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fb-spin {
  to {
    transform: rotate(360deg);
  }
}

@keyframes fb-shimmer {
  from {
    background-position: 200% 0;
  }
  to {
    background-position: -200% 0;
  }
}

@media (max-width: 960px) {
  .fb-grid {
    grid-template-columns: 1fr;
  }

  .fb-aside {
    grid-template-columns: minmax(0, 1fr) minmax(0, 0.72fr);
  }
}

@media (max-width: 720px) {
  .fb-shell {
    width: min(100% - 32px, 1180px);
  }

  .fb-hide-mobile {
    display: none;
  }

  .fb-hero {
    padding: 66px 0 44px;
  }

  .fb-hero-title {
    font-size: clamp(40px, 12vw, 58px);
  }

  .fb-hero-lead {
    font-size: 14px;
    line-height: 1.85;
  }

  .fb-main {
    padding-bottom: 64px;
  }

  .fb-form-card {
    border-radius: 20px;
    padding: 22px;
  }

  .fb-form-head {
    align-items: center;
  }

  .fb-form-title {
    font-size: 24px;
  }

  .fb-form-actions {
    align-items: stretch;
    flex-direction: column;
  }

  .fb-privacy-note {
    max-width: none;
  }

  .fb-submit {
    justify-content: center;
  }

  .fb-aside {
    grid-template-columns: 1fr;
  }

  .fb-adopted {
    padding: 52px 0 66px;
  }

  .fb-adopted-head {
    align-items: flex-start;
    flex-direction: column;
  }

  .fb-adopted-grid {
    grid-template-columns: 1fr;
  }

  .fb-footer-inner {
    align-items: flex-start;
    flex-direction: column;
  }
}

@media (max-width: 420px) {
  .fb-nav {
    height: 58px;
  }

  .fb-nav-cta {
    padding: 8px 10px;
    font-size: 11.5px;
  }

  .fb-hero-tags span {
    padding: 6px 10px;
    font-size: 11px;
  }

  .fb-form-card {
    padding: 18px;
  }

  .fb-optional {
    display: none;
  }
}

@media (prefers-reduced-motion: reduce) {
  .fb-rise,
  .fb-spin,
  .fb-skeleton span {
    animation: none;
  }

  .fb-quote-card,
  .fb-nav-cta,
  .fb-submit {
    transition: none;
  }
}
</style>
