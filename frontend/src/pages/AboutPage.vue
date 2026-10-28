<template>
  <div class="about-page">
    <div class="about-grain" aria-hidden="true"></div>

    <header class="about-header">
      <div class="about-shell about-nav">
        <router-link to="/" class="about-brand">
          <span class="about-brand-mark">
            <img src="/tapiito-icon-white.png" alt="Tapiito" class="h-5 w-5" draggable="false" />
          </span>
          <span>Tapiito</span>
        </router-link>

        <nav class="about-nav-links" aria-label="主导航">
          <router-link to="/showroom" class="about-nav-link about-nav-hide-mobile">展厅</router-link>
          <router-link to="/about" class="about-nav-link about-nav-link-active about-nav-hide-mobile">关于我们</router-link>
          <router-link to="/feedback" class="about-nav-link">我有意见</router-link>
          <router-link to="/consult" class="about-nav-cta">
            预约咨询
            <ArrowUpRight class="h-3.5 w-3.5" />
          </router-link>
        </nav>
      </div>
    </header>

    <main>
      <!-- ============ Hero ============ -->
      <section class="about-hero">
        <div class="about-hero-bg" aria-hidden="true">
          <div class="about-glow about-glow-orange"></div>
          <div class="about-glow about-glow-sun"></div>
          <div class="about-glow about-glow-mint"></div>
          <div class="about-hero-grid"></div>
        </div>

        <div class="about-shell about-hero-inner">
          <div class="about-hero-copy">
            <p class="about-kicker reveal" style="--d: 0s">ABOUT TAPIITO <span>·</span> 关于我们</p>

            <h1 class="about-hero-title reveal" style="--d: 0.08s">
              让每一次碰触，<br />
              都有<span class="about-serif">回应</span>。
            </h1>

            <p class="about-lead reveal" style="--d: 0.16s">
              Tapiito 把内容写进一张薄薄的 NFC 标签。名片、菜单、门店、作品、宠物铭牌——
              手机轻轻一碰，页面立刻打开。没有 App，没有扫码，也没有多余的等待。
            </p>

            <div class="about-hero-actions reveal" style="--d: 0.24s">
              <router-link to="/showroom" class="about-btn about-btn-dark">
                去展厅看看
                <ArrowUpRight class="h-4 w-4" />
              </router-link>
              <a href="#how" class="about-btn about-btn-ghost">
                看看怎么用
                <ArrowDown class="h-4 w-4" />
              </a>
            </div>

            <dl class="about-stats reveal" style="--d: 0.32s">
              <div>
                <dt>10</dt>
                <dd>种标签类型</dd>
              </div>
              <div>
                <dt>300</dt>
                <dd>种视觉风格</dd>
              </div>
              <div>
                <dt>0</dt>
                <dd>个需要下载的 App</dd>
              </div>
            </dl>
          </div>

          <div class="about-hero-visual reveal" style="--d: 0.2s" aria-hidden="true">
            <div class="about-orbit about-orbit-1"></div>
            <div class="about-orbit about-orbit-2"></div>
            <div class="about-orbit about-orbit-3"></div>

            <div class="about-specimen-wrap">
              <div class="about-specimen-back"></div>

              <div class="about-specimen">
                <div class="about-specimen-top">
                  <span class="about-specimen-brand">TAPIITO</span>
                  <span class="about-specimen-nfc">
                    <img src="/tapiito-icon-white.png" alt="Tapiito" class="h-[15px] w-[15px]" draggable="false" />
                    NFC
                  </span>
                </div>

                <p class="about-specimen-eyebrow">SMART IDENTITY TAG</p>
                <p class="about-specimen-title">碰一下，<br />它就活了。</p>

                <div class="about-specimen-bottom">
                  <span class="about-specimen-chip"><i></i></span>
                  <span class="about-specimen-id">NO. 9999</span>
                </div>
              </div>

              <span class="about-float about-float-a">NO APP</span>
              <span class="about-float about-float-b">TAP · OPEN</span>
              <span class="about-float about-float-c">ALWAYS ON</span>
            </div>

            <p class="about-visual-note">一张标签 / 一个入口 / 一直有效</p>
          </div>
        </div>
      </section>

      <!-- ============ Ticker ============ -->
      <div class="about-ticker" aria-hidden="true">
        <div class="about-ticker-track">
          <span v-for="(item, index) in tickerItems" :key="index" class="about-ticker-item">
            {{ item }}
            <i></i>
          </span>
        </div>
      </div>

      <!-- ============ 01 Platform ============ -->
      <section id="platform" class="about-section">
        <div class="about-shell about-platform-grid">
          <div class="about-section-label">
            <span>01</span>
            <p>PLATFORM</p>
          </div>

          <div>
            <h2 class="about-h2">
              我们做的，<br />
              <span class="about-serif">只是把距离缩短到一次碰触。</span>
            </h2>

            <div class="about-platform-body">
              <p>
                Tapiito 把一张普通的 NFC 标签，变成你的数字身份、门店内容，或一段关系的入口。
                它不需要充电，不需要打开应用，也不要求对方懂任何技术。
              </p>
              <p>
                你只需要决定「想让人看到什么」，剩下的交给标签。内容可以随时更新，链接保持不变；
                已经发出去的标签，不必因为改了一行字而作废。
              </p>
            </div>

            <div class="about-principles">
              <article v-for="item in principles" :key="item.num" class="about-principle">
                <span>{{ item.num }}</span>
                <h3>{{ item.title }}</h3>
                <p>{{ item.desc }}</p>
              </article>
            </div>
          </div>
        </div>
      </section>

      <!-- ============ 02 Features ============ -->
      <section id="features" class="about-features">
        <div class="about-shell">
          <div class="about-section-head">
            <div class="about-section-label">
              <span>02</span>
              <p>WHAT IT DOES</p>
            </div>
            <div>
              <h2 class="about-h2">Tapiito能做什么</h2>
              <p class="about-section-lead">
                从展示、内容，一张硬件标签背后是一套完整的自述。
              </p>
            </div>
          </div>

          <div class="about-feature-grid">
            <article
              v-for="(feature, index) in features"
              :key="feature.title"
              class="about-feature"
              :class="[feature.span, feature.tone ? `about-feature-${feature.tone}` : '']"
            >
              <div class="about-feature-top">
                <span class="about-feature-index">{{ String(index + 1).padStart(2, '0') }}</span>
                <component :is="feature.icon" class="about-feature-icon h-[18px] w-[18px]" />
              </div>

              <h3>{{ feature.title }}</h3>
              <p>{{ feature.desc }}</p>
              <span class="about-feature-tag">{{ feature.tag }}</span>
            </article>
          </div>

          <div class="about-empower">
            <div class="about-empower-intro">
              <p class="about-empower-kicker">EMPOWER · 赋能</p>
              <h3 class="about-empower-title">
                为每一种连接，<br />
                <span class="about-serif">赋能。</span>
              </h3>
              <p class="about-empower-lead">
                同一张标签，在不同的人手里，解决不同的问题。Tapiito 不只展示信息，也把一次碰面，变成一段可以继续的关系。
              </p>
            </div>

            <ul class="about-empower-list">
              <li
                v-for="(item, index) in empowerments"
                :key="item.title"
                class="about-empower-item"
              >
                <div class="about-empower-meta">
                  <span class="about-empower-num">{{ String(index + 1).padStart(2, '0') }}</span>
                  <span class="about-empower-icon">
                    <component :is="item.icon" class="h-4 w-4" :stroke-width="2.2" />
                  </span>
                </div>

                <div class="about-empower-copy">
                  <h4>{{ item.title }}</h4>
                  <p>{{ item.desc }}</p>
                </div>

                <span class="about-empower-tag">{{ item.tag }}</span>
              </li>
            </ul>
          </div>

          <p class="about-quote">
            <span class="about-serif">不是把纸换成屏幕，</span>
            而是把「找到你」这件事，变得没有摩擦力。
          </p>
        </div>
      </section>

      <!-- ============ 03 How ============ -->
      <section id="how" class="about-how">
        <div class="about-how-glow" aria-hidden="true"></div>

        <div class="about-shell about-how-grid">
          <div class="about-how-intro">
            <div class="about-section-label about-label-light">
              <span>03</span>
              <p>HOW IT WORKS</p>
            </div>

            <h2 class="about-h2 about-h2-light">
              从选择到碰触，<br />
              <span class="about-serif">只要四步。</span>
            </h2>

            <p class="about-how-lead">
              不需要学习一套新工具。选好、填好、收到，然后碰一下。
            </p>

            <div class="about-how-note">
              <Check class="h-4 w-4" :stroke-width="2.6" />
              购买即完成首次写入，含 1 次免费内容调整
            </div>

            <router-link to="/showroom" class="about-btn about-btn-light">
              开始挑选
              <ArrowUpRight class="h-4 w-4" />
            </router-link>
          </div>

          <ol class="about-steps">
            <li v-for="(step, index) in steps" :key="step.title" class="about-step">
              <span class="about-step-num">{{ String(index + 1).padStart(2, '0') }}</span>
              <div class="about-step-body">
                <div class="about-step-top">
                  <h3>{{ step.title }}</h3>
                  <span>{{ step.meta }}</span>
                </div>
                <p>{{ step.desc }}</p>
              </div>
              <MoveRight class="about-step-arrow h-5 w-5" />
            </li>
          </ol>
        </div>
      </section>

      <!-- ============ Closing ============ -->
      <section class="about-closing">
        <div class="about-shell">
          <div class="about-closing-card">
            <div class="about-closing-glow" aria-hidden="true"></div>

            <div class="about-closing-copy">
              <p class="about-closing-kicker">READY WHEN YOU ARE</p>
              <h2>
                把想说的话，<br />
                <span class="about-serif">放进一张标签里。</span>
              </h2>
            </div>

            <div class="about-closing-actions">
              <router-link to="/showroom" class="about-btn about-btn-dark">
                逛展厅
                <ArrowUpRight class="h-4 w-4" />
              </router-link>
              <router-link to="/consult" class="about-btn about-btn-paper">
                预约咨询
                <ArrowUpRight class="h-4 w-4" />
              </router-link>
            </div>
          </div>
        </div>
      </section>
    </main>

    <footer class="about-footer">
      <div class="about-shell about-footer-inner">
        <div class="about-footer-brand">
          <span class="about-brand-mark about-brand-mark-sm">
            <img src="/tapiito-icon-white.png" alt="Tapiito" class="h-[15px] w-[15px]" draggable="false" />
          </span>
          <span>Tapiito</span>
          <span class="about-footer-tagline">Tap it alive.</span>
        </div>

        <nav class="about-footer-links" aria-label="页脚导航">
          <router-link to="/showroom">展厅</router-link>
          <router-link to="/about">关于我们</router-link>
          <router-link to="/feedback">我有意见</router-link>
          <router-link to="/privacy">隐私政策</router-link>
          <router-link to="/disclaimer">免责声明</router-link>
          <router-link to="/consult">预约咨询</router-link>
        </nav>

        <p class="about-footer-copy">© {{ new Date().getFullYear() }} Tapiito</p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { onBeforeUnmount, onMounted } from 'vue'
