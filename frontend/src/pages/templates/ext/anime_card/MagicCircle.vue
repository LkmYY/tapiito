<template>
  <div class="ac-magic-root relative min-h-screen overflow-hidden px-5 py-8 text-[#f7ecd2]">
    <div class="ac-magic-nebula"></div>
    <div class="ac-magic-specks"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="ac-magic-top">
        <span>ARCANA / {{ roman }}</span>
        <span>NFC · {{ nfcId || '0006' }}</span>
      </header>

      <main class="ac-magic-altar">
        <div class="ac-magic-halo">
          <svg class="ac-magic-ring" viewBox="0 0 320 320" aria-hidden="true">
            <g fill="none" stroke="currentColor" stroke-width="0.8">
              <circle cx="160" cy="160" r="150" stroke-dasharray="2 7" />
              <circle cx="160" cy="160" r="134" opacity="0.8" />
              <circle cx="160" cy="160" r="118" stroke-dasharray="12 5 2 5" />
              <path d="M160 34 L269 222 L51 222 Z" opacity="0.65" />
              <path d="M160 286 L51 98 L269 98 Z" opacity="0.65" />
              <circle cx="160" cy="160" r="72" />
              <path d="M160 88 L222 196 L98 196 Z" opacity="0.9" />
            </g>
            <g fill="#f7d488">
              <circle cx="160" cy="10" r="2.2" />
              <circle cx="310" cy="160" r="1.8" />
              <circle cx="160" cy="310" r="1.5" />
              <circle cx="10" cy="160" r="2" />
              <circle cx="160" cy="42" r="1.6" />
              <circle cx="243" cy="198" r="1.4" />
              <circle cx="77" cy="198" r="1.4" />
            </g>
          </svg>
          <div class="ac-magic-orb">
            <span class="ac-magic-initial">{{ initial }}</span>
            <span class="ac-magic-glint"></span>
          </div>
          <span class="ac-magic-rune ac-magic-rune-a">✦</span>
          <span class="ac-magic-rune ac-magic-rune-b">☾</span>
          <span class="ac-magic-rune ac-magic-rune-c">✧</span>
          <span class="ac-magic-rune ac-magic-rune-d">✶</span>
        </div>

        <div class="ac-magic-titleblock">
          <p class="ac-magic-eyebrow">{{ p.series || '未知作品' }}</p>
          <h1 class="ac-magic-name">{{ p.name || '未命名角色' }}</h1>
          <div class="ac-magic-rule"><span></span><i>✦</i><span></span></div>
          <p class="ac-magic-title">{{ p.title }}</p>
          <span class="ac-magic-element">{{ p.element }}</span>
        </div>

        <blockquote class="ac-magic-quote">
          <span class="ac-magic-qmark">“</span>
          <p>{{ p.quote || '……星光尚未留下回音。' }}</p>
          <span class="ac-magic-qmark ac-magic-qmark-end">”</span>
        </blockquote>

        <footer class="ac-magic-foot">
          <span class="ac-magic-sigil">✧</span>
          <div>
            <small>CONTRACTED TO</small>
            <b>@{{ p.social }}</b>
          </div>
          <span class="ac-magic-sigil">✧</span>
        </footer>
      </main>

      <p class="ac-magic-caption">THE STARS REMEMBER YOUR NAME</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
const initial = computed(() => (p.value.name || '？').toString().trim().slice(0, 1) || '？')
const roman = computed(() => ['I', 'II', 'III', 'IV', 'V', 'VI', 'VII', 'VIII', 'IX', 'X'][(Number(props.nfcId) || 6) % 10] || 'VI')
</script>

