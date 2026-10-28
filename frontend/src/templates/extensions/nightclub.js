import NeonNight from '../../pages/templates/ext/nightclub/NeonNight.vue'
import ChromeFever from '../../pages/templates/ext/nightclub/ChromeFever.vue'
import Terminal303 from '../../pages/templates/ext/nightclub/Terminal303.vue'
import VinylAfterhours from '../../pages/templates/ext/nightclub/VinylAfterhours.vue'
import TicketBooth from '../../pages/templates/ext/nightclub/TicketBooth.vue'
import VelvetMember from '../../pages/templates/ext/nightclub/VelvetMember.vue'
import BassDrop from '../../pages/templates/ext/nightclub/BassDrop.vue'
import ChampagneList from '../../pages/templates/ext/nightclub/ChampagneList.vue'
import SprayKill from '../../pages/templates/ext/nightclub/SprayKill.vue'
import CoverStory from '../../pages/templates/ext/nightclub/CoverStory.vue'
import BlackoutMinimal from '../../pages/templates/ext/nightclub/BlackoutMinimal.vue'
import RisoRave from '../../pages/templates/ext/nightclub/RisoRave.vue'
import Hologram from '../../pages/templates/ext/nightclub/Hologram.vue'
import BlueprintClub from '../../pages/templates/ext/nightclub/BlueprintClub.vue'
import GlassLounge from '../../pages/templates/ext/nightclub/GlassLounge.vue'
import PixelPulse from '../../pages/templates/ext/nightclub/PixelPulse.vue'
import FilmNoir from '../../pages/templates/ext/nightclub/FilmNoir.vue'
import DecoAfterDark from '../../pages/templates/ext/nightclub/DecoAfterDark.vue'
import SmokeSignal from '../../pages/templates/ext/nightclub/SmokeSignal.vue'
import BarcodeEntry from '../../pages/templates/ext/nightclub/BarcodeEntry.vue'
import MirrorBall from '../../pages/templates/ext/nightclub/MirrorBall.vue'
import BackstagePass from '../../pages/templates/ext/nightclub/BackstagePass.vue'
import LateCheck from '../../pages/templates/ext/nightclub/LateCheck.vue'
import NeonMotel from '../../pages/templates/ext/nightclub/NeonMotel.vue'
import ConcreteRave from '../../pages/templates/ext/nightclub/ConcreteRave.vue'
import LaserGrid from '../../pages/templates/ext/nightclub/LaserGrid.vue'
import GoldBar from '../../pages/templates/ext/nightclub/GoldBar.vue'
import AcidHouse from '../../pages/templates/ext/nightclub/AcidHouse.vue'
import MidnightFM from '../../pages/templates/ext/nightclub/MidnightFM.vue'
import ClubIndex from '../../pages/templates/ext/nightclub/ClubIndex.vue'

