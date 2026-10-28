<template>
  <div class="bc-term-root min-h-screen px-4 py-7">
    <div class="bc-term-crt"></div>
    <div class="bc-term-wrap mx-auto flex min-h-[calc(100vh-3.5rem)] w-full max-w-md flex-col">
      <div class="bc-term-window">
        <div class="bc-term-bar">
          <span class="bc-term-dots"><i></i><i></i><i></i></span>
          <span>guest@yuntu: ~/identity</span>
          <span class="bc-term-bar-r">bash — 80×24</span>
        </div>

        <div class="bc-term-body">
          <p class="bc-term-line"><span class="bc-term-prompt">$</span> whoami</p>
          <h1 class="bc-term-name">{{ payload.name || '未命名' }}<span class="bc-term-caret"></span></h1>

          <p class="bc-term-line bc-term-gap"><span class="bc-term-prompt">$</span> cat profile.txt</p>
          <div class="bc-term-profile">
            <p><span>role</span>{{ payload.title }}</p>
            <p><span>org</span>{{ payload.company }}</p>
            <p><span>nfc</span>{{ nfcId || '—' }}</p>
          </div>

          <p class="bc-term-line bc-term-gap"><span class="bc-term-prompt">$</span> contact --list</p>
          <div class="bc-term-contacts">
            <a v-if="payload.phone" :href="'tel:' + payload.phone" class="bc-term-row">
              <span class="bc-term-key">[tel]</span>
              <span class="bc-term-val">{{ payload.phone }}</span>
              <span class="bc-term-ok">OK</span>
            </a>
            <div v-else class="bc-term-row bc-term-off">
              <span class="bc-term-key">[tel]</span>
              <span class="bc-term-val"></span>
              <span class="bc-term-ok">--</span>
            </div>
            <a v-if="payload.email" :href="'mailto:' + payload.email" class="bc-term-row">
              <span class="bc-term-key">[mail]</span>
              <span class="bc-term-val">{{ payload.email }}</span>
              <span class="bc-term-ok">OK</span>
            </a>
            <div v-else class="bc-term-row bc-term-off">
              <span class="bc-term-key">[mail]</span>
              <span class="bc-term-val"></span>
              <span class="bc-term-ok">--</span>
            </div>
            <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer" class="bc-term-row">
              <span class="bc-term-key">[www]</span>
              <span class="bc-term-val">{{ payload.website }}</span>
              <span class="bc-term-ok">OK</span>
            </a>
            <div v-else class="bc-term-row bc-term-off">
              <span class="bc-term-key">[www]</span>
              <span class="bc-term-val"></span>
              <span class="bc-term-ok">--</span>
            </div>
            <div class="bc-term-row bc-term-off">
              <span class="bc-term-key">[wechat]</span>
              <span class="bc-term-val">{{ payload.wechat }}</span>
              <span class="bc-term-ok">WX</span>
            </div>
          </div>

          <p class="bc-term-line bc-term-gap"><span class="bc-term-prompt">$</span> echo "let's connect"</p>
          <p class="bc-term-out">let's connect<span class="bc-term-caret"></span></p>
        </div>

        <div class="bc-term-status">
          <span>-- INSERT --</span>
          <span>UTF-8</span>
          <span>Ln 42, Col 7</span>
        </div>
      </div>

      <p class="bc-term-foot">SESSION SECURE · PGP NOT INCLUDED · {{ payload.company }}</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.bc-term-root {
  position: relative;
  overflow: hidden;
  background: #020604;
  font-family: 'SFMono-Regular', Consolas, 'Courier New', monospace;
  color: #86efac;
}
.bc-term-root::before {
  content: '';
  position: absolute;
  inset: 0;
  background-image: none;
  background: radial-gradient(ellipse at 50% 42%, rgba(22, 101, 52, 0.28), transparent 70%);
}
.bc-term-crt {
  position: absolute;
  inset: 0;
  background: repeating-linear-gradient(180deg, rgba(134, 239, 172, 0.035) 0 1px, transparent 1px 3px);
  pointer-events: none;
  z-index: 3;
}
.bc-term-wrap {
  position: relative;
  z-index: 2;
}
.bc-term-window {
  border: 1px solid #166534;
  border-radius: 5px;
  background: rgba(1, 12, 5, 0.92);
  box-shadow: 0 0 42px rgba(22, 163, 74, 0.16), inset 0 0 60px rgba(22, 101, 52, 0.12);
  overflow: hidden;
  animation: bc-term-on 0.55s steps(5, end) both;
}
.bc-term-bar {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 7px 10px;
  border-bottom: 1px solid #14532d;
  background: #052e16;
  font-size: 7px;
  letter-spacing: 0.04em;
  color: #4ade80;
}
.bc-term-dots {
  display: flex;
  gap: 4px;
}
.bc-term-dots i {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: #166534;
}
.bc-term-dots i:first-child { background: #ef4444; }
.bc-term-dots i:nth-child(2) { background: #eab308; }
.bc-term-dots i:nth-child(3) { background: #22c55e; }
.bc-term-bar-r {
  margin-left: auto;
  color: rgba(74, 222, 128, 0.45);
}
.bc-term-body {
  padding: 20px 16px 16px;
}
.bc-term-line {
  margin: 0;
  font-size: 9px;
  line-height: 1.7;
  letter-spacing: 0.02em;
  color: rgba(134, 239, 172, 0.55);
  animation: bc-term-type 0.3s both;
}
.bc-term-gap {
  margin-top: 20px;
}
.bc-term-prompt {
  margin-right: 6px;
  color: #22c55e;
  font-weight: 700;
}
.bc-term-name {
  margin: 6px 0 0;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: clamp(28px, 9vw, 38px);
  font-weight: 700;
  line-height: 1.15;
  letter-spacing: -0.02em;
  color: #dcfce7;
  text-shadow: 0 0 16px rgba(74, 222, 128, 0.35);
  overflow-wrap: anywhere;
  animation: bc-term-glow 0.8s 0.15s both;
}
.bc-term-caret {
  display: inline-block;
  width: 8px;
  height: 1em;
  margin-left: 4px;
  vertical-align: -0.13em;
  background: #4ade80;
  animation: bc-term-blink 1.1s steps(1, end) infinite;
}
.bc-term-profile {
  margin-top: 6px;
  padding-left: 14px;
  border-left: 1px solid rgba(34, 197, 94, 0.25);
}
.bc-term-profile p {
  display: flex;
  gap: 12px;
  font-size: 9px;
  line-height: 1.9;
  color: #bbf7d0;
}
.bc-term-profile span {
  width: 42px;
  flex: none;
  color: rgba(134, 239, 172, 0.4);
}
.bc-term-contacts {
  margin-top: 4px;
  border-top: 1px dashed rgba(34, 197, 94, 0.25);
}
.bc-term-row {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 2px;
  border-bottom: 1px dashed rgba(34, 197, 94, 0.18);
  text-decoration: none;
  color: inherit;
  transition: background 0.15s ease, padding-left 0.15s ease;
  animation: bc-term-row-in 0.4s both;
}
a.bc-term-row:hover {
  background: rgba(34, 197, 94, 0.12);
  padding-left: 7px;
}
a.bc-term-row:hover .bc-term-val {
  color: #ffffff;
}
a.bc-term-row:active {
  background: rgba(34, 197, 94, 0.2);
}
.bc-term-key {
  flex: none;
  font-size: 7.5px;
  color: #22c55e;
}
.bc-term-val {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 9px;
  color: #bbf7d0;
  transition: color 0.15s ease;
}
.bc-term-ok {
  flex: none;
  font-size: 6.5px;
  letter-spacing: 0.1em;
  color: rgba(74, 222, 128, 0.45);
}
.bc-term-off {
  opacity: 0.42;
}
.bc-term-out {
  font-size: 10px;
  color: #eafff1;
}
.bc-term-status {
  display: flex;
  gap: 16px;
  padding: 6px 10px;
  border-top: 1px solid #14532d;
  background: #052e16;
  font-size: 6.5px;
  letter-spacing: 0.06em;
  color: #86efac;
}
.bc-term-status span:first-child {
  color: #fde047;
}
.bc-term-status span:last-child {
  margin-left: auto;
  color: rgba(134, 239, 172, 0.5);
}
.bc-term-foot {
  margin-top: auto;
  padding-top: 16px;
  text-align: center;
  font-size: 6px;
  letter-spacing: 0.12em;
  color: rgba(74, 222, 128, 0.3);
}
@keyframes bc-term-on {
  from { opacity: 0; transform: scaleY(0.92); filter: brightness(2); }
  to { opacity: 1; transform: none; filter: none; }
}
@keyframes bc-term-type {
  from { opacity: 0; transform: translateX(-6px); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-term-glow {
  from { opacity: 0; text-shadow: 0 0 40px rgba(74, 222, 128, 0.9); }
  to { opacity: 1; }
}
@keyframes bc-term-row-in {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes bc-term-blink {
  0%, 49% { opacity: 1; }
  50%, 100% { opacity: 0; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-term-root *,
  .bc-term-root *::before,
  .bc-term-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
