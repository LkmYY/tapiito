<template>
  <div class="ac-villain-root relative min-h-screen overflow-hidden px-4 py-8 text-[#f4e6e6]">
    <div class="ac-villain-bleed"></div>
    <div class="ac-villain-thorns"></div>
    <div class="ac-villain-ash"></div>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <article class="ac-villain-poster">
        <span class="ac-villain-arch"></span>
        <header class="ac-villain-head">
          <span>VILLAIN DOSSIER · XIII</span>
          <span class="ac-villain-threat">THREAT SSS</span>
        </header>

        <div class="ac-villain-crest">
          <span class="ac-villain-crest-ring"></span>
          <Skull :size="34" />
          <Crown :size="16" class="ac-villain-crown" />
        </div>

        <p class="ac-villain-kicker">THE FALLEN ONE</p>
        <h1 class="ac-villain-name">{{ p.name || '未命名角色' }}</h1>
        <p class="ac-villain-title">「{{ p.title }}」</p>

        <div class="ac-villain-dossier">
          <div>
            <span>ORIGIN / 所属</span>
            <b>{{ p.series || '未知作品' }}</b>
          </div>
          <div>
            <span>CURSE / 灾厄属性</span>
            <b class="ac-villain-curse">{{ p.element }}</b>
          </div>
        </div>

        <blockquote class="ac-villain-quote">
          <span class="ac-villain-glyph">“</span>
          <p>{{ p.quote || '……别回头。' }}</p>
          <cite>—— 记录于第 13 夜</cite>
        </blockquote>

        <div class="ac-villain-rule"><span></span><i>✝</i><span></span></div>

        <footer class="ac-villain-foot">
          <div>
            <small>FOLLOW THE SIGNAL</small>
            <b><AtSign :size="12" />{{ p.social }}</b>
          </div>
          <span class="ac-villain-seal">封<br />印</span>
        </footer>

        <p class="ac-villain-warning">IF YOU ARE READING THIS, RUN.</p>
      </article>

      <div class="ac-villain-ticker">
        <span>REPENT · 不可直视 · REPENT · 不可直视 ·</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { AtSign, Crown, Skull } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-villain-root {
  background: #060406;
  font-family: 'Cinzel', 'Songti SC', 'Noto Serif SC', serif;
}
.ac-villain-bleed {
  position: absolute;
  top: -16%;
  left: 50%;
  width: 520px;
  height: 520px;
  transform: translateX(-50%);
  background: radial-gradient(circle, rgba(159, 18, 57, 0.42), transparent 62%);
  filter: blur(30px);
  animation: ac-villain-breathe 7s ease-in-out infinite;
}
.ac-villain-thorns {
  position: absolute;
  inset: 0;
  background:
    repeating-conic-gradient(from 200deg at 0% 0%, rgba(159, 18, 57, 0.14) 0deg 0.45deg, transparent 0.45deg 6deg),
    repeating-conic-gradient(from 20deg at 100% 100%, rgba(159, 18, 57, 0.12) 0deg 0.45deg, transparent 0.45deg 7deg);
  mask-image: radial-gradient(ellipse at 50% 40%, transparent 24%, #000 90%);
}
.ac-villain-ash {
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(248, 113, 113, 0.4) 0.6px, transparent 0.7px);
  background-size: 37px 41px;
  opacity: 0.25;
  animation: ac-villain-ash 12s linear infinite;
}
.ac-villain-poster {
  position: relative;
  border: 1px solid rgba(159, 18, 57, 0.55);
  background:
    linear-gradient(180deg, rgba(27, 7, 11, 0.94), rgba(8, 4, 6, 0.99) 55%, rgba(20, 4, 8, 0.98));
  padding: 22px 22px 18px;
  box-shadow:
    0 0 0 1px #000,
    0 0 50px rgba(159, 18, 57, 0.16),
    inset 0 0 70px rgba(88, 4, 20, 0.42);
  animation: ac-villain-enter 1s cubic-bezier(0.16, 1, 0.3, 1) both;
  transition: transform 0.5s ease, box-shadow 0.5s ease;
}
.ac-villain-poster:hover {
  transform: translateY(-4px);
  box-shadow:
    0 0 0 1px #000,
    0 0 80px rgba(190, 18, 60, 0.3),
    inset 0 0 80px rgba(120, 8, 30, 0.5);
}
.ac-villain-poster::before {
  content: '';
  position: absolute;
  inset: 7px;
  border: 1px solid rgba(159, 18, 57, 0.24);
  pointer-events: none;
}
.ac-villain-arch {
  position: absolute;
  top: -1px;
  left: 50%;
  width: 200px;
  height: 200px;
  transform: translateX(-50%);
  border: 1px solid rgba(190, 18, 60, 0.16);
  border-bottom: none;
  border-radius: 100px 100px 0 0;
  pointer-events: none;
}
.ac-villain-head {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-family: Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: rgba(248, 113, 113, 0.42);
}
.ac-villain-threat {
  border: 1px solid rgba(248, 113, 113, 0.45);
  padding: 3px 6px;
  color: #f87171;
  background: rgba(127, 29, 29, 0.25);
}
.ac-villain-crest {
  position: relative;
  display: flex;
  width: 92px;
  height: 104px;
  margin: 30px auto 18px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(190, 18, 60, 0.55);
  border-radius: 46px 46px 8px 8px;
  background: radial-gradient(circle at 50% 40%, rgba(127, 29, 29, 0.65), rgba(8, 4, 6, 0.9) 72%);
  color: #fca5a5;
  box-shadow: 0 0 30px rgba(159, 18, 57, 0.3), inset 0 0 22px rgba(159, 18, 57, 0.35);
}
.ac-villain-crest::before,
.ac-villain-crest::after {
  content: '';
  position: absolute;
  bottom: -13px;
  width: 1px;
  height: 13px;
  background: rgba(190, 18, 60, 0.45);
}
.ac-villain-crest::before { left: 16px; transform: rotate(24deg); }
.ac-villain-crest::after { right: 16px; transform: rotate(-24deg); }
.ac-villain-crest-ring {
  position: absolute;
  inset: 9px;
  border: 1px dashed rgba(190, 18, 60, 0.35);
  border-radius: 40px 40px 5px 5px;
  animation: ac-villain-pulse 4s ease-in-out infinite;
}
.ac-villain-crown {
  position: absolute;
  top: -9px;
  color: #fbbf24;
  filter: drop-shadow(0 0 8px rgba(251, 191, 36, 0.5));
}
.ac-villain-kicker {
  text-align: center;
  font-family: Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.44em;
  color: rgba(248, 113, 113, 0.55);
}
.ac-villain-name {
  margin-top: 9px;
  text-align: center;
  font-size: 47px;
  font-weight: 700;
  line-height: 1.02;
  letter-spacing: 0.08em;
  color: #fef2f2;
  text-shadow: 0 0 30px rgba(190, 18, 60, 0.55), 0 0 2px #fff;
  overflow-wrap: anywhere;
}
.ac-villain-title {
  margin-top: 9px;
  text-align: center;
  font-size: 13px;
  font-style: italic;
  letter-spacing: 0.18em;
  color: #f87171;
}
.ac-villain-dossier {
  display: grid;
  gap: 1px;
  margin-top: 26px;
  background: rgba(190, 18, 60, 0.22);
  border: 1px solid rgba(190, 18, 60, 0.22);
}
.ac-villain-dossier > div {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  gap: 12px;
  background: #0b0508;
  padding: 10px 12px;
}
.ac-villain-dossier span {
  flex: none;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.16em;
  color: rgba(248, 113, 113, 0.45);
}
.ac-villain-dossier b {
  min-width: 0;
  text-align: right;
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 0.06em;
  color: rgba(254, 242, 242, 0.88);
  overflow-wrap: anywhere;
}
.ac-villain-curse { color: #fca5a5 !important; }
.ac-villain-quote {
  position: relative;
  margin-top: 24px;
  border-left: 2px solid #be123c;
  padding: 8px 12px 8px 26px;
}
.ac-villain-glyph {
  position: absolute;
  top: -2px;
  left: 5px;
  font-family: Georgia, serif;
  font-size: 34px;
  line-height: 1;
  color: rgba(190, 18, 60, 0.55);
}
.ac-villain-quote p {
  font-size: 15.5px;
  font-style: italic;
  line-height: 1.85;
  color: rgba(254, 226, 226, 0.9);
  overflow-wrap: anywhere;
}
.ac-villain-quote cite {
  display: block;
  margin-top: 7px;
  font-family: Consolas, monospace;
  font-size: 7.5px;
  font-style: normal;
  letter-spacing: 0.14em;
  color: rgba(248, 113, 113, 0.4);
}
.ac-villain-rule {
  display: flex;
  align-items: center;
  gap: 10px;
  margin: 22px 0 16px;
}
.ac-villain-rule span { flex: 1; height: 1px; background: linear-gradient(90deg, transparent, rgba(190, 18, 60, 0.4)); }
.ac-villain-rule span:last-child { background: linear-gradient(90deg, rgba(190, 18, 60, 0.4), transparent); }
.ac-villain-rule i { font-style: normal; font-size: 8px; color: rgba(248, 113, 113, 0.5); }
.ac-villain-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.ac-villain-foot small {
  display: block;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.24em;
  color: rgba(248, 113, 113, 0.36);
}
.ac-villain-foot b {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  margin-top: 3px;
  font-family: Consolas, monospace;
  font-size: 11.5px;
  font-weight: 400;
  color: #fca5a5;
  overflow-wrap: anywhere;
}
.ac-villain-seal {
  display: flex;
  width: 42px;
  height: 42px;
  align-items: center;
  justify-content: center;
  border: 2px solid #be123c;
  border-radius: 4px;
  color: #be123c;
  font-size: 8px;
  font-weight: 700;
  line-height: 1.2;
  letter-spacing: 0.1em;
  text-align: center;
  transform: rotate(-8deg);
  opacity: 0.8;
}
.ac-villain-warning {
  margin-top: 18px;
  border-top: 1px solid rgba(190, 18, 60, 0.2);
  padding-top: 9px;
  text-align: center;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.34em;
  color: rgba(248, 113, 113, 0.28);
}
.ac-villain-ticker {
  margin-top: 20px;
  overflow: hidden;
  border-top: 1px solid rgba(190, 18, 60, 0.18);
  border-bottom: 1px solid rgba(190, 18, 60, 0.18);
  padding: 6px 0;
  font-family: Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.28em;
  color: rgba(248, 113, 113, 0.32);
  white-space: nowrap;
}
.ac-villain-ticker span { display: inline-block; animation: ac-villain-marquee 18s linear infinite; }
@keyframes ac-villain-enter {
  from { opacity: 0; transform: translateY(30px) scale(0.97); filter: blur(6px); }
  to { opacity: 1; transform: none; filter: blur(0); }
}
@keyframes ac-villain-breathe { 0%, 100% { opacity: 0.7; } 50% { opacity: 1; } }
@keyframes ac-villain-pulse { 0%, 100% { opacity: 0.35; } 50% { opacity: 1; } }
@keyframes ac-villain-ash { to { background-position: 37px 41px; } }
@keyframes ac-villain-marquee { from { transform: translateX(0); } to { transform: translateX(-50%); } }
@media (prefers-reduced-motion: reduce) {
  .ac-villain-root *,
  .ac-villain-root *::before,
  .ac-villain-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