export default [
  {
    value: 'neon_night',
    label: '霓虹夜色',
    preview: '深紫霓虹与超大姓名',
    gradient: 'linear-gradient(135deg,#1e1b4b,#a21caf)',
    component: NeonNight,
    ambient: 'amb-dark',
  },
  {
    value: 'chrome_fever',
    label: '铬金狂潮',
    preview: '液态铬金属分屏海报',
    gradient: 'linear-gradient(135deg,#0c0d10,#c9ccd3 55%,#4a4d55)',
    component: ChromeFever,
    ambient: 'amb-dark',
  },
  {
    value: 'terminal_303',
    label: '绿码终端',
    preview: '黑客终端命令行名片',
    gradient: 'linear-gradient(135deg,#020a06,#0f3d22,#39ff88)',
    component: Terminal303,
    ambient: 'amb-dark-green',
  },
  {
    value: 'vinyl_afterhours',
    label: '黑胶深夜',
    preview: '奶油纸黑胶唱片档案',
    gradient: 'linear-gradient(135deg,#17130f,#d8572c)',
    component: VinylAfterhours,
    ambient: 'amb-warm',
  },
  {
    value: 'ticket_booth',
    label: '票根入场',
    preview: '复古票根虚线撕口',
    gradient: 'linear-gradient(135deg,#0b1930,#e5d3b2 60%,#a33a2b)',
    component: TicketBooth,
    ambient: 'amb-dark',
  },
  {
    value: 'velvet_member',
    label: '丝绒会员',
    preview: '黑金会员卡烫金质感',
    gradient: 'linear-gradient(135deg,#0a0806,#d4af6a)',
    component: VelvetMember,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'bass_drop',
    label: '低音炮台',
    preview: '荧光绿波形DJ台',
    gradient: 'linear-gradient(135deg,#070a05,#1e2e05,#c8ff2f)',
    component: BassDrop,
    ambient: 'amb-dark-green',
  },
  {
    value: 'champagne_list',
    label: '香槟酒单',
    preview: '象牙纸香槟菜单体',
    gradient: 'linear-gradient(135deg,#f5eede,#b08a3e)',
    component: ChampagneList,
    ambient: 'amb-warm',
  },
  {
    value: 'spray_kill',
    label: '荧光喷漆',
    preview: '涂鸦喷漆撞色海报',
    gradient: 'linear-gradient(135deg,#0c0c0b,#ff2d87 55%,#d6ff2f)',
    component: SprayKill,
    ambient: 'amb-dark',
  },
  {
    value: 'cover_story',
    label: '封面人物',
    preview: '红黑杂志封面大片',
    gradient: 'linear-gradient(135deg,#f3f0e8,#d92b1c 55%,#0d0b12)',
    component: CoverStory,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'blackout_minimal',
    label: '暗黑极简',
    preview: '纯黑留白极简排版',
    gradient: 'linear-gradient(135deg,#000000,#2a2a2a)',
    component: BlackoutMinimal,
    ambient: 'amb-dark',
  },
  {
    value: 'riso_rave',
    label: '丝网狂欢',
    preview: '新闻纸红蓝丝网印',
    gradient: 'linear-gradient(135deg,#e9e0cb,#e23b2e 55%,#1e2a49)',
    component: RisoRave,
    ambient: 'amb-warm',
  },
  {
    value: 'hologram',
    label: '全息幻彩',
    preview: '彩虹全息玻璃卡片',
    gradient: 'linear-gradient(135deg,#4cc9f0,#b5179e 45%,#ffd166 75%,#06d6a0)',
    component: Hologram,
    ambient: 'amb-cool',
  },
  {
    value: 'blueprint_club',
    label: '蓝图俱乐部',
    preview: '青色工程蓝图标注',
    gradient: 'linear-gradient(135deg,#041f33,#0a5273,#7fdcff)',
    component: BlueprintClub,
    ambient: 'amb-cool',
  },
  {
    value: 'glass_lounge',
    label: '玻璃酒廊',
    preview: '磨砂玻璃珊瑚光晕',
    gradient: 'linear-gradient(135deg,#071318,#18a8a0 50%,#ff7a66)',
    component: GlassLounge,
    ambient: 'amb-cool',
  },
  {
    value: 'pixel_pulse',
    label: '像素脉冲',
    preview: '街机像素霓虹HUD',
    gradient: 'linear-gradient(135deg,#05050c,#f7ff3c 50%,#ff2e88)',
    component: PixelPulse,
    ambient: 'amb-dark',
  },
  {
    value: 'film_noir',
    label: '胶片夜行',
    preview: '琥珀胶片齿孔长条',
    gradient: 'linear-gradient(135deg,#0b0906,#6b3d1b,#d9a05b)',
    component: FilmNoir,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'deco_after_dark',
    label: '装饰夜宴',
    preview: '祖母绿金对称装饰',
    gradient: 'linear-gradient(135deg,#06231c,#0c4a39,#c9a959)',
    component: DecoAfterDark,
    ambient: 'amb-dark-green',
  },
  {
    value: 'smoke_signal',
    label: '烟雾信号',
    preview: '紫灰烟雾留白字',
    gradient: 'linear-gradient(135deg,#0b0910,#3a2d55,#b7a5db)',
    component: SmokeSignal,
    ambient: 'amb-dark',
  },
  {
    value: 'barcode_entry',
    label: '条码入场',
    preview: '警示黄工业通行证',
    gradient: 'linear-gradient(135deg,#f0ede4,#ffd400 55%,#111110)',
    component: BarcodeEntry,
    ambient: 'amb-plain',
  },
  {
    value: 'mirror_ball',
    label: '镜球时刻',
    preview: '银色镜球粉色光斑',
    gradient: 'linear-gradient(135deg,#120716,#6d4c7a,#ffa8d4)',
    component: MirrorBall,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'backstage_pass',
    label: '后台通行',
    preview: '酒红挂绳后台证',
    gradient: 'linear-gradient(135deg,#24070f,#d23a2e 55%,#f4e7d3)',
    component: BackstagePass,
    ambient: 'amb-dark-red',
  },
  {
    value: 'late_check',
    label: '深夜小票',
    preview: '热敏纸小票清单',
    gradient: 'linear-gradient(135deg,#0f2620,#f3f0e4)',
    component: LateCheck,
    ambient: 'amb-dark-green',
  },
  {
    value: 'neon_motel',
    label: '霓虹旅馆',
    preview: '复古明信片汽车旅馆',
    gradient: 'linear-gradient(135deg,#0d2f30,#ff8b78 55%,#f9ecd4)',
    component: NeonMotel,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'concrete_rave',
    label: '混凝土狂欢',
    preview: '水泥灰橙色粗野海报',
    gradient: 'linear-gradient(135deg,#b7b1a5,#ff4d00 55%,#171614)',
    component: ConcreteRave,
    ambient: 'amb-plain',
  },
  {
    value: 'laser_grid',
    label: '激光矩阵',
    preview: '红色激光透视网格',
    gradient: 'linear-gradient(135deg,#060507,#ff2436 45%,#2a0a0e)',
    component: LaserGrid,
    ambient: 'amb-dark-red',
  },
  {
    value: 'gold_bar',
    label: '金条夜场',
    preview: '黑底金条浮雕质感',
    gradient: 'linear-gradient(135deg,#0a0805,#d4af6a 55%,#f5e3b7)',
    component: GoldBar,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'acid_house',
    label: '酸性浩室',
    preview: '荧光黄笑脸锐舞传单',
    gradient: 'linear-gradient(135deg,#d9ff00,#0b0b06)',
    component: AcidHouse,
    ambient: 'amb-warm-green',
  },
  {
    value: 'midnight_fm',
    label: '午夜电台',
    preview: '深蓝琥珀收音机刻度',
    gradient: 'linear-gradient(135deg,#081226,#ffb347)',
    component: MidnightFM,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'club_index',
    label: '俱乐部索引',
    preview: '瑞士白纸红黑索引',
    gradient: 'linear-gradient(135deg,#efede6,#e63322 55%,#111110)',
    component: ClubIndex,
    ambient: 'amb-plain',
  },
]
