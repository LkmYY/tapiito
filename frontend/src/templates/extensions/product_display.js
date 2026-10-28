import GraphiteSpec from '../../pages/templates/ext/product_display/GraphiteSpec.vue'
import AtelierNoir from '../../pages/templates/ext/product_display/AtelierNoir.vue'
import NeonPulse from '../../pages/templates/ext/product_display/NeonPulse.vue'
import EditorialFold from '../../pages/templates/ext/product_display/EditorialFold.vue'
import FrostGlass from '../../pages/templates/ext/product_display/FrostGlass.vue'
import BrutalCart from '../../pages/templates/ext/product_display/BrutalCart.vue'
import RisoPrint from '../../pages/templates/ext/product_display/RisoPrint.vue'
import ClayStudio from '../../pages/templates/ext/product_display/ClayStudio.vue'
import TrailGear from '../../pages/templates/ext/product_display/TrailGear.vue'
import VelvetBloom from '../../pages/templates/ext/product_display/VelvetBloom.vue'
import PixelDrop from '../../pages/templates/ext/product_display/PixelDrop.vue'
import InkBottle from '../../pages/templates/ext/product_display/InkBottle.vue'
import BauhausBlock from '../../pages/templates/ext/product_display/BauhausBlock.vue'
import Datasheet from '../../pages/templates/ext/product_display/Datasheet.vue'
import FilmRoll from '../../pages/templates/ext/product_display/FilmRoll.vue'
import TicketStub from '../../pages/templates/ext/product_display/TicketStub.vue'
import SplitLab from '../../pages/templates/ext/product_display/SplitLab.vue'
import PosterBold from '../../pages/templates/ext/product_display/PosterBold.vue'
import StackCards from '../../pages/templates/ext/product_display/StackCards.vue'
import ArchiveIndex from '../../pages/templates/ext/product_display/ArchiveIndex.vue'
import ChromeY2K from '../../pages/templates/ext/product_display/ChromeY2K.vue'
import TerrazzoTile from '../../pages/templates/ext/product_display/TerrazzoTile.vue'
import MonolithNoir from '../../pages/templates/ext/product_display/MonolithNoir.vue'
import Botanical from '../../pages/templates/ext/product_display/Botanical.vue'
import Blueprint from '../../pages/templates/ext/product_display/Blueprint.vue'

