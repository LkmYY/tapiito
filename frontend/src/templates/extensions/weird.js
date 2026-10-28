/**
 * 「奇怪的想法」30 款猎奇横屏风格
 *
 * 全部复用同一个 WeirdStage 组件（Canvas 渲染引擎），
 * 通过 styleKey 选择具体绘制器，避免 30 个重复的 Vue 文件。
 */
import { defineComponent, h } from 'vue'
import WeirdStage from '../../pages/templates/ext/weird/WeirdStage.vue'

function make(key, label, preview, gradient) {
  const Comp = defineComponent({
    name: `Weird_${key}`,
    props: {
      payload: { type: Object, default: () => ({}) },
      nfcId: { type: String, default: '' },
    },
    setup(props) {
      return () => h(WeirdStage, {
        payload: props.payload,
        nfcId: props.nfcId,
        styleKey: key,
      })
    },
  })
  return {
    value: key,
    label,
    preview,
    gradient,
    component: Comp,
    ambient: 'amb-dark',
  }
}

export default [
  /* A. 眼睛与凝视 */
  make('sharingan', '血轮眼', '红色眼球 + 三勾玉旋转', 'linear-gradient(135deg,#450a0a,#ef4444)'),
  make('kaleido_eye', '万花筒之眼', '六重对称几何瞳', 'linear-gradient(135deg,#2e1065,#a855f7)'),
  make('infinite_pupil', '无尽瞳孔', '瞳孔无限向内坠落', 'linear-gradient(135deg,#082f49,#38bdf8)'),
  make('watcher', '凝视者', '一只眼跟着你的指针', 'linear-gradient(135deg,#422006,#fbbf24)'),
  make('compound_eye', '复眼阵列', '60 只小眼随机眨动', 'linear-gradient(135deg,#083344,#22d3ee)'),

  /* B. 时间 */
  make('flip_clock', '巨型翻页钟', '全屏翻牌数字每秒跳动', 'linear-gradient(135deg,#0a0a0a,#525252)'),
  make('countdown_bomb', '倒计时炸弹', '红色倒计时越近越抖', 'linear-gradient(135deg,#450a0a,#f87171)'),
  make('broken_clock', '失控指针钟', '指针乱转、偶尔卡住', 'linear-gradient(135deg,#111827,#9ca3af)'),
  make('doomsday_clock', '世界末日钟', '永远停在 23:59:59', 'linear-gradient(135deg,#1c1206,#f97316)'),
  make('hourglass_reverse', '沙漏倒流', '沙粒自下而上流动', 'linear-gradient(135deg,#291a06,#fbbf24)'),

  /* C. 文字 */
  make('marquee_warning', '循环警告跑马灯', '警示带 + 横向跑马灯', 'linear-gradient(135deg,#0a0a0a,#facc15)'),
  make('danmaku_wall', '弹幕墙', '多条弹幕不同速度飘过', 'linear-gradient(135deg,#0b0b12,#e2e8f0)'),
  make('neon_vertical', '竖排霓虹字', '霓虹字自上而下滚动', 'linear-gradient(135deg,#04121a,#22d3ee)'),
  make('typewriter', '打字机', '一句话写完又删掉', 'linear-gradient(135deg,#07130c,#4ade80)'),
  make('matrix_rain', '乱码雨', '竖排乱码字符下落', 'linear-gradient(135deg,#02120a,#22c55e)'),

  /* D. 伪系统 */
  make('fake_terminal', '假终端', '命令行逐行输出后卡住', 'linear-gradient(135deg,#0b0f0c,#86efac)'),
  make('fake_monitor', '假监控面板', 'CPU/内存条疯狂跳动', 'linear-gradient(135deg,#070b12,#38bdf8)'),
  make('blue_screen', '蓝屏', '经典 BSOD，进度永远 0%', 'linear-gradient(135deg,#0b57d0,#9ecbff)'),
  make('radar_scan', '雷达扫描', '绿色雷达旋转 + 光点', 'linear-gradient(135deg,#03150c,#22c55e)'),
  make('spectrum_bars', '频谱分析', '声波条形上下跳动', 'linear-gradient(135deg,#1a0714,#f472b6)'),

  /* E. 故障与复古 */
  make('crt_snow', 'CRT 雪花屏', '电视无信号噪点', 'linear-gradient(135deg,#0b0b0d,#a1a1aa)'),
  make('scanline_glitch', '扫描线撕裂', 'RGB 错位 + 横向撕裂', 'linear-gradient(135deg,#0a0508,#f43f5e)'),
  make('vhs_tape', 'VHS 录像带', '磁带噪点 + REC 红点', 'linear-gradient(135deg,#120b0b,#f87171)'),
  make('oscilloscope', '老式示波器', '绿色波形横向滚动', 'linear-gradient(135deg,#04120a,#4ade80)'),
  make('frame_glitch', '跳帧故障', '画面被色块错位切割', 'linear-gradient(135deg,#0a0710,#a855f7)'),

  /* F. 奇趣 */
  make('infinite_mirror', '无限镜像', '画面把自己缩进自己里', 'linear-gradient(135deg,#04050a,#cbd5e1)'),
  make('kaleidoscope', '万花筒', '八重镜面花瓣缓慢旋转', 'linear-gradient(135deg,#05040a,#818cf8)'),
  make('heartbeat', '心跳失控', '心电波形 + BPM 疯涨', 'linear-gradient(135deg,#040a0c,#22d3ee)'),
  make('jellyfish', '深海生物', '发光水母缓缓漂浮', 'linear-gradient(135deg,#02060f,#67e8f9)'),
  make('cosmos', '宇宙坐标', '星图旋转 + 坐标滚动', 'linear-gradient(135deg,#02030a,#c4b5fd)'),
]
