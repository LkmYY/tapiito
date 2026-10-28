<template>
  <div class="space-y-5">
    <template v-for="field in visibleFields" :key="field.key || field.label">
      <!-- 分组 -->
      <section v-if="field.type === 'group'" class="rounded-xl border border-border/70 bg-muted/20 p-4">
        <h4 class="mb-3 text-[13px] font-semibold tracking-tight text-muted-foreground">{{ field.label }}</h4>
        <PayloadEditor
          :fields="field.fields"
          :model-value="modelValue"
          :active-style="activeStyle"
          :path-prefix="pathPrefix + (field.key || 'g') + '/'"
          @update:model-value="(v) => emit('update:modelValue', v)"
        />
      </section>

      <!-- 列表（可嵌套） -->
      <section v-else-if="field.type === 'list'" class="space-y-2.5">
        <div class="flex items-center justify-between">
          <label class="text-sm font-medium">{{ field.label }}</label>
          <span class="text-xs text-muted-foreground">{{ listValue(field).length }} 项</span>
        </div>

        <div v-if="listValue(field).length === 0" class="rounded-lg border border-dashed border-border py-6 text-center">
          <p class="text-xs text-muted-foreground">还没有内容，点击下方按钮添加</p>
        </div>

        <div
          v-for="(item, idx) in listValue(field)"
          :key="idx"
          class="overflow-hidden rounded-xl border border-border bg-card"
        >
          <div class="flex items-center gap-2 px-3 py-2.5">
            <button
              type="button"
              class="flex min-w-0 flex-1 items-center gap-2 text-left"
              @click="toggleItem(field, idx)"
            >
              <ChevronRight
                class="h-3.5 w-3.5 flex-shrink-0 text-muted-foreground transition-transform"
                :class="isOpen(field, idx) ? 'rotate-90' : ''"
              />
              <span class="flex-shrink-0 text-[11px] font-mono text-muted-foreground/70">{{ String(idx + 1).padStart(2, '0') }}</span>
              <span class="truncate text-[13px] font-medium">{{ itemTitle(field, item, idx) }}</span>
            </button>
            <div class="flex flex-shrink-0 items-center gap-0.5">
              <button type="button" class="row-btn" title="上移" :disabled="idx === 0" @click="moveItem(field, idx, -1)">
                <ArrowUp class="h-3.5 w-3.5" />
              </button>
              <button type="button" class="row-btn" title="下移" :disabled="idx === listValue(field).length - 1" @click="moveItem(field, idx, 1)">
                <ArrowDown class="h-3.5 w-3.5" />
              </button>
              <button type="button" class="row-btn hover:!text-red-600" title="删除" @click="removeItem(field, idx)">
                <Trash2 class="h-3.5 w-3.5" />
              </button>
            </div>
          </div>

          <div v-if="isOpen(field, idx)" class="border-t border-border/70 bg-muted/10 p-3.5">
            <PayloadEditor
              :fields="field.fields"
              :model-value="item"
              :active-style="activeStyle"
              :path-prefix="pathPrefix + field.key + ':' + idx + '/'"
              @update:model-value="(v) => replaceItem(field, idx, v)"
            />
          </div>
        </div>

        <button type="button" class="add-btn" @click="addItem(field)">
          <Plus class="h-3.5 w-3.5" />
          {{ field.addLabel || '添加一项' }}
        </button>
      </section>

      <!-- 图片 -->
      <div v-else-if="field.type === 'image'" class="space-y-2">
        <label class="text-sm font-medium">{{ field.label }}</label>
        <div class="flex gap-3">
          <div class="h-16 w-16 flex-shrink-0 overflow-hidden rounded-lg border border-border bg-muted">
            <img
              v-if="strValue(field)"
              :src="strValue(field)"
              alt=""
              class="h-full w-full object-cover"
              @error="(e) => (e.target.style.opacity = 0.25)"
            />
            <div v-else class="flex h-full w-full items-center justify-center">
              <ImageIcon class="h-5 w-5 text-muted-foreground/50" />
            </div>
          </div>
          <div class="min-w-0 flex-1">
            <input
              :value="strValue(field)"
              type="url"
              :placeholder="field.placeholder || '粘贴图片地址 https://…'"
              class="field-input"
              @input="(e) => setField(field, e.target.value)"
            />
            <p class="mt-1 text-[11px] text-muted-foreground">支持外链图片；留空则显示占位图</p>
          </div>
        </div>
      </div>

      <!-- 标签数组 -->
      <div v-else-if="field.type === 'tags'" class="space-y-2">
        <label class="text-sm font-medium">{{ field.label }}</label>
        <div v-if="arrValue(field).length > 0" class="flex flex-wrap gap-1.5">
          <span
            v-for="(tag, ti) in arrValue(field)"
            :key="ti"
            class="inline-flex items-center gap-1 rounded-md bg-secondary px-2 py-1 text-xs text-secondary-foreground"
          >
            <span class="max-w-[220px] truncate">{{ tag }}</span>
            <button type="button" class="text-muted-foreground hover:text-red-600" @click="removeTag(field, ti)">
              <X class="h-3 w-3" />
            </button>
          </span>
        </div>
        <div class="flex gap-2">
          <input
            v-model="tagDraft[field.key]"
            :placeholder="field.placeholder || '输入后按回车添加'"
            class="field-input"
            @keydown.enter.prevent="addTag(field)"
          />
          <button type="button" class="flex-shrink-0 rounded-md border border-input px-3 text-xs font-medium hover:bg-accent" @click="addTag(field)">
            添加
          </button>
        </div>
      </div>

      <!-- 键值对 -->
      <div v-else-if="field.type === 'keyValue'" class="space-y-2">
        <div class="flex items-center justify-between">
          <label class="text-sm font-medium">{{ field.label }}</label>
          <span class="text-xs text-muted-foreground">{{ kvRows(field).length }} 项</span>
        </div>
        <div v-for="(row, ri) in kvRows(field)" :key="ri" class="flex gap-2">
          <input
            :value="row.key"
            :placeholder="field.keyPlaceholder || '名称'"
            class="field-input flex-1"
            @input="(e) => renameKv(field, ri, e.target.value)"
          />
          <input
            :value="row.value"
            :placeholder="field.valuePlaceholder || '内容'"
            class="field-input flex-1"
            @input="(e) => setKv(field, ri, e.target.value)"
          />
          <button type="button" class="row-btn hover:!text-red-600" title="删除" @click="removeKv(field, ri)">
            <Trash2 class="h-3.5 w-3.5" />
          </button>
        </div>
        <button type="button" class="add-btn" @click="addKv(field)">
          <Plus class="h-3.5 w-3.5" /> 添加一项
        </button>
      </div>

      <!-- 开关 -->
      <div v-else-if="field.type === 'toggle'" class="flex items-center justify-between rounded-lg border border-border/70 px-3.5 py-2.5">
        <label class="text-sm font-medium">{{ field.label }}</label>
        <button
          type="button"
          role="switch"
          :aria-checked="Boolean(valueOf(field))"
          class="relative h-5 w-9 flex-shrink-0 rounded-full transition-colors"
          :class="valueOf(field) ? 'bg-primary' : 'bg-muted-foreground/30'"
          @click="setField(field, !valueOf(field))"
        >
          <span
            class="absolute top-0.5 h-4 w-4 rounded-full bg-white shadow-sm transition-all"
            :class="valueOf(field) ? 'left-[18px]' : 'left-0.5'"
          />
        </button>
      </div>

      <!-- 下拉 -->
      <div v-else-if="field.type === 'select'" class="space-y-2">
        <label class="text-sm font-medium">{{ field.label }}</label>
        <select
          :value="valueOf(field)"
          class="field-input"
          @change="(e) => setField(field, coerceSelect(field, e.target.value))"
        >
          <option v-for="opt in normalizedOptions(field)" :key="String(opt.value)" :value="opt.value">{{ opt.label }}</option>
        </select>
      </div>

      <!-- 多行文本 -->
      <div v-else-if="field.type === 'textarea'" class="space-y-2">
        <label class="text-sm font-medium">
          {{ field.label }}
          <span v-if="field.required" class="text-red-500">*</span>
        </label>
        <textarea
          :value="strValue(field)"
          :rows="field.rows || 3"
          :placeholder="field.placeholder || ''"
          class="field-input resize-y leading-relaxed"
          @input="(e) => setField(field, e.target.value)"
        ></textarea>
      </div>

      <!-- 单行文本（含 tel / email / url / price） -->
      <div v-else class="space-y-2">
        <label class="text-sm font-medium">
          {{ field.label }}
          <span v-if="field.required" class="text-red-500">*</span>
        </label>
        <div class="relative">
          <span v-if="field.type === 'price'" class="pointer-events-none absolute left-3 top-1/2 -translate-y-1/2 text-sm text-muted-foreground">¥</span>
          <input
            :value="strValue(field)"
            :type="inputType(field)"
            :placeholder="field.placeholder || ''"
            :class="['field-input', field.type === 'price' ? 'pl-7 tabular-nums' : '']"
            @input="(e) => setField(field, e.target.value)"
          />
        </div>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import {
  ChevronRight, ArrowUp, ArrowDown, Trash2, Plus, X, Image as ImageIcon,
} from 'lucide-vue-next'
import { emptyFromFields } from './schema.js'

