<template>
  <div class="nc-cs-root relative min-h-screen px-4 py-6 text-[#101010]">
    <div class="nc-cs-paper"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-3rem)] w-full max-w-md flex-col">
      <article class="nc-cs-cover">
        <header class="nc-cs-masthead">
          <h2>NIGHT<span>/</span></h2>
          <div class="nc-cs-mast-meta">
            <span>ISSUE 24</span>
            <span>{{ payload?.city || 'SHANGHAI' }}</span>
          </div>
        </header>

        <div class="nc-cs-cover-lines">
          <span>THE RESIDENT ISSUE</span>
          <span>PRINTED AFTER MIDNIGHT</span>
        </div>

        <div class="nc-cs-visual">
          <div class="nc-cs-orb"></div>
          <div class="nc-cs-visual-grid"></div>
          <span class="nc-cs-vertical">AFTER HOURS · VOL. 24</span>
          <div class="nc-cs-hero-copy">
            <p class="nc-cs-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
            <h1 class="nc-cs-name">{{ payload?.name }}</h1>
            <p class="nc-cs-role">{{ payload?.role }}</p>
          </div>
          <div class="nc-cs-stamp">
            <span>N</span>
            <small>24</small>
          </div>
        </div>

        <div class="nc-cs-teaser">
          <p>“{{ payload?.venue }} keeps the door open until the city forgets the time.”</p>
          <span>— COVER STORY, P. 24</span>
        </div>

        <section class="nc-cs-details">
          <div class="nc-cs-detail">
            <span class="nc-cs-key">VENUE</span>
            <span class="nc-cs-val">{{ payload?.venue }} · {{ payload?.city || '—' }}</span>
          </div>
          <div class="nc-cs-detail">
            <span class="nc-cs-key">WECHAT</span>
            <span class="nc-cs-val">{{ payload?.wechat }}</span>
          </div>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-cs-detail nc-cs-detail-link">
            <span class="nc-cs-key">PHONE</span>
            <span class="nc-cs-val">{{ payload.phone }}</span>
            <ArrowUpRight :size="14" />
          </a>
          <div v-else class="nc-cs-detail">
            <span class="nc-cs-key">PHONE</span>
            <span class="nc-cs-val nc-cs-empty"></span>
          </div>
        </section>

        <footer class="nc-cs-foot">
          <div class="nc-cs-barcode" aria-hidden="true"></div>
          <span>NFC {{ nfcId || '0000' }} · ¥0.00</span>
          <span class="nc-cs-foot-mark">NIGHT/24</span>
        </footer>
      </article>

      <p class="nc-cs-outside">COVER STORY · SHOT ON LOCATION · NO RETOUCH</p>
    </div>
  </div>
</template>

