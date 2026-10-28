import CyberTerminal from '../../pages/templates/ext/anime_card/CyberTerminal.vue'
import GachaUR from '../../pages/templates/ext/anime_card/GachaUR.vue'
import GameboyScreen from '../../pages/templates/ext/anime_card/GameboyScreen.vue'
import GlassMoe from '../../pages/templates/ext/anime_card/GlassMoe.vue'
import HoloPrism from '../../pages/templates/ext/anime_card/HoloPrism.vue'
import IdolStage from '../../pages/templates/ext/anime_card/IdolStage.vue'
import MagicCircle from '../../pages/templates/ext/anime_card/MagicCircle.vue'
import MangaPanel from '../../pages/templates/ext/anime_card/MangaPanel.vue'
import MechaSuit from '../../pages/templates/ext/anime_card/MechaSuit.vue'
import PixelQuest from '../../pages/templates/ext/anime_card/PixelQuest.vue'
import RetroCassette from '../../pages/templates/ext/anime_card/RetroCassette.vue'
import SakuraTicket from '../../pages/templates/ext/anime_card/SakuraTicket.vue'
import SchoolDays from '../../pages/templates/ext/anime_card/SchoolDays.vue'
import ShrineOmikuji from '../../pages/templates/ext/anime_card/ShrineOmikuji.vue'
import SumiInk from '../../pages/templates/ext/anime_card/SumiInk.vue'
import Vaporwave from '../../pages/templates/ext/anime_card/Vaporwave.vue'
import VillainThrone from '../../pages/templates/ext/anime_card/VillainThrone.vue'
import NeoTokyo from '../../pages/templates/ext/anime_card/NeoTokyo.vue'
import StarRail from '../../pages/templates/ext/anime_card/StarRail.vue'
import MaidCafe from '../../pages/templates/ext/anime_card/MaidCafe.vue'
import NinjaScroll from '../../pages/templates/ext/anime_card/NinjaScroll.vue'
import DragonShrine from '../../pages/templates/ext/anime_card/DragonShrine.vue'
import BackstagePass from '../../pages/templates/ext/anime_card/BackstagePass.vue'
import MagicalGirl from '../../pages/templates/ext/anime_card/MagicalGirl.vue'
import DetectiveFile from '../../pages/templates/ext/anime_card/DetectiveFile.vue'
import SportsClub from '../../pages/templates/ext/anime_card/SportsClub.vue'
import GhostStory from '../../pages/templates/ext/anime_card/GhostStory.vue'
import SummerFestival from '../../pages/templates/ext/anime_card/SummerFestival.vue'
import CardBattle from '../../pages/templates/ext/anime_card/CardBattle.vue'
import ForestSpirit from '../../pages/templates/ext/anime_card/ForestSpirit.vue'

