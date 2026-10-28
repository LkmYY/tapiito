import NeonNight from '../../pages/templates/ext/business_card/NeonNight.vue'
import Porcelain from '../../pages/templates/ext/business_card/Porcelain.vue'
import EditorialCover from '../../pages/templates/ext/business_card/EditorialCover.vue'
import LetterpressRetro from '../../pages/templates/ext/business_card/LetterpressRetro.vue'
import GlassAurora from '../../pages/templates/ext/business_card/GlassAurora.vue'
import ChromeAtelier from '../../pages/templates/ext/business_card/ChromeAtelier.vue'
import PaperCollage from '../../pages/templates/ext/business_card/PaperCollage.vue'
import LiquidGradient from '../../pages/templates/ext/business_card/LiquidGradient.vue'
import PixelArcade from '../../pages/templates/ext/business_card/PixelArcade.vue'
import InkScroll from '../../pages/templates/ext/business_card/InkScroll.vue'
import BlueprintStudio from '../../pages/templates/ext/business_card/BlueprintStudio.vue'
import TerminalGreen from '../../pages/templates/ext/business_card/TerminalGreen.vue'
import DecoGold from '../../pages/templates/ext/business_card/DecoGold.vue'
import SwissGrid from '../../pages/templates/ext/business_card/SwissGrid.vue'
import CineFilm from '../../pages/templates/ext/business_card/CineFilm.vue'
import BrutalPop from '../../pages/templates/ext/business_card/BrutalPop.vue'
import Botanical from '../../pages/templates/ext/business_card/Botanical.vue'
import KyotoWashi from '../../pages/templates/ext/business_card/KyotoWashi.vue'
import HoloTicket from '../../pages/templates/ext/business_card/HoloTicket.vue'
import ClayMono from '../../pages/templates/ext/business_card/ClayMono.vue'
import SpaceOdyssey from '../../pages/templates/ext/business_card/SpaceOdyssey.vue'
import RisoPrint from '../../pages/templates/ext/business_card/RisoPrint.vue'
import Varsity from '../../pages/templates/ext/business_card/Varsity.vue'
import DataGrid from '../../pages/templates/ext/business_card/DataGrid.vue'
import PressProof from '../../pages/templates/ext/business_card/PressProof.vue'

