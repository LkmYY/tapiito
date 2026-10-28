/**
 * 首页预渲染：把 Vue 渲染后的 DOM 注入到 dist/index.html 的 #app 里。
 *
 * 为什么需要：百度等爬虫不执行 JS，SPA 的空 <div id="app"></div> 等于零内容。
 * 注入后爬虫能直接读到全部文案；用户访问时 Vue 挂载会接管并替换这份静态内容。
 *
 * 用法：先起本地静态服务（dist），再跑本脚本。
 */
import fs from 'node:fs'
import path from 'node:path'
import http from 'node:http'

const SOURCE = process.argv[2] || 'http://127.0.0.1:4180/'
const DIST = process.argv[3] || 'H:/project/tapii/tapiito/frontend/dist'
const PORT = 9235
const sleep = (ms) => new Promise((r) => setTimeout(r, ms))

class Cdp {
  constructor(ws) {
    this.id = 0
    this.p = new Map()
    this.ws = new WebSocket(ws)
    this.ready = new Promise((res) => this.ws.addEventListener('open', () => res()))
    this.ws.addEventListener('message', (ev) => {
      const m = JSON.parse(ev.data)
      if (m.id && this.p.has(m.id)) {
        const { res, rej } = this.p.get(m.id)
        this.p.delete(m.id)
        m.error ? rej(new Error(JSON.stringify(m.error))) : res(m.result)
      }
    })
  }
  async send(method, params = {}) {
    await this.ready
    const id = ++this.id
    return new Promise((res, rej) => {
      this.p.set(id, { res, rej })
      this.ws.send(JSON.stringify({ id, method, params }))
    })
  }
}

// 脚本内起静态服务（外部 python 服务在子 shell 里经常被回收）
const MIME = {
  '.html': 'text/html; charset=utf-8', '.js': 'text/javascript', '.css': 'text/css',
  '.json': 'application/json', '.webp': 'image/webp', '.jpg': 'image/jpeg',
  '.png': 'image/png', '.svg': 'image/svg+xml', '.woff2': 'font/woff2',
}
const server = http.createServer((req, res) => {
  let p = decodeURIComponent(req.url.split('?')[0])
  if (p === '/') p = '/index.html'
  let file = path.join(DIST, p)
  if (!fs.existsSync(file) || fs.statSync(file).isDirectory()) file = path.join(DIST, 'index.html')
  const ext = path.extname(file)
  res.writeHead(200, { 'Content-Type': MIME[ext] || 'application/octet-stream' })
  fs.createReadStream(file).pipe(res)
})
await new Promise((r) => server.listen(4180, '127.0.0.1', r))
console.log('静态服务已起: http://127.0.0.1:4180')

const tab = await (await fetch(`http://127.0.0.1:${PORT}/json/new?${encodeURIComponent('about:blank')}`, { method: 'PUT' })).json()
const cdp = new Cdp(tab.webSocketDebuggerUrl)
await cdp.send('Page.enable')
await cdp.send('Runtime.enable')
await cdp.send('Emulation.setDeviceMetricsOverride', { width: 1280, height: 900, deviceScaleFactor: 1, mobile: false })
await cdp.send('Page.navigate', { url: SOURCE })
// 等 Vue 挂载 + 图片占位完成（本地无后端，部分数据会空，但文案齐全）
await sleep(6000)

const r = await cdp.send('Runtime.evaluate', {
  expression: `
(() => {
  const app = document.querySelector('#app');
  if (!app) return { error: 'no #app', href: location.href, title: document.title, head: document.body.innerHTML.slice(0, 200) };
  // 去掉 v-motion 留下的内联过渡样式，避免静态版本闪一下
  const clone = app.cloneNode(true);
  clone.querySelectorAll('[style*="opacity"]').forEach(el => el.removeAttribute('style'));
  return {
    html: clone.innerHTML,
    textLen: (app.innerText || '').replace(/\\s+/g, '').length,
  };
})()
`,
  returnByValue: true,
})

const data = r.result.value
if (!data || data.error) {
  console.error('抓取失败:', data)
  process.exit(1)
}

const indexPath = path.join(DIST, 'index.html')
let html = fs.readFileSync(indexPath, 'utf8')

const EMPTY = '<div id="app"></div>'
if (!html.includes(EMPTY)) {
  console.error('index.html 里没有空的 #app 容器，可能已经预渲染过或结构变了')
  process.exit(1)
}

html = html.replace(EMPTY, `<div id="app">${data.html}</div>`)

// 预渲染内容不需要参与动画，直接显示
html = html.replace(
  '</head>',
  '<style>[data-prerendered] [style*="opacity: 0"]{opacity:1!important;transform:none!important}</style></head>',
)

fs.writeFileSync(indexPath, html, 'utf8')

// 回读校验：静态 HTML 里的可见字数
const textOf = (s) =>
  s
    .replace(/<script[\s\S]*?<\/script>/gi, '')
    .replace(/<style[\s\S]*?<\/style>/gi, '')
    .replace(/<[^>]+>/g, '')
    .replace(/&[a-z]+;/gi, '')
    .replace(/\s+/g, '')

const finalHtml = fs.readFileSync(indexPath, 'utf8')
console.log(
  JSON.stringify({
    渲染时字数: data.textLen,
    注入后HTML字节: finalHtml.length,
    静态可见字数: textOf(finalHtml).length,
  }),
)

await fetch(`http://127.0.0.1:${PORT}/json/close/${tab.id}`).catch(() => {})
server.close()