defineOptions({ name: 'PayloadEditor' })

const props = defineProps({
  fields: { type: Array, default: () => [] },
  modelValue: { type: Object, default: () => ({}) },
  activeStyle: { type: String, default: '' },
  pathPrefix: { type: String, default: '' },
})
const emit = defineEmits(['update:modelValue'])

const tagDraft = reactive({})
const expanded = reactive({})

/* ---------- 可见字段（按风格过滤分组） ---------- */
const visibleFields = computed(() =>
  (props.fields || []).filter((f) => {
    if (!f) return false
    if (f.type === 'group' && Array.isArray(f.onlyStyles) && f.onlyStyles.length > 0) {
      return f.onlyStyles.includes(props.activeStyle)
    }
    return true
  })
)

/* ---------- 取值 / 赋值 ---------- */
const getPath = (obj, path) => {
  if (!obj || !path) return undefined
  return path.split('.').reduce((acc, k) => (acc == null ? undefined : acc[k]), obj)
}
const setPath = (obj, path, val) => {
  const keys = path.split('.')
  let cur = obj
  for (let i = 0; i < keys.length - 1; i++) {
    const k = keys[i]
    if (cur[k] == null || typeof cur[k] !== 'object') cur[k] = {}
    cur = cur[k]
  }
  cur[keys[keys.length - 1]] = val
  return obj
}