export default [
  {
    value: 'neon_night',
    label: '霓虹夜色',
    preview: '深紫霓虹 + 超大姓名',
    gradient: 'linear-gradient(135deg,#0f0c29,#a21caf)',
    component: NeonNight,
    ambient: 'amb-dark',
  },
  {
    value: 'porcelain',
    label: '素白纸感',
    preview: '暖白留白 + 衬线大字',
    gradient: 'linear-gradient(135deg,#fffdf8,#e7e5e4)',
    component: Porcelain,
    ambient: 'amb-plain',
  },
  {
    value: 'editorial',
    label: '杂志封面',
    preview: '杂志大片 + 错位色块',
    gradient: 'linear-gradient(135deg,#f3efe6,#c2410c)',
    component: EditorialCover,
    ambient: 'amb-plain',
  },
  {
    value: 'letterpress',
    label: '复古票根',
    preview: '活版印刷 + 撕票存根',
    gradient: 'linear-gradient(135deg,#efe0c8,#9a3412)',
    component: LetterpressRetro,
    ambient: 'amb-warm',
  },
  {
    value: 'glass_aurora',
    label: '极光玻璃',
    preview: '极光渐变 + 磨砂玻璃',
    gradient: 'linear-gradient(135deg,#4338ca,#22d3ee)',
    component: GlassAurora,
    ambient: 'amb-cool',
  },
  {
    value: 'chrome_atelier',
    label: '液态金属',
    preview: '黑银铬面 + 巨幅海报',
    gradient: 'linear-gradient(135deg,#09090b,#a1a1aa)',
    component: ChromeAtelier,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'paper_collage',
    label: '手账拼贴',
    preview: '胶带票根 + 层叠纸片',
    gradient: 'linear-gradient(135deg,#fde68a,#fb7185)',
    component: PaperCollage,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'liquid_gradient',
    label: '流体渐变',
    preview: '流动色团 + 悬浮胶囊',
    gradient: 'linear-gradient(135deg,#fb7185,#6366f1)',
    component: LiquidGradient,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'pixel_arcade',
    label: '像素街机',
    preview: '像素边框 + 街机面板',
    gradient: 'linear-gradient(135deg,#111827,#22c55e)',
    component: PixelArcade,
    ambient: 'amb-dark-green',
  },
  {
    value: 'ink_scroll',
    label: '水墨卷轴',
    preview: '宣纸墨晕 + 朱红印章',
    gradient: 'linear-gradient(135deg,#f8fafc,#475569)',
    component: InkScroll,
    ambient: 'amb-plain',
  },
  {
    value: 'blueprint',
    label: '蓝图工坊',
    preview: '网格蓝图 + 技术标注',
    gradient: 'linear-gradient(135deg,#0c4a6e,#38bdf8)',
    component: BlueprintStudio,
    ambient: 'amb-cool',
  },
  {
    value: 'terminal',
    label: '终端绿屏',
    preview: '黑客终端 + 命令分行',
    gradient: 'linear-gradient(135deg,#020617,#16a34a)',
    component: TerminalGreen,
    ambient: 'amb-dark-green',
  },
  {
    value: 'deco_gold',
    label: '装饰金线',
    preview: '黑金对称 + 扇形纹样',
    gradient: 'linear-gradient(135deg,#0a0a0a,#b8860b)',
    component: DecoGold,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'swiss_grid',
    label: '瑞士网格',
    preview: '红黑网格 + 编号信息',
    gradient: 'linear-gradient(135deg,#fafaf9,#dc2626)',
    component: SwissGrid,
    ambient: 'amb-plain',
  },
  {
    value: 'cine_film',
    label: '胶片放映',
    preview: '齿孔胶片 + 琥珀光晕',
    gradient: 'linear-gradient(135deg,#111827,#f59e0b)',
    component: CineFilm,
    ambient: 'amb-dark',
  },
  {
    value: 'brutal_pop',
    label: '野兽波普',
    preview: '厚框撞色 + 硬质阴影',
    gradient: 'linear-gradient(135deg,#facc15,#2563eb)',
    component: BrutalPop,
    ambient: 'amb-warm',
  },
  {
    value: 'botanical',
    label: '植物标本',
    preview: '鼠尾草绿 + 叶片标本',
    gradient: 'linear-gradient(135deg,#ecfdf5,#166534)',
    component: Botanical,
    ambient: 'amb-warm-green',
  },
  {
    value: 'kyoto_washi',
    label: '和纸竖排',
    preview: '竖排姓名 + 朱印落款',
    gradient: 'linear-gradient(135deg,#f5f0e6,#9f1239)',
    component: KyotoWashi,
    ambient: 'amb-warm',
  },
  {
    value: 'holo_ticket',
    label: '全息票券',
    preview: '镭射箔面 + 撕口票券',
    gradient: 'linear-gradient(135deg,#c084fc,#22d3ee)',
    component: HoloTicket,
    ambient: 'amb-cool',
  },
  {
    value: 'clay_mono',
    label: '陶土软塑',
    preview: '暖灰陶土 + 软塑卡片',
    gradient: 'linear-gradient(135deg,#e7e5e4,#a8a29e)',
    component: ClayMono,
    ambient: 'amb-warm',
  },
  {
    value: 'space_odyssey',
    label: '星际漫游',
    preview: '星图轨道 + 任务徽章',
    gradient: 'linear-gradient(135deg,#020617,#6d28d9)',
    component: SpaceOdyssey,
    ambient: 'amb-dark',
  },
  {
    value: 'riso_print',
    label: '丝网错版',
    preview: '荧光错版 + 颗粒油墨',
    gradient: 'linear-gradient(135deg,#fb7185,#38bdf8)',
    component: RisoPrint,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'varsity',
    label: '学院徽章',
    preview: '校园徽章 + 酒红条纹',
    gradient: 'linear-gradient(135deg,#7f1d1d,#fef3c7)',
    component: Varsity,
    ambient: 'amb-dark-red',
  },
  {
    value: 'data_grid',
    label: '数据档案',
    preview: '深色仪表 + 指标卡片',
    gradient: 'linear-gradient(135deg,#0f172a,#0ea5e9)',
    component: DataGrid,
    ambient: 'amb-dark',
  },
  {
    value: 'press_proof',
    label: '报纸号外',
    preview: '新闻铅字 + 分栏排版',
    gradient: 'linear-gradient(135deg,#f5f5f4,#292524)',
    component: PressProof,
    ambient: 'amb-plain',
  },
]