<style scoped>
.ac-magic-root {
  background: radial-gradient(circle at 50% 30%, #241348 0%, #0d0820 48%, #05030c 100%);
  font-family: 'Cormorant Garamond', 'Songti SC', 'Noto Serif SC', serif;
}
.ac-magic-nebula {
  position: absolute;
  inset: -20%;
  background:
    radial-gradient(circle at 20% 25%, rgba(124, 58, 237, 0.3), transparent 45%),
    radial-gradient(circle at 78% 65%, rgba(190, 24, 93, 0.18), transparent 42%),
    radial-gradient(circle at 46% 90%, rgba(8, 145, 178, 0.18), transparent 46%);
  filter: blur(30px);
  animation: ac-magic-drift 18s ease-in-out infinite;
}
.ac-magic-specks {
  position: absolute;
  inset: 0;
  background-image:
    radial-gradient(#fde68a 0.7px, transparent 0.8px),
    radial-gradient(rgba(255, 255, 255, 0.7) 0.6px, transparent 0.7px);
  background-size: 73px 73px, 113px 113px;
  background-position: 0 0, 41px 27px;
  opacity: 0.45;
  animation: ac-magic-specks 24s linear infinite;
}
.ac-magic-top {
  display: flex;
  justify-content: space-between;
  font-family: Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.26em;
  color: rgba(247, 212, 136, 0.5);
  animation: ac-magic-fade 0.8s both;
}
.ac-magic-altar {
  position: relative;
  display: flex;
  flex: 1;
  flex-direction: column;
  align-items: center;
  padding-top: 18px;
  animation: ac-magic-enter 1.1s cubic-bezier(0.2, 0.9, 0.2, 1) both;
  transition: transform 0.6s ease;
}
.ac-magic-altar:hover { transform: translateY(-3px); }
.ac-magic-halo {
  position: relative;
  width: 296px;
  height: 296px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.ac-magic-halo::before {
  content: '';
  position: absolute;
  inset: 20%;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(251, 191, 36, 0.22), rgba(124, 58, 237, 0.1) 52%, transparent 70%);
  filter: blur(18px);
}
.ac-magic-ring {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  color: rgba(247, 212, 136, 0.6);
  filter: drop-shadow(0 0 10px rgba(251, 191, 36, 0.24));
  animation: ac-magic-spin 42s linear infinite;
}
.ac-magic-orb {
  position: relative;
  display: flex;
  width: 116px;
  height: 116px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(247, 212, 136, 0.75);
  border-radius: 50%;
  background:
    radial-gradient(circle at 38% 30%, rgba(253, 230, 138, 0.5), transparent 42%),
    radial-gradient(circle at 50% 55%, #3b1d63, #120a26 72%);
  box-shadow:
    0 0 0 6px rgba(247, 212, 136, 0.06),
    0 0 44px rgba(251, 191, 36, 0.28),
    inset 0 0 30px rgba(251, 191, 36, 0.16);
}
.ac-magic-orb::before {
  content: '';
  position: absolute;
  inset: -13px;
  border: 1px dashed rgba(247, 212, 136, 0.36);
  border-radius: 50%;
  animation: ac-magic-spin 20s linear infinite reverse;
}
.ac-magic-initial {
  font-size: 48px;
  font-weight: 600;
  color: #fff7e2;
  text-shadow: 0 0 24px rgba(251, 191, 36, 0.8);
}
.ac-magic-glint {
  position: absolute;
  top: 20%;
  left: 22%;
  width: 15px;
  height: 15px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.8);
  filter: blur(5px);
  animation: ac-magic-glint 4s ease-in-out infinite;
}
.ac-magic-rune {
  position: absolute;
  color: #f7d488;
  font-size: 12px;
  text-shadow: 0 0 10px rgba(251, 191, 36, 0.8);
  animation: ac-magic-float 5s ease-in-out infinite;
}
.ac-magic-rune-a { top: 32px; left: 52px; }
.ac-magic-rune-b { top: 74px; right: 30px; font-size: 15px; animation-delay: 1s; }
.ac-magic-rune-c { bottom: 42px; right: 74px; animation-delay: 2s; }
.ac-magic-rune-d { bottom: 84px; left: 24px; font-size: 10px; animation-delay: 3s; }
.ac-magic-titleblock { text-align: center; margin-top: 2px; }
.ac-magic-eyebrow {
  font-family: Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.4em;
  color: rgba(247, 212, 136, 0.6);
}
.ac-magic-name {
  margin-top: 10px;
  font-size: 44px;
  font-weight: 600;
  line-height: 1.05;
  letter-spacing: 0.12em;
  background: linear-gradient(180deg, #fffaf0 10%, #f7d488 55%, #a16207 100%);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  filter: drop-shadow(0 0 20px rgba(251, 191, 36, 0.25));
  overflow-wrap: anywhere;
}
.ac-magic-rule {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 9px;
  margin: 12px auto 9px;
  width: 72%;
}
.ac-magic-rule span { height: 1px; flex: 1; background: linear-gradient(90deg, transparent, rgba(247, 212, 136, 0.65)); }
.ac-magic-rule span:last-child { background: linear-gradient(90deg, rgba(247, 212, 136, 0.65), transparent); }
.ac-magic-rule i { font-style: normal; color: #f7d488; font-size: 8px; }
.ac-magic-title { font-size: 13px; letter-spacing: 0.34em; color: #e9d5ff; }
.ac-magic-element {
  display: inline-block;
  margin-top: 13px;
  border: 1px solid rgba(247, 212, 136, 0.38);
  border-radius: 999px;
  padding: 4px 14px;
  font-size: 10px;
  letter-spacing: 0.2em;
  color: #f7d488;
  background: rgba(251, 191, 36, 0.06);
}
.ac-magic-quote {
  position: relative;
  width: 100%;
  margin-top: 24px;
  padding: 18px 28px;
  text-align: center;
}
.ac-magic-quote::before,
.ac-magic-quote::after {
  content: '';
  position: absolute;
  left: 50%;
  width: 44%;
  height: 1px;
  transform: translateX(-50%);
  background: linear-gradient(90deg, transparent, rgba(247, 212, 136, 0.4), transparent);
}
.ac-magic-quote::before { top: 0; }
.ac-magic-quote::after { bottom: 0; }
.ac-magic-quote p {
  font-size: 16px;
  font-style: italic;
  line-height: 1.85;
  color: rgba(255, 247, 226, 0.88);
  overflow-wrap: anywhere;
}
.ac-magic-qmark {
  position: absolute;
  top: 6px;
  left: 12px;
  font-size: 38px;
  line-height: 1;
  color: rgba(247, 212, 136, 0.35);
}
.ac-magic-qmark-end { top: auto; left: auto; right: 12px; bottom: 0; }
.ac-magic-foot {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 16px;
  margin-top: auto;
  padding-top: 26px;
  text-align: center;
}
.ac-magic-foot small {
  display: block;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.3em;
  color: rgba(247, 212, 136, 0.4);
}
.ac-magic-foot b {
  display: block;
  margin-top: 3px;
  font-family: Consolas, monospace;
  font-size: 12px;
  font-weight: 400;
  letter-spacing: 0.08em;
  color: #f7d488;
  overflow-wrap: anywhere;
}
.ac-magic-sigil { color: rgba(247, 212, 136, 0.45); font-size: 11px; }
.ac-magic-caption {
  padding: 24px 0 8px;
  text-align: center;
  font-family: Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.3em;
  color: rgba(247, 212, 136, 0.24);
  animation: ac-magic-fade 1.3s 0.6s both;
}
@keyframes ac-magic-enter {
  from { opacity: 0; transform: scale(0.92) translateY(24px); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-magic-fade { from { opacity: 0; } to { opacity: 1; } }
@keyframes ac-magic-spin { to { transform: rotate(360deg); } }
@keyframes ac-magic-drift {
  0%, 100% { transform: translate3d(0, 0, 0) scale(1); }
  50% { transform: translate3d(16px, -12px, 0) scale(1.06); }
}
@keyframes ac-magic-specks { to { background-position: 73px 146px, -72px 140px; } }
@keyframes ac-magic-glint { 0%, 100% { opacity: 0.3; transform: scale(0.7); } 50% { opacity: 1; transform: scale(1.2); } }
@keyframes ac-magic-float { 0%, 100% { transform: translateY(0); opacity: 0.55; } 50% { transform: translateY(-7px); opacity: 1; } }
@media (prefers-reduced-motion: reduce) {
  .ac-magic-root *,
  .ac-magic-root *::before,
  .ac-magic-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
