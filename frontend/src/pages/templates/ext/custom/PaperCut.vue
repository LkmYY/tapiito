<template>
  <div class="cu-paper-root min-h-screen">
    <div class="cu-paper-ambient" aria-hidden="true">
      <span class="cu-paper-blob cu-paper-blob-a"></span>
      <span class="cu-paper-blob cu-paper-blob-b"></span>
      <span class="cu-paper-stitch cu-paper-stitch-a"></span>
      <span class="cu-paper-stitch cu-paper-stitch-b"></span>
    </div>

    <div class="cu-paper-wrap mx-auto w-full max-w-md">
      <header class="cu-paper-top cu-paper-in">
        <span class="cu-paper-mark"><Scissors :size="12" /> PAPER CUT</span>
        <span class="cu-paper-serial">CUT No. {{ serial }}</span>
      </header>

      <section class="cu-paper-hero">
        <div class="cu-paper-visual cu-paper-in cu-paper-d1">
          <span class="cu-paper-mat cu-paper-mat-a" aria-hidden="true"></span>
          <span class="cu-paper-mat cu-paper-mat-b" aria-hidden="true"></span>

          <figure class="cu-paper-frame">
            <img v-if="images.length" :src="images[0]" :alt="companyName + ' 纸艺作品'" class="cu-paper-photo" />
            <div v-else class="cu-paper-cutscene" aria-hidden="true">
              <span class="cu-paper-sun"></span>
              <span class="cu-paper-arch cu-paper-arch-a"></span>
              <span class="cu-paper-arch cu-paper-arch-b"></span>
              <span class="cu-paper-hill cu-paper-hill-a"></span>
              <span class="cu-paper-hill cu-paper-hill-b"></span>
              <svg class="cu-paper-botanical" viewBox="0 0 180 180" fill="none">
                <path d="M88 176C84 140 82 100 88 54" stroke="currentColor" stroke-width="2.4" stroke-linecap="round" />
                <path d="M88 96C66 90 51 71 49 48c23 2 38 20 39 48Z" fill="currentColor" opacity=".8" />
                <path d="M90 76c18-8 28-25 28-45-19 4-29 22-28 45Z" fill="currentColor" opacity=".58" />
                <path d="M86 132c-16-4-27-17-29-33 18 1 29 14 29 33Z" fill="currentColor" opacity=".68" />
                <path d="M90 120c15-5 24-17 25-31-15 3-24 14-25 31Z" fill="currentColor" opacity=".9" />
              </svg>
              <span class="cu-paper-cutlabel"><Scissors :size="11" /> 手工剪影</span>
            </div>
            <figcaption class="cu-paper-caption">
              {{ images.length ? '作品一瞥 / A GLIMPSE' : '层叠纸境 / LAYERED PAPER' }}
            </figcaption>
          </figure>

          <span class="cu-paper-tape cu-paper-tape-a" aria-hidden="true"></span>
          <span class="cu-paper-tape cu-paper-tape-b" aria-hidden="true"></span>
        </div>

        <div class="cu-paper-heading">
          <p class="cu-paper-eyebrow cu-paper-in cu-paper-d2">PAPER CUT · COLLAGE · BY HAND</p>
          <h1 class="cu-paper-name cu-paper-in cu-paper-d2">{{ companyName }}</h1>
          <div class="cu-paper-rule cu-paper-in cu-paper-d3">
            <span></span>
            <Scissors :size="13" />
            <span></span>
          </div>
          <p class="cu-paper-desc cu-paper-in cu-paper-d3">{{ description }}</p>
        </div>
      </section>

      <section class="cu-paper-services cu-paper-in cu-paper-d3">
        <div class="cu-paper-sec-head">
          <h2>手作服务 <span>WHAT WE CUT</span></h2>
          <span class="cu-paper-count">{{ pad2(services.length) }} 项</span>
        </div>
        <div v-if="services.length" class="cu-paper-scraps">
          <article v-for="(service, i) in services" :key="i" class="cu-paper-scrap" :class="'cu-paper-tone-' + (i % 4)">
            <span class="cu-paper-scrap-num">{{ pad2(i + 1) }}</span>
            <h3>{{ service.name }}</h3>
            <span class="cu-paper-scrap-cut" aria-hidden="true"></span>
          </article>
        </div>
      </section>

      <section class="cu-paper-contact cu-paper-in cu-paper-d4">
        <span class="cu-paper-tape cu-paper-tape-c" aria-hidden="true"></span>
        <div class="cu-paper-contact-inner">
          <p class="cu-paper-contact-label">预约一张纸的相遇 / SAY HELLO</p>
          <p class="cu-paper-contact-note">剪一张纸，留一段空白。</p>
          <a v-if="phone" :href="'tel:' + phone" class="cu-paper-phone">
            <Phone :size="15" />
            <span>{{ phone }}</span>
          </a>
          <a v-if="email" :href="'mailto:' + email" class="cu-paper-mail">{{ email }}</a>
          <span class="cu-paper-stamp" aria-hidden="true">{{ initial }}</span>
        </div>
      </section>

      <footer class="cu-paper-foot cu-paper-in cu-paper-d4">
        <span>{{ companyName }} · 手工纸艺</span>
        <span>CUT No. {{ serial }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Scissors } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const clean = (value) => (typeof value === 'string' ? value.trim() : value != null ? String(value).trim() : '')