const valueOf = (field) => getPath(props.modelValue, field.key)
const strValue = (field) => {
  const v = valueOf(field)
  return v == null ? '' : String(v)
}
const arrValue = (field) => {
  const v = valueOf(field)
  return Array.isArray(v) ? v : []
}

const setField = (field, val) => {
  setPath(props.modelValue, field.key, val)
  emit('update:modelValue', props.modelValue)
}

/* ---------- 下拉 ---------- */
const normalizedOptions = (field) =>
  (field.options || []).map((o) => (typeof o === 'object' ? o : { value: o, label: String(o) }))

const coerceSelect = (field, raw) => {
  const first = (field.options || [])[0]
  const sample = typeof first === 'object' ? first.value : first
  if (typeof sample === 'number') return Number(raw)
  if (typeof sample === 'boolean') return raw === 'true'
  return raw
}

const inputType = (field) => {
  if (field.type === 'tel') return 'tel'
  if (field.type === 'email') return 'email'
  if (field.type === 'url') return 'url'
  if (field.type === 'price') return 'text'
  return 'text'
}

/* ---------- 标签数组 ---------- */
const addTag = (field) => {
  const draft = (tagDraft[field.key] || '').trim()
  if (!draft) return
  const list = [...arrValue(field)]
  if (!list.includes(draft)) list.push(draft)
  setField(field, list)
  tagDraft[field.key] = ''
}
const removeTag = (field, idx) => {
  const list = [...arrValue(field)]
  list.splice(idx, 1)
  setField(field, list)
}

/* ---------- 键值对 ---------- */
const kvRows = (field) => {
  const v = valueOf(field)
  if (!v || typeof v !== 'object' || Array.isArray(v)) return []
  return Object.entries(v).map(([key, value]) => ({ key, value }))
}
const rebuildKv = (field, rows) => {
  const obj = {}
  for (const r of rows) {
    if (r.key !== '') obj[r.key] = r.value
  }
  setField(field, obj)
}
const renameKv = (field, idx, newKey) => {
  const rows = kvRows(field)
  rows[idx] = { ...rows[idx], key: newKey }
  rebuildKv(field, rows)
}
const setKv = (field, idx, value) => {
  const rows = kvRows(field)
  rows[idx] = { ...rows[idx], value }
  rebuildKv(field, rows)
}
const addKv = (field) => {
  const rows = kvRows(field)
  rows.push({ key: '', value: '' })
  // 空键不入库，先写入一个占位键以便渲染
  const obj = {}
  for (const r of rows) obj[r.key || `__new_${Date.now()}`] = r.value
  setField(field, obj)
}
const removeKv = (field, idx) => {
  const rows = kvRows(field)
  rows.splice(idx, 1)
  rebuildKv(field, rows)
}

/* ---------- 列表 ---------- */
const listValue = (field) => arrValue(field)

