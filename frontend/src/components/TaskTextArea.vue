<template>
  <div class="textarea-wrapper" :class="{ 'has-error': error }">
    <div class="textarea-container">
      <textarea
          ref="textareaRef"
          @input="handleInput"
          @focus="handleFocus"
          @blur="handleBlur"
          :maxlength="maxLength"
          :placeholder="placeholder"
          :aria-label="label"
          :aria-invalid="!!error"
          :aria-describedby="error ? 'error-message' : undefined"
          class="custom-textarea"
          :class="{ 'has-content': modelValue.length > 0 }"
      ></textarea>
      <label class="floating-label">{{ label }}</label>
      <div class="textarea-border"></div>
    </div>

    <div class="textarea-footer">
      <span
          v-if="error"
          class="error-message"
          id="error-message"
      >
        {{ error }}
      </span>
      <span class="char-count" :class="{ 'near-limit': isNearLimit }">
        {{ textLength.length ?? 0}} / {{ maxLength }}
      </span>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, onMounted, computed, defineEmits, defineProps } from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  },
  label: {
    type: String,
    required: true
  },
  placeholder: {
    type: String,
    default: ''
  },
  maxLength: {
    type: Number,
    default: 500
  },
  error: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue'])

const textareaRef = ref(null)
const isFocused = ref(false)

const isNearLimit = computed(() => {
  return props.modelValue.length > props.maxLength * 0.9
})

const count = ref(0)
const textLength = computed({
  get: () => count.value,
  set: val => {
    count.value = val
  }
})

const handleInput = (event) => {
  textLength.value = event.target.value
  emit('update:modelValue', event.target.value)
  adjustHeight()
}

const handleFocus = () => {
  isFocused.value = true
}

const handleBlur = () => {
  isFocused.value = false
}

const adjustHeight = () => {
  const textarea = textareaRef.value
  if (textarea) {
    textarea.style.height = 'auto'
    textarea.style.height = `${textarea.scrollHeight}px`
  }
}

onMounted(() => {
  adjustHeight()
})

watch(() => props.modelValue, () => {
  adjustHeight()
})
</script>

<style scoped>
.textarea-wrapper {
  position: relative;
  width: 100%;
  margin: 1.5rem 0;
  font-family: system-ui, -apple-system, sans-serif;
}

.textarea-container {
  position: relative;
  background: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.textarea-container:hover {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
}

.custom-textarea {
  width: 100%;
  min-height: 400px;
  padding: 1.5rem 1rem 0.75rem;
  font-size: 1rem;
  line-height: 1.5;
  color: #2c3e50;
  background: transparent;
  border: 2px solid #e0e0e0;
  border-radius: 8px;
  resize: none;
  transition: all 0.3s ease;
}

.custom-textarea:focus {
  outline: none;
  border-color: #3498db;
  box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.2);
}

/* Floating Label */
.floating-label {
  position: absolute;
  top: 1.2rem;
  left: 1rem;
  font-size: 1rem;
  color: #666;
  pointer-events: none;
  transition: all 0.25s ease;
  transform-origin: left top;
  background: linear-gradient(180deg, rgba(255,255,255,0) 0%, rgba(255,255,255,1) 50%);
  padding: 0 0.25rem;
}

.custom-textarea:focus ~ .floating-label,
.custom-textarea.has-content ~ .floating-label {
  transform: translateY(-1rem) scale(0.85);
  color: #3498db;
}

/* Border Animation */
.textarea-border {
  position: absolute;
  bottom: 0;
  left: 50%;
  width: 0;
  height: 2px;
  background: #3498db;
  transition: all 0.3s ease;
  transform: translateX(-50%);
}

.custom-textarea:focus ~ .textarea-border {
  width: 100%;
}

/* Custom Scrollbar */
.custom-textarea {
  padding-top: 2%;
  scrollbar-width: thin;
  scrollbar-color: #3498db #e0e0e0;
}

.custom-textarea::-webkit-scrollbar {
  width: 6px;
}

.custom-textarea::-webkit-scrollbar-track {
  background: #e0e0e0;
  border-radius: 3px;
}

.custom-textarea::-webkit-scrollbar-thumb {
  background-color: #3498db;
  border-radius: 3px;
  transition: background-color 0.3s ease;
}

.custom-textarea::-webkit-scrollbar-thumb:hover {
  background-color: #2980b9;
}

/* Footer Styles */
.textarea-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 0.5rem;
  padding: 0 0.25rem;
}

.char-count {
  font-size: 0.875rem;
  color: #666;
  transition: color 0.3s ease;
}

.char-count.near-limit {
  color: #e67e22;
  font-weight: 500;
}

/* Error States */
.has-error .custom-textarea {
  border-color: #e74c3c;
}

.has-error .floating-label {
  color: #e74c3c;
}

.error-message {
  font-size: 0.875rem;
  color: #e74c3c;
}

/* Placeholder Animation */
.custom-textarea::placeholder {
  color: transparent;
  transition: color 0.3s ease;
}

.custom-textarea:focus::placeholder {
  color: #999;
}

/* Focus Ring for Accessibility */
.custom-textarea:focus-visible {
  outline: none;
  box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.4);
}

/* Disabled State */
.custom-textarea:disabled {
  background-color: #f5f5f5;
  cursor: not-allowed;
  opacity: 0.7;
}

/* Dark Mode Support */
@media (prefers-color-scheme: dark) {
  .textarea-container {
    background: #1a1a1a;
  }

  .custom-textarea {
    color: #e0e0e0;
    border-color: #333;
  }

  .floating-label {
    background: linear-gradient(180deg, rgba(26,26,26,0) 0%, rgba(26,26,26,1) 50%);
    color: #999;
  }

  .char-count {
    color: #999;
  }

  .custom-textarea::placeholder {
    color: #666;
  }
}

/* Reduced Motion */
@media (prefers-reduced-motion: reduce) {
  .textarea-container,
  .custom-textarea,
  .floating-label,
  .textarea-border {
    transition: none;
  }
}
</style>