const pad2 = (n) => String(n).padStart(2, '0')

const companyName = computed(() => clean(props.payload?.companyName) || '未命名纸艺工作室')
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
.cu-paper-root { position: relative; overflow: hidden; background: radial-gradient(circle at 6% 2%, rgba(190, 105, 69, 0.18), transparent 32%), radial-gradient(circle at 96% 26%, rgba(174, 139, 77, 0.2), transparent 36%), linear-gradient(165deg, #f0dfc5 0%, #e8d2b3 52%, #e2c9a7 100%); color: #3f2c1f; font-family: Georgia, 'Songti SC', 'Noto Serif SC', serif; }
.cu-paper-root::before { content: ''; position: absolute; inset: 0; pointer-events: none; opacity: 0.5; background-image: repeating-linear-gradient(0deg, rgba(92, 63, 42, 0.035) 0 1px, transparent 1px 4px), repeating-linear-gradient(90deg, rgba(255, 255, 255, 0.08) 0 1px, transparent 1px 6px); }
.cu-paper-ambient { position: absolute; inset: 0; pointer-events: none; }
.cu-paper-blob { position: absolute; border-radius: 46% 54% 58% 42%; filter: blur(3px); opacity: 0.6; }
.cu-paper-blob-a { top: -110px; right: -80px; width: 260px; height: 230px; background: rgba(217, 139, 106, 0.22); transform: rotate(18deg); animation: cu-paper-drift 12s ease-in-out infinite alternate; }
.cu-paper-blob-b { bottom: 4%; left: -100px; width: 250px; height: 250px; background: rgba(145, 150, 104, 0.18); animation: cu-paper-drift 15s ease-in-out infinite alternate-reverse; }
.cu-paper-stitch { position: absolute; top: 0; bottom: 0; width: 1px; background: repeating-linear-gradient(180deg, rgba(120, 82, 53, 0.22) 0 5px, transparent 5px 10px); }
.cu-paper-stitch-a { left: 9px; }
.cu-paper-stitch-b { right: 9px; opacity: 0.55; }
.cu-paper-wrap { position: relative; z-index: 2; padding: 20px 21px 30px; }
.cu-paper-top, .cu-paper-foot { display: flex; align-items: center; justify-content: space-between; gap: 12px; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.14em; color: #7a5a43; }
.cu-paper-top { padding-bottom: 10px; border-bottom: 1px solid rgba(96, 64, 42, 0.28); }
.cu-paper-mark { display: inline-flex; align-items: center; gap: 6px; color: #9a4b2e; }
.cu-paper-serial { flex-shrink: 0; color: #7a5a43; }
.cu-paper-hero { margin-top: 22px; }
.cu-paper-visual { position: relative; padding: 5px 2px 0; transform: rotate(-0.7deg); }
.cu-paper-mat { position: absolute; inset: 14px 10px 6px; border-radius: 2px; box-shadow: 0 9px 16px -12px rgba(79, 48, 28, 0.65); }
.cu-paper-mat-a { background: #c7c8a3; transform: rotate(3.6deg) translate(5px, -4px); }
.cu-paper-mat-b { background: #d18a68; transform: rotate(-2.4deg) translate(-5px, 3px); }
.cu-paper-frame { position: relative; z-index: 2; margin: 0; padding: 8px 8px 0; background: #fbf3e3; border: 1px solid rgba(118, 80, 50, 0.3); box-shadow: 0 14px 26px -18px rgba(70, 41, 23, 0.75), inset 0 0 0 1px rgba(255, 255, 255, 0.6); animation: cu-paper-sway 8s ease-in-out infinite; }
.cu-paper-photo, .cu-paper-cutscene { display: block; width: 100%; aspect-ratio: 5 / 4; object-fit: cover; }
.cu-paper-photo { filter: sepia(0.14) saturate(0.92) contrast(1.02); }
.cu-paper-cutscene { position: relative; overflow: hidden; background: linear-gradient(180deg, #f5e5c8, #f0d5b0); }
.cu-paper-sun { position: absolute; top: 14%; right: 18%; width: 74px; height: 74px; border-radius: 50%; background: #e9b26a; box-shadow: 0 0 0 7px rgba(233, 178, 106, 0.16); }
.cu-paper-arch { position: absolute; bottom: 20%; border-radius: 70px 70px 0 0; }
.cu-paper-arch-a { left: 13%; width: 48%; height: 45%; background: #d48663; box-shadow: inset -8px 0 0 rgba(112, 53, 37, 0.12); }
.cu-paper-arch-b { left: 30%; width: 30%; height: 31%; background: #aeb08a; border-radius: 48px 48px 0 0; }
.cu-paper-hill { position: absolute; bottom: -9%; border-radius: 50% 50% 0 0 / 94% 94% 0 0; }
.cu-paper-hill-a { left: -16%; width: 86%; height: 38%; background: #8e956c; }
.cu-paper-hill-b { right: -20%; width: 72%; height: 29%; background: #6e7653; }
.cu-paper-botanical { position: absolute; left: 7%; bottom: 6%; width: 43%; color: #5d6646; filter: drop-shadow(3px 4px 0 rgba(70, 45, 30, 0.14)); }
.cu-paper-cutlabel { position: absolute; right: 9px; top: 9px; display: inline-flex; align-items: center; gap: 4px; padding: 4px 7px; background: rgba(251, 243, 227, 0.9); border: 1px solid rgba(118, 80, 50, 0.2); font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.1em; color: #6e513b; }
.cu-paper-caption { padding: 6px 2px 7px; font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.13em; color: #86654b; }
.cu-paper-tape { position: absolute; z-index: 3; height: 21px; background: rgba(231, 205, 159, 0.68); border-left: 1px dashed rgba(255, 255, 255, 0.4); border-right: 1px dashed rgba(255, 255, 255, 0.4); box-shadow: 0 1px 3px rgba(80, 50, 30, 0.18); }
.cu-paper-tape-a { top: 0; left: 23px; width: 68px; transform: rotate(-5deg); }
.cu-paper-tape-b { top: 8px; right: 15px; width: 54px; transform: rotate(4deg); }
.cu-paper-heading { margin-top: 20px; }
.cu-paper-eyebrow { font-family: 'Courier New', monospace; font-size: 7.5px; letter-spacing: 0.22em; color: #9a4b2e; }
.cu-paper-name { margin-top: 8px; font-size: clamp(35px, 11.5vw, 45px); font-weight: 600; line-height: 1.05; letter-spacing: 0.01em; color: #3b291d; text-shadow: 2px 2px 0 rgba(255, 247, 232, 0.7); overflow-wrap: anywhere; }
.cu-paper-rule { display: flex; align-items: center; gap: 9px; margin-top: 14px; color: #a55435; }
.cu-paper-rule span { flex: 1; height: 1px; background: repeating-linear-gradient(90deg, rgba(154, 75, 46, 0.45) 0 3px, transparent 3px 7px); }
.cu-paper-desc { margin-top: 13px; max-width: 35ch; font-size: 12.5px; line-height: 1.9; color: #6b4e39; }
.cu-paper-services { margin-top: 26px; }
.cu-paper-sec-head { display: flex; align-items: baseline; justify-content: space-between; gap: 10px; padding-bottom: 8px; border-bottom: 1px solid rgba(96, 64, 42, 0.28); }
.cu-paper-sec-head h2 { font-size: 14px; font-weight: 600; letter-spacing: 0.06em; }
.cu-paper-sec-head h2 span { font-family: 'Courier New', monospace; font-size: 7px; font-weight: 400; letter-spacing: 0.12em; color: #9a4b2e; }
.cu-paper-count { font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.1em; color: #8a6a50; }
.cu-paper-scraps { display: grid; grid-template-columns: repeat(2, minmax(0, 1fr)); gap: 11px; margin-top: 14px; }
.cu-paper-scrap { position: relative; min-height: 88px; padding: 13px 12px 12px; border: 1px solid rgba(96, 64, 42, 0.16); clip-path: polygon(0 0, 91% 0, 100% 11%, 100% 100%, 9% 100%, 0 89%); filter: drop-shadow(0 7px 6px rgba(80, 50, 30, 0.13)); transform: rotate(-0.7deg); transition: transform 0.25s ease, filter 0.25s ease; }
.cu-paper-scrap:nth-child(even) { transform: rotate(0.9deg); }
.cu-paper-scrap:last-child:nth-child(odd) { grid-column: span 2; }
.cu-paper-scrap:hover { transform: rotate(0deg) translateY(-2px); filter: drop-shadow(0 11px 9px rgba(80, 50, 30, 0.18)); }
.cu-paper-tone-0 { background: #f8ecd6; }
.cu-paper-tone-1 { background: #efd3bd; }
.cu-paper-tone-2 { background: #dfe0c4; }
.cu-paper-tone-3 { background: #f0ddb2; }
.cu-paper-scrap-num { display: block; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.16em; color: #9a4b2e; }
.cu-paper-scrap h3 { margin-top: 8px; padding-right: 14px; font-size: 15px; font-weight: 600; line-height: 1.45; letter-spacing: 0.02em; color: #44301f; overflow-wrap: anywhere; }
.cu-paper-scrap-cut { position: absolute; right: 10px; bottom: 11px; width: 7px; height: 7px; border-radius: 50%; background: rgba(236, 217, 189, 0.9); box-shadow: inset 0 1px 2px rgba(96, 64, 42, 0.35); }
.cu-paper-empty { margin-top: 13px; padding: 18px 12px; border: 1px dashed rgba(96, 64, 42, 0.36); text-align: center; font-size: 11px; color: #7d6048; }
.cu-paper-contact { position: relative; margin-top: 28px; padding: 20px 17px 17px; background: #faf1df; border: 1px solid rgba(118, 80, 50, 0.28); box-shadow: 0 1px 0 rgba(255, 255, 255, 0.7) inset, 0 15px 24px -19px rgba(70, 41, 23, 0.8); transform: rotate(-0.45deg); }
.cu-paper-tape-c { top: -10px; left: 50%; width: 74px; margin-left: -37px; transform: rotate(-2deg); }
.cu-paper-contact-inner { position: relative; }
.cu-paper-contact-label { font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.17em; color: #9a4b2e; }
.cu-paper-contact-note { margin-top: 7px; font-size: 12px; font-style: italic; color: #7a5a43; }
.cu-paper-phone { display: inline-flex; align-items: center; gap: 7px; max-width: calc(100% - 66px); margin-top: 13px; padding-bottom: 3px; border-bottom: 1px dashed rgba(154, 75, 46, 0.55); font-size: 21px; font-weight: 600; letter-spacing: 0.01em; color: #3b291d; transition: color 0.2s ease; }
.cu-paper-phone span { overflow-wrap: anywhere; }
.cu-paper-phone:hover { color: #9a4b2e; }
.cu-paper-mail { display: inline-block; max-width: calc(100% - 66px); margin-top: 8px; font-family: 'Courier New', monospace; font-size: 9.5px; letter-spacing: 0.02em; color: #6f513b; text-decoration: underline; text-decoration-style: dotted; text-underline-offset: 4px; overflow-wrap: anywhere; transition: color 0.2s ease; }
.cu-paper-mail:hover { color: #9a4b2e; }
.cu-paper-missing { margin-top: 8px; font-family: 'Courier New', monospace; font-size: 9px; letter-spacing: 0.04em; color: #95765a; }
.cu-paper-stamp { position: absolute; right: 0; bottom: -1px; display: grid; place-items: center; width: 52px; height: 52px; border: 1px solid rgba(154, 75, 46, 0.6); border-radius: 50%; color: #9a4b2e; font-size: 23px; transform: rotate(-9deg); }
.cu-paper-stamp::before { content: ''; position: absolute; inset: 4px; border: 1px dashed rgba(154, 75, 46, 0.35); border-radius: 50%; }
.cu-paper-foot { margin-top: 18px; padding-top: 10px; border-top: 1px solid rgba(96, 64, 42, 0.25); font-size: 7px; color: #806149; }
.cu-paper-in { animation: cu-paper-in 0.85s cubic-bezier(0.22, 1, 0.36, 1) both; }
.cu-paper-d1 { animation-delay: 0.08s; }
.cu-paper-d2 { animation-delay: 0.17s; }
.cu-paper-d3 { animation-delay: 0.26s; }
.cu-paper-d4 { animation-delay: 0.35s; }
@keyframes cu-paper-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@keyframes cu-paper-sway {
  0%, 100% { transform: rotate(0deg) translateY(0); }
  50% { transform: rotate(0.35deg) translateY(2px); }
}
@keyframes cu-paper-drift {
  from { transform: translate3d(0, 0, 0) rotate(0deg); }
  to { transform: translate3d(-16px, 18px, 0) rotate(6deg); }
}
@media (prefers-reduced-motion: reduce) {
  .cu-paper-in, .cu-paper-frame, .cu-paper-blob { animation: none; }
  .cu-paper-scrap, .cu-paper-phone, .cu-paper-mail { transition: none; }
}
</style>