import {
  Nfc,
  ArrowUpRight,
  ArrowDown,
  Layers,
  RefreshCw,
  ScanLine,
  BarChart2,
  LayoutDashboard,
  Compass,
  Check,
  MoveRight,
  Building2,
  Store,
  Megaphone,
  HeartHandshake,
  ShieldCheck
} from 'lucide-vue-next'

const tickerBase = [
  '碰触即达',
  '无需 App',
  '链接永远不变',
  '10 种类型',
  '300 种风格',
  '内容随时更新'
]
const tickerItems = [...tickerBase, ...tickerBase, ...tickerBase]

const principles = [
  {
    num: '01',
    title: '轻',
    desc: '一张卡片的大小，装得下一整页内容。'
  },
  {
    num: '02',
    title: '快',
    desc: '手机靠近，页面打开，路径短到没有步骤。'
  },
  {
    num: '03',
    title: '久',
    desc: '链接不变，内容常新，标签可以一直用下去。'
  }
]

const features = [
  {
    icon: ScanLine,
    title: '碰触即达',
    desc: 'NFC 靠近即打开，不下载 App，不扫码。把「找到你」的路径，缩短到一次碰触。',
    tag: 'NATIVE NFC',
    span: 'sm:col-span-1 lg:col-span-2'
  },
  {
    icon: Layers,
    title: '300 种风格',
    desc: '10 种类型，每种 30 种独立设计。名片、菜单、官网、宠物铭牌——总有一种像你。',
    tag: '10 × 30',
    span: 'sm:col-span-1 lg:col-span-4',
    tone: 'dark'
  },
  {
    icon: RefreshCw,
    title: '内容联系随时定制更新',
    desc: '改文案、调整价格，链接保持不变。已经发出去的标签，继续有效。',
    tag: 'LINK STAYS',
    span: 'sm:col-span-1 lg:col-span-4',
    tone: 'sun'
  },
  {
    icon: Compass,
    title: '10 类真实场景',
    desc: '商务名片、儿童信息、紧急联系、企业官网、产品展示、菜谱、夜店社交、二次元闪卡、宠物铭牌、自定义企业。',
    tag: 'USE CASES',
    span: 'sm:col-span-1 lg:col-span-2'
  }
]

