<template>
  <div class="select-container" :class="{ 'is-open': isOpen }">
    <!-- Select Button -->
    <button
        type="button"
        class="select-button"
        :class="{ 'has-value': !!selectedOption }"
        @click="toggleDropdown"
        @blur="handleBlur"
    >
      <div class="select-value">
        <div v-if="selectedOption" class="selected-option">
          {{ selectedOption.name }}
        </div>
        <div v-else class="placeholder">Выберите нейронную сеть</div>
      </div>

      <div class="select-icon">
        <svg
            class="chevron-icon"
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
        >
          <polyline points="6 9 12 15 18 9"></polyline>
        </svg>
      </div>
    </button>

    <!-- Dropdown -->
    <Transition name="dropdown">
      <div v-if="isOpen" class="select-dropdown">
        <!-- Options List -->
        <div class="options-container">
          <button
              v-for="option in props.options"
              :key="option.id"
              class="option-item"
              :class="{ 'is-selected': isSelected(option) }"
              @click="selectOption(option)"
          >
            {{ option.name}}
            <svg
                v-if="isSelected(option)"
                class="check-icon"
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
            >
              <polyline points="20 6 9 17 4 12"></polyline>
            </svg>
          </button>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, defineProps, defineEmits } from 'vue'

const props = defineProps({
  options: {
    type: Array,
    required: true
  }
})

const emit = defineEmits(['select'])

// Состояния
const isOpen = ref(false)
const selectedOption = ref(null)

// Методы
const toggleDropdown = () => {
  isOpen.value = !isOpen.value
}

const handleBlur = (event) => {
  const relatedTarget = event.relatedTarget
  if (!relatedTarget || !event.currentTarget.contains(relatedTarget)) {
    isOpen.value = false
  }
}

const selectOption = (option) => {
  selectedOption.value = option
  emit('select', option)
  isOpen.value = false
}

const isSelected = (option) => {
  return selectedOption.value?.value === option.value
}
</script>

<style scoped>
.select-container {
  position: absolute;
  top: 80px;
  width: 100%;
}

.select-button {
  width: 100%;
  min-height: 2.5rem;
  padding: 0.5rem 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 0.5rem;
  color: #1f2937;
  cursor: pointer;
  transition: all 0.2s ease;
}

.select-button:hover {
  border-color: #d1d5db;
}

.select-button:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.select-value {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  min-width: 0;
}

.placeholder {
  color: #9ca3af;
}

.selected-option {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.option-icon {
  width: 1.25rem;
  height: 1.25rem;
  object-fit: contain;
}

.chevron-icon {
  color: #6b7280;
  transition: transform 0.2s ease;
}

.is-open .chevron-icon {
  transform: rotate(180deg);
}

.select-dropdown {
  position: absolute;
  top: calc(100% + 0.5rem);
  left: 0;
  width: 100%;
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 0.5rem;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1),
  0 4px 6px -2px rgba(0, 0, 0, 0.05);
  z-index: 50;
  overflow: hidden;
}

.options-container {
  max-height: 15rem;
  overflow-y: auto;
}

.option-item {
  width: 100%;
  padding: 0.75rem 1rem;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background: transparent;
  border: none;
  color: #1f2937;
  cursor: pointer;
  text-align: left;
  transition: background-color 0.2s ease;
}

.option-item:hover,
.option-item:focus {
  background-color: #f3f4f6;
  outline: none;
}

.option-item.is-selected {
  background-color: #eff6ff;
  color: #3b82f6;
  font-weight: 500;
}

.check-icon {
  margin-left: auto;
  color: #3b82f6;
}

/* Анимации */
.dropdown-enter-active,
.dropdown-leave-active {
  transition: all 0.2s ease;
}

.dropdown-enter-from,
.dropdown-leave-to {
  opacity: 0;
  transform: translateY(-0.5rem);
}

/* Scrollbar Styling */
.options-container::-webkit-scrollbar {
  width: 6px;
}

.options-container::-webkit-scrollbar-track {
  background: #f1f1f1;
}

.options-container::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 3px;
}

.options-container::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}

@media (max-width: 640px) {
  .select-button {
    min-height: 2.75rem;
  }

  .option-item {
    padding: 0.875rem 1rem;
  }
}
</style>