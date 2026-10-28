<template>
  <div class="bc-press-root min-h-screen px-4 py-6">
    <div class="bc-press-wrap mx-auto w-full max-w-md">
      <header class="bc-press-masthead">
        <p class="bc-press-vol">VOL. I — NO. 07</p>
        <h1 class="bc-press-title">The Daily Card</h1>
        <p class="bc-press-date">TUESDAY, OCTOBER 20, 2026 · PRICE: ONE INTRODUCTION</p>
      </header>

      <div class="bc-press-rule bc-press-rule-thick"></div>

      <div class="bc-press-ticker">
        <span>SPECIAL EDITION</span>
        <span>NFC {{ nfcId || '0007' }}</span>
        <span>ALL THE CONNECTIONS FIT TO PRINT</span>
      </div>

      <section class="bc-press-lead">
        <p class="bc-press-kicker">PROFILE · PEOPLE YOU SHOULD KNOW</p>
        <h2 class="bc-press-headline">{{ payload.name || '未命名' }}</h2>
        <p class="bc-press-deck">
          {{ payload.title }}<template v-if="payload.company"> at {{ payload.company }}</template>,
          available for good conversations and better collaborations.
        </p>
        <div class="bc-press-byline">
          <span>BY THE EDITORIAL DESK</span>
          <span>PHOTO ILLUSTRATION / TYPESET</span>
        </div>
      </section>

      <div class="bc-press-rule"></div>

      <section class="bc-press-columns">
        <article class="bc-press-col">
          <p class="bc-press-col-head">THE DETAILS</p>
          <p class="bc-press-body">
            A designer of systems and moments, currently shaping digital products at
            {{ payload.company || 'an independent studio' }}. Reachable by all channels below.
          </p>
          <div class="bc-press-contact">
            <span>TEL</span>
            <a v-if="payload.phone" :href="'tel:' + payload.phone">{{ payload.phone }}</a>
          </div>
          <div class="bc-press-contact">
            <span>MAIL</span>
            <a v-if="payload.email" :href="'mailto:' + payload.email">{{ payload.email }}</a>
          </div>
        </article>

        <article class="bc-press-col bc-press-col-right">
          <div class="bc-press-boxed">
            <span class="bc-press-box-label">CLASSIFIEDS</span>
            <p class="bc-press-box-title">OPEN TO<br />NEW WORK</p>
            <p class="bc-press-box-text">Freelance, full-time, or a coffee. Inquire within.</p>
          </div>
          <div class="bc-press-contact">
            <span>WEB</span>
            <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer">{{ payload.website }}</a>
          </div>
          <div class="bc-press-contact">
            <span>WECHAT</span>
            <span :class="{ 'bc-press-empty': !payload.wechat }">{{ payload.wechat }}</span>
          </div>
        </article>
      </section>

      <div class="bc-press-rule bc-press-rule-thick"></div>

      <div class="bc-press-bottom">
        <div class="bc-press-weather">
          <span class="bc-press-weather-icon">☼</span>
          <div>
            <p>FAIR SKIES</p>
            <span>HIGH 24° / LOW 18°</span>
          </div>
        </div>
        <div class="bc-press-index">
          <span>INDEX</span>
          <p>NAME A1 · WORK A2 · CONTACT B1</p>
        </div>
        <div class="bc-press-edition">
          <span class="bc-press-edition-mark">{{ initial }}</span>
          <p>FINAL<br />EDITION</p>
        </div>
      </div>

      <p class="bc-press-foot">PRINTED ON RECYCLED PAPER · THE DAILY CARD IS A PUBLICATION OF YUNTU PRESS</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.bc-press-root {
  background: #e9e4d6;
  font-family: 'Times New Roman', Georgia, 'Songti SC', 'SimSun', serif;
  color: #1c1917;
}
.bc-press-wrap {
  animation: bc-press-in 0.9s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-press-masthead {
  text-align: center;
}
.bc-press-vol {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.24em;
  color: #57534e;
}
.bc-press-title {
  margin-top: 2px;
  font-family: 'Old English Text MT', 'Times New Roman', Georgia, serif;
  font-size: clamp(34px, 11vw, 46px);
  font-weight: 700;
  line-height: 1;
  letter-spacing: -0.02em;
}
.bc-press-date {
  margin-top: 6px;
  padding-bottom: 8px;
  border-bottom: 1px solid #1c1917;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.14em;
  color: #57534e;
}
.bc-press-rule {
  height: 1px;
  background: #1c1917;
}
.bc-press-rule-thick {
  height: 4px;
  border-top: 1px solid #1c1917;
  border-bottom: 1px solid #1c1917;
  background: transparent;
}
.bc-press-ticker {
  display: flex;
  justify-content: space-between;
  gap: 8px;
  padding: 5px 2px;
  border-bottom: 1px solid #1c1917;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5px;
  letter-spacing: 0.1em;
}
.bc-press-ticker span:first-child {
  font-weight: 700;
  color: #991b1b;
}
.bc-press-ticker span:nth-child(2) {
  color: #57534e;
}
.bc-press-lead {
  padding: 18px 2px 14px;
  text-align: center;
  animation: bc-press-lead 0.9s 0.15s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-press-kicker {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.22em;
  color: #991b1b;
}
.bc-press-headline {
  margin-top: 7px;
  font-family: 'Times New Roman', Georgia, 'Songti SC', serif;
  font-size: clamp(46px, 15vw, 64px);
  font-weight: 700;
  line-height: 0.9;
  letter-spacing: -0.045em;
  overflow-wrap: anywhere;
}
.bc-press-deck {
  max-width: 32em;
  margin: 12px auto 0;
  font-size: 10.5px;
  line-height: 1.5;
  font-style: italic;
  color: #44403c;
}
.bc-press-byline {
  display: flex;
  justify-content: space-between;
  gap: 10px;
  margin-top: 12px;
  padding-top: 6px;
  border-top: 1px solid rgba(28, 25, 23, 0.35);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5px;
  letter-spacing: 0.1em;
  color: #78716c;
}
.bc-press-columns {
  display: grid;
  grid-template-columns: 1.15fr 1fr;
  gap: 0;
  padding: 14px 0;
}
.bc-press-col {
  min-width: 0;
  padding-right: 12px;
}
.bc-press-col-right {
  border-left: 1px solid #a8a29e;
  padding-right: 0;
  padding-left: 12px;
}
.bc-press-col-head {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  font-weight: 700;
  letter-spacing: 0.16em;
  color: #991b1b;
}
.bc-press-body {
  margin-top: 6px;
  font-size: 9.5px;
  line-height: 1.55;
  text-align: justify;
  color: #292524;
}
.bc-press-body::first-letter {
  float: left;
  margin-right: 3px;
  font-size: 27px;
  line-height: 0.8;
  font-weight: 700;
}
.bc-press-contact {
  display: flex;
  align-items: baseline;
  gap: 6px;
  margin-top: 8px;
  padding-top: 6px;
  border-top: 1px dotted rgba(28, 25, 23, 0.35);
}
.bc-press-contact > span:first-child {
  flex: none;
  width: 34px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5px;
  letter-spacing: 0.1em;
  color: #78716c;
}
.bc-press-contact a,
.bc-press-contact span:last-child,
.bc-press-contact em {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 8.5px;
  font-style: normal;
  color: #1c1917;
  text-decoration: none;
  transition: color 0.2s ease;
}
.bc-press-contact a:hover {
  color: #991b1b;
  text-decoration: underline;
}
.bc-press-contact a:active {
  opacity: 0.6;
}
.bc-press-contact em,
.bc-press-empty {
  color: #c7c1b4;
}
.bc-press-boxed {
  border: 2px solid #1c1917;
  padding: 8px;
  background: rgba(255, 255, 255, 0.2);
}
.bc-press-box-label {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5px;
  letter-spacing: 0.14em;
  color: #991b1b;
}
.bc-press-box-title {
  margin-top: 4px;
  font-family: 'Times New Roman', Georgia, serif;
  font-size: 17px;
  font-weight: 700;
  line-height: 0.95;
  letter-spacing: -0.02em;
}
.bc-press-box-text {
  margin-top: 5px;
  font-size: 7.5px;
  line-height: 1.4;
  font-style: italic;
  color: #57534e;
}
.bc-press-bottom {
  display: grid;
  grid-template-columns: 1.1fr 1fr 58px;
  gap: 10px;
  align-items: center;
  padding: 12px 0;
}
.bc-press-weather {
  display: flex;
  align-items: center;
  gap: 7px;
}
.bc-press-weather-icon {
  font-size: 20px;
  color: #b45309;
}
.bc-press-weather p,
.bc-press-index span {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5px;
  font-weight: 700;
  letter-spacing: 0.1em;
}
.bc-press-weather span,
.bc-press-index p {
  display: block;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 4.5px;
  letter-spacing: 0.08em;
  color: #78716c;
}
.bc-press-index {
  border-left: 1px solid #a8a29e;
  padding-left: 10px;
}
.bc-press-index p {
  margin-top: 3px;
  line-height: 1.5;
}
.bc-press-edition {
  border: 1px solid #1c1917;
  padding: 4px;
  text-align: center;
}
.bc-press-edition-mark {
  display: block;
  font-family: 'Times New Roman', Georgia, serif;
  font-size: 42px;
  line-height: 1;
  font-weight: 700;
  transition: color 0.25s ease;
}
.bc-press-edition:hover .bc-press-edition-mark {
  color: #991b1b;
}
.bc-press-edition p {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 4.5px;
  line-height: 1.4;
  letter-spacing: 0.08em;
}
.bc-press-foot {
  padding-top: 8px;
  border-top: 1px solid #1c1917;
  text-align: center;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 4.5px;
  letter-spacing: 0.1em;
  color: #78716c;
}
@keyframes bc-press-in {
  from { opacity: 0; transform: translateY(18px) rotate(-0.4deg); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-press-lead {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-press-root *,
  .bc-press-root *::before,
  .bc-press-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