const empowerments = [
  {
    icon: Building2,
    title: '为企业',
    desc: '把品牌、产品与联系方式收进一张标签。一次碰触完成介绍，也让后续联系更自然。',
    tag: 'BRAND · GROWTH'
  },
  {
    icon: Store,
    title: '为街边小店',
    desc: '菜单、营业时间、优惠活动随时改。门口一张标签，就能把路过的人留下来。',
    tag: 'LOCAL · RETAIN'
  },
  {
    icon: Megaphone,
    title: '为地推留客',
    desc: '把一次派样、一次路演，变成长期连接。碰一下进入活动页，后续内容还能继续更新。',
    tag: 'FIELD · CONVERT'
  },
  {
    icon: HeartHandshake,
    title: '为交友',
    desc: '不用现场找二维码。碰一下打开你的个人主页，兴趣、作品和联系方式都在里面。',
    tag: 'SOCIAL · CONNECT'
  },
  {
    icon: ShieldCheck,
    title: '为老人小孩防走失',
    desc: '把紧急联系人和必要的健康信息放在随手可碰的位置。需要时，路人碰一下就能联系家人。',
    tag: 'SAFETY · CARE'
  }
]

const steps = [
  {
    title: '逛展厅，选一个像你的',
    meta: 'CHOOSE',
    desc: '10 种类型、300 种风格，先看真实渲染效果，再决定。'
  },
  {
    title: '填内容，留下收件信息',
    meta: 'FILL',
    desc: '选好类型和风格，把要展示的内容填进去；确认后下单。'
  },
  {
    title: '我们制作并写入',
    meta: 'MAKE',
    desc: '购买即完成首次写入，标签寄出前已经准备好；含 1 次免费内容调整。'
  },
  {
    title: '手机碰一碰',
    meta: 'TAP',
    desc: '无需 App，靠近即可打开。之后想改内容，链接不变，随时更新。'
  }
]

onMounted(() => {
  document.title = '关于我们 · Tapiito'
})

onBeforeUnmount(() => {
  document.title = 'Tapiito'
})
</script>

<style scoped>
.about-page {
  --ink: #14181d;
  --paper: #fbfaf7;
  --paper-deep: #f3f0e9;
  --orange: #ff6b35;
  --sun: #ffd166;
  --mint: #7bdff2;
  position: relative;
  min-height: 100vh;
  overflow-x: hidden;
  background: var(--paper);
  color: var(--ink);
  font-family: 'Geist', 'PingFang SC', 'Microsoft YaHei', 'Noto Sans SC', system-ui, sans-serif;
}

.about-grain {
  position: fixed;
  inset: 0;
  z-index: 80;
  pointer-events: none;
  opacity: 0.035;
  mix-blend-mode: multiply;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='140' height='140'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='.85' numOctaves='3' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)'/%3E%3C/svg%3E");
}

.about-shell {
  width: 100%;
  max-width: 1400px;
  margin: 0 auto;
  padding-left: 1rem;
  padding-right: 1rem;
}

@media (min-width: 640px) {
  .about-shell {
    padding-left: 1.5rem;
    padding-right: 1.5rem;
  }
}

/* ---------- Nav ---------- */
.about-header {
  position: sticky;
  top: 0;
  z-index: 40;
  border-bottom: 1px solid rgba(20, 24, 29, 0.07);
  background: rgba(251, 250, 247, 0.88);
  backdrop-filter: blur(14px);
  -webkit-backdrop-filter: blur(14px);
}

.about-nav {
  display: flex;
  height: 60px;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
}

.about-brand {
  display: inline-flex;
  min-width: 0;
  align-items: center;
  gap: 0.6rem;
  font-size: 18px;
  font-weight: 650;
  letter-spacing: -0.025em;
  color: var(--ink);
}

