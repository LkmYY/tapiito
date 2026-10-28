<template>
  <div class="cu-neon-root min-h-screen">
    <div class="cu-neon-ambient" aria-hidden="true">
      <span class="cu-neon-glow cu-neon-glow-a"></span>
      <span class="cu-neon-glow cu-neon-glow-b"></span>
      <span class="cu-neon-tube cu-neon-tube-a"></span>
      <span class="cu-neon-tube cu-neon-tube-b"></span>
      <span class="cu-neon-scan"></span>
    </div>

    <div class="cu-neon-wrap mx-auto w-full max-w-md">
      <header class="cu-neon-top cu-neon-in">
        <div class="cu-neon-brand">
          <span class="cu-neon-brand-mark">
            <svg viewBox="0 0 24 24" width="17" height="17" fill="none" aria-hidden="true">
              <path d="M12 2l2.1 7.9L22 12l-7.9 2.1L12 22l-2.1-7.9L2 12l7.9-2.1L12 2Z" stroke="currentColor" stroke-width="1.2" stroke-linejoin="round" />
            </svg>
          </span>
          <span class="cu-neon-brand-text">NEON STUDIO<em>夜行创意事务所</em></span>
        </div>
        <span class="cu-neon-serial">CARD / {{ serial }}</span>
      </header>

      <p class="cu-neon-status cu-neon-in cu-neon-d1">
        <span class="cu-neon-pulse"></span>
        OPEN AFTER DARK
        <span class="cu-neon-status-sep">·</span>
        22:00 — 02:00
      </p>

      <section class="cu-neon-hero">
        <div class="cu-neon-stage cu-neon-in cu-neon-d1">
          <img v-if="images.length" :src="images[0]" :alt="companyName + ' 工作室'" class="cu-neon-photo" />
          <div v-else class="cu-neon-sign" aria-hidden="true">
            <span class="cu-neon-sign-orbit cu-neon-sign-orbit-a"></span>
            <span class="cu-neon-sign-orbit cu-neon-sign-orbit-b"></span>
            <span class="cu-neon-sign-chip">OPEN LATE</span>
            <span class="cu-neon-sign-initial">{{ initial }}</span>
            <span class="cu-neon-sign-word">CREATIVE / NIGHT SHIFT</span>
          </div>
          <span class="cu-neon-stage-shine" aria-hidden="true"></span>
          <span class="cu-neon-stage-tag">{{ images.length ? 'SELECTED WORK' : 'AFTER HOURS' }}</span>
        </div>

        <div class="cu-neon-title">
          <p class="cu-neon-kicker cu-neon-in cu-neon-d2">Creative practice · 紫青之间</p>
          <h1 class="cu-neon-name cu-neon-in cu-neon-d2">{{ companyName }}</h1>
          <div class="cu-neon-desc-row cu-neon-in cu-neon-d3">
            <span class="cu-neon-desc-bar"></span>
            <p>{{ description }}</p>
          </div>
        </div>
      </section>

      <section class="cu-neon-services cu-neon-in cu-neon-d3">
        <div class="cu-neon-sec-head">
          <h2>服务清单 <span>/ SERVICES</span></h2>
          <span class="cu-neon-count">{{ pad2(services.length) }} ITEMS</span>
        </div>
        <div v-if="services.length" class="cu-neon-list">
          <article v-for="(service, i) in services" :key="i" class="cu-neon-item">
            <span class="cu-neon-item-num">{{ pad2(i + 1) }}</span>
            <span class="cu-neon-item-line"></span>
            <h3>{{ service.name }}</h3>
            <span class="cu-neon-item-dot"></span>
          </article>
        </div>
      </section>

      <section class="cu-neon-contact cu-neon-in cu-neon-d4">
        <p class="cu-neon-contact-kicker">CALL THE STUDIO / 夜间联络</p>
        <a v-if="phone" :href="'tel:' + phone" class="cu-neon-phone">
          <Phone :size="15" />
          <span>{{ phone }}</span>
          <ArrowUpRight :size="16" class="cu-neon-arrow" />
        </a>
        <a v-if="email" :href="'mailto:' + email" class="cu-neon-mail">{{ email }}</a>

        <div class="cu-neon-card">
          <span class="cu-neon-card-chip"></span>
          <div class="cu-neon-card-id">
            <span>NFC CARD</span>
            <strong>{{ serial }}</strong>
          </div>
          <p class="cu-neon-card-note">TAP · TALK · MAKE<br />BEYOND MIDNIGHT</p>
          <span class="cu-neon-card-bars" aria-hidden="true"></span>
        </div>
      </section>

      <footer class="cu-neon-foot cu-neon-in cu-neon-d4">
        <span>{{ companyName }} · 夜间营业</span>
        <span>N° {{ serial }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, Phone } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const clean = (value) => (typeof value === 'string' ? value.trim() : value != null ? String(value).trim() : '')