const itemTitle = (field, item, idx) => {
  if (!item || typeof item !== 'object') return `${field.itemLabel || '项目'} ${idx + 1}`
  const nameKeys = ['name', 'title', 'shopName', 'companyName', 'salonName', 'gymName', 'schoolName', 'label']
  for (const k of nameKeys) {
    if (item[k]) return String(item[k])
  }
  const firstText = Object.values(item).find((v) => typeof v === 'string' && v.trim())
  return firstText || `${field.itemLabel || '项目'} ${idx + 1}`
}

const openKey = (field, idx) => `${props.pathPrefix}${field.key}:${idx}`
const isOpen = (field, idx) => {
  const k = openKey(field, idx)
  return expanded[k] === undefined ? idx === 0 : expanded[k]
}
const toggleItem = (field, idx) => {
  const k = openKey(field, idx)
  expanded[k] = !isOpen(field, idx)
}

const addItem = (field) => {
  const list = [...listValue(field)]
  const base = field.defaultItem
    ? (typeof field.defaultItem === 'function' ? field.defaultItem() : JSON.parse(JSON.stringify(field.defaultItem)))
    : emptyFromFields(field.fields)
  list.push(base)
  setField(field, list)
  expanded[openKey(field, list.length - 1)] = true
}
const removeItem = (field, idx) => {
  const list = [...listValue(field)]
  list.splice(idx, 1)
  setField(field, list)
}
const moveItem = (field, idx, delta) => {
  const list = [...listValue(field)]
  const target = idx + delta
  if (target < 0 || target >= list.length) return
  const [moved] = list.splice(idx, 1)
  list.splice(target, 0, moved)
  setField(field, list)
}
const replaceItem = (field, idx, val) => {
  const list = [...listValue(field)]
  list[idx] = val
  setField(field, list)
}

/* ---------- 校验（供父组件调用） ---------- */
const validate = () => {
  const errors = []
  const walk = (fields, obj, pathLabel) => {
    for (const f of fields || []) {
      if (!f) continue
      if (f.type === 'group') {
        if (Array.isArray(f.onlyStyles) && f.onlyStyles.length > 0 && !f.onlyStyles.includes(props.activeStyle)) continue
        walk(f.fields, obj, pathLabel)
        continue
      }
      if (f.type === 'list') {
        const list = getPath(obj, f.key)
        if (Array.isArray(list)) {
          list.forEach((it, i) => walk(f.fields, it, `${pathLabel}${f.label} ${i + 1} · `))
        }
        continue
      }
      if (!f.required) continue
      const v = getPath(obj, f.key)
      if (v == null || String(v).trim() === '') {
        errors.push(`${pathLabel}${f.label} 不能为空`)
      }
    }
  }
  walk(visibleFields.value, props.modelValue, '')
  return errors
}

defineExpose({ validate })
</script>

<style scoped>
.field-input {
  display: flex;
  width: 100%;
  border-radius: 0.375rem;
  border: 1px solid hsl(var(--input, 214 32% 91%));
  background-color: hsl(var(--background, 0 0% 100%));
  padding: 0.5rem 0.75rem;
  font-size: 0.875rem;
  line-height: 1.25rem;
  outline: none;
  transition: border-color 0.15s, box-shadow 0.15s;
}
.field-input::placeholder {
  color: hsl(var(--muted-foreground, 215 20% 65%));
}
.field-input:focus-visible {
  border-color: hsl(var(--ring, 222 84% 5%));
  box-shadow: 0 0 0 2px hsl(var(--ring, 222 84% 5%) / 0.12);
}
.row-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 1.75rem;
  width: 1.75rem;
  border-radius: 0.375rem;
  color: hsl(var(--muted-foreground, 215 20% 65%));
  transition: background-color 0.15s, color 0.15s;
}
.row-btn:hover:not(:disabled) {
  background-color: hsl(var(--accent, 210 40% 96%));
}
.row-btn:disabled {
  opacity: 0.35;
}
.add-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.375rem;
  width: 100%;
  border-radius: 0.5rem;
  border: 1px dashed hsl(var(--border, 214 32% 91%));
  padding: 0.5rem;
  font-size: 0.8125rem;
  font-weight: 500;
  color: hsl(var(--muted-foreground, 215 20% 65%));
  transition: background-color 0.15s, color 0.15s, border-color 0.15s;
}
.add-btn:hover {
  border-color: hsl(var(--primary, 222 47% 11%) / 0.4);
  background-color: hsl(var(--accent, 210 40% 96%));
  color: hsl(var(--foreground, 222 84% 5%));
}
</style>