.about-brand-mark {
  display: grid;
  width: 30px;
  height: 30px;
  flex: none;
  place-items: center;
  border-radius: 9px;
  background: var(--ink);
  color: var(--orange);
  transition: transform 0.3s cubic-bezier(0.22, 1, 0.36, 1);
}

.about-brand:hover .about-brand-mark {
  transform: rotate(-6deg) scale(1.04);
}

.about-nav-links {
  display: flex;
  align-items: center;
  gap: 0.35rem;
}

.about-nav-link {
  display: inline-flex;
  align-items: center;
  height: 34px;
  padding: 0 0.8rem;
  border-radius: 999px;
  font-size: 13px;
  font-weight: 500;
  color: rgba(20, 24, 29, 0.55);
  transition: color 0.2s ease, background-color 0.2s ease;
}

.about-nav-link:hover {
  background: rgba(20, 24, 29, 0.045);
  color: var(--ink);
}

.about-nav-link-active {
  background: var(--ink);
  color: #fff;
}

.about-nav-link-active:hover {
  background: var(--ink);
  color: #fff;
}

.about-nav-cta {
  display: inline-flex;
  align-items: center;
  gap: 0.3rem;
  height: 34px;
  padding: 0 0.85rem;
  border-radius: 999px;
  background: var(--orange);
  color: #fff;
  font-size: 12.5px;
  font-weight: 650;
  box-shadow: 0 8px 18px -12px rgba(255, 107, 53, 0.9);
  transition: transform 0.22s cubic-bezier(0.22, 1, 0.36, 1), background-color 0.22s ease;
}

.about-nav-cta:hover {
  background: #f25a22;
  transform: translateY(-1px);
}

@media (max-width: 639px) {
  .about-nav-hide-mobile {
    display: none;
  }
}

/* ---------- Hero ---------- */
.about-hero {
  position: relative;
  overflow: hidden;
  border-bottom: 1px solid rgba(20, 24, 29, 0.07);
}

.about-hero-bg {
  position: absolute;
  inset: 0;
  pointer-events: none;
}

.about-glow {
  position: absolute;
  border-radius: 999px;
  filter: blur(100px);
}

.about-glow-orange {
  top: -180px;
  left: -120px;
  width: 460px;
  height: 460px;
  background: rgba(255, 107, 53, 0.17);
}

.about-glow-sun {
  top: -130px;
  right: -100px;
  width: 420px;
  height: 420px;
  background: rgba(255, 209, 102, 0.28);
}

.about-glow-mint {
  bottom: -280px;
  left: 33%;
  width: 460px;
  height: 460px;
  background: rgba(123, 223, 242, 0.17);
}

.about-hero-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(to right, rgba(20, 24, 29, 0.045) 1px, transparent 1px),
    linear-gradient(to bottom, rgba(20, 24, 29, 0.045) 1px, transparent 1px);
  background-size: 64px 64px;
  mask-image: radial-gradient(circle at 32% 24%, black, transparent 72%);
  -webkit-mask-image: radial-gradient(circle at 32% 24%, black, transparent 72%);
}

.about-hero-inner {
  position: relative;
  z-index: 1;
  display: grid;
  gap: 4rem;
  padding-top: 4.5rem;
  padding-bottom: 4.5rem;
}

@media (min-width: 900px) {
  .about-hero-inner {
    grid-template-columns: minmax(0, 1.1fr) minmax(320px, 0.9fr);
    align-items: center;
    padding-top: 6rem;
    padding-bottom: 6rem;
  }
}

.about-kicker {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 10.5px;
  font-weight: 600;
  letter-spacing: 0.2em;
  color: var(--orange);
}

.about-kicker span {
  color: rgba(20, 24, 29, 0.25);
}

.about-hero-title {
  margin: 1.35rem 0 0;
  font-size: clamp(2.8rem, 7vw, 5.4rem);
  font-weight: 650;
  line-height: 1.02;
  letter-spacing: -0.05em;
  text-wrap: balance;
}

.about-serif {
  font-family: 'Playfair Display', 'Songti SC', 'STSong', 'SimSun', serif;
  font-style: italic;
  font-weight: 500;
  letter-spacing: -0.02em;
}

.about-hero-title .about-serif {
  color: var(--orange);
}

.about-lead {
  max-width: 35rem;
  margin: 1.5rem 0 0;
  font-size: 15.5px;
  line-height: 1.85;
  color: rgba(20, 24, 29, 0.58);
}

@media (min-width: 768px) {
  .about-lead {
    font-size: 16.5px;
  }
}

.about-hero-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  margin-top: 2rem;
}

.about-btn {
  display: inline-flex;
  align-items: center;
  gap: 0.45rem;
  height: 44px;
  padding: 0 1.15rem;
  border-radius: 11px;
  font-size: 13.5px;
  font-weight: 650;
  line-height: 1;
  transition:
    transform 0.22s cubic-bezier(0.22, 1, 0.36, 1),
    background-color 0.22s ease,
    border-color 0.22s ease,
    box-shadow 0.22s ease;
}

.about-btn:hover {
  transform: translateY(-2px);
}

.about-btn-dark {
  background: var(--ink);
  color: #fff;
}

.about-btn-dark:hover {
  background: var(--orange);
  box-shadow: 0 14px 26px -14px rgba(255, 107, 53, 0.75);
}

.about-btn-ghost {
  border: 1px solid rgba(20, 24, 29, 0.14);
  background: rgba(255, 255, 255, 0.58);
  color: rgba(20, 24, 29, 0.7);
}

.about-btn-ghost:hover {
  border-color: rgba(20, 24, 29, 0.32);
  color: var(--ink);
}

.about-btn-light {
  background: var(--paper);
  color: var(--ink);
}

.about-btn-light:hover {
  background: var(--sun);
}

.about-btn-paper {
  border: 1px solid rgba(20, 24, 29, 0.25);
  background: rgba(251, 250, 247, 0.35);
  color: var(--ink);
}

