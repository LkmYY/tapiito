/**
 * Schema DSL —— 描述「可视化表单」的结构。
 *
 * 设计目标：新增模板类型时，只需写一个 manifest（字段结构 + 风格 + 示例），
 * 表单、校验、JSON 互转全部自动生成，无需改动渲染器。
 *
 * 每个字段是一个纯对象：
 *   { key, type, label, ...options }
 * 渲染器 PayloadEditor.vue 按 type 分发到具体控件。
 */

/** 单行文本 */
export const text = (key, label, opts = {}) => ({ key, label, type: 'text', ...opts })

/** 多行文本 */
export const textarea = (key, label, opts = {}) => ({ key, label, type: 'textarea', ...opts })

/** 电话（tel 链接 + 宽松校验） */
export const tel = (key, label, opts = {}) => ({ key, label, type: 'tel', ...opts })

/** 邮箱 */
export const email = (key, label, opts = {}) => ({ key, label, type: 'email', ...opts })

/** 网址（校验 http/https） */
export const url = (key, label, opts = {}) => ({ key, label, type: 'url', ...opts })

/** 价格：以字符串保存，兼容现有模板组件（它们直接拼接 ¥{{ price }}） */
export const price = (key, label, opts = {}) => ({ key, label, type: 'price', ...opts })

/** 开关 */
export const toggle = (key, label, opts = {}) => ({ key, label, type: 'toggle', ...opts })

/** 下拉选择：options 支持 [1,2,3] 或 [{value,label}] */
export const select = (key, label, options, opts = {}) => ({ key, label, type: 'select', options, ...opts })

/** 图片地址：URL 输入 + 缩略图预览 */
export const image = (key, label, opts = {}) => ({ key, label, type: 'image', ...opts })

/** 字符串数组：标签式增删（特点、设施、相册） */
export const tags = (key, label, opts = {}) => ({ key, label, type: 'tags', ...opts })

/** 键值对（产品规格：{ 屏幕: '1.78英寸' }） */
export const keyValue = (key, label, opts = {}) => ({ key, label, type: 'keyValue', ...opts })

/**
 * 对象数组（联系人、菜单分类、课程…），支持任意嵌套。
 * opts.fields  子字段 schema
 * opts.itemLabel 单项标题，用于折叠卡片头部
 * opts.addLabel  「新增」按钮文案
 * opts.defaultItem 新增一项时的初始值（函数或对象）
 */
export const list = (key, label, opts = {}) => ({ key, label, type: 'list', ...opts })

/** 分组：把若干字段收纳成一个小节（纯视觉分组，不改变数据结构） */
export const group = (label, fields, opts = {}) => ({ type: 'group', label, fields, ...opts })

/**
 * 辣度选项：菜谱类型用
 */
export const SPICY_OPTIONS = [
  { value: 0, label: '不辣' },
  { value: 1, label: '微辣' },
  { value: 2, label: '中辣' },
  { value: 3, label: '特辣' },
]

/** 按 schema 生成一个空对象（用于「新增一项」） */
export function emptyFromFields(fields = []) {
  const out = {}
  for (const f of fields) {
    if (f.type === 'group') {
      Object.assign(out, emptyFromFields(f.fields))
      continue
    }
    if (f.type === 'list') {
      out[f.key] = []
    } else if (f.type === 'toggle') {
      out[f.key] = false
    } else if (f.type === 'select') {
      const first = f.options?.[0]
      out[f.key] = first && typeof first === 'object' ? first.value : first ?? ''
    } else if (f.type === 'keyValue') {
      out[f.key] = {}
    } else {
      out[f.key] = ''
    }
  }
  return out
}