export default [
  {
    value: 'graphite_spec',
    label: '石墨极简',
    preview: '石墨灰网格 + 规格清单',
    gradient: 'linear-gradient(135deg,#fbfbfa,#d6d3cd)',
    component: GraphiteSpec,
    ambient: 'amb-plain',
  },
  {
    value: 'atelier_noir',
    label: '黑金订制',
    preview: '暖黑金细线 + 居中表盘',
    gradient: 'linear-gradient(135deg,#0b0a08,#c9a96a)',
    component: AtelierNoir,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'neon_pulse',
    label: '霓虹脉冲',
    preview: '暗紫霓虹 + 终端面板',
    gradient: 'linear-gradient(135deg,#06050d,#5eead4)',
    component: NeonPulse,
    ambient: 'amb-dark',
  },
  {
    value: 'editorial_fold',
    label: '杂志折页',
    preview: '米色衬线 + 编辑网格',
    gradient: 'linear-gradient(135deg,#f6f1e6,#9a3412)',
    component: EditorialFold,
    ambient: 'amb-warm',
  },
  {
    value: 'frost_glass',
    label: '霜感玻璃',
    preview: '冰蓝玻璃 + 悬浮卡片',
    gradient: 'linear-gradient(135deg,#dfe8f4,#99f6e4)',
    component: FrostGlass,
    ambient: 'amb-cool',
  },
  {
    value: 'brutal_cart',
    label: '粗野货架',
    preview: '亮黄粗黑块 + 硬阴影',
    gradient: 'linear-gradient(135deg,#f5e900,#ff2d00)',
    component: BrutalCart,
    ambient: 'amb-plain',
  },
  {
    value: 'riso_print',
    label: '孔版印刷',
    preview: '米纸双色 + 错位印刷',
    gradient: 'linear-gradient(135deg,#f4ead7,#e04a2f)',
    component: RisoPrint,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'clay_studio',
    label: '陶土手作',
    preview: '赤陶奶油 + 手工弧线',
    gradient: 'linear-gradient(135deg,#f2e8da,#b9673c)',
    component: ClayStudio,
    ambient: 'amb-warm',
  },
  {
    value: 'trail_gear',
    label: '山野装备',
    preview: '森林深绿 + 等高线图',
    gradient: 'linear-gradient(135deg,#101b15,#b9d17a)',
    component: TrailGear,
    ambient: 'amb-dark-green',
  },
  {
    value: 'velvet_bloom',
    label: '丝绒花瓣',
    preview: '粉雾拱门 + 香氛瓶身',
    gradient: 'linear-gradient(135deg,#fbeef1,#c08497)',
    component: VelvetBloom,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'pixel_drop',
    label: '像素上架',
    preview: '像素星空 + 游戏面板',
    gradient: 'linear-gradient(135deg,#141026,#7cf29b)',
    component: PixelDrop,
    ambient: 'amb-dark',
  },
  {
    value: 'ink_bottle',
    label: '墨色留白',
    preview: '宣纸墨器 + 竖排标题',
    gradient: 'linear-gradient(135deg,#f3efe6,#8a3b2e)',
    component: InkBottle,
    ambient: 'amb-plain',
  },
  {
    value: 'bauhaus_block',
    label: '包豪斯块',
    preview: '红蓝黄几何 + 构成排版',
    gradient: 'linear-gradient(135deg,#f1ece1,#1d4ed8)',
    component: BauhausBlock,
    ambient: 'amb-warm',
  },
  {
    value: 'datasheet',
    label: '参数档案',
    preview: '工程白底 + 参数表格',
    gradient: 'linear-gradient(135deg,#e9edf2,#0e7490)',
    component: Datasheet,
    ambient: 'amb-cool',
  },
  {
    value: 'film_roll',
    label: '胶片接触',
    preview: '暖黑胶片 + 接触印样',
    gradient: 'linear-gradient(135deg,#17130f,#d9b27c)',
    component: FilmRoll,
    ambient: 'amb-dark',
  },
  {
    value: 'ticket_stub',
    label: '票根商店',
    preview: '米色票根 + 打孔虚线',
    gradient: 'linear-gradient(135deg,#e7e2d6,#c2410c)',
    component: TicketStub,
    ambient: 'amb-plain',
  },
  {
    value: 'split_lab',
    label: '分屏实验',
    preview: '荧光绿分屏 + 实验标签',
    gradient: 'linear-gradient(135deg,#d7f34a,#101208)',
    component: SplitLab,
    ambient: 'amb-cool',
  },
  {
    value: 'poster_bold',
    label: '巨幅海报',
    preview: '橙红海报 + 超大价格',
    gradient: 'linear-gradient(135deg,#f4e7d3,#d84f1e)',
    component: PosterBold,
    ambient: 'amb-warm',
  },
  {
    value: 'stack_cards',
    label: '卡片堆叠',
    preview: '奶油卡片 + 层叠纸感',
    gradient: 'linear-gradient(135deg,#e9e4d8,#a16207)',
    component: StackCards,
    ambient: 'amb-warm',
  },
  {
    value: 'archive_index',
    label: '档案索引',
    preview: '档案纸卡 + 红色印章',
    gradient: 'linear-gradient(135deg,#e8e1d3,#9a3412)',
    component: ArchiveIndex,
    ambient: 'amb-plain',
  },
  {
    value: 'chrome_y2k',
    label: '千禧镜面',
    preview: '银紫镜面 + 未来球体',
    gradient: 'linear-gradient(135deg,#dfe3ef,#6d5bd0)',
    component: ChromeY2K,
    ambient: 'amb-cool',
  },
  {
    value: 'terrazzo_tile',
    label: '水磨石砖',
    preview: '彩石瓷砖 + 拼贴网格',
    gradient: 'linear-gradient(135deg,#f7f1e7,#9d5c63)',
    component: TerrazzoTile,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'monolith_noir',
    label: '暗夜巨碑',
    preview: '纯黑极简 + 超大价格',
    gradient: 'linear-gradient(135deg,#0a0a0a,#3a3a3a)',
    component: MonolithNoir,
    ambient: 'amb-dark',
  },
  {
    value: 'botanical',
    label: '植萃自然',
    preview: '草木淡绿 + 植物手绘',
    gradient: 'linear-gradient(135deg,#f0f2e6,#4d7c45)',
    component: Botanical,
    ambient: 'amb-warm-green',
  },
  {
    value: 'blueprint',
    label: '蓝图工程',
    preview: '深蓝图纸 + 工程制图',
    gradient: 'linear-gradient(135deg,#0e3a5c,#7dd3fc)',
    component: Blueprint,
    ambient: 'amb-cool',
  },
]