.about-btn-paper:hover {
  background: rgba(251, 250, 247, 0.7);
}

.about-stats {
  display: flex;
  flex-wrap: wrap;
  gap: 1.5rem 2.75rem;
  margin: 2.5rem 0 0;
  padding-top: 1.4rem;
  border-top: 1px solid rgba(20, 24, 29, 0.11);
}

.about-stats dt {
  font-size: 26px;
  font-weight: 650;
  line-height: 1;
  letter-spacing: -0.04em;
  font-variant-numeric: tabular-nums;
}

.about-stats dd {
  margin: 0.4rem 0 0;
  font-size: 11px;
  letter-spacing: 0.04em;
  color: rgba(20, 24, 29, 0.42);
}

/* ---------- Hero specimen ---------- */
.about-hero-visual {
  position: relative;
  display: none;
  min-height: 450px;
  align-items: center;
  justify-content: center;
}

@media (min-width: 900px) {
  .about-hero-visual {
    display: flex;
  }
}

.about-orbit {
  position: absolute;
  border: 1px solid rgba(20, 24, 29, 0.1);
  border-radius: 999px;
}

.about-orbit::before {
  content: '';
  position: absolute;
  top: -3px;
  left: 50%;
  width: 6px;
  height: 6px;
  border-radius: 999px;
  background: var(--orange);
  box-shadow: 0 0 0 4px rgba(255, 107, 53, 0.12);
}

.about-orbit-1 {
  width: 280px;
  height: 280px;
  animation: about-orbit-spin 16s linear infinite;
}

.about-orbit-2 {
  width: 360px;
  height: 360px;
  border-color: rgba(20, 24, 29, 0.06);
  animation: about-orbit-spin 24s linear infinite reverse;
}

.about-orbit-2::before {
  background: var(--sun);
  box-shadow: 0 0 0 4px rgba(255, 209, 102, 0.18);
}

.about-orbit-3 {
  width: 440px;
  height: 440px;
  border-style: dashed;
  border-color: rgba(20, 24, 29, 0.055);
}

.about-orbit-3::before {
  display: none;
}

@keyframes about-orbit-spin {
  to {
    transform: rotate(360deg);
  }
}

.about-specimen-wrap {
  position: relative;
  z-index: 2;
  width: 320px;
  height: 400px;
}

.about-specimen-back {
  position: absolute;
  inset: 38px 6px 38px 12px;
  border-radius: 22px;
  background: var(--ink);
  transform: rotate(5deg);
  opacity: 0.92;
}

.about-specimen {
  position: absolute;
  inset: 38px 14px 38px 4px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 1.4rem;
  border: 1px solid rgba(20, 24, 29, 0.1);
  border-radius: 22px;
  background:
    radial-gradient(circle at 88% 0%, rgba(255, 209, 102, 0.35), transparent 38%),
    linear-gradient(145deg, #ffffff 0%, #f5f1e9 100%);
  box-shadow:
    0 2px 4px rgba(20, 24, 29, 0.06),
    0 34px 60px -30px rgba(20, 24, 29, 0.5);
  transform: rotate(-4deg);
  transition: transform 0.5s cubic-bezier(0.22, 1, 0.36, 1);
}

.about-hero-visual:hover .about-specimen {
  transform: rotate(-1.5deg) translateY(-4px);
}

.about-specimen-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.about-specimen-brand {
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.22em;
}

.about-specimen-nfc {
  display: inline-flex;
  align-items: center;
  gap: 0.25rem;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 9px;
  letter-spacing: 0.1em;
  color: rgba(20, 24, 29, 0.45);
}

.about-specimen-eyebrow {
  margin: 2.8rem 0 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 9px;
  font-weight: 600;
  letter-spacing: 0.18em;
  color: var(--orange);
}

.about-specimen-title {
  margin: 0.8rem 0 0;
  font-size: 38px;
  font-weight: 650;
  line-height: 1.07;
  letter-spacing: -0.055em;
}

.about-specimen-bottom {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  padding-top: 1rem;
  border-top: 1px solid rgba(20, 24, 29, 0.1);
}

.about-specimen-chip {
  display: grid;
  width: 34px;
  height: 24px;
  place-items: center;
  border-radius: 6px;
  background: linear-gradient(135deg, var(--orange), var(--sun));
  box-shadow: inset 0 0 0 1px rgba(20, 24, 29, 0.12);
}

.about-specimen-chip i {
  width: 13px;
  height: 9px;
  border: 1px solid rgba(20, 24, 29, 0.28);
  border-radius: 2px;
}

.about-specimen-id {
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 9px;
  letter-spacing: 0.16em;
  color: rgba(20, 24, 29, 0.38);
}

.about-float {
  position: absolute;
  z-index: 3;
  display: inline-flex;
  align-items: center;
  height: 26px;
  padding: 0 0.65rem;
  border: 1px solid rgba(20, 24, 29, 0.09);
  border-radius: 999px;
  background: rgba(251, 250, 247, 0.9);
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 8.5px;
  font-weight: 600;
  letter-spacing: 0.12em;
  color: rgba(20, 24, 29, 0.52);
  box-shadow: 0 10px 22px -16px rgba(20, 24, 29, 0.5);
  backdrop-filter: blur(8px);
  animation: about-float 7s ease-in-out infinite;
}

.about-float-a {
  top: 4px;
  left: -6px;
}

.about-float-b {
  top: 178px;
  right: -26px;
  animation-delay: 1.2s;
}

.about-float-c {
  bottom: 6px;
  left: -6px;
  animation-delay: 2.1s;
}

@keyframes about-float {
  0%,
  100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-7px);
  }
}

.about-visual-note {
  position: absolute;
  bottom: 0;
  left: 50%;
  margin: 0;
  transform: translateX(-50%);
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 9px;
  letter-spacing: 0.16em;
  color: rgba(20, 24, 29, 0.34);
  white-space: nowrap;
}