<script setup>
import { ArrowUpRight } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-cs-root {
  background: #d8d3c9;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-cs-paper {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(ellipse at 20% 0%, rgba(255, 255, 255, 0.7), transparent 44%),
    repeating-linear-gradient(90deg, rgba(0, 0, 0, 0.02) 0 1px, transparent 1px 3px);
  pointer-events: none;
}
.nc-cs-cover {
  display: flex;
  flex: 1;
  flex-direction: column;
  padding: 18px 18px 14px;
  border: 1px solid rgba(16, 16, 16, 0.3);
  background: #f3f0e8;
  box-shadow: 0 28px 60px -30px rgba(0, 0, 0, 0.55);
  animation: nc-cs-open 0.9s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-cs-masthead {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  gap: 14px;
  padding-bottom: 8px;
  border-bottom: 3px solid #101010;
}
.nc-cs-masthead h2 {
  margin: 0;
  font-size: clamp(54px, 17vw, 74px);
  font-weight: 900;
  line-height: 0.76;
  letter-spacing: -0.075em;
  color: #101010;
}
.nc-cs-masthead h2 span {
  color: #d92b1c;
}
.nc-cs-mast-meta {
  display: flex;
  flex-direction: column;
  gap: 3px;
  padding-bottom: 3px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.14em;
  color: #101010;
  text-align: right;
}
.nc-cs-cover-lines {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 6px 0;
  border-bottom: 1px solid #101010;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: #4b4b4b;
}
.nc-cs-visual {
  position: relative;
  display: flex;
  flex: 1;
  min-height: 340px;
  flex-direction: column;
  justify-content: flex-end;
  overflow: hidden;
  margin-top: 10px;
  padding: 20px;
  background:
    linear-gradient(160deg, rgba(217, 43, 28, 0.94), rgba(115, 12, 52, 0.98) 46%, #0d0b12 100%);
  animation: nc-cs-visual 1.1s 0.1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-cs-orb {
  position: absolute;
  top: -52px;
  right: -56px;
  width: 240px;
  height: 240px;
  border-radius: 50%;
  background: radial-gradient(circle at 34% 30%, #ffd2c4, #e4482f 34%, rgba(120, 10, 50, 0.1) 68%);
  filter: blur(2px);
  opacity: 0.9;
  animation: nc-cs-float 9s ease-in-out infinite;
}
.nc-cs-visual-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(255, 255, 255, 0.08) 1px, transparent 1px),
    linear-gradient(90deg, rgba(255, 255, 255, 0.08) 1px, transparent 1px);
  background-size: 44px 44px;
  mask-image: linear-gradient(180deg, transparent 10%, #000 70%);
  pointer-events: none;
}
.nc-cs-visual::after {
  content: '';
  position: absolute;
  inset: 0;
  background: radial-gradient(rgba(255, 255, 255, 0.14) 0.5px, transparent 0.5px);
  background-size: 4px 4px;
  opacity: 0.4;
  pointer-events: none;
}
.nc-cs-vertical {
  position: absolute;
  top: 18px;
  left: 16px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.3em;
  color: rgba(255, 255, 255, 0.62);
  writing-mode: vertical-rl;
}
.nc-cs-hero-copy {
  position: relative;
  z-index: 2;
  animation: nc-cs-in 0.9s 0.35s both;
}
.nc-cs-alias {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.4em;
  color: #ffd9cc;
  text-transform: uppercase;
}
.nc-cs-name {
  margin-top: 6px;
  font-size: clamp(58px, 18.5vw, 78px);
  font-weight: 900;
  line-height: 0.82;
  letter-spacing: -0.06em;
  text-transform: uppercase;
  color: #fff6ed;
  text-shadow: 0 3px 0 rgba(0, 0, 0, 0.18);
  overflow-wrap: anywhere;
}
.nc-cs-role {
  margin-top: 12px;
  font-size: 9.5px;
  font-weight: 600;
  letter-spacing: 0.26em;
  color: rgba(255, 246, 237, 0.84);
  text-transform: uppercase;
}
.nc-cs-stamp {
  position: absolute;
  top: 18px;
  right: 18px;
  display: flex;
  width: 48px;
  height: 48px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(255, 246, 237, 0.6);
  color: #fff6ed;
  font-family: Georgia, serif;
}
.nc-cs-stamp span {
  font-size: 24px;
  line-height: 1;
}
.nc-cs-stamp small {
  margin-left: 2px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  align-self: flex-end;
  padding-bottom: 10px;
}
.nc-cs-teaser {
  padding: 10px 0 8px;
  border-bottom: 1px solid #101010;
}
.nc-cs-teaser p {
  font-family: Georgia, 'Songti SC', serif;
  font-size: 12px;
  font-style: italic;
  line-height: 1.45;
  color: #222;
}
.nc-cs-teaser span {
  display: block;
  margin-top: 3px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.16em;
  color: #777;
}
.nc-cs-details {
  display: flex;
  flex-direction: column;
  padding-top: 4px;
}
.nc-cs-detail {
  display: flex;
  align-items: baseline;
  gap: 10px;
  padding: 6px 0;
  border-bottom: 1px dotted rgba(16, 16, 16, 0.28);
  text-decoration: none;
  color: inherit;
}
.nc-cs-detail-link {
  margin: 0 -5px;
  padding: 6px 5px;
  transition: background 0.25s ease, padding 0.25s ease;
}
.nc-cs-detail-link:hover {
  background: #f1d9b9;
  padding-left: 11px;
}
.nc-cs-detail-link:active {
  background: #ecd0a6;
}
.nc-cs-key {
  width: 45px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.14em;
  color: #777;
}
.nc-cs-val {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 10.5px;
  font-weight: 600;
  letter-spacing: 0.03em;
  color: #161616;
}
.nc-cs-detail-link svg {
  flex: none;
  transition: transform 0.25s ease;
}
.nc-cs-detail-link:hover svg {
  transform: translate(2px, -2px);
}
.nc-cs-empty {
  color: #aaa;
}
.nc-cs-foot {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 10px;
  padding-top: 8px;
  border-top: 2px solid #101010;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.1em;
  color: #555;
}
.nc-cs-barcode {
  width: 62px;
  height: 17px;
  flex: none;
  background: repeating-linear-gradient(90deg, #101010 0 1px, transparent 1px 3px, #101010 3px 4px, transparent 4px 6px, #101010 6px 9px, transparent 9px 11px);
}
.nc-cs-foot-mark {
  margin-left: auto;
  font-weight: 700;
  letter-spacing: 0.08em;
  color: #d92b1c;
}
.nc-cs-outside {
  padding: 14px 0 4px;
  text-align: center;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(16, 16, 16, 0.5);
  animation: nc-cs-in 1s 0.6s both;
}
@keyframes nc-cs-open {
  from { opacity: 0; transform: translateY(24px) rotate(-0.8deg); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-cs-visual {
  from { opacity: 0; clip-path: inset(0 0 100% 0); }
  to { opacity: 1; clip-path: inset(0 0 0 0); }
}
@keyframes nc-cs-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-cs-float {
  0%, 100% { transform: translate(0, 0); }
  50% { transform: translate(-12px, 12px); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-cs-root *,
  .nc-cs-root *::before,
  .nc-cs-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