const pad2 = (n) => String(n).padStart(2, '0')

const companyName = computed(() => clean(props.payload?.companyName) || '未命名工作室')
const description = computed(() => clean(props.payload?.description))
const phone = computed(() => clean(props.payload?.phone))
const email = computed(() => clean(props.payload?.email))
const initial = computed(() => companyName.value.slice(0, 1))
const serial = computed(() => clean(props.nfcId).slice(-8).toUpperCase() || '—')

const images = computed(() => {
  const list = props.payload?.images
  return Array.isArray(list) ? list.filter((src) => typeof src === 'string' && src.trim()) : []
})

const services = computed(() => {
  const list = props.payload?.services
  return Array.isArray(list) ? list.filter((item) => item && clean(item.name)) : []
})
</script>

<style scoped>
.cu-neon-root { position: relative; overflow: hidden; background: #06030f; color: #f5f1ff; font-family: 'PingFang SC', 'Microsoft YaHei', 'Helvetica Neue', sans-serif; }
.cu-neon-ambient { position: absolute; inset: 0; pointer-events: none; }
.cu-neon-glow { position: absolute; border-radius: 50%; filter: blur(12px); }
.cu-neon-glow-a { top: -90px; left: -80px; width: 280px; height: 280px; background: radial-gradient(circle, rgba(139, 61, 255, 0.34), transparent 68%); animation: cu-neon-drift 11s ease-in-out infinite alternate; }
.cu-neon-glow-b { top: 38%; right: -110px; width: 300px; height: 300px; background: radial-gradient(circle, rgba(28, 211, 238, 0.19), transparent 68%); animation: cu-neon-drift 14s ease-in-out infinite alternate-reverse; }
.cu-neon-tube { position: absolute; top: 0; bottom: 0; width: 1px; background: linear-gradient(180deg, transparent, rgba(192, 132, 252, 0.5), rgba(103, 232, 249, 0.35), transparent); box-shadow: 0 0 14px rgba(192, 132, 252, 0.5); }
.cu-neon-tube-a { left: 10px; }
.cu-neon-tube-b { right: 10px; opacity: 0.45; }
.cu-neon-scan { position: absolute; inset: 0; background: repeating-linear-gradient(180deg, rgba(255, 255, 255, 0.025) 0 1px, transparent 1px 5px); opacity: 0.55; }
.cu-neon-wrap { position: relative; z-index: 2; padding: 20px 20px 30px; }
.cu-neon-top, .cu-neon-foot { display: flex; align-items: center; justify-content: space-between; gap: 12px; }
.cu-neon-top { border-bottom: 1px solid rgba(192, 132, 252, 0.22); padding-bottom: 12px; }
.cu-neon-brand { display: flex; align-items: center; gap: 9px; min-width: 0; }
.cu-neon-brand-mark { display: grid; place-items: center; width: 34px; height: 34px; flex-shrink: 0; border: 1px solid rgba(103, 232, 249, 0.5); border-radius: 9px; background: linear-gradient(145deg, rgba(192, 132, 252, 0.14), rgba(103, 232, 249, 0.06)); color: #a5f3fc; box-shadow: 0 0 18px rgba(103, 232, 249, 0.15), inset 0 0 10px rgba(192, 132, 252, 0.15); }
.cu-neon-brand-text { display: flex; flex-direction: column; font-family: 'Courier New', monospace; font-size: 9px; line-height: 1.35; letter-spacing: 0.18em; color: #e9d5ff; }
.cu-neon-brand-text em { font-style: normal; font-size: 8px; letter-spacing: 0.12em; color: #9d8bbd; }
.cu-neon-serial { flex-shrink: 0; font-family: 'Courier New', monospace; font-size: 7.5px; letter-spacing: 0.17em; color: #67e8f9; }
.cu-neon-status { display: flex; align-items: center; gap: 7px; margin-top: 13px; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.16em; color: #9d8bbd; }
.cu-neon-pulse { width: 5px; height: 5px; border-radius: 50%; background: #67e8f9; box-shadow: 0 0 9px #67e8f9; animation: cu-neon-pulse 2.4s ease-in-out infinite; }
.cu-neon-status-sep { color: rgba(192, 132, 252, 0.5); }
.cu-neon-hero { margin-top: 20px; }
.cu-neon-stage { position: relative; overflow: hidden; aspect-ratio: 4 / 3; border: 1px solid rgba(103, 232, 249, 0.35); border-radius: 17px; background: #0b0618; box-shadow: 0 0 0 4px rgba(192, 132, 252, 0.04), 0 0 34px -10px rgba(139, 61, 255, 0.65), inset 0 0 30px rgba(3, 0, 12, 0.8); }
.cu-neon-photo { position: absolute; inset: 0; width: 100%; height: 100%; object-fit: cover; filter: saturate(0.9) contrast(1.06) brightness(0.85); }
.cu-neon-stage::after { content: ''; position: absolute; inset: 0; background: linear-gradient(180deg, rgba(6, 3, 15, 0.05) 40%, rgba(6, 3, 15, 0.84)), linear-gradient(90deg, rgba(76, 29, 149, 0.22), transparent 60%); pointer-events: none; }
.cu-neon-sign { position: absolute; inset: 0; display: flex; flex-direction: column; align-items: center; justify-content: center; gap: 14px; background: radial-gradient(circle at 24% 16%, rgba(139, 61, 255, 0.32), transparent 54%), radial-gradient(circle at 80% 84%, rgba(28, 211, 238, 0.18), transparent 56%), linear-gradient(145deg, #130a28, #06030f 64%); }
.cu-neon-sign::before { content: ''; position: absolute; inset: 17px; border: 1px solid rgba(192, 132, 252, 0.55); border-radius: 12px; box-shadow: 0 0 18px rgba(192, 132, 252, 0.24), inset 0 0 24px rgba(103, 232, 249, 0.07); }
.cu-neon-sign-orbit { position: absolute; border: 1px solid rgba(103, 232, 249, 0.24); border-radius: 50%; transform: rotate(-18deg); }
.cu-neon-sign-orbit-a { width: 220px; height: 106px; }
.cu-neon-sign-orbit-b { width: 132px; height: 230px; transform: rotate(34deg); opacity: 0.55; }
.cu-neon-sign-chip { position: relative; z-index: 1; padding: 4px 10px; border: 1px solid rgba(103, 232, 249, 0.45); border-radius: 99px; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.24em; color: #a5f3fc; box-shadow: 0 0 14px rgba(103, 232, 249, 0.22); }
.cu-neon-sign-initial { position: relative; z-index: 1; font-family: 'Arial Black', 'PingFang SC', sans-serif; font-size: 92px; font-weight: 900; line-height: 0.9; background: linear-gradient(160deg, #f5f1ff 8%, #d8b4fe 44%, #67e8f9 98%); -webkit-background-clip: text; background-clip: text; color: transparent; filter: drop-shadow(0 0 18px rgba(168, 85, 247, 0.55)) drop-shadow(0 0 34px rgba(34, 211, 238, 0.28)); }
.cu-neon-sign-word { position: relative; z-index: 1; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.26em; color: #9d8bbd; }
.cu-neon-stage-shine { position: absolute; inset: 0; z-index: 2; background: linear-gradient(116deg, transparent 32%, rgba(255, 255, 255, 0.13) 46%, rgba(255, 255, 255, 0.02) 54%, transparent 64%); mix-blend-mode: screen; pointer-events: none; }
.cu-neon-stage-tag { position: absolute; z-index: 3; left: 13px; bottom: 12px; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.18em; color: rgba(245, 241, 255, 0.62); }
.cu-neon-title { padding-top: 22px; }
.cu-neon-kicker { font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.26em; text-transform: uppercase; color: #67e8f9; }
.cu-neon-name { margin-top: 10px; font-family: 'Arial Black', 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: clamp(40px, 13vw, 54px); font-weight: 900; line-height: 0.96; letter-spacing: -0.045em; color: #fbf8ff; text-shadow: 0 0 26px rgba(168, 85, 247, 0.42), 0 0 48px rgba(34, 211, 238, 0.18); overflow-wrap: anywhere; }
.cu-neon-desc-row { display: flex; align-items: flex-start; gap: 12px; margin-top: 17px; }
.cu-neon-desc-bar { width: 2px; height: 38px; flex-shrink: 0; border-radius: 99px; background: linear-gradient(180deg, #c084fc, #67e8f9); box-shadow: 0 0 12px rgba(103, 232, 249, 0.55); }
.cu-neon-desc-row p { max-width: 34ch; font-size: 12.5px; line-height: 1.85; color: #bdb0d4; }
.cu-neon-services { margin-top: 27px; }
.cu-neon-sec-head { display: flex; align-items: baseline; justify-content: space-between; padding-bottom: 9px; border-bottom: 1px solid rgba(192, 132, 252, 0.25); }
.cu-neon-sec-head h2 { font-size: 13px; font-weight: 600; letter-spacing: 0.08em; }
.cu-neon-sec-head h2 span { font-family: 'Courier New', monospace; font-size: 8px; font-weight: 400; letter-spacing: 0.16em; color: #8f7aad; }
.cu-neon-count { font-family: 'Courier New', monospace; font-size: 7.5px; letter-spacing: 0.16em; color: #67e8f9; }
.cu-neon-list { display: flex; flex-direction: column; gap: 7px; margin-top: 11px; }
.cu-neon-item { display: flex; align-items: center; gap: 11px; padding: 11px 12px; border: 1px solid rgba(192, 132, 252, 0.16); border-radius: 10px; background: linear-gradient(100deg, rgba(139, 61, 255, 0.08), rgba(28, 211, 238, 0.025)); transition: border-color 0.25s ease, box-shadow 0.25s ease, transform 0.25s ease; }
.cu-neon-item:hover { transform: translateX(2px); border-color: rgba(103, 232, 249, 0.55); box-shadow: 0 0 24px -8px rgba(103, 232, 249, 0.55); }
.cu-neon-item-num { flex-shrink: 0; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.1em; color: rgba(103, 232, 249, 0.7); }
.cu-neon-item-line { width: 14px; height: 1px; flex-shrink: 0; background: linear-gradient(90deg, #c084fc, #67e8f9); opacity: 0.55; }
.cu-neon-item h3 { min-width: 0; flex: 1; font-size: 14px; font-weight: 500; line-height: 1.45; letter-spacing: 0.02em; color: #ece4f8; overflow-wrap: anywhere; }
.cu-neon-item-dot { width: 5px; height: 5px; flex-shrink: 0; border-radius: 50%; background: #67e8f9; box-shadow: 0 0 9px rgba(103, 232, 249, 0.9); transition: transform 0.25s ease; }
.cu-neon-item:hover .cu-neon-item-dot { transform: scale(1.6); }
.cu-neon-empty { margin-top: 11px; padding: 20px 12px; border: 1px dashed rgba(192, 132, 252, 0.28); border-radius: 10px; text-align: center; font-size: 11px; color: #907fa9; }
.cu-neon-empty-tight { margin-top: 7px; padding: 5px 0; border: 0; text-align: left; }
.cu-neon-contact { margin-top: 28px; padding: 18px 15px 15px; border: 1px solid rgba(103, 232, 249, 0.25); border-radius: 15px; background: linear-gradient(155deg, rgba(139, 61, 255, 0.11), rgba(255, 255, 255, 0.025) 48%, rgba(28, 211, 238, 0.05)); box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.08), 0 0 28px -14px rgba(139, 61, 255, 0.7); }
.cu-neon-contact-kicker { font-family: 'Courier New', monospace; font-size: 7.5px; letter-spacing: 0.18em; color: #67e8f9; }
.cu-neon-phone { display: flex; align-items: center; gap: 8px; margin-top: 10px; padding-bottom: 4px; border-bottom: 1px solid rgba(192, 132, 252, 0.4); width: fit-content; max-width: 100%; font-size: 24px; font-weight: 700; letter-spacing: 0.01em; color: #fbf8ff; text-shadow: 0 0 18px rgba(103, 232, 249, 0.28); transition: color 0.2s ease, border-color 0.2s ease; }
.cu-neon-phone span { overflow-wrap: anywhere; }
.cu-neon-phone:hover { color: #a5f3fc; border-color: #67e8f9; }
.cu-neon-arrow { flex-shrink: 0; color: #67e8f9; }
.cu-neon-mail { display: inline-block; margin-top: 9px; font-family: 'Courier New', monospace; font-size: 10px; letter-spacing: 0.04em; color: #bfb1d6; text-decoration: underline; text-decoration-color: rgba(103, 232, 249, 0.4); text-underline-offset: 4px; overflow-wrap: anywhere; transition: color 0.2s ease; }
.cu-neon-mail:hover { color: #a5f3fc; }
.cu-neon-card { position: relative; display: flex; align-items: center; gap: 11px; margin-top: 16px; padding: 11px 12px; border-top: 1px solid rgba(255, 255, 255, 0.1); font-family: 'Courier New', monospace; }
.cu-neon-card-chip { width: 25px; height: 19px; flex-shrink: 0; border: 1px solid rgba(103, 232, 249, 0.45); border-radius: 4px; background: linear-gradient(135deg, rgba(255, 255, 255, 0.14), transparent 60%); box-shadow: inset 0 0 8px rgba(103, 232, 249, 0.12); }
.cu-neon-card-id { display: flex; flex-direction: column; gap: 2px; }
.cu-neon-card-id span { font-size: 6.5px; letter-spacing: 0.2em; color: #8f7aad; }
.cu-neon-card-id strong { font-size: 11px; letter-spacing: 0.13em; color: #e9d5ff; }
.cu-neon-card-note { margin-left: auto; text-align: right; font-size: 6.5px; line-height: 1.6; letter-spacing: 0.13em; color: #8f7aad; }
.cu-neon-card-bars { position: absolute; right: 0; bottom: -6px; width: 42px; height: 8px; opacity: 0.35; background: repeating-linear-gradient(90deg, #a5f3fc 0 1px, transparent 1px 3px); }
.cu-neon-foot { margin-top: 18px; padding-top: 11px; border-top: 1px solid rgba(255, 255, 255, 0.08); font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.14em; color: #756690; }
.cu-neon-in { animation: cu-neon-in 0.8s cubic-bezier(0.22, 1, 0.36, 1) both; }
.cu-neon-d1 { animation-delay: 0.08s; }
.cu-neon-d2 { animation-delay: 0.17s; }
.cu-neon-d3 { animation-delay: 0.26s; }
.cu-neon-d4 { animation-delay: 0.35s; }
@keyframes cu-neon-in {
  from { opacity: 0; transform: translateY(15px); }
  to { opacity: 1; transform: none; }
}
@keyframes cu-neon-pulse {
  0%, 100% { opacity: 0.45; transform: scale(0.9); }
  50% { opacity: 1; transform: scale(1.15); }
}
@keyframes cu-neon-drift {
  from { transform: translate3d(0, 0, 0); }
  to { transform: translate3d(18px, 22px, 0); }
}
@media (prefers-reduced-motion: reduce) {
  .cu-neon-in, .cu-neon-pulse, .cu-neon-glow { animation: none; }
  .cu-neon-item, .cu-neon-item-dot, .cu-neon-phone, .cu-neon-mail { transition: none; }
}
</style>