/* ---------- Ticker ---------- */
.about-ticker {
  overflow: hidden;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  background: var(--ink);
  padding: 0.8rem 0;
}

.about-ticker-track {
  display: flex;
  width: max-content;
  animation: about-ticker 42s linear infinite;
}

.about-ticker-item {
  display: inline-flex;
  align-items: center;
  gap: 1rem;
  padding-right: 1rem;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 10px;
  letter-spacing: 0.18em;
  color: rgba(255, 255, 255, 0.48);
  white-space: nowrap;
}

.about-ticker-item i {
  width: 3px;
  height: 3px;
  border-radius: 999px;
  background: var(--orange);
}

@keyframes about-ticker {
  to {
    transform: translateX(-33.333%);
  }
}

/* ---------- Sections ---------- */
.about-section {
  padding-top: 5.5rem;
  padding-bottom: 5.5rem;
}

@media (min-width: 768px) {
  .about-section {
    padding-top: 7rem;
    padding-bottom: 7rem;
  }
}

.about-platform-grid {
  display: grid;
  gap: 2rem;
}

@media (min-width: 900px) {
  .about-platform-grid {
    grid-template-columns: 180px minmax(0, 1fr);
    gap: 4rem;
  }
}

.about-section-label {
  display: flex;
  align-items: baseline;
  gap: 0.65rem;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
}

.about-section-label span {
  font-size: 11px;
  color: var(--orange);
}

.about-section-label p {
  margin: 0;
  font-size: 9.5px;
  font-weight: 600;
  letter-spacing: 0.2em;
  color: rgba(20, 24, 29, 0.35);
}

.about-h2 {
  margin: 0;
  font-size: clamp(2rem, 4.6vw, 3.6rem);
  font-weight: 650;
  line-height: 1.1;
  letter-spacing: -0.045em;
  text-wrap: balance;
}

.about-h2 .about-serif {
  color: var(--orange);
}

.about-platform-body {
  display: grid;
  gap: 1.2rem;
  margin-top: 2rem;
  max-width: 62rem;
}

@media (min-width: 768px) {
  .about-platform-body {
    grid-template-columns: 1fr 1fr;
    gap: 2.5rem;
  }
}

.about-platform-body p {
  margin: 0;
  font-size: 14px;
  line-height: 1.9;
  color: rgba(20, 24, 29, 0.55);
}

.about-principles {
  display: grid;
  gap: 1.8rem;
  margin-top: 3.2rem;
  padding-top: 2rem;
  border-top: 1px solid rgba(20, 24, 29, 0.1);
}

@media (min-width: 768px) {
  .about-principles {
    grid-template-columns: repeat(3, minmax(0, 1fr));
    gap: 2rem;
  }
}

.about-principle > span {
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 9px;
  letter-spacing: 0.14em;
  color: rgba(20, 24, 29, 0.28);
}

.about-principle h3 {
  margin: 0.7rem 0 0;
  font-family: 'Playfair Display', 'Songti SC', 'STSong', 'SimSun', serif;
  font-size: 32px;
  font-style: italic;
  font-weight: 500;
  line-height: 1;
  color: var(--orange);
}

.about-principle p {
  margin: 0.8rem 0 0;
  max-width: 18rem;
  font-size: 13px;
  line-height: 1.75;
  color: rgba(20, 24, 29, 0.5);
}

/* ---------- Features ---------- */
.about-features {
  border-top: 1px solid rgba(20, 24, 29, 0.07);
  border-bottom: 1px solid rgba(20, 24, 29, 0.07);
  background: var(--paper-deep);
  padding-top: 5.5rem;
  padding-bottom: 5.5rem;
}

@media (min-width: 768px) {
  .about-features {
    padding-top: 7rem;
    padding-bottom: 7rem;
  }
}

.about-section-head {
  display: grid;
  gap: 1.4rem;
  margin-bottom: 2.6rem;
}

@media (min-width: 900px) {
  .about-section-head {
    grid-template-columns: 180px minmax(0, 1fr);
    gap: 4rem;
  }
}

.about-section-lead {
  max-width: 34rem;
  margin: 1rem 0 0;
  font-size: 14px;
  line-height: 1.8;
  color: rgba(20, 24, 29, 0.52);
}

.about-feature-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 0.85rem;
}