export default [
  {
    value: 'cyber_terminal',
    label: '绿码终端',
    preview: '黑客终端命令行档案',
    gradient: 'linear-gradient(135deg,#020604,#0f3d22,#39ff88)',
    component: CyberTerminal,
    ambient: 'amb-dark-green',
  },
  {
    value: 'gacha_ur',
    label: '抽卡 SSR',
    preview: '黑金卡池召唤界面',
    gradient: 'linear-gradient(135deg,#060402,#3b1c05,#fbbf24)',
    component: GachaUR,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'gameboy_screen',
    label: '掌机像素',
    preview: '灰绿点阵掌机屏',
    gradient: 'linear-gradient(135deg,#131b16,#4b5563,#a3e635)',
    component: GameboyScreen,
    ambient: 'amb-dark-green',
  },
  {
    value: 'glass_moe',
    label: '玻璃萌系',
    preview: '柔光玻璃拟态档案',
    gradient: 'linear-gradient(135deg,#e0e7ff,#fae8ff,#cffafe)',
    component: GlassMoe,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'holo_prism',
    label: '镭射棱镜',
    preview: '全息镭射收藏卡',
    gradient: 'linear-gradient(135deg,#07050f,#7c3aed,#22d3ee)',
    component: HoloPrism,
    ambient: 'amb-dark',
  },
  {
    value: 'idol_stage',
    label: '偶像舞台',
    preview: '打歌舞台直拍小卡',
    gradient: 'linear-gradient(135deg,#0b0614,#db2777,#facc15)',
    component: IdolStage,
    ambient: 'amb-dark',
  },
  {
    value: 'magic_circle',
    label: '魔法契约',
    preview: '星阵魔法契约书',
    gradient: 'linear-gradient(135deg,#05030c,#241348,#a78bfa)',
    component: MagicCircle,
    ambient: 'amb-dark',
  },
  {
    value: 'manga_panel',
    label: '漫画分镜',
    preview: '网点纸黑白分镜',
    gradient: 'linear-gradient(135deg,#e9e2d4,#4b5563)',
    component: MangaPanel,
    ambient: 'amb-plain',
  },
  {
    value: 'mecha_suit',
    label: '机甲座舱',
    preview: '机甲 HUD 驾驶员档案',
    gradient: 'linear-gradient(135deg,#050a10,#0e7490,#22d3ee)',
    component: MechaSuit,
    ambient: 'amb-dark',
  },
  {
    value: 'pixel_quest',
    label: '像素冒险',
    preview: '像素 RPG 队伍面板',
    gradient: 'linear-gradient(135deg,#0b0b22,#4338ca,#a3e635)',
    component: PixelQuest,
    ambient: 'amb-dark',
  },
  {
    value: 'retro_cassette',
    label: '复古磁带',
    preview: '磁带 A 面曲目卡',
    gradient: 'linear-gradient(135deg,#2b1710,#d97706,#fde68a)',
    component: RetroCassette,
    ambient: 'amb-warm',
  },
  {
    value: 'sakura_ticket',
    label: '樱花票根',
    preview: '粉色剧场入场券',
    gradient: 'linear-gradient(135deg,#f7dfe7,#fb7185)',
    component: SakuraTicket,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'school_days',
    label: '校园手帐',
    preview: '生徒手帐放学记录',
    gradient: 'linear-gradient(135deg,#fbeff2,#f9a8d4,#93c5fd)',
    component: SchoolDays,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'shrine_omikuji',
    label: '神社御签',
    preview: '朱红神社御神签',
    gradient: 'linear-gradient(135deg,#7f1d1d,#dc2626,#fde68a)',
    component: ShrineOmikuji,
    ambient: 'amb-dark-red',
  },
  {
    value: 'sumi_ink',
    label: '水墨卷轴',
    preview: '宣纸水墨留白',
    gradient: 'linear-gradient(135deg,#ded3bd,#78716c)',
    component: SumiInk,
    ambient: 'amb-warm',
  },
  {
    value: 'vaporwave',
    label: '蒸汽波',
    preview: '98 复古桌面窗口',
    gradient: 'linear-gradient(135deg,#1a0733,#db2777,#22d3ee)',
    component: Vaporwave,
    ambient: 'amb-dark',
  },
  {
    value: 'villain_throne',
    label: '反派档案',
    preview: '黑红反派通缉档案',
    gradient: 'linear-gradient(135deg,#060406,#7f1d1d,#ef4444)',
    component: VillainThrone,
    ambient: 'amb-dark-red',
  },
  {
    value: 'neo_tokyo',
    label: '赛博东京',
    preview: '霓虹街区招牌档案',
    gradient: 'linear-gradient(135deg,#050510,#a21caf,#22d3ee)',
    component: NeoTokyo,
    ambient: 'amb-dark',
  },
  {
    value: 'star_rail',
    label: '星际列车',
    preview: '深空列车登车票',
    gradient: 'linear-gradient(135deg,#050b1f,#4338ca,#fbbf24)',
    component: StarRail,
    ambient: 'amb-dark',
  },
  {
    value: 'maid_cafe',
    label: '女仆咖啡',
    preview: '奶油粉咖啡菜单',
    gradient: 'linear-gradient(135deg,#fff1f2,#f9a8d4,#7c2d12)',
    component: MaidCafe,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'ninja_scroll',
    label: '忍者密卷',
    preview: '墨黑和纸密卷',
    gradient: 'linear-gradient(135deg,#0c0a09,#57534e,#b91c1c)',
    component: NinjaScroll,
    ambient: 'amb-dark',
  },
  {
    value: 'dragon_shrine',
    label: '龙纹神殿',
    preview: '青金神殿龙纹',
    gradient: 'linear-gradient(135deg,#052e2b,#0f766e,#fbbf24)',
    component: DragonShrine,
    ambient: 'amb-dark-green',
  },
  {
    value: 'backstage_pass',
    label: '后台通行',
    preview: '偶像后台工作证',
    gradient: 'linear-gradient(135deg,#111827,#db2777,#facc15)',
    component: BackstagePass,
    ambient: 'amb-dark',
  },
  {
    value: 'magical_girl',
    label: '魔法少女',
    preview: '变身星光卡',
    gradient: 'linear-gradient(135deg,#fdf2f8,#ec4899,#a855f7)',
    component: MagicalGirl,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'detective_file',
    label: '侦探档案',
    preview: '牛皮纸案件卷宗',
    gradient: 'linear-gradient(135deg,#d6c7a1,#78350f,#111827)',
    component: DetectiveFile,
    ambient: 'amb-warm',
  },
  {
    value: 'sports_club',
    label: '运动社团',
    preview: '青春运动社团证',
    gradient: 'linear-gradient(135deg,#ecfdf5,#059669,#f59e0b)',
    component: SportsClub,
    ambient: 'amb-warm-green',
  },
  {
    value: 'ghost_story',
    label: '怪谈夜话',
    preview: '暗夜怪谈手记',
    gradient: 'linear-gradient(135deg,#050505,#1e1b4b,#84cc16)',
    component: GhostStory,
    ambient: 'amb-dark',
  },
  {
    value: 'summer_festival',
    label: '夏日祭典',
    preview: '浴衣祭典摊位卡',
    gradient: 'linear-gradient(135deg,#1e1b4b,#dc2626,#fbbf24)',
    component: SummerFestival,
    ambient: 'amb-dark',
  },
  {
    value: 'card_battle',
    label: '卡牌对战',
    preview: 'TCG 对战卡面',
    gradient: 'linear-gradient(135deg,#0c0a09,#b45309,#ef4444)',
    component: CardBattle,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'forest_spirit',
    label: '森林精灵',
    preview: '森系精灵图鉴',
    gradient: 'linear-gradient(135deg,#ecfdf5,#15803d,#a3e635)',
    component: ForestSpirit,
    ambient: 'amb-warm-green',
  },
]