@media (min-width: 640px) {
  .about-feature-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (min-width: 1024px) {
  .about-feature-grid {
    grid-template-columns: repeat(6, minmax(0, 1fr));
  }
}

.about-feature {
  position: relative;
  display: flex;
  min-height: 235px;
  flex-direction: column;
  padding: 1.4rem;
  border: 1px solid rgba(20, 24, 29, 0.08);
  border-radius: 18px;
  background: #fff;
  transition:
    transform 0.28s cubic-bezier(0.22, 1, 0.36, 1),
    box-shadow 0.28s ease,
    border-color 0.28s ease;
}

.about-feature:hover {
  transform: translateY(-4px);
  border-color: rgba(20, 24, 29, 0.18);
  box-shadow: 0 22px 40px -26px rgba(20, 24, 29, 0.38);
}

.about-feature-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.about-feature-index {
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 9px;
  letter-spacing: 0.14em;
  color: rgba(20, 24, 29, 0.3);
}

.about-feature-icon {
  color: var(--orange);
}

.about-feature h3 {
  margin: 1.8rem 0 0;
  font-size: 18px;
  font-weight: 650;
  letter-spacing: -0.025em;
}

.about-feature p {
  margin: 0.7rem 0 1.4rem;
  font-size: 12.5px;
  line-height: 1.75;
  color: rgba(20, 24, 29, 0.52);
}

.about-feature-tag {
  margin-top: auto;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 8.5px;
  font-weight: 600;
  letter-spacing: 0.15em;
  color: rgba(20, 24, 29, 0.3);
}

.about-feature-dark {
  border-color: var(--ink);
  background:
    radial-gradient(circle at 100% 0%, rgba(255, 107, 53, 0.24), transparent 44%),
    var(--ink);
  color: #fff;
}

.about-feature-dark .about-feature-index,
.about-feature-dark .about-feature-tag {
  color: rgba(255, 255, 255, 0.38);
}

.about-feature-dark .about-feature-icon {
  color: var(--sun);
}

.about-feature-dark p {
  color: rgba(255, 255, 255, 0.58);
}

.about-feature-sun {
  border-color: #f2c64f;
  background:
    radial-gradient(circle at 96% 4%, rgba(255, 255, 255, 0.45), transparent 34%),
    var(--sun);
}

.about-feature-sun .about-feature-icon {
  color: var(--ink);
}

.about-feature-sun .about-feature-index,
.about-feature-sun .about-feature-tag {
  color: rgba(20, 24, 29, 0.4);
}

.about-feature-sun p {
  color: rgba(20, 24, 29, 0.62);
}

/* ---------- Empowerment ---------- */
.about-empower {
  display: grid;
  gap: 1.8rem;
  margin-top: 3rem;
  padding: clamp(1.4rem, 3vw, 2.4rem);
  border: 1px solid rgba(20, 24, 29, 0.09);
  border-radius: 24px;
  background:
    radial-gradient(circle at 0% 0%, rgba(255, 107, 53, 0.08), transparent 36%),
    #fff;
  box-shadow: 0 24px 50px -42px rgba(20, 24, 29, 0.45);
}

@media (min-width: 960px) {
  .about-empower {
    grid-template-columns: minmax(0, 0.82fr) minmax(0, 1.18fr);
    gap: 3.5rem;
    align-items: start;
    padding: 2.4rem 2.6rem;
  }
}

.about-empower-kicker {
  margin: 0 0 1rem;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 9px;
  font-weight: 600;
  letter-spacing: 0.16em;
  color: var(--orange);
}

.about-empower-title {
  margin: 0;
  font-size: clamp(1.65rem, 3.2vw, 2.15rem);
  font-weight: 650;
  line-height: 1.15;
  letter-spacing: -0.03em;
}

.about-empower-title .about-serif {
  color: var(--orange);
}

.about-empower-lead {
  max-width: 25rem;
  margin: 1rem 0 0;
  font-size: 13px;
  line-height: 1.8;
  color: rgba(20, 24, 29, 0.55);
}

.about-empower-list {
  margin: 0;
  padding: 0;
  border-top: 1px solid rgba(20, 24, 29, 0.09);
  list-style: none;
}

.about-empower-item {
  display: grid;
  grid-template-columns: auto minmax(0, 1fr) auto;
  gap: 0.9rem;
  align-items: start;
  padding: 1.05rem 0.35rem;
  border-bottom: 1px solid rgba(20, 24, 29, 0.09);
  transition: background 0.25s ease;
}

.about-empower-item:hover {
  background: rgba(255, 107, 53, 0.045);
}

.about-empower-meta {
  display: flex;
  align-items: center;
  gap: 0.55rem;
  padding-top: 0.08rem;
}

.about-empower-num {
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 8.5px;
  letter-spacing: 0.12em;
  color: rgba(20, 24, 29, 0.3);
}

.about-empower-icon {
  display: grid;
  width: 30px;
  height: 30px;
  place-items: center;
  border-radius: 50%;
  background: rgba(255, 107, 53, 0.1);
  color: var(--orange);
}

.about-empower-copy h4 {
  margin: 0;
  font-size: 14.5px;
  font-weight: 650;
  letter-spacing: -0.02em;
}

.about-empower-copy p {
  margin: 0.3rem 0 0;
  font-size: 12px;
  line-height: 1.7;
  color: rgba(20, 24, 29, 0.52);
}

.about-empower-tag {
  padding-top: 0.28rem;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 8px;
  font-weight: 600;
  letter-spacing: 0.12em;
  color: rgba(20, 24, 29, 0.28);
  white-space: nowrap;
}

@media (max-width: 639px) {
  .about-empower {
    margin-top: 2.2rem;
    padding: 1.25rem;
    border-radius: 20px;
  }

  .about-empower-item {
    grid-template-columns: auto minmax(0, 1fr);
    gap: 0.75rem;
    padding: 0.95rem 0.15rem;
  }

  .about-empower-tag {
    grid-column: 2;
    padding-top: 0;
    margin-top: -0.05rem;
  }
}

.about-quote {
  max-width: 56rem;
  margin: 3.2rem 0 0 auto;
  font-size: clamp(1.1rem, 2.3vw, 1.55rem);
  font-weight: 500;
  line-height: 1.55;
  letter-spacing: -0.02em;
  color: rgba(20, 24, 29, 0.72);
}

.about-quote .about-serif {
  color: var(--orange);
}

/* ---------- How ---------- */
.about-how {
  position: relative;
  overflow: hidden;
  background: var(--ink);
  color: var(--paper);
  padding-top: 5.5rem;
  padding-bottom: 5.5rem;
}

@media (min-width: 768px) {
  .about-how {
    padding-top: 7rem;
    padding-bottom: 7rem;
  }
}

.about-how-glow {
  position: absolute;
  top: -280px;
  left: 12%;
  width: 560px;
  height: 560px;
  border-radius: 999px;
  background: rgba(255, 107, 53, 0.2);
  filter: blur(120px);
  pointer-events: none;
}

.about-how-grid {
  position: relative;
  display: grid;
  gap: 3rem;
}

@media (min-width: 1024px) {
  .about-how-grid {
    grid-template-columns: minmax(0, 0.85fr) minmax(0, 1.15fr);
    gap: 5rem;
  }
}

.about-label-light p {
  color: rgba(255, 255, 255, 0.35);
}

.about-h2-light {
  margin-top: 1.5rem;
  color: var(--paper);
}

.about-h2-light .about-serif {
  color: var(--sun);
}

.about-how-lead {
  max-width: 26rem;
  margin: 1.4rem 0 0;
  font-size: 14px;
  line-height: 1.85;
  color: rgba(255, 255, 255, 0.48);
}

.about-how-note {
  display: flex;
  align-items: flex-start;
  gap: 0.55rem;
  max-width: 22rem;
  margin-top: 1.8rem;
  padding-top: 1.2rem;
  border-top: 1px solid rgba(255, 255, 255, 0.12);
  font-size: 11.5px;
  line-height: 1.7;
  color: rgba(255, 255, 255, 0.55);
}

.about-how-note svg {
  flex: none;
  margin-top: 0.15rem;
  color: var(--sun);
}

.about-how-intro .about-btn {
  margin-top: 2rem;
}

.about-steps {
  margin: 0;
  padding: 0;
  list-style: none;
}

.about-step {
  display: grid;
  grid-template-columns: 2.5rem minmax(0, 1fr) 1.5rem;
  gap: 1rem;
  align-items: start;
  padding: 1.5rem 0;
  border-top: 1px solid rgba(255, 255, 255, 0.13);
}

.about-step:last-child {
  border-bottom: 1px solid rgba(255, 255, 255, 0.13);
}

.about-step-num {
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 10px;
  letter-spacing: 0.14em;
  color: var(--orange);
}

.about-step-top {
  display: flex;
  flex-wrap: wrap;
  align-items: baseline;
  justify-content: space-between;
  gap: 0.4rem 1rem;
}

.about-step h3 {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  letter-spacing: -0.02em;
  color: var(--paper);
}

.about-step-top span {
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 8.5px;
  letter-spacing: 0.16em;
  color: rgba(255, 255, 255, 0.28);
}

.about-step p {
  max-width: 34rem;
  margin: 0.55rem 0 0;
  font-size: 12.5px;
  line-height: 1.75;
  color: rgba(255, 255, 255, 0.48);
}

.about-step-arrow {
  margin-top: 0.1rem;
  color: rgba(255, 255, 255, 0.2);
  transition: color 0.22s ease, transform 0.22s ease;
}

.about-step:hover .about-step-arrow {
  color: var(--orange);
  transform: translateX(4px);
}

/* ---------- Closing ---------- */
.about-closing {
  padding: 4.5rem 0;
}

@media (min-width: 768px) {
  .about-closing {
    padding: 6.5rem 0;
  }
}

.about-closing-card {
  position: relative;
  display: flex;
  flex-direction: column;
  gap: 2.2rem;
  overflow: hidden;
  padding: 2rem;
  border-radius: 24px;
  background: var(--orange);
  color: var(--ink);
}

@media (min-width: 900px) {
  .about-closing-card {
    flex-direction: row;
    align-items: flex-end;
    justify-content: space-between;
    padding: 3.2rem;
  }
}

.about-closing-glow {
  position: absolute;
  top: -120px;
  right: -80px;
  width: 360px;
  height: 360px;
  border-radius: 999px;
  background: rgba(255, 209, 102, 0.55);
  filter: blur(80px);
  pointer-events: none;
}

.about-closing-copy {
  position: relative;
}

.about-closing-kicker {
  margin: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 9px;
  font-weight: 600;
  letter-spacing: 0.2em;
  color: rgba(20, 24, 29, 0.48);
}

.about-closing h2 {
  margin: 1rem 0 0;
  font-size: clamp(2rem, 4.5vw, 3.2rem);
  font-weight: 650;
  line-height: 1.08;
  letter-spacing: -0.045em;
}

.about-closing h2 .about-serif {
  color: var(--paper);
}

.about-closing-actions {
  position: relative;
  display: flex;
  flex-wrap: wrap;
  gap: 0.7rem;
}

/* ---------- Footer ---------- */
.about-footer {
  border-top: 1px solid rgba(20, 24, 29, 0.08);
}

.about-footer-inner {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  padding-top: 1.8rem;
  padding-bottom: 1.8rem;
  font-size: 11.5px;
  color: rgba(20, 24, 29, 0.4);
}

@media (min-width: 900px) {
  .about-footer-inner {
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
  }
}

.about-footer-brand {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: rgba(20, 24, 29, 0.7);
  font-weight: 600;
}

.about-brand-mark-sm {
  width: 20px;
  height: 20px;
  border-radius: 6px;

  background: #15151B;}

.about-footer-tagline {
  font-weight: 400;
  color: rgba(20, 24, 29, 0.32);
}

.about-footer-links {
  display: flex;
  flex-wrap: wrap;
  gap: 1.2rem;
}

.about-footer-links a {
  transition: color 0.2s ease;
}

.about-footer-links a:hover {
  color: var(--orange);
}

.about-footer-copy {
  margin: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 9.5px;
  letter-spacing: 0.08em;
  color: rgba(20, 24, 29, 0.28);
}

/* ---------- Motion ---------- */
.reveal {
  opacity: 0;
  animation: about-rise 0.8s cubic-bezier(0.22, 1, 0.36, 1) forwards;
  animation-delay: var(--d, 0s);
}

@keyframes about-rise {
  from {
    opacity: 0;
    transform: translateY(16px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (prefers-reduced-motion: reduce) {
  .reveal {
    opacity: 1;
    animation: none;
  }

  .about-orbit,
  .about-float,
  .about-ticker-track {
    animation: none;
  }

  .about-specimen,
  .about-btn,
  .about-feature,
  .about-step-arrow {
    transition: none;
  }
}
</style>
